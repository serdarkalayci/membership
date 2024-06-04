package application

// MockContext mimics the DataContext struct, which is used to test the application layer
type MockContext struct {
	userRepository         UserRepository
	MemberRepository	   MemberRepository
	LookupRepository	   LookupRepository
}

func (mc *MockContext) SetRepositories(ur UserRepository, mr MemberRepository, lr LookupRepository) {
	mc.userRepository = ur
}

func (mc *MockContext) GetUserRepository() UserRepository {
	return mc.userRepository
}

func (mc *MockContext) GetMemberRepository() MemberRepository {
	return mc.MemberRepository
}

func (mc *MockContext) GetLookupRepository() LookupRepository {
	return mc.LookupRepository
}