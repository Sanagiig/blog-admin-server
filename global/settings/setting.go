package settings

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	// sys.go
	LogDir string
	// server
	AppMode   string
	HttpPort  string
	Domain    string
	StaticDir string

	// upload
	QiNiuPubBucket    string
	QiNiuPriBucket    string
	QiNiuPubKVDomain  string
	QiNiuPriKVDomain  string
	QiNiuAccessKey    string
	QiNiuSecretKey    string
	QiNiuTokenExpires int
	// auth
	JWTKey   string
	TokenExp int
	// redis
	RedisHost     string
	RedisPassword string
	// db
	DB         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	cfg, err := ini.Load("./config/config.ini")

	if err != nil {
		fmt.Println("ini errMsg:", err.Error())
		return
	}
	initSys(cfg)
	initServer(cfg)
	initAuth(cfg)
	initRedis(cfg)
	initDB(cfg)
}

func initSys(cfg *ini.File) {
	LogDir = cfg.Section("sys.go").Key("LogDir").MustString("")
	QiNiuAccessKey = cfg.Section("upload").Key("QiNiuAccessKey").MustString("")
	QiNiuSecretKey = cfg.Section("upload").Key("QiNiuSecretKey").MustString("")
	QiNiuTokenExpires = cfg.Section("upload").Key("QiNiuTokenExpires").MustInt(10)
	QiNiuPubBucket = cfg.Section("upload").Key("QiNiuPubBucket").MustString("")
	QiNiuPriBucket = cfg.Section("upload").Key("QiNiuPriBucket").MustString("")
	QiNiuPubKVDomain = cfg.Section("upload").Key("QiNiuPubKVDomain").MustString("")
	QiNiuPriKVDomain = cfg.Section("upload").Key("QiNiuPriKVDomain").MustString("")
}

func initServer(cfg *ini.File) {
	AppMode = cfg.Section("server").Key("AppMode").MustString("debug")
	HttpPort = cfg.Section("server").Key("HttpPort").MustString(":3000")
	Domain = cfg.Section("server").Key("Domain").MustString("localhost")
	StaticDir = cfg.Section("server").Key("StaticDir").MustString("g:\\tmp")
}

func initAuth(cfg *ini.File) {
	JWTKey = cfg.Section("auth").Key("JWTKey").MustString("123456")
	TokenExp = cfg.Section("auth").Key("JWTKey").MustInt(10)
}

func initRedis(cfg *ini.File) {
	RedisHost = cfg.Section("redis").Key("Host").MustString("localhost:6379")
	RedisPassword = cfg.Section("redis").Key("Password").MustString("")

}

func initDB(cfg *ini.File) {
	DB = cfg.Section("database").Key("DB").MustString("mysql")
	DbHost = cfg.Section("database").Key("DbHost").MustString("localhost")
	DbPort = cfg.Section("database").Key("DbPort").MustString("3306")
	DbUser = cfg.Section("database").Key("DbUser").MustString("")
	DbPassword = cfg.Section("database").Key("DbPassword").MustString("")
	DbName = cfg.Section("database").Key("DbName").MustString("")
}
