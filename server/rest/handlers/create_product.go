package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shohag/github.com/database"
	"shohag/github.com/util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)
	
	if err != nil {
		fmt.Println("error", err)
		http.Error(w, "Plz give valid json", 400)
		return
	}

	product := database.Store(newProduct)

	util.SendData(w, product, 201)
}
