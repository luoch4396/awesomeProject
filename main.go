package main

import (
	"awesomeProject/infrastructure/repository"
	"awesomeProject/interfaces"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	//获取数据库配置
	dbDriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	//redis details
	//redis_host := os.Getenv("REDIS_HOST")
	//redis_port := os.Getenv("REDIS_PORT")
	//redis_password := os.Getenv("REDIS_PASSWORD")

	//创建services
	repositories, err := repository.NewRepositories(dbDriver, user, password, port, host, dbname)
	if err != nil {
		panic(err)
	}

	//关闭资源
	defer repositories.Close()
	//增加没有的字段
	//repositories.AutoMigrate()

	records := interfaces.NewCloudhubBillRecordsController(repositories.CloudhubBillRecords)

	r := gin.Default()
	//r.Use(middleware.CORSMiddleware()) //For CORS

	//定义rest接口
	r.POST("/records/get/:id", records.GetById)

	//开启web服务
	appPort := os.Getenv("PORT")
	if appPort == "" {
		appPort = "8888" //localhost
	}
	log.Fatal(r.Run(":" + appPort))
}

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}
