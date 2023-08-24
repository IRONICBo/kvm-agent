package conn

import (
	"fmt"
	"kvm-agent/internal/config"
	"kvm-agent/internal/log"
	"time"

	dm "github.com/nfjBill/gorm-driver-dm"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var d *gorm.DB

type writer struct{}

// Write implement log writer interface.
func (w writer) Printf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

// InitDMDB init mysql connection.
func InitDMDB(config config.DM, debug bool) {
	// try to use default database [mysql]
	dsn := fmt.Sprintf("dm://%s:%s@%s:%d?compatibleMode=mysql",
		config.Username,
		config.Password,
		config.Ip,
		config.Port,
		// "SYSDBA", // sys database
	)

	cfg := logger.Config{
		IgnoreRecordNotFoundError: true, // Ignore ErrRecordNotFound error for logger
		Colorful:                  true, // Disable color
	}
	if debug {
		cfg = logger.Config{
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true, // Ignore ErrRecordNotFound error for logger
			Colorful:                  true, // Disable color
		}
	}

	logger := logger.New(
		writer{},
		cfg,
	)

	// connect to dm
	db, err := gorm.Open(dm.Open(dsn), &gorm.Config{
		Logger: logger,
	})
	if err != nil {
		// retry
		time.Sleep(time.Duration(10) * time.Second)
		db, err = gorm.Open(dm.Open(dsn), &gorm.Config{
			Logger: logger,
		})
		if err != nil {
			log.Panic("DM", err.Error(), " open failed ", dsn)
		}
	}

	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	db.Set("gorm:table_options", "collation=utf8_unicode_ci")

	d = db
	log.Info("DM", "connect ok", dsn)
}

// GetDMDB get dm connection.
func GetDMDB() *gorm.DB {
	return d
}

// CloseDMDB close dm connection.
func CloseDMDB() {
	if d == nil {
		return
	}

	sqlDB, err := d.DB()
	if err != nil {
		log.Error("DM", err.Error(), " db.DB() failed ")
	} else {
		err = sqlDB.Close()
		if err != nil {
			log.Error("DM", err.Error(), " sqlDB.Close() failed ")
		}
	}
}
