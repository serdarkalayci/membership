package mappers

import (
	"strconv"

	"github.com/serdarkalayci/membership/api/adapters/data/postgres/dao"
	"github.com/serdarkalayci/membership/api/domain"
)

// MapCitydao2City maps a city dao to a city domain object
func MapCitydao2City(daoCity dao.City) domain.City {
	city := domain.City{
		ID:   string(daoCity.ID),
		Name: daoCity.Name,
	}
	return city
}

// MapCitydaos2Cities maps a slice of city daos to a slice of city domain objects
func MapCitydaos2Cities(daos []dao.City) []domain.City {
	cities := make([]domain.City, len(daos))
	for i, dao := range daos {
		cities[i] = MapCitydao2City(dao)
	}
	return cities
}

// MapCity2Citydao maps a city domain object to a city dao
func MapCity2Citydao(city domain.City) dao.City {
	id, err := strconv.Atoi(city.ID)
	if err != nil {
		id = 0
	}
	daoCity := dao.City{
		ID:   id,
		Name: city.Name,
	}
	return daoCity
}

// MapCities2Citydaos maps a slice of city domain objects to a slice of city daos
func MapCities2Citydaos(cities []domain.City) []dao.City {
	daos := make([]dao.City, len(cities))
	for i, city := range cities {
		daos[i] = MapCity2Citydao(city)
	}
	return daos
}

// MapProvincedao2Province maps a province dao to a province domain object
func MapProvincedao2Province(daoProvince dao.Province) domain.Province {
	province := domain.Province{
		ID:   string(daoProvince.ID),
		Name: daoProvince.Name,
	}
	return province
}

// MapProvincedaos2Provinces maps a slice of province daos to a slice of province domain objects
func MapProvincedaos2Provinces(daos []dao.Province) []domain.Province {
	provinces := make([]domain.Province, len(daos))
	for i, dao := range daos {
		provinces[i] = MapProvincedao2Province(dao)
	}
	return provinces
}

// MapProvince2Provincedao maps a province domain object to a province dao
func MapProvince2Provincedao(province domain.Province) dao.Province {
	id, err := strconv.Atoi(province.ID)
	if err != nil {
		id = 0
	}
	daoProvince := dao.Province{
		ID:   id,
		Name: province.Name,
	}
	return daoProvince
}

// MapProvinces2Provincedaos maps a slice of province domain objects to a slice of province daos
func MapProvinces2Provincedaos(provinces []domain.Province) []dao.Province {
	daos := make([]dao.Province, len(provinces))
	for i, province := range provinces {
		daos[i] = MapProvince2Provincedao(province)
	}
	return daos
}

// MapAreadao2Area maps an area dao to an area domain object
func MapAreadao2Area(daoArea dao.Area) domain.Area {
	area := domain.Area{
		ID:   string(daoArea.ID),
		Name: daoArea.Name,
	}
	return area
}

// MapAreadaos2Areas maps a slice of area daos to a slice of area domain objects
func MapAreadaos2Areas(daos []dao.Area) []domain.Area {
	areas := make([]domain.Area, len(daos))
	for i, dao := range daos {
		areas[i] = MapAreadao2Area(dao)
	}
	return areas
}

// MapArea2Areadao maps an area domain object to an area dao
func MapArea2Areadao(area domain.Area) dao.Area {
	id, err := strconv.Atoi(area.ID)
	if err != nil {
		id = 0
	}
	daoArea := dao.Area{
		ID:   id,
		Name: area.Name,
	}
	return daoArea
}

// MapAreas2Areadaos maps a slice of area domain objects to a slice of area daos
func MapAreas2Areadaos(areas []domain.Area) []dao.Area {
	daos := make([]dao.Area, len(areas))
	for i, area := range areas {
		daos[i] = MapArea2Areadao(area)
	}
	return daos
}

// MapMembershipTypedao2MembershipType maps a membership type dao to a membership type domain object
func MapMembershipTypedao2MembershipType(daoMembershipType dao.MembershipType) domain.MembershipType {
	membershipType := domain.MembershipType{
		ID:   string(daoMembershipType.ID),
		Name: daoMembershipType.Name,
	}
	return membershipType
}

// MapMembershipTypedaos2MembershipTypes maps a slice of membership type daos to a slice of membership type domain objects
func MapMembershipTypedaos2MembershipTypes(daos []dao.MembershipType) []domain.MembershipType {
	membershipTypes := make([]domain.MembershipType, len(daos))
	for i, dao := range daos {
		membershipTypes[i] = MapMembershipTypedao2MembershipType(dao)
	}
	return membershipTypes
}

// MapMembershipType2MembershipTypedao maps a membership type domain object to a membership type dao
func MapMembershipType2MembershipTypedao(membershipType domain.MembershipType) dao.MembershipType {
	id, err := strconv.Atoi(membershipType.ID)
	if err != nil {
		id = 0
	}
	daoMembershipType := dao.MembershipType{
		ID:   id,
		Name: membershipType.Name,
	}
	return daoMembershipType
}

// MapMembershipTypes2MembershipTypedaos maps a slice of membership type domain objects to a slice of membership type daos
func MapMembershipTypes2MembershipTypedaos(membershipTypes []domain.MembershipType) []dao.MembershipType {
	daos := make([]dao.MembershipType, len(membershipTypes))
	for i, membershipType := range membershipTypes {
		daos[i] = MapMembershipType2MembershipTypedao(membershipType)
	}
	return daos
}