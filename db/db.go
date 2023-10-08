package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func NewDB() *gorm.DB {
	var user, password, host, port, dbname string
	
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	if os.Getenv("GO_ENV") == "dev" {
		user = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		dbname = os.Getenv("DB_NAME")
	} else {
		user = getParameter("/recruit-service-board/rds/user")
		password = getParameter("/recruit-service-board/rds/password")
		host = getParameter("/recruit-service-board/rds/host")
		port = getParameter("/recruit-service-board/rds/port")
		dbname = getParameter("/recruit-service-board/rds/name")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("mysql can not open: ",err)
	}
	log.Println("Connceted")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}

// private functions
func getParameter(name string) string {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            aws.Config{Region: aws.String("ap-northeast-1")},
	}))

	svc := ssm.New(sess)
	param, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(name),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		log.Fatalf("Failed to get parameter %s: %v", name, err)
	}
	return *param.Parameter.Value
}
