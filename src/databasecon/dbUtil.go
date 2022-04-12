package databasecon

import (
	"fmt"

	"sdbgo.in/sdbgo/src/models"
)

//CreateTables is the function which is used to create db tables
func (d *DatabaseConn) CreateTables() error {
	err := d.Data.AutoMigrate(&models.Institution{})
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	fmt.Println("table created")

	return nil
}
