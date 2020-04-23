package main

import (
	"github.com/Dadard29/go-api-utils/API"
	"github.com/Dadard29/go-api-utils/database"
	"github.com/Dadard29/go-api-utils/service"
	"github.com/Dadard29/go-music-library/api"
	"github.com/Dadard29/go-music-library/controllers"
	"github.com/Dadard29/go-music-library/models"
	"github.com/Dadard29/go-subscription-connector/subChecker"
	"net/http"
)

// -- env
// CORS_ORIGIN
// DB_USER, DB_PASSWORD
// HOST_SUB
func main() {

	var routes = service.RouteMapping{
		"/library": service.Route{
			Description:   "manage library",
			MethodMapping: service.MethodMapping{
				http.MethodGet: controllers.LibraryGet,
				http.MethodPost: controllers.LibraryPost,
				http.MethodDelete: controllers.LibraryDelete,
			},
		},
	}

	api.Api = API.NewAPI("music-library", "config/config.json", routes, true)

	c, err := api.Api.Config.GetSubcategoryFromFile("api", "db")
	if err != nil {
		api.Api.Logger.CheckErr(err)
	}

	api.Api.Database = database.NewConnector(c, true, []interface{}{
		models.MusicEntity{},
	})

	controllers.Sc = subChecker.NewSubChecker(api.Api.Config.GetEnv("HOST_SUB"))

	api.Api.Service.Start()
	api.Api.Service.Stop()
}
