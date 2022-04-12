package config

//AppConfig is where we embed the structs together that can be exported and used for configuration after loading from confuration file
type AppConfig struct {
	Title    string
	Server   Server
	Database Database
	// LdapAuth  LdapAuth
	// JwtConfig JwtConfig
}

//Server struct contains the information which enables the webserver connection
type Server struct {
	Host string
	Port int64
}

//Database struct holds the information which enables the connection of the database
type Database struct {
	Host     string
	Port     int64
	User     string
	Password string
	Dbname   string
	Sslmode  bool
}
