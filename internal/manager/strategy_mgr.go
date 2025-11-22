package manager

import (
	"fmt"
	"sync"

	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/internal/domain"
	strategy "github.com/kamalpratik/Uber-Ola-Low-Level-Design/internal/usecase"
	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/pkg/common"
	"github.com/kamalpratik/Uber-Ola-Low-Level-Design/pkg/interfaces"
)

// StrategyMgr manages strategies in the system (Singleton pattern)
type StrategyMgr struct {
	driverMgr *DriverMgr
}

var (
	strategyMgrInstance *StrategyMgr
	strategyMgrOnce     sync.Once
)

// GetStrategyMgr returns the singleton instance of StrategyMgr
func GetStrategyMgr() *StrategyMgr {
	strategyMgrOnce.Do(func() {
		strategyMgrInstance = &StrategyMgr{
			driverMgr: GetDriverMgr(),
		}
	})
	return strategyMgrInstance
}

// DeterminePricingStrategy determines and returns a pricing strategy
func (s *StrategyMgr) DeterminePricingStrategy(metaData *domain.TripMetaData) interfaces.PricingStrategy {
	fmt.Println("Based on location and other factors, setting pricing strategy")
	//strategyPattern
	// Use rating-based pricing if rider has a valid rating
	// Use default pricing if rating is unassigned
	riderRating := metaData.GetRiderRating()
	if riderRating != common.Unassigned {
		fmt.Println("Using rating-based pricing strategy")
		return strategy.NewRatingBasedPricingStrategy()
	}

	fmt.Println("Using default pricing strategy")
	return strategy.NewDefaultPricingStrategy()
}

// DetermineMatchingStrategy determines and returns a driver matching strategy
func (s *StrategyMgr) DetermineMatchingStrategy(metaData *domain.TripMetaData) interfaces.DriverMatchingStrategy {
	fmt.Println("Based on location and other factors, setting matching strategy")
	return strategy.NewLeastTimeBasedMatchingStrategy(s.driverMgr)
}
