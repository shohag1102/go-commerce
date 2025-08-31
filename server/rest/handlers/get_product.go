package handlers

import (
	"net/http"
	"shohag/github.com/database"
	"shohag/github.com/util"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Something went wrong", 404)
		return
	}

	product := database.Get(id)

	if product == nil {
		util.SendError(w, "product not found", 404)
		return
	}
	util.SendData(w, product, 200)
}
