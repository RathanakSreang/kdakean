package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/joho/godotenv"
	"github.com/kdakean/kdakean/db"
	"github.com/kdakean/kdakean/model"
)

func CustomValidation(obj interface{}) { // (bool, map[string]string){
	_, err := govalidator.ValidateStruct(obj)
	errs := err.(govalidator.Errors)
	for _, e := range errs {
		fmt.Println(e)
	}
	// return result, errs
}
func main() {
	user := model.User{}
	// result, err := govalidator.ValidateStruct(user)
	// fmt.Println(result)
	// fmt.Println(err)
	CustomValidation(user)
	fmt.Println(user)
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}

	s := &http.Server{
		Addr:         ":8080",
		Handler:      initRoutes(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Println("Listening on ", "8080")
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
