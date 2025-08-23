package handlers

import (
	"net/http"
	"shohag/github.com/database"
	"shohag/github.com/util"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.ProductList, 200)
}
