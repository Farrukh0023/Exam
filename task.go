package main

import "fmt"


//Task 1: Barcha mijozlar ma'lumotlarini chop eting
func printCustomerDetails(customers []Customer) {
	fmt.Println("Task 1: Barcha customerlarni umumiy mablagini va qancha summa harid qilganini hisoblang va ko'rsating:")
	totalCash := 0.0
	totalSpent := 0.0

	for _, customer := range customers {
		printCustomerInfo(customer)
		totalCash += customer.Cash
		totalSpent += customer.Basket.Total
	}
	fmt.Printf("Umumiy mijoz naqd puli: %.2f\n", totalCash)
	fmt.Printf("Sarflangan umumiy miqdor: %.2f\n", totalSpent)
}

// Eng ko'p pul sarflaydigan mijozni toping
func findTopSpender(customers []Customer) Customer {
	fmt.Println("\nTask 2: Eng ko'p pul sarflagan mijoz")
	if len(customers) == 0 {
		// Eğer müşteri yoksa, nil bir müşteri döndür
		return Customer{}
	}

	topSpender := customers[0]

	for _, customer := range customers {
		if customer.Basket.Total > topSpender.Basket.Total {
			topSpender = customer
		}
	}

	return topSpender
}

// Eng qimmat mahsulotni topish
func findMostExpensiveProduct(products []Product) Product {
	if len(products) == 0 {
		return Product{}
	}

	mostExpensive := products[0]

	for _, product := range products {
		if product.Price > mostExpensive.Price {
			mostExpensive = product
		}
	}

	return mostExpensive
}

// Barcha mahsulotlarni olib keling
func allProducts(customers []Customer) []Product {
    var allProds []Product
    for _, customer := range customers {
        allProds = append(allProds, customer.Basket.Products...)
    }
    return allProds
}



// Mahsulot haqida ma'lumotni chop eting
func printProductInfo(product Product) {
	fmt.Printf("Toifa: %s\n", product.Category)
	fmt.Printf("Mahsulot ismi: %s\n", product.Name)
	fmt.Printf("Narh: %.0f\n", product.Price)
	fmt.Printf("Miqdor: %d\n", product.Quantity)
	fmt.Println("------------------------------")
}

// Barcha mahsulotlarning o'rtacha narxini hisoblang va ko'rsating
func calculateAndPrintAveragePrice(allProducts []Product) {
	fmt.Println("\nTask 4: Barcha maxsulotlarning ortacha narxi")
    if len(allProducts) == 0 {
        fmt.Println("Maxsulot topilmadi.")
        return
    }

    var total float64
    for _, product := range allProducts {
        total += product.Price
    }

    average := total / float64(len(allProducts))
    fmt.Printf("Barcha maxsulotlarning ortalama narxi: %.0f\n", average)
}




// Jami eng past narxda sotib olgan mijozni topish
func findLowestSpender(customers []Customer) Customer {
    if len(customers) == 0 {
        return Customer{}
    }

    lowestSpender := customers[0]

    for _, customer := range customers {
        if customer.Basket.Total < lowestSpender.Basket.Total {
            lowestSpender = customer
        }
    }

    return lowestSpender
}

// Eng past narxda jami xarid qilgan mijozni toping va chop eting
func printLowestSpender(customers []Customer) {
    lowestSpender := findLowestSpender(customers)
	fmt.Println("\nTask 5: Eng past narxda xarid qilgan mijoz")
    if lowestSpender.ID == "" {
        fmt.Println("Mijoz topimadi")
        return
    }

    fmt.Println("Eng past narxda xarid qilgan mijoz:")
    printCustomerInfo(lowestSpender)
}

//Eng ko'p sotiladigan mahsulot toifasini aniqlash
func findBestSellingCategory(customers []Customer) string {
	fmt.Println("\nTask 6:")
	if len(customers) == 0 {
		return ""
	}

	categoryCounts := make(map[string]int)

	for _, customer := range customers {
		for _, product := range customer.Basket.Products {
			categoryCounts[product.Category] += product.Quantity
		}
	}

	bestSellingCategory := ""
	maxQuantity := 0
	for category, quantity := range categoryCounts {
		if quantity > maxQuantity {
			bestSellingCategory = category
			maxQuantity = quantity
		}
	}

	return bestSellingCategory
}

// Eng yaxshi va eng kam sotilgan mahsulotlarni toping
func findMinMaxSoldProducts(customers []Customer) (Product, Product) {
	allProds := allProducts(customers)
	if len(allProds) == 0 {
		return Product{}, Product{}
	}

	maxSold := allProds[0]
	minSold := allProds[0]

	for _, product := range allProds {
		if product.Quantity > maxSold.Quantity {
			maxSold = product
		} else if product.Quantity < minSold.Quantity {
			minSold = product
		}
	}

	return maxSold, minSold
}

func printMinMaxSoldProducts(customers []Customer) {
	maxSold, minSold := findMinMaxSoldProducts(customers)

	if maxSold.ID == "" || minSold.ID == "" {
		fmt.Println("Maxsulot topilmadi.")
		return
	}

	fmt.Println("---------------------------------------")
	fmt.Println("Task 7:\nEng kop va eng kam sotilgan maxsulot")
	fmt.Println("Eng kop sotilgan maxsulot:")
	printProductInfo(maxSold)
	fmt.Println("Eng kam sotigan maxsulot:")
	printProductInfo(minSold)

}

// Task 8: Har bir savdo uchun o'rtacha mahsulot miqdorini hisoblang va ko'rsating
func calculateAndPrintAverageQuantitySold(customers []Customer) {
	mahsulotMiqdori:=0
	var  OrtaMiqdor float64 = 1.1
	sum:=0
	for _, customer := range customers{
	  for _, product := range customer.Basket.Products{
			 mahsulotMiqdori+=product.Quantity
		 sum++
	  }
	}
	fmt.Println(sum)
	OrtaMiqdor  = float64(mahsulotMiqdori)/float64(sum) 
	fmt.Printf("O'rtacha mahsulot miqdori: %.3f \n", OrtaMiqdor)
  }


// Savatdagi narsalarning umumiy miqdorini hisoblang
func calculateBasketTotalQuantity(basket Basket) int {
    totalQuantity := 0
    for _, product := range basket.Products {
        totalQuantity += product.Quantity
    }
    return totalQuantity
}

// Task 9: Eng ko'p mahsulot sotib olgan mijozni topadigan funksiya
func findTopCustomerByProductQuantity(customers []Customer) {
	fmt.Println("\nTask 9: ")
	Çokas := 0
	var customerName string
	var customerLastName string
	for _, customer := range customers {
	  totalProducts:=0
	  for _, products := range customer.Basket.Products {
		totalProducts+=products.Quantity  
	  if totalProducts > Çokas {
		Çokas = totalProducts
		customerName = customer.FirstName
		customerLastName = customer.LastName
	  }
	 }
	}
	fmt.Printf("Eng ko'p mahsulot sotib olgan mijoz: %s %s (%d ta)\n", customerName, customerLastName, Çokas)
  }

// Task 10: Sotilgan mahsulotlarda eng keng tarqalgan mahsulotni toping va chop eting
func findMostSoldProduct(customers []Customer) {
	fmt.Println("\nTask 10: ")
    MahsulotSoni := make(map[string]int)
  
    for _, customer := range customers {
          for _, product := range customer.Basket.Products {
        MahsulotSoni[product.Name]++
      }
    }
    var CokGorunenUrunAdi string
    CokGorununen := 0
    for productName, frequency := range MahsulotSoni {
      if frequency > CokGorununen {
        CokGorununen = frequency
        CokGorunenUrunAdi = productName
      }
    }
  
  if CokGorununen > 0 {
      fmt.Printf("Eng ko'p savdolarda ko'rinadigan mahsulot: %s (%d ta)\n", CokGorunenUrunAdi, CokGorununen)
    } else {
      fmt.Println("Mijozlar savatchalarida hech qanday mahsulot topilmadi.")
    }
  }

// ID bo'yicha mahsulotlarni topish
func findProductByID(products []Product, productID string) Product {
    for _, product := range products {
        if product.ID == productID {
            return product
        }
    }
    return Product{}
}

// Task 11
// O'rtacha savdo mablag'i bo'yicha eng ko'p mablag'ga ega customerni aniqlang va ko'rsating
func calculateAndPrintAverageSpending(customers []Customer) {
	fmt.Println("\nTask 11: ")
    maxOrtachaXarajat := 0
    var enCokHarcayan, enCokHarcayanSoyad string

    for _, mijoz := range customers {
        jamiXarajat := 0
        for _, customer := range mijoz.Basket.Products {
            jamiXarajat += int(customer.Price) * customer.Quantity
        }

        ortachaXarajat := jamiXarajat / len(mijoz.Basket.Products)
        if ortachaXarajat > maxOrtachaXarajat {
            maxOrtachaXarajat = ortachaXarajat
            enCokHarcayan = mijoz.FirstName
            enCokHarcayanSoyad = mijoz.LastName
        }
    }

    fmt.Printf("O'rtacha savdo mablag'i: %d so'm\n", maxOrtachaXarajat)
    fmt.Printf("Eng kop xarajatga ega bo'lgan mijoz: %s %s\n", enCokHarcayan, enCokHarcayanSoyad)
}

// Task 12: Eng koʻp umumiy daromad keltiradigan toifani koʻrsatish (quantity*price)
func findMostProfitableCategory(customers []Customer) {
	fmt.Println("\nTask 12: ")
	if len(customers) == 0 {
		fmt.Println("Mijoz topilmadi")
		return
	}

	categoryProfits := make(map[string]float64)

	for _, customer := range customers {
		for _, product := range customer.Basket.Products {
			profit := float64(product.Quantity) * product.Price
			categoryProfits[product.Category] += profit
		}
	}

	mostProfitableCategory := ""
	maxProfit := 0.0
	for category, profit := range categoryProfits {
		if profit > maxProfit {
			mostProfitableCategory = category
			maxProfit = profit
		}
	}

	fmt.Printf("Eng ko'p umumiy daromad keltiradigan toifa: %s (Jami daromad: %.2f)\n", mostProfitableCategory, maxProfit)
}



// Task 13: Har bir foydalanuvchi tomonidan sotib olingan eng qimmat buyumni ko'rsatish
func findMostExpensivePurchaseByCustomer(customers []Customer) {
	fmt.Println("\nTask 13: ")
	if len(customers) == 0 {
		fmt.Println("Mijoz topilmadi")
		return
	}

	for _, customer := range customers {
		mostExpensiveProduct := findMostExpensiveProduct(customer.Basket.Products)

		if mostExpensiveProduct.ID != "" {
			fmt.Printf("%s %s isimli xaridor tomonidan sotib olingan eng qimmat buyum:\n", customer.FirstName, customer.LastName)
			printProductInfo(mostExpensiveProduct)
		} else {
			fmt.Printf("%s %s Xaridor tomonidan sotib olingan eng qimmat buyum topilmadi.\n", customer.FirstName, customer.LastName)
		}
	}
}




// Task 14: Har bir foydalanuvchi eng ko'p sotib olgan toifani va ushbu turkumda qancha pul sarflaganini ko'rsating.
func findMostExpensiveCategoryByCustomer(customers []Customer) {
	fmt.Println("\nTask 14: ")
	if len(customers) == 0 {
		fmt.Println("Mijoz topilmadi")
		return
	}

	for _, customer := range customers {
		categorySpending := make(map[string]float64)

		for _, product := range customer.Basket.Products {
			categorySpending[product.Category] += product.Price * float64(product.Quantity)
		}

		mostExpensiveCategory := ""
		maxSpending := 0.0

		for category, spending := range categorySpending {
			if spending > maxSpending {
				mostExpensiveCategory = category
				maxSpending = spending
			}
		}

		if mostExpensiveCategory != "" {
			fmt.Println("--------------------------------------------------------")
			fmt.Printf("Task 14: %s %s mijozning eng kop maxsulot olgan toifasi: %s\n", customer.FirstName, customer.LastName, mostExpensiveCategory)
			fmt.Printf("Bu toifada sotib olgan hamma maxsulot narxi: %.2f\n", maxSpending)
			fmt.Println("--------------------------------------------------------")
		} else {
			fmt.Printf("Task 14: %s %s mijozning eng kop maxsulot olgan toifasi topimadi.\n", customer.FirstName, customer.LastName)
		}
	}
}

// Task 15: Har bir mahsulotning nechtasi jami sotilganini va jami nechta mahsulot sotilganini ko'rsating.
func printTotalSoldQuantity(products []Product) {
	fmt.Println("\nTask 15: ")
	if len(products) == 0 {
		fmt.Println("Sotilgan maxsulot topilmadi.")
		return
	}

	productSoldQuantity := make(map[string]int)
	totalSoldQuantity := 0

	for _, product := range products {
		productSoldQuantity[product.Name] += product.Quantity
		totalSoldQuantity += product.Quantity
	}

	fmt.Println("Task 15: Har bir mahsulotning sotilgan umumiy miqdori:")
	for productName, quantity := range productSoldQuantity {
		fmt.Printf("%s: %d ta\n", productName, quantity)
	}

	fmt.Printf("Sotilgan mahsulotlarning umumiy miqdori: %d ta\n", totalSoldQuantity)
}
