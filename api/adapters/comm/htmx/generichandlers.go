package htmx

import (
	"github.com/gin-gonic/gin"
	"github.com/serdarkalayci/membership/api/adapters/comm/htmx/common"
	"github.com/serdarkalayci/membership/api/adapters/comm/htmx/middleware"
)

func (ws WebServer) GetHomePage(c *gin.Context) {
	c.HTML(200, "main.html", gin.H{})
}

func (ws WebServer) GetEmptyPage(c *gin.Context) {
	c.HTML(200, "emptyr.html", gin.H{})
}

func (ws WebServer) GetNavigationPage(c *gin.Context) {
	claims, err := middleware.CheckAuthentication(c)
	if err != nil {
		c.HTML(401, "navigation.html", nil)
		return
	}
	links := []common.Link{
		{URL: "/home", Text: "Home"},
		{URL: "/memberpage", Text: "Üyeler"},
		{URL: "/logout", Text: "Çıkış"},
	}
	
	c.HTML(200, "navigation.html", gin.H{
		"Links" : links,
		"UserName" : claims.UserName,
	})
	return
}