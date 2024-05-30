package mappers

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/serdarkalayci/membership/api/adapters/data/postgres/dao"
	"github.com/serdarkalayci/membership/api/domain"
)

// MapMemberdao2Member maps a member dao to a member domain object
func MapMemberdao2Member(daoMember dao.Member) domain.Member {
	city := domain.City{
		ID:   string(daoMember.CityID),
		Name: daoMember.CityName,
	}
	area := domain.Area{
		ID:   string(daoMember.AreaID),
		Name: daoMember.AreaName,
	}
	membershipType := domain.MembershipType{
		ID:   string(daoMember.MembershipTypeID),
		Name: daoMember.MembershipTypeName,
	}
	member := domain.Member{
		ID:                  daoMember.ID.String(),
		Email:               daoMember.Email,
		FirstName:           daoMember.Firstname,
		LastName:            daoMember.Lastname,
		City:                city,
		Area:                area,
		Phone:               daoMember.Phone,
		Notes:               daoMember.Notes,
		MembershipType:      membershipType,
		MembershipStartDate: daoMember.MembershipStartDate,
		LastContactDate:     daoMember.LastContactDate,
		Occupation:          daoMember.Occupation,
		Education:           daoMember.Education,
		DateOfBirth:         daoMember.DateOfBirth,
		Passive:             daoMember.Passive,
	}
	return member
}

// MapMember2Memberdao maps a member domain object to a member dao
func MapMember2Memberdao(member domain.Member) dao.Member {
	id, err := uuid.Parse(member.ID)
	if err != nil {
		id = uuid.New()
	}
	cityID,_ := strconv.Atoi(member.City.ID)
	areaID,_ := strconv.Atoi(member.Area.ID)
	membershipTypeID,_ := strconv.Atoi(member.MembershipType.ID)
	daoMember := dao.Member{
		ID:        id,
		Firstname: member.FirstName,
		Lastname:  member.LastName,
		Email:     member.Email,
		CityID:    cityID,
		CityName: member.City.Name,
		Phone:     member.Phone,
		Passive:   member.Passive,
		AreaID:      areaID,
		AreaName: member.Area.Name,
		Notes:     member.Notes,
		MembershipTypeID: membershipTypeID,
		MembershipTypeName: member.MembershipType.Name,
		MembershipStartDate: member.MembershipStartDate,
		LastContactDate: member.LastContactDate,
		Occupation: member.Occupation,
		Education: member.Education,
		DateOfBirth: member.DateOfBirth,
	}
	return daoMember
}

// MapListMemberdao2Member maps a list member dao to a member domain object
func MapListMemberdao2Member(daoMember *dao.ListMember) domain.Member {
	city := domain.City{
		ID:   string(daoMember.CityID),
		Name: daoMember.CityName,
	}
	member := domain.Member{
		ID:        daoMember.ID.String(),
		FirstName: daoMember.Firstname,
		LastName:  daoMember.Lastname,
		Email:     daoMember.Email,
		City:      city,
		Phone:     daoMember.Phone,

	}
	return member
}

// MapMemberdaos2Members maps a slice of member daos to a slice of member domain objects
func MapMemberdaos2Members(daos []dao.Member) []domain.Member {
	members := make([]domain.Member, len(daos))
	for i, daoMember := range daos {
		members[i] = MapMemberdao2Member(daoMember)
	}
	return members
}

// MapListMemberdaos2Members maps a slice of list member daos to a slice of member domain objects
func MapListMemberdaos2Members(daos []dao.ListMember) []domain.Member {
	members := make([]domain.Member, len(daos))
	for i, daoMember := range daos {
		members[i] = MapListMemberdao2Member(&daoMember)
	}
	return members
}

// MapMembers2Memberdaos maps a slice of member domain objects to a slice of member daos
func MapMembers2Memberdaos(members []domain.Member) []dao.Member {
	daos := make([]dao.Member, len(members))
	for i, member := range members {
		daos[i] = MapMember2Memberdao(member)
	}
	return daos
}

