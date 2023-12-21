package service

import (
	"fmt"
	"kvm-agent/internal/log"
	"kvm-agent/internal/runner/models/request"
	"os/exec"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	dm "github.com/nfjBill/gorm-driver-dm"
	"gorm.io/gorm"
)

type DefaultService struct {
	context *gin.Context
}

func NewDefaultService(c *gin.Context) *DefaultService {
	return &DefaultService{
		context: c,
	}
}

func Ping(host string, count, size, interval int) (string, error) {
	cmd := exec.Command("ping", "-c", strconv.Itoa(count), "-s", strconv.Itoa(size), "-i", strconv.Itoa(interval), host)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	result := string(output)

	// result := &PingResult{}
	// result.PacketsSent, result.PacketsReceived, result.PacketLoss, result.AvgRtt, result.MinRtt, result.MaxRtt = parsePingOutput(string(output))

	return result, nil
}

func (s *DefaultService) HandlePing(pingPlug request.PingPlug) string {
	result, err := Ping(pingPlug.Host, pingPlug.Count, pingPlug.Size, pingPlug.Interval)
	if err != nil {
		log.Errorf("HandlePing", "Ping error: %v", err)

		return fmt.Sprintf("Ping error: %v", err)
	}

	return result
}

func FIO(dir, batch, size string) (string, error) {
	cmd := exec.Command("fio", "-directory", dir, "-direct", "1", "-iodepth", "1", "-thread", "-ioengine", "libaio", "-randrepeat", "0", "-bs", batch,
		"-size", size, "-group_reporting", "-rw", "read", "-name", "4k-write")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	result := string(output)

	// result := &PingResult{}
	// result.PacketsSent, result.PacketsReceived, result.PacketLoss, result.AvgRtt, result.MinRtt, result.MaxRtt = parsePingOutput(string(output))

	return result, nil
}

func (s *DefaultService) HandleFIO(fioPlug request.FIOPlug) string {
	result, err := FIO(fioPlug.Dir, fioPlug.Batch, fioPlug.Size)
	if err != nil {
		log.Errorf("HandleFIO", "FIO error: %v", err)

		return fmt.Sprintf("FIO error: %v", err)
	}

	return result
}

func PTP4L(iface string) (string, error) {
	cmd := exec.Command("ptp4l", "-m", "-i", iface)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	result := string(output)

	return result, nil
}

func (s *DefaultService) HandlePTP4L(ptp4lPlug request.PTP4LPlug) string {
	result, err := PTP4L(ptp4lPlug.Interface)
	if err != nil {
		log.Errorf("HandlePTP4L", "PTP4L error: %v", err)

		return fmt.Sprintf("PTP4L error: %v", err)
	}

	return result
}

func (s *DefaultService) HandleDBTest(dbTestPlug request.DBTestPlug) string {
	dsn := fmt.Sprintf("dm://%s:%s@%s:%d?compatibleMode=mysql",
		dbTestPlug.Username,
		dbTestPlug.Password,
		dbTestPlug.Host,
		dbTestPlug.Port,
	)
	db, err := gorm.Open(dm.Open(dsn))
	if err != nil {
		log.Panic("DM", err.Error(), " open failed ", dsn)
	}
	defer func(db *gorm.DB) {
		sqlDB, err := db.DB()
		if err != nil {
			log.Error("DM", err.Error(), " db.DB() failed ")
		} else {
			err = sqlDB.Close()
			if err != nil {
				log.Error("DM", err.Error(), " sqlDB.Close() failed ")
			}
		}
	}(db)

	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	db.Set("gorm:table_options", "collation=utf8_unicode_ci")

	// Create
	type Test struct {
		gorm.Model
		Column1 string
		Column2 int
		Column3 float64
	}
	if db.Migrator().HasTable(&Test{}) && db.Migrator().DropTable(&Test{}) != nil {
		log.Panicf("KVM-Agent Table Migration", "Drop table %T... failed", Test{})
	}
	db.AutoMigrate(&Test{})

	start := time.Now()
	for i := 0; i < dbTestPlug.Count; i++ {
		test := Test{
			Column1: "Value1",
			Column2: 123,
			Column3: 3.14,
		}
		db.Create(&test)
	}
	createElapsed := time.Since(start)

	start = time.Now()
	for i := 0; i < dbTestPlug.Count; i++ {
		var test Test
		db.First(&test, i+1)
	}
	readElapsed := time.Since(start)

	start = time.Now()
	var wg sync.WaitGroup
	count := dbTestPlug.Count
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var test Test
			db.First(&test)
		}()
	}
	wg.Wait()
	concurrentElapsed := time.Since(start)

	start = time.Now()
	for i := 0; i < dbTestPlug.Count; i++ {
		var test Test
		db.Delete(&test, i+1)
	}
	deleteElapsed := time.Since(start)

	return fmt.Sprintf("Create: %v, Read: %v, Delete: %v, ConcurrentElapsed: %v", createElapsed, readElapsed, deleteElapsed, concurrentElapsed)
}
