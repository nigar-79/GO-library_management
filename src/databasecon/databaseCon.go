package databasecon

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sdbgo.in/sdbgo/src/config"
)

//DatabaseConn is the struct used to initilaise the Appconfig struct from the config file which holds the configuration structs.
type DatabaseConn struct {
	AppConfiguration *config.Database
	Data             *gorm.DB
}

//DatabaseConnection is the function which enables the database connection after loading the configuration from toml file.
func (d *DatabaseConn) DatabaseConnection() (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		d.AppConfiguration.Host, d.AppConfiguration.Port, d.AppConfiguration.User, d.AppConfiguration.Password, d.AppConfiguration.Dbname)

	d.Data, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return d.Data, nil

}
