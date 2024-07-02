package htmx

import (
	"os"
	"path"

	"github.com/rs/zerolog/log"

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
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatal().Err(err).Msg("Error while getting current path")
	}
	// currentPath = "./adapters/comm/htmxq"
	engine.Static("/assets", path.Join(currentPath, "assets"))
	// templatePath := path.Join(currentPath, "templates", "*")
	engine.LoadHTMLGlob(path.Join(currentPath, "templates", "*"))
	engine.GET("/memberpage", ws.GetMemberPage)
	engine.GET("/member", ws.GetMemberList)
	engine.GET("/member/:id", ws.GetMemberDetail)
	engine.PUT("/member/:id", ws.UpdateMember)
	engine.GET("/member/:id/edit", ws.EditMemberDetail)
	engine.GET("/memberform", ws.GetMemberForm)
	engine.POST("/member", ws.CreateMember)
	engine.GET("/cities", ws.GetCities)
}

