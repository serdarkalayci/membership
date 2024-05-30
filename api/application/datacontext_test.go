package application

// MockContext mimics the DataContext struct, which is used to test the application layer
type MockContext struct {
	userRepository         UserRepository
	MemberRepository	   MemberRepository
}

func (mc *MockContext) SetRepositories(ur UserRepository, mr MemberRepository) {
	mc.userRepository = ur
}

func (mc *MockContext) GetUserRepository() UserRepository {
	return mc.userRepository
}

func (mc *MockContext) GetMemberRepository() MemberRepository {
	return mc.MemberRepository
}