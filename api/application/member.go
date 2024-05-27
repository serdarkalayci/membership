// Package application is the package that holds the application logic between database and communication layers
package application

import (
	"github.com/serdarkalayci/membership/api/domain"
)

// MemberRepository is the interface to interact with Member domain object
type MemberRepository interface {
	ListMembers(pageSize, pageNum int) ([]domain.Member, int64, error)
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
func (ms MemberService) ListMembers(pageSize, pageNum int) ([]domain.Member, int64, error) {
	switch {
		case pageSize <= 0 : 
			pageSize = 10
		case pageSize > 100 :
			pageSize = 100
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	return ms.dc.GetMemberRepository().ListMembers(pageSize, pageNum)
}

// GetMember simply returns the member with the given id or an error that is returned from the repository
func (ms MemberService) GetMember(id string) (domain.Member, error) {
	return ms.dc.GetMemberRepository().GetMember(id)
}