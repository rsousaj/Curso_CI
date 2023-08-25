package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	val, ok := os.LookupEnv("DB_PASSWORD")
	var stringDeConexao string
	if !ok {
		stringDeConexao = "host=localhost user=user password=root dbname=appdb port=5433 sslmode=disable"
	} else {
		stringDeConexao = "host=localhost user=user password=" + val + " dbname=appdb port=5433 sslmode=disable"
	}

	log.Print(stringDeConexao)

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
