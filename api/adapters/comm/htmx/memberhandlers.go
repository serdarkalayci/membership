package htmx

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/serdarkalayci/membership/api/application"
	"github.com/serdarkalayci/membership/api/domain"
)

func (ws WebServer) GetMemberPage(c *gin.Context) {
	ls := application.NewLookupService(ws.dbContext)
	cities, err := ls.ListCities()
	if err != nil {
		log.Error().Err(err).Msg("Error while getting cities for member create")
		c.HTML(500, "membercreate.html", nil)
		return
	}
	areas, err := ls.ListAreas()
	if err != nil {
		log.Error().Err(err).Msg("Error while getting areas for member create")
		c.HTML(500, "membercreate.html", nil)
		return
	}
	c.HTML(200, "member.html", gin.H{
		"title": "Member Page",
		"Cities": cities,
		"Areas": areas,
	})
}

func (ws WebServer) GetMemberList(c *gin.Context) {
	pageSize := 10
	pageNum := 1
	searchName := c.Query("searchname")
	var searchCity, searchArea int
	searchCity, _ = strconv.Atoi(c.Query("searchcity")) 
	searchArea, _ = strconv.Atoi(c.Query("searcharea"))
	if c.Query("pageSize") != "" {
		size, err := strconv.Atoi(c.Query("pageSize")); if err == nil {
			pageSize = size
		}
	}
	if c.Query("pageNum") != "" {
		num, err := strconv.Atoi(c.Query("pageNum")); if err == nil {
			pageNum = num
		}
	}
	ms := application.NewMemberService(ws.dbContext)
	members, count, err := ms.ListMembers(pageSize, pageNum, searchName, searchCity, searchArea)

	if err != nil {
		log.Error().Err(err).Msg("Error while listing members")
		c.HTML(500, "memberlist.html", nil)
		return
	}
	pageInfo := calculatePageInfo(pageSize, pageNum, count)
	c.HTML(200, "memberlist.html", gin.H{
		"Members": members,
		"PageInfo": pageInfo,
	})

}

func (ws WebServer) GetMemberDetail(c *gin.Context) {
	id := c.Param("id")
	ms := application.NewMemberService(ws.dbContext)
	member, err := ms.GetMember(id)
	if err != nil {
		log.Error().Err(err).Msg("Error while getting member detail")
		c.HTML(500, "memberdetail.html", nil)
		return
	}
	c.HTML(200, "memberdetail.html", gin.H{
		"Member": member,
	})
}

func (ws WebServer) EditMemberDetail(c *gin.Context) {
	id := c.Param("id")
	ms := application.NewMemberService(ws.dbContext)
	member, err := ms.GetMember(id)
	if err != nil {
		log.Error().Err(err).Msg("Error while getting member for edit")
		c.HTML(500, "memberedit.html", nil)
		return
	}
	ls := application.NewLookupService(ws.dbContext)
	provinces, err := ls.ListProvinces()
	if err != nil {
		log.Error().Err(err).Msg("Error while getting provinces for member edit")
		c.HTML(500, "memberedit.html", nil)
		return
	}
	cities, err := ls.ListProvinceCities(member.City.Province.ID)
	if err != nil {
		log.Error().Err(err).Msg("Error while getting cities for member edit")
		c.HTML(500, "memberedit.html", nil)
		return
	}
	areas, err := ls.ListAreas()
	if err != nil {
		log.Error().Err(err).Msg("Error while getting areas for member edit")
		c.HTML(500, "memberedit.html", nil)
		return
	}
	membershipTypes, err := ls.ListMembershipTypes()
	if err != nil {
		log.Error().Err(err).Msg("Error while getting membership types for member edit")
		c.HTML(500, "memberedit.html", nil)
		return
	}
	c.HTML(200, "memberedit.html", gin.H{
		"Member": member,
		"Cities": cities,
		"Provinces": provinces,
		"Areas": areas,
		"MembershipTypes": membershipTypes,
	})
}

func (ws WebServer) UpdateMember(c *gin.Context) {
	id := c.Param("id")
	ms := application.NewMemberService(ws.dbContext)
	member, err := ms.GetMember(id)
	if err != nil {
		log.Error().Err(err).Msg("Error while getting member for update")
		c.HTML(500, "memberedit.html", nil)
		return
	}
	member.Email = c.PostForm("email")
	member.FirstName = c.PostForm("firstName")
	member.LastName = c.PostForm("lastName")
	member.City.ID = c.PostForm("city")
	member.Area.ID = c.PostForm("area")
	member.Phone = c.PostForm("phone")
	member.Notes = c.PostForm("notes")
	member.MembershipType.ID = c.PostForm("membershipType")
	member.MembershipStartDate, _ = time.Parse("2006-01-022", c.PostForm("membershipStartDate"))
	member.LastContactDate, _ = time.Parse("2006-01-022", c.PostForm("lastContactDate"))
	member.Occupation = c.PostForm("occupation")
	member.Education = c.PostForm("education")
	member.DateOfBirth, _ = time.Parse("2006-01-022", c.PostForm("dateOfBirth"))
	err = ms.UpdateMember(member)
	if err != nil {
		log.Error().Err(err).Msg("Error while updating member")
		c.HTML(500, "memberedit.html", nil)
		return
	}
	c.Redirect(303, "/memberpage")
}

func (ws WebServer) GetMemberForm(c *gin.Context) {
	ls := application.NewLookupService(ws.dbContext)
	provinces, err := ls.ListProvinces()
	if err != nil {
		log.Error().Err(err).Msg("Error while getting provinces for member create")
		c.HTML(500, "membercreate.html", nil)
		return
	}
	cities, err := ls.ListCities()
	if err != nil {
		log.Error().Err(err).Msg("Error while getting cities for member create")
		c.HTML(500, "membercreate.html", nil)
		return
	}
	areas, err := ls.ListAreas()
	if err != nil {
		log.Error().Err(err).Msg("Error while getting areas for member create")
		c.HTML(500, "membercreate.html", nil)
		return
	}
	membershipTypes, err := ls.ListMembershipTypes()
	if err != nil {
		log.Error().Err(err).Msg("Error while getting membership types for member create")
		c.HTML(500, "membercreate.html", nil)
		return
	}
	c.HTML(200, "membercreate.html", gin.H{
		"Cities": cities,
		"Provinces": provinces,
		"Areas": areas,
		"MembershipTypes": membershipTypes,
	})
}

// CreateMember creates a new member with the given data
func (ws WebServer) CreateMember(c *gin.Context) {
	msd, _ := time.Parse("2006-01-022", c.PostForm("membershipStartDate"))
	lcd, _ := time.Parse("2006-01-022", c.PostForm("lastContactDate"))
	dob, _ := time.Parse("2006-01-022", c.PostForm("dateOfBirth"))
	member := domain.Member{
		FirstName: c.PostForm("firstName"),
		LastName: c.PostForm("lastName"),
		Email: c.PostForm("email"),
		Phone: c.PostForm("phone"),
		City: domain.City{ID: c.PostForm("city")},
		Area: domain.Area{ID: c.PostForm("area")},
		MembershipType: domain.MembershipType{ID: c.PostForm("membershipType")},
		MembershipStartDate: msd,
		LastContactDate: lcd,
		Occupation: c.PostForm("occupation"),
		Education: c.PostForm("education"),
		DateOfBirth: dob,
		Notes: c.PostForm("notes"),
	}
	ms := application.NewMemberService(ws.dbContext)
	_, err := ms.CreateMember(member)
	if err != nil {
		log.Error().Err(err).Msg("Error while creating member")
		c.HTML(500, "membercreate.html", nil)
		return
	}
	c.Redirect(303, "/memberpage")
}