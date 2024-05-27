package postgres

import (
	"github.com/serdarkalayci/membership/api/adapters/data/postgres/dto"
	"github.com/serdarkalayci/membership/api/adapters/data/postgres/mappers"
	"github.com/serdarkalayci/membership/api/domain"
	"gorm.io/gorm"
)

// MemberRepository holds the arangodb client and database name for methods to use
type MemberRepository struct {
	db *gorm.DB
}

func newMemberRepository(database *gorm.DB) MemberRepository {
	return MemberRepository{
		db: database,
	}
}

func (mr MemberRepository) ListMembers(pageSize, pageNum int) ([]domain.Member, int64, error) {
	var members []dto.ListMember
	if err := mr.db.Scopes(Paginate(pageSize, pageNum)).Model(&dto.Member{}).Joins("City").Find(&members).Error; err != nil {
		return nil, 0, err
	}
	var count int64 = 0
	if err := mr.db.Model(&dto.Member{}).Joins("City").Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return mappers.MapListMemberDTOs2Members(members), count, nil
}

func (mr MemberRepository) GetMember(id string) (domain.Member, error) {
	var member dto.Member
	if err := mr.db.First(&member, "ID = ?", id).Error; err != nil {
		return domain.Member{}, err
	}
	return mappers.MapMemberDTO2Member(member), nil
}