package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shohag/github.com/database"
	"shohag/github.com/util"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "Something went wrong", 404)
		return
	}

	var updatedProd database.Product

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&updatedProd)
	if err != nil {
		fmt.Println("error", err)
		http.Error(w, "Plz give valid json", 400)
		return
	}

	updatedProd.ID = id

	database.Update(updatedProd)

	util.SendData(w, "product updated", 201)
}
