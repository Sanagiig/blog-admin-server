package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisDB *redis.Client
var Log *logrus.Logger
