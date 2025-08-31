package database

var productList []Product

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImgURL      string `json:"imageUrl"`
}


func Store(p Product) Product {
	p.ID = len(productList) + 1

	productList = append(productList, p)
	return p
}

func Get(productId int) *Product {
	
	for _, product := range productList {
		if product.ID == productId {
			return &product
		}
	}
	return nil
}

func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}

func Delete(productId int) {
	var tempList []Product
	for _, p := range productList {
		if p.ID != productId {
			tempList = append(tempList, p)
		}
	}
	productList = tempList
}

func List() []Product {
	return productList
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

	prod4 := Product{
		ID: 4, Title: "Wireless Headphones", Description: "High-quality noise-cancelling headphones.", Price: "129.99", ImgURL: "https://www.lovefoodhatewaste.com/sites/default/files/styles/twitter_card_image/public/2022-07/Citrus%20fruits.jpg.webp?itok=H1j9CCCS"}

	prod5 := Product{
		ID: 5, Title: "Smart Watch", Description: "Stylish smart watch with health tracking.", Price: "199.99", ImgURL: "https://i0.wp.com/post.healthline.com/wp-content/uploads/2021/05/apples-1296x728-header.jpg?w=1155&h=1528",
	}

	prod6 := Product{
		ID: 6, Title: "Running Shoes", Description: "Lightweight shoes for everyday running.", Price: "89.50", ImgURL: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcRZbMOVB8a8wRQ6e-UKZggiu7-edRAN1GolPQ&s",
	}

	// productList = append(productList, prod1, prod2, prod3, prod4, prod5, prod6)
	productList = append(productList, prod1, prod2, prod3, prod4, prod5, prod6)
	// fmt.Println("products", productList)

}