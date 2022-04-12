package controller

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"sdbgo.in/sdbgo/src/databasecon"
	"sdbgo.in/sdbgo/src/models"
)

//DataController contails the connection details to the db
type DataController struct {
	DatabaseCon *databasecon.DatabaseConn
}

//GetInstitutionDetails controller
//context act as middle man between the api calls and methods
//Fetches all institution details
func (dc *DataController) GetInstitutionDetails(ctx iris.Context) {

	var institution []models.Institution

	dc.DatabaseCon.Data.Find(&institution)
	fmt.Println("institute:", institution)
	statuscode, err := ctx.JSON(institution)
	if err != nil {
		ctx.StopWithError(statuscode, err)
	}

}

// CreateInstitution - creates new institution
func (dc *DataController) CreateInstitution(ctx iris.Context) {

}
