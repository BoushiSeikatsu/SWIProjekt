package initializers

import (
	"log"
	"os"
	"swi-warehouse/models"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var PORT string
var IP string
var DBName string
var SECRET string

func ConnectToDB() {
	db, err := gorm.Open(sqlite.Open(DBName), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db

	DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Manufacturer{}, &models.Added{}, &models.Removed{}, &models.Storage{}, &models.Stores{})
}

type Configuration struct {
	Port     string `json:"port"`
	IP	 	string `json:"ip"`
	DB      string `json:"db"`
	SECRET string `json:"secret"`
}

func Load() {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Failed to open config file")
	}

	configuration := Configuration{
		DB: "register.db",
	}
	err = yaml.Unmarshal(file, &configuration)
	if err != nil {
		log.Fatal("Failed to read yaml file")
	}

	PORT = configuration.Port
	IP = configuration.IP
	DBName = configuration.DB
	SECRET = configuration.SECRET
}
