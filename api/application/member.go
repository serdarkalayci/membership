// Package application is the package that holds the application logic between database and communication layers
package application

import (
	"github.com/serdarkalayci/membership/api/domain"
)

// MemberRepository is the interface to interact with Member domain object
type MemberRepository interface {
	ListMembers() ([]domain.Member, error)
	GetMember(id string) (domain.Member, error)
}

// MemberService is the struct to let outer layers to interact to the Member Applicatopn
type MemberService struct {
	dc DataContextCarrier
}

// NewMemberService creates a new MemberService instance and sets its repository
func NewMemberService(dc DataContextCarrier) MemberService {
	return MemberService{
		dc: dc,
	}
}

// ListMembers simply returns the whole list of member or an error that is returned from the repository
func (ms MemberService) ListMembers() ([]domain.Member, error) {
	return ms.dc.GetMemberRepository().ListMembers()
}

// GetMember simply returns the member with the given id or an error that is returned from the repository
func (ms MemberService) GetMember(id string) (domain.Member, error) {
	return ms.dc.GetMemberRepository().GetMember(id)
}