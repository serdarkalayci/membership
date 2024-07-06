package htmx

import (
	"path"

	"github.com/nicholasjackson/env"

	"github.com/gin-gonic/gin"
	"github.com/serdarkalayci/membership/api/application"
)

type WebServer struct {
	dbContext *application.DataContext
}

var basePath = env.String("BASE_PATH", false, "./adapters/comm/htmx", "Base path for the static files")

func SetWebRoutes(engine *gin.Engine, dbContext *application.DataContext) {
	env.Parse()
	ws := &WebServer{
		dbContext: dbContext,
	}
	engine.Static("/assets", path.Join(*basePath, "assets"))
	engine.LoadHTMLGlob(path.Join(*basePath, "templates", "*"))
	engine.GET("/memberpage", ws.GetMemberPage)
	engine.GET("/member", ws.GetMemberList)
	engine.GET("/member/:id", ws.GetMemberDetail)
	engine.PUT("/member/:id", ws.UpdateMember)
	engine.GET("/member/:id/edit", ws.EditMemberDetail)
	engine.GET("/memberform", ws.GetMemberForm)
	engine.POST("/member", ws.CreateMember)
	engine.GET("/cities", ws.GetCities)
}

