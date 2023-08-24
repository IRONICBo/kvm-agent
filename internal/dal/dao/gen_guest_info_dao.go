package dao

import (
	"context"
	"time"

	"kvm-agent/internal/conn"
	"kvm-agent/internal/dal/cache"
	"kvm-agent/internal/dal/gen"
	models "kvm-agent/internal/models"
)

// GuestInfoDao guestinfo dao.
type GuestInfoDao struct {
	Dao
}

// NewGuestInfoDao return a guestinfo dao.
func NewGuestInfoDao() *GuestInfoDao {
	query := gen.Use(conn.GetDMDB())
	cache := cache.Use(conn.GetRedis())

	return &GuestInfoDao{
		Dao: Dao{
			ctx:   context.Background(),
			query: query,
			cache: &cache,
		},
	}
}

// Create create one or multi models.
func (d *GuestInfoDao) Create(m ...*models.GuestInfo) error {
	return d.query.WithContext(d.ctx).GuestInfo.Create(m...)
}

// First get first matched result.
func (d *GuestInfoDao) First() (*models.GuestInfo, error) {
	return d.query.WithContext(d.ctx).GuestInfo.First()
}

// FindAll get all matched results.
func (d *GuestInfoDao) FindAll() ([]*models.GuestInfo, error) {
	return d.query.WithContext(d.ctx).GuestInfo.Find()
}

// FindFirstById get first matched result by id.
func (d *GuestInfoDao) FindFirstById(id uint) (*models.GuestInfo, error) {
	m := d.query.GuestInfo

	return m.WithContext(d.ctx).Where(m.Id.Eq(id)).First()
}

// FindByIdPage get page by Id.
func (d *GuestInfoDao) FindByIdPage(id uint, offset int, limit int) ([]*models.GuestInfo, int64, error) {
	m := d.query.GuestInfo

	result, count, err := m.WithContext(d.ctx).Where(m.Id.Eq(id)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByCreatedAt get first matched result by createdat.
func (d *GuestInfoDao) FindFirstByCreatedAt(createdat time.Time) (*models.GuestInfo, error) {
	m := d.query.GuestInfo

	return m.WithContext(d.ctx).Where(m.CreatedAt.Eq(createdat)).First()
}

// FindByCreatedAtPage get page by CreatedAt.
func (d *GuestInfoDao) FindByCreatedAtPage(createdat time.Time, offset int, limit int) ([]*models.GuestInfo, int64, error) {
	m := d.query.GuestInfo

	result, count, err := m.WithContext(d.ctx).Where(m.CreatedAt.Eq(createdat)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByUpdatedAt get first matched result by updatedat.
func (d *GuestInfoDao) FindFirstByUpdatedAt(updatedat time.Time) (*models.GuestInfo, error) {
	m := d.query.GuestInfo

	return m.WithContext(d.ctx).Where(m.UpdatedAt.Eq(updatedat)).First()
}

// FindByUpdatedAtPage get page by UpdatedAt.
func (d *GuestInfoDao) FindByUpdatedAtPage(updatedat time.Time, offset int, limit int) ([]*models.GuestInfo, int64, error) {
	m := d.query.GuestInfo

	result, count, err := m.WithContext(d.ctx).Where(m.UpdatedAt.Eq(updatedat)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByDeletedAt get first matched result by deletedat.
func (d *GuestInfoDao) FindFirstByDeletedAt(deletedat time.Time) (*models.GuestInfo, error) {
	m := d.query.GuestInfo

	return m.WithContext(d.ctx).Where(m.DeletedAt.Eq(deletedat)).First()
}

// FindByDeletedAtPage get page by DeletedAt.
func (d *GuestInfoDao) FindByDeletedAtPage(deletedat time.Time, offset int, limit int) ([]*models.GuestInfo, int64, error) {
	m := d.query.GuestInfo

	result, count, err := m.WithContext(d.ctx).Where(m.DeletedAt.Eq(deletedat)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByUUID get first matched result by uuid.
func (d *GuestInfoDao) FindFirstByUUID(uuid string) (*models.GuestInfo, error) {
	m := d.query.GuestInfo

	return m.WithContext(d.ctx).Where(m.UUID.Eq(uuid)).First()
}

// FindByUUIDPage get page by UUID.
func (d *GuestInfoDao) FindByUUIDPage(uuid string, offset int, limit int) ([]*models.GuestInfo, int64, error) {
	m := d.query.GuestInfo

	result, count, err := m.WithContext(d.ctx).Where(m.UUID.Eq(uuid)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByHostDesc get first matched result by hostdesc.
func (d *GuestInfoDao) FindFirstByHostDesc(hostdesc string) (*models.GuestInfo, error) {
	m := d.query.GuestInfo

	return m.WithContext(d.ctx).Where(m.HostDesc.Eq(hostdesc)).First()
}

// FindByHostDescPage get page by HostDesc.
func (d *GuestInfoDao) FindByHostDescPage(hostdesc string, offset int, limit int) ([]*models.GuestInfo, int64, error) {
	m := d.query.GuestInfo

	result, count, err := m.WithContext(d.ctx).Where(m.HostDesc.Eq(hostdesc)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByCpuDesc get first matched result by cpudesc.
func (d *GuestInfoDao) FindFirstByCpuDesc(cpudesc string) (*models.GuestInfo, error) {
	m := d.query.GuestInfo

	return m.WithContext(d.ctx).Where(m.CpuDesc.Eq(cpudesc)).First()
}

// FindByCpuDescPage get page by CpuDesc.
func (d *GuestInfoDao) FindByCpuDescPage(cpudesc string, offset int, limit int) ([]*models.GuestInfo, int64, error) {
	m := d.query.GuestInfo

	result, count, err := m.WithContext(d.ctx).Where(m.CpuDesc.Eq(cpudesc)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByMemDesc get first matched result by memdesc.
func (d *GuestInfoDao) FindFirstByMemDesc(memdesc string) (*models.GuestInfo, error) {
	m := d.query.GuestInfo

	return m.WithContext(d.ctx).Where(m.MemDesc.Eq(memdesc)).First()
}

// FindByMemDescPage get page by MemDesc.
func (d *GuestInfoDao) FindByMemDescPage(memdesc string, offset int, limit int) ([]*models.GuestInfo, int64, error) {
	m := d.query.GuestInfo

	result, count, err := m.WithContext(d.ctx).Where(m.MemDesc.Eq(memdesc)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByDiskDesc get first matched result by diskdesc.
func (d *GuestInfoDao) FindFirstByDiskDesc(diskdesc string) (*models.GuestInfo, error) {
	m := d.query.GuestInfo

	return m.WithContext(d.ctx).Where(m.DiskDesc.Eq(diskdesc)).First()
}

// FindByDiskDescPage get page by DiskDesc.
func (d *GuestInfoDao) FindByDiskDescPage(diskdesc string, offset int, limit int) ([]*models.GuestInfo, int64, error) {
	m := d.query.GuestInfo

	result, count, err := m.WithContext(d.ctx).Where(m.DiskDesc.Eq(diskdesc)).FindByPage(offset, limit)

	return result, count, err
}

// FindFirstByNetDesc get first matched result by netdesc.
func (d *GuestInfoDao) FindFirstByNetDesc(netdesc string) (*models.GuestInfo, error) {
	m := d.query.GuestInfo

	return m.WithContext(d.ctx).Where(m.NetDesc.Eq(netdesc)).First()
}

// FindByNetDescPage get page by NetDesc.
func (d *GuestInfoDao) FindByNetDescPage(netdesc string, offset int, limit int) ([]*models.GuestInfo, int64, error) {
	m := d.query.GuestInfo

	result, count, err := m.WithContext(d.ctx).Where(m.NetDesc.Eq(netdesc)).FindByPage(offset, limit)

	return result, count, err
}

// Update update model.
func (d *GuestInfoDao) Update(m *models.GuestInfo) error {
	q := d.query.GuestInfo
	res, err := d.query.WithContext(d.ctx).GuestInfo.Where(q.UUID.Eq(m.UUID)).Updates(m)
	if err != nil && res.Error != nil {
		return err
	}

	return nil
}

// Delete delete model.
func (d *GuestInfoDao) Delete(m ...*models.GuestInfo) error {
	res, err := d.query.WithContext(d.ctx).GuestInfo.Delete(m...)
	if err != nil && res.Error != nil {
		return err
	}

	return nil
}

// Count count matched records.
func (d *GuestInfoDao) Count() (int64, error) {
	return d.query.WithContext(d.ctx).GuestInfo.Count()
}

///////////////////////////////////////////////////////////
//              Append your code here.                   //
///////////////////////////////////////////////////////////
