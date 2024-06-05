package htmx

import (
	"github.com/gin-gonic/gin"
	"github.com/serdarkalayci/membership/api/application"
)

type WebServer struct {
	dbContext *application.DataContext
}

func SetWebRoutes(engine *gin.Engine, dbContext *application.DataContext) {
	ws := &WebServer{
		dbContext: dbContext,
	}
	engine.Static("/assets", "./adapters/comm/htmx/assets")
	engine.LoadHTMLGlob("./adapters/comm/htmx/templates/*")
	engine.GET("/memberpage", ws.GetMemberPage)
	engine.GET("/member", ws.GetMemberList)
	engine.GET("/member/:id", ws.GetMemberDetail)
	engine.PUT("/member/:id", ws.UpdateMember)
	engine.GET("/member/:id/edit", ws.EditMemberDetail)
	engine.GET("/cities", ws.GetCities)
}

