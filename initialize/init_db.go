package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go-blog/global"
	"go-blog/global/settings"
	"go-blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func InitDB() {
	var err error
	if global.DB != nil {
		fmt.Println("db was initial")
		return
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)

	global.DB, err = gorm.Open(
		mysql.Open(
			fmt.Sprintf(
				"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
				settings.DbUser,
				settings.DbPassword,
				settings.DbHost,
				settings.DbPort,
				settings.DbName,
			),
		),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
			Logger: newLogger,
		},
	)

	if err != nil {
		fmt.Printf("连接数据库失败: ", err)
	}

	sqldb, err := global.DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}

	err = global.DB.AutoMigrate(
		&model.User{},
		&model.Role{},
		&model.UserInfo{},
		&model.Article{},
		&model.Category{},
		&model.Tag{},
		&model.Dictionary{},
	)
	if err != nil {
		panic(err)
	}

	//pwd, err := os.Getwd()
	//if err != nil {
	//	panic(err)
	//}
	//sqlFIle, err := os.OpenFile(path.Join(pwd, "./initialize/sql/procedule.sql"), os.O_RDONLY, 0666)
	//if err != nil {
	//	panic(err)
	//}
	//
	//if sqlData, err := io.ReadAll(sqlFIle); err == nil {
	//	err = global.DB.Exec(string(sqlData)).Error
	//	if err != nil {
	//		panic(err)
	//	}
	//} else {
	//	panic(err)
	//}

	fmt.Println("初始化存储过程完成.")
	sqldb.SetMaxIdleConns(10)
	sqldb.SetMaxOpenConns(100)
	sqldb.SetConnMaxLifetime(time.Second * 10)
}
