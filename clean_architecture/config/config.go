package config

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

type AppConfig struct {
	DBPort    uint
	DBUser    string
	DBPwd     string
	DBHost    string
	DBName    string
	JWTSecret string
}

func NewConfig() *AppConfig {
	cfg := initConfig()
	if cfg == nil {
		log.Fatal("Cannot run configuration setup")
		return nil
	}

	return cfg
}

func initConfig() *AppConfig {
	var app AppConfig

	err := godotenv.Load("config.env")
	if err != nil {
		log.Error("config error :", err.Error())
		return nil
	}

	app.DBUser = os.Getenv("DB_USER")
	app.DBPwd = os.Getenv("DB_PWD")
	app.DBHost = os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Error("config error :", err.Error())
		return nil
	}
	app.DBPort = uint(port)
	app.DBName = os.Getenv("DB_NAME")
	app.JWTSecret = os.Getenv("JWT_SECRET")

	return &app
}

func GenerateToken(id uint) string {
	claim := make(jwt.MapClaims)
	claim["autorized"] = true
	claim["id"] = id
	claim["exp"] = time.Now().Add(time.Hour * 1).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	str, err := token.SignedString([]byte("kh@k1m"))
	if err != nil {
		log.Error(err.Error())
		return ""
	}
	return str
}
