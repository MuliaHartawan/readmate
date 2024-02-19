package main

import (
	"fmt"
	"os"
	"readmate/delivery/http"
	"readmate/domain"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func main() {
	domain := domain.DomainInit()
	app := http.NewHttpDelivery(domain)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT_APP")))
}
