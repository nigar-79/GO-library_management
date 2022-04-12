package main

import (
	"fmt"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
	"sdbgo.in/sdbgo/src/config"
	controller "sdbgo.in/sdbgo/src/controllers"
	"sdbgo.in/sdbgo/src/databasecon"
	"sdbgo.in/sdbgo/src/utils"
)

func main() {

	var dbconn *gorm.DB
	var appConfiguration *config.AppConfig

	appConfiguration, err := utils.LoadAppConfig("app.dev.toml")
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	var dbCon *databasecon.DatabaseConn = &databasecon.DatabaseConn{AppConfiguration: &appConfiguration.Database, Data: dbconn}

	dbconn, err = dbCon.DatabaseConnection()
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	err = dbCon.CreateTables()
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	sdblib := iris.New()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders:   []string{"*"},
	})

	sdblib.UseRouter(crs)

	var DC *controller.DataController = &controller.DataController{DatabaseCon: &databasecon.DatabaseConn{Data: dbconn}}

	apicalls := sdblib.Party("/baseapis")
	{
		apicalls.Get("/", DC.GetInstitutionDetails)
		apicalls.Post("/", DC.CreateInstitution)
	}

	err = sdblib.Listen(":8087")
	if err != nil {
		fmt.Printf("cannot start the server on the TCP network address host:port")
	}

}
