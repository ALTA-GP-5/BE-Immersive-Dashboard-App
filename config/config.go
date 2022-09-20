package config

import (
	"log"
	"os"
	"strconv"
	"sync"
)

type AppConfig struct {
	DB_NAME     string
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASS     string
	SERVER_PORT string
	JWT_SECRET  string
	BASE_URL    string
}

var mutex = sync.Mutex{}

func GetConfig() *AppConfig {
	mutex.Lock()
	defer mutex.Unlock()
	return loadConfig()
}

func loadConfig() *AppConfig {
	DB_PORT, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err.Error())
	}

	return &AppConfig{
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     DB_PORT,
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASS:     os.Getenv("DB_PASS"),
		SERVER_PORT: os.Getenv("SERVER_PORT"),
		JWT_SECRET:  os.Getenv("JWT_SECRET"),
		BASE_URL:    os.Getenv("BASE_URL"),
	}
}
