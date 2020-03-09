package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"users/driver"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	userHandler "users/user/handler"
	userRepo "users/user/repo"
	userUseCase "users/user/usecase"
)

func init() {
	gotenv.Load()
}

func main() {
	//Open port for connect to database
	PORT := os.Getenv("PORT")
	db, err := driver.Connect()
	if err != nil {
		fmt.Println("Error when connect to database")
		log.Fatal(err)
	}
	// di gunakan untuk, bahwa akan di jalankan terakhir
	defer db.Close()
	// StrictSlah itu untuk mendeklarasikan ketika path example.com/user/ tidak akan err
	router := mux.NewRouter().StrictSlash(true)

	//Depedency Injection
	userRepo := userRepo.CreateUserRepoImpl(db)
	userUseCase := userUseCase.CreateUserUseCaseImpl(userRepo)
	userHandler.CreateUserHandler(router, userUseCase)

	fmt.Println("Web Start on Port :", PORT)

	err = http.ListenAndServe(":"+PORT, router)
	if err != nil {
		fmt.Println("Error on http Listen")
		log.Fatal(err)
	}
}
