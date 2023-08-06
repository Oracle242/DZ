package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BankClient interface {
	Deposit(amount int)
	Withdrawal(amount int) error
	Balance() int
}

// type Client struct {
// 	balance int
// }

// func (c *Client) Balance() int {
// 	return c.balance
// }

var mu = sync.Mutex{}

func main() {
	rand.Seed(time.Now().UnixNano())
	var amount int
	menu(amount)
}

func deposit(amount *int) error {
	for {
		time.Sleep(time.Second / 5)
		mu.Lock()
		result := *amount
		result += rand.Intn(10) + 1
		*amount = result
		mu.Unlock()
	}
}

func withdrawal(amount *int) error {
	for {
		time.Sleep(time.Second / 5)
		mu.Lock()
		result := *amount
		result -= rand.Intn(5) + 1
		*amount = result
		fmt.Println(result, "tw")
		mu.Unlock()
	}

}

func menu(amount int) {
	for i := 0; i < 10; i++ {
		go deposit(&amount)
	}
	for i := 0; i < 5; i++ {
		go withdrawal(&amount)
	}
	for {
		command := ""
		fmt.Scanln(&command)
		switch command {
		case "balance":
		case "deposit":
		case "withdrawal":
		case "exit":
			return
		default:
			fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit")
		}
	}
}
