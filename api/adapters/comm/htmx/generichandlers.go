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
	c.HTML(200, "empty.html", gin.H{})
}

func (ws WebServer) GetNavigationPage(c *gin.Context) {
	claims, err := middleware.CheckAuthentication(c)
	if err != nil {
		links := []common.Link{
			{URL: "/loginpage", Text: "Giriş Yap"},
		}
		c.HTML(200, "navigation.html", gin.H{
			"Links" : links,
		})
		return
	}
	links := []common.Link{
		{URL: "/home", Text: "Home"},
		{URL: "/memberpage", Text: "Üyeler"},
	}
	
	c.HTML(200, "navigation.html", gin.H{
		"Links" : links,
		"UserName" : claims.UserName,
	})
	return
}