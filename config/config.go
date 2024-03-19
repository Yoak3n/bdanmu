package config

import (
	"time"

	"github.com/spf13/viper"
)

type (
	Configuration struct {
		RoomId    int       `yaml:"room_id"`
		Proxy     string    `yaml:"proxy"`
		Extension bool      `yaml:"extension"`
		Auth      Auth      `yaml:"wbi"`
		Database  *Database `yaml:"database"`
	}
	Auth struct {
		Cookie       string `yaml:"cookie"`
		RefreshToken string `yaml:"refresh_token"`
		ImgKey       string `yaml:"img_key"`
		SubKey       string `yaml:"sub_key"`
		LastUpdate   int64  `yaml:"last_update"`
	}
	Database struct {
		Type     string `yaml:"type"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"database"`
	}
)

var (
	Conf *Configuration
	v    *viper.Viper
)

func init() {
	v = viper.New()
	Conf = &Configuration{
		Auth:     Auth{},
		Database: &Database{},
	}
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		return
	}
	v.SetDefault("database.type", "sqlite")
	v.SetDefault("database.name", "bliveDB")
	getConfigFromFile()
}

func getConfigFromFile() {
	Conf.RoomId = v.GetInt("room_id")
	Conf.Proxy = v.GetString("proxy")
	Conf.Extension = v.GetBool("extension")

	Conf.Auth.Cookie = v.GetString("auth.cookie")
	Conf.Auth.RefreshToken = v.GetString("auth.refresh_token")
	Conf.Auth.ImgKey = v.GetString("auth.img_key")
	Conf.Auth.SubKey = v.GetString("auth.sub_key")
	Conf.Auth.LastUpdate = v.GetInt64("auth.last_update")

	Conf.Database.Type = v.GetString("database.type")
	Conf.Database.Host = v.GetString("database.host")
	Conf.Database.Port = v.GetInt("database.port")
	Conf.Database.User = v.GetString("database.user")
	Conf.Database.Password = v.GetString("database.password")
	Conf.Database.Name = v.GetString("database.name")

}

func SetWBIKey(img string, sub string) {
	v.Set("auth.img_key", img)
	v.Set("auth.sub_key", sub)
	Conf.Auth.ImgKey = img
	Conf.Auth.SubKey = sub
	t := time.Now().Unix()
	Conf.Auth.LastUpdate = t
	v.Set("wbi.last_update", t)
	v.WriteConfig()
}

func SetCookieRefresh() {
	v.Set("cookie", Conf.Auth.Cookie)
	v.Set("refresh_token", Conf.Auth.RefreshToken)
	v.WriteConfig()
}
