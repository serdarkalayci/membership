package mappers

import (
	"strconv"

	"github.com/google/uuid"
	"github.com/serdarkalayci/membership/api/adapters/data/postgres/dto"
	"github.com/serdarkalayci/membership/api/domain"
)

// MapMemberDTO2Member maps a member dto to a member domain object
func MapMemberDTO2Member(dtoMember dto.Member) domain.Member {
	city := MapCityDTO2City(dtoMember.City)
	area := MapAreaDTO2Area(dtoMember.Area)
	membershipType := MapMembershipTypeDTO2MembershipType(dtoMember.MembershipType)
	member := domain.Member{
		ID:                  dtoMember.ID.String(),
		Email:               dtoMember.Email,
		FirstName:           dtoMember.Firstname,
		LastName:            dtoMember.Lastname,
		City:                city,
		Area:                area,
		Phone:               dtoMember.Phone,
		Notes:               dtoMember.Notes,
		MembershipType:      membershipType,
		MembershipStartDate: dtoMember.MembershipStartDate,
		LastContactDate:     dtoMember.LastContactDate,
		Occupation:          dtoMember.Occupation,
		Education:           dtoMember.Education,
		DateOfBirth:         dtoMember.DateOfBirth,
		Passive:             dtoMember.Passive,
	}
	return member
}

// MapMember2MemberDTO maps a member domain object to a member dto
func MapMember2MemberDTO(member domain.Member) dto.Member {
	id, err := uuid.Parse(member.ID)
	if err != nil {
		id = uuid.New()
	}
	cityDto := MapCity2CityDTO(member.City)
	areaDto := MapArea2AreaDTO(member.Area)
	membershipTypeDto := MapMembershipType2MembershipTypeDTO(member.MembershipType)
	dtoMember := dto.Member{
		ID:        id,
		Firstname: member.FirstName,
		Lastname:  member.LastName,
		Email:     member.Email,
		City:      cityDto,
		Phone:     member.Phone,
		Passive:   member.Passive,
		Area:      areaDto,
		Notes:     member.Notes,
		MembershipType: membershipTypeDto,
		MembershipStartDate: member.MembershipStartDate,
		LastContactDate: member.LastContactDate,
		Occupation: member.Occupation,
		Education: member.Education,
		DateOfBirth: member.DateOfBirth,
	}
	return dtoMember
}

// MapListMemberDTO2Member maps a list member dto to a member domain object
func MapListMemberDTO2Member(dtoMember dto.ListMember) domain.Member {
	city := MapCityDTO2City(dtoMember.City)
	member := domain.Member{
		ID:        dtoMember.ID.String(),
		FirstName: dtoMember.Firstname,
		LastName:  dtoMember.Lastname,
		Email:     dtoMember.Email,
		City:      city,
		Phone:     dtoMember.Phone,

	}
	return member
}

// MapMemberDTOs2Members maps a slice of member dtos to a slice of member domain objects
func MapMemberDTOs2Members(dtos []dto.Member) []domain.Member {
	members := make([]domain.Member, len(dtos))
	for i, dtoMember := range dtos {
		members[i] = MapMemberDTO2Member(dtoMember)
	}
	return members
}

// MapListMemberDTOs2Members maps a slice of list member dtos to a slice of member domain objects
func MapListMemberDTOs2Members(dtos []dto.ListMember) []domain.Member {
	members := make([]domain.Member, len(dtos))
	for i, dtoMember := range dtos {
		members[i] = MapListMemberDTO2Member(dtoMember)
	}
	return members
}

// MapMembers2MemberDTOs maps a slice of member domain objects to a slice of member dtos
func MapMembers2MemberDTOs(members []domain.Member) []dto.Member {
	dtos := make([]dto.Member, len(members))
	for i, member := range members {
		dtos[i] = MapMember2MemberDTO(member)
	}
	return dtos
}

// MapCityDTO2City maps a city dto to a city domain object
func MapCityDTO2City(dtoCity dto.City) domain.City {
	city := domain.City{
		ID:   string(dtoCity.ID),
		Name: dtoCity.Name,
	}
	return city
}

// MapCity2CityDTO maps a city domain object to a city dto
func MapCity2CityDTO(city domain.City) dto.City {
	id, err := strconv.Atoi(city.ID)
	if err != nil {
		id = 0
	}
	dtoCity := dto.City{
		ID:   id,
		Name: city.Name,
	}
	return dtoCity
}

// MapProvinceDTO2Province maps a province dto to a province domain object
func MapProvinceDTO2Province(dtoProvince dto.Province) domain.Province {
	province := domain.Province{
		ID:   string(dtoProvince.ID),
		Name: dtoProvince.Name,
	}
	return province
}

// MapProvince2ProvinceDTO maps a province domain object to a province dto
func MapProvince2ProvinceDTO(province domain.Province) dto.Province {
	id, err := strconv.Atoi(province.ID)
	if err != nil {
		id = 0
	}
	dtoProvince := dto.Province{
		ID:   id,
		Name: province.Name,
	}
	return dtoProvince
}

// MapAreaDTO2Area maps an area dto to an area domain object
func MapAreaDTO2Area(dtoArea dto.Area) domain.Area {
	area := domain.Area{
		ID:   string(dtoArea.ID),
		Name: dtoArea.Name,
	}
	return area
}

// MapArea2AreaDTO maps an area domain object to an area dto
func MapArea2AreaDTO(area domain.Area) dto.Area {
	id, err := strconv.Atoi(area.ID)
	if err != nil {
		id = 0
	}
	dtoArea := dto.Area{
		ID:   id,
		Name: area.Name,
	}
	return dtoArea
}

// MapMembershipTypeDTO2MembershipType maps a membership type dto to a membership type domain object
func MapMembershipTypeDTO2MembershipType(dtoMembershipType dto.MembershipType) domain.MembershipType {
	membershipType := domain.MembershipType{
		ID:   string(dtoMembershipType.ID),
		Name: dtoMembershipType.Name,
	}
	return membershipType
}

// MapMembershipType2MembershipTypeDTO maps a membership type domain object to a membership type dto
func MapMembershipType2MembershipTypeDTO(membershipType domain.MembershipType) dto.MembershipType {
	id, err := strconv.Atoi(membershipType.ID)
	if err != nil {
		id = 0
	}
	dtoMembershipType := dto.MembershipType{
		ID:   id,
		Name: membershipType.Name,
	}
	return dtoMembershipType
}