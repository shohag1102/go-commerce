package handlers

import (
	"net/http"
	"shohag/github.com/database"
	"shohag/github.com/util"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Something went wrong", 404)
		return
	}

	database.Delete(id)

	util.SendData(w, "product deleted", 201)
}
