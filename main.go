package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit"
)

type Item struct {
	Name  string
	Price float64
}

type Customer struct {
	Name  string
	Items []Item
}

func generateItems() []Item {
	items := []Item{
		{"Laptop", 15000},
		{"Mouse", 500},
		{"Keyboard", 1000},
		{"Monitor", 3000},
		{"Headphone", 2000},
	}
	return items
}

func generateCustomers(num int, items []Item) []Customer {
	customers := []Customer{}
	for i := 0; i < num; i++ {
		gofakeit.Seed(time.Now().UnixNano())
		name := gofakeit.Name()
		customerItems := []Item{items[rand.Intn(len(items))], items[rand.Intn(len(items))]}
		customers = append(customers, Customer{Name: name, Items: customerItems})
	}
	return customers
}

func cashier(id int, customers <-chan Customer, wg *sync.WaitGroup, mu *sync.Mutex, totalRevenue *float64) {
	defer wg.Done()
	for customer := range customers {
		time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
		total := 0.0
		fmt.Printf("Kasir %d melayani pelanggan %s\n", id, customer.Name)
		for _, item := range customer.Items {
			fmt.Printf("  - %s: Rp%.2f\n", item.Name, item.Price)
			total += item.Price
		}
		fmt.Printf("Total belanja %s: Rp%.2f\n\n", customer.Name, total)
		mu.Lock()
		*totalRevenue += total
		mu.Unlock()
	}
}

func main() {
	items := generateItems()
	customers := generateCustomers(100, items)
	totalRevenue := 0.0
	var mu sync.Mutex
	var wg sync.WaitGroup
	customerChannel := make(chan Customer, len(customers))

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go cashier(i, customerChannel, &wg, &mu, &totalRevenue)
	}

	for _, customer := range customers {
		customerChannel <- customer
	}
	close(customerChannel)

	wg.Wait()

	fmt.Printf("Total pemasukan toko: Rp%.2f\n", totalRevenue)
}
