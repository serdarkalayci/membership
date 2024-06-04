package htmx

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/serdarkalayci/membership/api/application"
)

func (ws WebServer) GetMemberPage(c *gin.Context) {
	c.HTML(200, "member.html", gin.H{
		"title": "Member Page",
	})
}

func (ws WebServer) GetMemberList(c *gin.Context) {
	pageSize := 10
	pageNum := 1
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
	members, count, err := ms.ListMembers(pageSize, pageNum)

	if err != nil {
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
		c.HTML(500, "memberedit.html", nil)
		return
	}
	ls := application.NewLookupService(ws.dbContext)
	provinces, err := ls.ListProvinces()
	if err != nil {
		c.HTML(500, "memberedit.html", nil)
		return
	}
	cities, err := ls.ListProvinceCities(member.City.Province.ID)
	if err != nil {
		c.HTML(500, "memberedit.html", nil)
		return
	}
	areas, err := ls.ListAreas()
	if err != nil {
		c.HTML(500, "memberedit.html", nil)
		return
	}
	membershipTypes, err := ls.ListMembershipTypes()
	if err != nil {
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