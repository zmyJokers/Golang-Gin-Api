package conf

type Config struct {
	Host     string
	Username string
	Password string
	Port     string
	Database string
	//Charset     string
	//TablePrefix string
}

var MysqlConf = map[string]*Config{
	"db": {
		Host:     "127.0.0.1",
		Username: "root",
		Password: "root",
		Port:     "3306",
		Database: "test",
	},
	"db1": {
		Host:     "127.0.0.1",
		Username: "root",
		Password: "root",
		Port:     "3306",
		Database: "es",
	},
}
