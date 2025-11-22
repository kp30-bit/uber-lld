package interfaces

import "github.com/kamalpratik/Uber-Ola-Low-Level-Design/internal/domain"

// DriverProvider provides access to drivers for matching strategies
// This interface breaks the circular dependency between strategy and manager packages
type DriverProvider interface {
	GetDriversMap() map[string]*domain.Driver
}

