package main

import  (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

// Ürün (Product) struct'ı
type Product struct {
	ID       string  `json:"id"`
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

// Alışveriş Sepeti (Basket) struct'ı
type Basket struct {
	ID       string    `json:"id"`
	Products []Product `json:"products"`
	Total    float64   `json:"total"`
}

// Müşteri (Customer) struct'ı
type Customer struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Cash      float64 `json:"cash"`
	Basket    Basket  `json:"basket"`
}

// JSON'dan bilgileri okuma
func readData(filename string) ([]Customer, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var customers []Customer
	err = json.Unmarshal(data, &customers)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

// Müşteri ve alışveriş sepeti bilgilerini yazdırma
func printCustomerInfo(customer Customer) {
	fmt.Printf("Isim: %s, Familya: %s, Mijoz puli: %.2f\n",
		 customer.FirstName, customer.LastName, customer.Cash)

	// Alışveriş sepetini yazdırma
	for _, product := range customer.Basket.Products {
		fmt.Printf("   Toifa: %s, Ism: %s, Narh: %.2f, Miqdor: %d\n",
		         product.Category, product.Name, product.Price, product.Quantity)
	}

	fmt.Printf("   Umumiy hisob: %.2f\n", customer.Basket.Total)
	fmt.Println("------------------------------")
}

// Tüm müşteri bilgilerini yazdırma
func printAllCustomers(customers []Customer) {
	for _, customer := range customers {
		printCustomerInfo(customer)
	}
}