package htmx

import "github.com/gin-gonic/gin"

func (ws WebServer) GetEmptyPage(c *gin.Context) {
	c.HTML(200, "emptyr.html", gin.H{})
}