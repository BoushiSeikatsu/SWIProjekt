package initializers

import (
	"log"
	"os"

	"github.com/BoushiSeikatsu/SWIProjekt/models"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Configuration struct {
	Port string `json:"port"`
	IP  string `json:"ip"`
	Adminpass string `json:"adminpass"`
	Secret string `json:"secret"`
}

var Config Configuration = Configuration{
	Port: "8080",
	IP: "localhost",
	Adminpass: "admin123",
	Secret: "123",
}

func ConnectToDB() {
	db, err := gorm.Open(sqlite.Open("warehouse.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to load DB")
		return
	}

	DB = db

	DB.AutoMigrate(&models.Product{}, &models.Storage{}, &models.Store{}, &models.Icoming{}, &models.Distribution{})
}

func Load() {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Failed to load config file")
		return
	}

	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		log.Fatal("Failed to parse config file")
		return
	}

	log.Println("Config loaded")
}