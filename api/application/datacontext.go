// Package application is the package that holds the application logic between database and communication layers
package application

// DataContextCarrier is the interface to be passed to the application layer
type DataContextCarrier interface {
	SetRepositories(ur UserRepository, mr MemberRepository)
	GetUserRepository() UserRepository
	GetMemberRepository() MemberRepository
}

// DataContext represents a struct that holds concrete repositories
type DataContext struct {
	userRepository         UserRepository
	memberRepository       MemberRepository
}

// SetRepositories sets the repositories of the datacontext
func (dc *DataContext) SetRepositories(ur UserRepository, mr MemberRepository) {
	dc.userRepository = ur
	dc.memberRepository = mr
}

// GetUserRepository returns the user repository
func (dc *DataContext) GetUserRepository() UserRepository {
	return dc.userRepository
}

// GetMemberRepository returns the member repository
func (dc *DataContext) GetMemberRepository() MemberRepository {
	return dc.memberRepository
}