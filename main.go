package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BankClient struct {
	balance int
}

var mu = sync.RWMutex{}

func main() {
	client := &BankClient{}
	rand.Seed(time.Now().UnixNano())
	amount := 0
	for i := 0; i < 10; i++ {
		go func() {

			for {
				amount := rand.Intn(10) + 1
				client.Deposit(&amount)
			}
		}()
	}
	for i := 0; i < 5; i++ {
		go func() {
			for {
				amount := rand.Intn(5) + 1
				err := client.Withdrawal(&amount)
				if err != nil {
					fmt.Println(err)
				}
			}
		}()
	}
	for {
		command := ""
		fmt.Scanln(&command)
		switch command {
		case "balance":
			fmt.Println(client.Balance())
		case "deposit":
			fmt.Scanln(&amount)
			client.Deposit(&amount)
			fmt.Println(client.Balance())
		case "withdrawal":
			fmt.Scanln(&amount)
			client.Withdrawal(&amount)
			fmt.Println(client.Balance())
		case "exit":
			return
		default:
			fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit")
		}
	}
}

func (bc *BankClient) Deposit(amount *int) {
	time.Sleep(time.Second / 5)
	mu.RLock()
	bc.balance += *amount
	mu.RUnlock()
}

func (bc *BankClient) Withdrawal(amount *int) error {
	time.Sleep(time.Second / 5)
	mu.RLock()
	if bc.balance < *amount {
		return fmt.Errorf("not enough balance")
	}
	bc.balance -= *amount
	mu.RUnlock()
	return nil

}

func (bc *BankClient) Balance() int {
	return bc.balance
}
