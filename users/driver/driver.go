package driver

import (
	"fmt"
	"log"
	"os"
	"users/model"

	// butuh di install terlebih dahulu pake go get
	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

//Pascal case function  menandakan bahwa si function itu typenya public
// () yang pertama itu untuk parameter () yang ke dua itu return
func Connect() (*gorm.DB, error) {
	URL := os.Getenv("URL")
	db, error := gorm.Open("postgres", URL)
	if error != nil {
		fmt.Println("[driver.Connect] got error when excute gorm")
		log.Fatal(error)
		return nil, error
	}
	// Pointer & kalo ada perubahan hanya meraplace
	// Ketika di generate gorm otomatis membuat 4 field (id, create_at, update_at, delete_At)
	db.Debug().AutoMigrate(&model.User{})
	return db, nil

}
