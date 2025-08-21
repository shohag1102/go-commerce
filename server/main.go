package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImgURL      string `json:"imageUrl"`
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about page")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightRequest(w, r)

	sendData(w, productList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightRequest(w, r)

	var newProduct Product

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println("error", err)
		http.Error(w, "Plz give valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1

	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)
}

func handleCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, habib")
	w.Header().Set("Content-Type", "application/json")
}

func handlePreflightRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)

	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux() // router
	mux.Handle("GET /hello", http.HandlerFunc(helloHandler))
	mux.Handle("GET /about", http.HandlerFunc(aboutHandler))
	mux.Handle("GET /products", http.HandlerFunc(getProducts))
	mux.Handle("POST /create-products", http.HandlerFunc(createProduct))

	fmt.Println("Server is running on port: 8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Error starting the server, ", err)
	}
}

func init() {
	prod1 := Product{
		ID: 1, Title: "Wireless Headphones", Description: "High-quality noise-cancelling headphones.", Price: "129.99", ImgURL: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS",
	}
	prod2 := Product{
		ID: 2, Title: "Smart Watch", Description: "Stylish smart watch with health tracking.", Price: "199.99", ImgURL: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528",
	}

	prod3 := Product{
		ID: 3, Title: "Running Shoes", Description: "Lightweight shoes for everyday running.", Price: "89.50", ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s",
	}

	// prod4 := Product{
	// ID: 4, Title: "Wireless Headphones", Description: "High-quality noise-cancelling headphones.", Price: 129.99, ImgURL: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS", }

	// prod5 := Product{
	// 	ID: 5, Title: "Smart Watch", Description: "Stylish smart watch with health tracking.", Price: 199.99, ImgURL: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528",
	// }

	// prod6 := Product{
	// 	ID: 6, Title: "Running Shoes", Description: "Lightweight shoes for everyday running.", Price: 89.50, ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s",
	// }

	// productList = append(productList, prod1, prod2, prod3, prod4, prod5, prod6)
	productList = append(productList, prod1, prod2, prod3)
	// fmt.Println("products", productList)

}
