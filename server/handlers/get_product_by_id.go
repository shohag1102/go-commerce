package handlers

import (
	"net/http"
	"shohag/github.com/database"
	"shohag/github.com/util"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Something went wrong", 404)
		return
	}

	for _, product := range database.ProductList {
		if product.ID == id {
			util.SendData(w, product, 200)
			return
		}
	}
	util.SendData(w, "Please provide valid product id", 404)
}
