package handlers

import (
	"log"
	"net/http"
	"shohag/github.com/database"
	"shohag/github.com/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	log.Println("I am product start")

	util.SendData(w, database.ProductList, 200)
	log.Println("I am product end")
}
