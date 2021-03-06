package config

import (
	"fmt"
	"net/url"
)

//Database is config for MySQL instance
type Database struct {
	Host     string
	Port     uint
	User     string
	Password string
	Name     string
	Charset  string
}

//GetURL return MySQL connection URL
func (db Database) GetURL() string {
	var userInfo = db.User + ":" + db.Password
	return fmt.Sprintf("%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=UTC&time_zone=%s",
		userInfo, db.Host, db.Port, db.Name, db.Charset, url.QueryEscape(`"+00:00"`))
}

//Config is global config for the app
type Config struct {
	LogLevel  string // debug, info, warn, error, fatal or panic
	LogFormat string // json or console
	Env       string
	Addr      string
	Database  Database
}
