package manager

import (
	"sync"
	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/internal/domain"
)

// DriverMgr manages drivers in the system (Singleton pattern)
type DriverMgr struct {
	driversMap map[string]*domain.Driver
	mu         sync.RWMutex
}

var (
	driverMgrInstance *DriverMgr
	driverMgrOnce     sync.Once
)

// GetDriverMgr returns the singleton instance of DriverMgr
func GetDriverMgr() *DriverMgr {
	driverMgrOnce.Do(func() {
		driverMgrInstance = &DriverMgr{
			driversMap: make(map[string]*domain.Driver),
		}
	})
	return driverMgrInstance
}

// AddDriver adds a driver to the map
func (d *DriverMgr) AddDriver(driverName string, driver *domain.Driver) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.driversMap[driverName] = driver
}

// GetDriver retrieves a driver by name
func (d *DriverMgr) GetDriver(driverName string) *domain.Driver {
	d.mu.RLock()
	defer d.mu.RUnlock()
	return d.driversMap[driverName]
}

// GetDriversMap returns the drivers map
func (d *DriverMgr) GetDriversMap() map[string]*domain.Driver {
	d.mu.RLock()
	defer d.mu.RUnlock()
	// Return a copy to prevent external modifications
	result := make(map[string]*domain.Driver)
	for k, v := range d.driversMap {
		result[k] = v
	}
	return result
}

