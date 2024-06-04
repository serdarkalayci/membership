package htmx

import (
	"github.com/gin-gonic/gin"
	"github.com/serdarkalayci/membership/api/application"
	"github.com/serdarkalayci/membership/api/domain"
)

func (ws WebServer) GetCities(c *gin.Context) {
	provinceID := c.Query("provinceId")
	var cities []domain.City
	var err error
	ls := application.NewLookupService(ws.dbContext)
	if provinceID == "" {
		cities, err = ls.ListCities()
		if err != nil {
			c.HTML(500, "memberedit-cities.html", nil)
			return
		}
	} else {
		cities, err = ls.ListProvinceCities(provinceID)
		if err != nil {
			c.HTML(500, "memberedit-cities.html", nil)
			return
		}
	}
	c.HTML(200, "memberedit-cities.html", gin.H{
		"Cities": cities,
	})
}