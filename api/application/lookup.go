// Package application is the package that holds the application logic between database and communication layers
package application

import (
	"github.com/serdarkalayci/membership/api/domain"
)

// LookupRepository is the interface to interact with Lookup objects
type LookupRepository interface {
	ListCities() ([]domain.City, error)
	ListAreas() ([]domain.Area, error)
	ListProvinces() ([]domain.Province, error)
	ListMembershipTypes() ([]domain.MembershipType, error)	
}

// LookupService is the struct to let outer layers to interact to the Lookup Applicatopn
type LookupService struct {
	dc DataContextCarrier
}

// NewLookupService creates a new LookupService instance and sets its repository
func NewLookupService(dc DataContextCarrier) LookupService {
	return LookupService{
		dc: dc,
	}
}

// ListCities simply returns the whole list of cities or an error that is returned from the repository
func (ls LookupService) ListCities() ([]domain.City, error) {
	return ls.dc.GetLookupRepository().ListCities()
}

// ListAreas simply returns the whole list of areas or an error that is returned from the repository
func (ls LookupService) ListAreas() ([]domain.Area, error) {
	return ls.dc.GetLookupRepository().ListAreas()
}

// ListProvinces simply returns the whole list of provinces or an error that is returned from the repository
func (ls LookupService) ListProvinces() ([]domain.Province, error) {
	return ls.dc.GetLookupRepository().ListProvinces()
}

// ListMembershipTypes simply returns the whole list of membership types or an error that is returned from the repository
func (ls LookupService) ListMembershipTypes() ([]domain.MembershipType, error) {
	return ls.dc.GetLookupRepository().ListMembershipTypes()
}
