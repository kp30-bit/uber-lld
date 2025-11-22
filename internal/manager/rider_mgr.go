package manager

import (
	"sync"
	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/internal/domain"
)

// RiderMgr manages riders in the system (Singleton pattern)
type RiderMgr struct {
	ridersMap map[string]*domain.Rider
	mu        sync.RWMutex
}

var (
	riderMgrInstance *RiderMgr
	riderMgrOnce     sync.Once
)

// GetRiderMgr returns the singleton instance of RiderMgr
func GetRiderMgr() *RiderMgr {
	riderMgrOnce.Do(func() {
		riderMgrInstance = &RiderMgr{
			ridersMap: make(map[string]*domain.Rider),
		}
	})
	return riderMgrInstance
}

// AddRider adds a rider to the map
func (r *RiderMgr) AddRider(riderName string, rider *domain.Rider) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.ridersMap[riderName] = rider
}

// GetRider retrieves a rider by name
func (r *RiderMgr) GetRider(riderName string) *domain.Rider {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.ridersMap[riderName]
}

