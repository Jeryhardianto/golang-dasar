package database

var connection string

func init() {
	connection = "mysql"
}

func GetDatabase() string {
	return connection
}