package helpers

import (
	"fmt"
	"log"
	// "os"
	// "strconv"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
)

var SqlSession *gorm.DB

func NewSqlSession() *gorm.DB{
	return SqlSession
}

func InitMySql()(err error , db *gorm.DB)  {
	envErr := godotenv.Load(".env")
    if envErr != nil {
        log.Fatal("Error loading .env file")
    }
	url := GetEnvStr("DB.URL")
  	userName := GetEnvStr("DB.USERNAME")
	passWord := GetEnvStr("DB.PASSWORD")
	dbName := GetEnvStr("DB.NAME")
	port := GetEnvStr("DB.PORT")
	maxIdleConn,_ := GetEnvInt("DB.MAX_IDLE_CONN")
	maxPoolConn,_ := GetEnvInt("DB.MAX_POOL_CONN")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userName,
		passWord,
		url,
		port,
		dbName,
	)

	SqlSession,err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil{
		return err,SqlSession
	}
	sqlDB, err := SqlSession.DB()
	if err != nil {
		return err,SqlSession
	}
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxPoolConn)
	return err,SqlSession
}
