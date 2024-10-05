package htmx

import (
	"path"

	"github.com/nicholasjackson/env"

	"github.com/gin-gonic/gin"
	"github.com/serdarkalayci/membership/api/adapters/comm/htmx/middleware"
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
	authorized := engine.Group("/")
	authorized.Use(middleware.Authenticate())
	authorized.GET("/memberpage", ws.GetMemberPage)
	authorized.GET("/member", ws.GetMemberList)
	authorized.GET("/member/:id", ws.GetMemberDetail)
	authorized.PUT("/member/:id", ws.UpdateMember)
	authorized.GET("/member/:id/edit", ws.EditMemberDetail)
	authorized.GET("/memberform", ws.GetMemberForm)
	authorized.POST("/member", ws.CreateMember)
	authorized.GET("/cities", ws.GetCities)

	engine.GET("/loginpage", ws.GetLoginPage)
	engine.POST("/login", ws.Login)
	engine.GET("/userform", ws.GetUserForm)
	engine.POST("/user", ws.UpsertUser)

	engine.GET("/home", ws.GetHomePage)
	engine.GET("/", ws.GetHomePage)
	engine.GET("/empty", ws.GetEmptyPage)
	engine.GET("/navigation", ws.GetNavigationPage)
}

