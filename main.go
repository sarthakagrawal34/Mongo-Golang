package main

import (
	"fmt"
	"log"
	"mongo-golang/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/users/:id", uc.GetUser)
	r.POST("/users", uc.CreateUser)
	r.DELETE("/users/:id", uc.DeleteUser)

	fmt.Println("Server running at port: 9000")
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	fmt.Println(s)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer s.Close()
	return s
}
