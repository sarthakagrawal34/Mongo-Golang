package controllers

import (
	"encoding/json"
	"fmt"
	"mongo-golang/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController{
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	oid := bson.ObjectIdHex(id)

	u := models.User{}

	if err := uc.session.DB("mongo-golang").C("users").FindId(oid).One(u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(uj)
	w.WriteHeader(200)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()

	if err := uc.session.DB("mongo-golang").C("users").Insert(u); err != nil {
		w.WriteHeader(500)
		return
	}

	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "pkglication/json")
	w.Write(uj)
	w.WriteHeader(201)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("mongo-golang").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "pkglication/json")
	fmt.Fprintf(w, "Deleted User with oid as: %v\n", oid)
	
	w.WriteHeader(200)
}