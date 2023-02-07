package initialize

import _ "go-blog/global/settings"

func InitAll() {
	InitLogger()
	InitRedis()
	InitDB()
}
