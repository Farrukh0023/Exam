package main

import "fmt"

func main() {
	// JSON dosyasını okuma
	filename := "store_data.json"
	customers, err := readData(filename)
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	// Task 1: Barcha mijozlar ma'lumotlarini va jami mablag'larni chop eting
	 //printCustomerDetails(customers)

	// Task 2: Eng ko'p pul sarflaydigan mijozni toping va chop eting
	 //topSpender := findTopSpender(customers)
	 //printCustomerInfo(topSpender)

	// Task 3: Eng qimmat mahsulotni toping va chop eting
	 //fmt.Println("\nTask 3: Eng qimmat mahsulot")
	 allProds := allProducts(customers)
	 //mostExpensiveProduct := findMostExpensiveProduct(allProds)
	 //printProductInfo(mostExpensiveProduct)

	// Task 4: Barcha mahsulotlarning o'rtacha narxini hisoblang va ko'rsating
	 calculateAndPrintAveragePrice(allProds)

	// Task 5: Eng past narxda xarid qilgan mijozni toping va chop eting
	 printLowestSpender(customers)

	// Task 6: Eng kop sotilgan toifa
	 //bestSellingCategory := findBestSellingCategory(customers)
	 //fmt.Println("Eng kop sotilgan toifa:", bestSellingCategory)

	// Task 7: Eng kop va eng kam sotilgan maxsulotni chop eting
     //printMinMaxSoldProducts(customers)

	// Task 8: Har bir savdo uchun oʻrtacha mahsulot miqdorini hisoblang va ko'rsating:
	 calculateAndPrintAverageQuantitySold(customers)

	// Task 9: Eng ko'p mahsulot sotib olgan mijozni va jami qancha mahsulot sotib olganini ko'rsating.
	 //findTopCustomerByProductQuantity(customers)

	// Task 10: Sotilgan mahsulotlarda eng keng tarqalgan mahsulotni toping va chop eting
	 //findMostSoldProduct(customers)

	// Task 11: Oʻrtacha savdo mablagʻi boʻyicha eng koʻp mablagʻaga ega boʻlgan customerni aniqlang va ko'rsating:
	 //calculateAndPrintAverageSpending(customers)

	// Task 12: Eng ko'p umumiy daromad keltiradigan toifani ko'rsating
	 //findMostProfitableCategory(customers)

	// Task 13: Har bir foydalanuvchi tomonidan sotib olingan eng qimmat buyumni ko'rsatish
	 //findMostExpensivePurchaseByCustomer(customers)

	// Task 14: Har bir foydalanuvchi eng ko'p sotib olgan toifani va ushbu turkumda qancha pul sarflaganini ko'rsating.
	 //findMostExpensiveCategoryByCustomer(customers)

	// Task 15: Her bir üründen toplamda kaç tane satıldığını ve toplamda kaç tane ürün satıldığını ekrana çıkart
	 //printTotalSoldQuantity(allProds)
}