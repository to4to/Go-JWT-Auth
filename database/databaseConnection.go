package database

import (
	"log"

	"github.com/joho/godotenv"
)

func DBinstance() *mongo.Client{

err:=godotenv.Load(".env")

if err!=nil{
	log.Fatal("Error Loading .env File")
}

}