package main

import (
	"errors"
	"fmt"
	"sync"
)

type User struct {
	ID      string
	Name    string
	Balance float64
	mu      sync.RWMutex
}

func (u *User) Deposit(amount float64) {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.Balance += amount
}

func (u *User) WithDraw(amount float64) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	if u.Balance >= amount {
		u.Balance -= amount
		return nil
	}
	return errors.New("not enough money on the user ID: " + u.ID)
}

type Transaction struct {
	FromID string
	ToID   string
	Amount float64
}

type PaymentSystem struct {
	Users            map[string]*User
	TransactionQueue []Transaction
}

func NewPaymentSystem() *PaymentSystem {
	return &PaymentSystem{
		Users:            make(map[string]*User),
		TransactionQueue: make([]Transaction, 0),
	}
}

func (ps *PaymentSystem) AddUser(id string, name string, balance float64) {
	ps.Users[id] = &User{
		ID:      id,
		Name:    name,
		Balance: balance,
	}
}

func (ps *PaymentSystem) AddTransaction(fromID, toID string, amount float64) {
	ps.TransactionQueue = append(ps.TransactionQueue, Transaction{
		FromID: fromID,
		ToID:   toID,
		Amount: amount,
	})
}

func (ps *PaymentSystem) ProcessingTransactions(tr Transaction) error {
	_, ok := ps.Users[tr.FromID]
	if !ok {
		err := fmt.Sprintf("user with id: \"%s\" is not found!", tr.FromID)
		return errors.New(err)
	}

	_, ok = ps.Users[tr.ToID]
	if !ok {
		err := fmt.Sprintf("user with id: \"%s\" is not found!", tr.FromID)
		return errors.New(err)
	}

	err := ps.Users[tr.FromID].WithDraw(tr.Amount)
	if err != nil {
		return err
	}
	ps.Users[tr.ToID].Deposit(tr.Amount)

	return nil
}

func (ps *PaymentSystem) Worker(ch <-chan Transaction, wg *sync.WaitGroup) error {
	defer wg.Done()

	for v := range ch {
		err := ps.ProcessingTransactions(v)
		if err != nil {
			return err
		}
	}
	return nil

}

func main() {

	ps := NewPaymentSystem()

	ps.AddUser("111", "Григорий", 50000)
	ps.AddUser("222", "Петр", 100000)
	ps.AddUser("333", "Василий", 15000)

	ps.AddTransaction("111", "222", 10000)
	ps.AddTransaction("333", "222", 5000)
	ps.AddTransaction("222", "333", 100000)

	ps.Users["111"].Deposit(100)
	ps.Users["222"].Deposit(500)
	ps.Users["333"].Deposit(1000)

	err := ps.Users["111"].WithDraw(100)
	if err != nil {
		fmt.Println(err)
	}
	err = ps.Users["222"].WithDraw(100)
	if err != nil {
		fmt.Println(err)
	}
	err = ps.Users["333"].WithDraw(100)
	if err != nil {
		fmt.Println(err)
	}

	ch := make(chan Transaction, len(ps.TransactionQueue))

	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go ps.Worker(ch, &wg)
	}

	for _, v := range ps.TransactionQueue {
		ch <- v
	}
	close(ch)

	wg.Wait()

	fmt.Printf("Айди: %s, Имя: %s, Баланс: %.2f руб.\n", ps.Users["111"].ID, ps.Users["111"].Name, ps.Users["111"].Balance)
	fmt.Printf("Айди: %s, Имя: %s, Баланс: %.2f руб.\n", ps.Users["222"].ID, ps.Users["222"].Name, ps.Users["222"].Balance)
	fmt.Printf("Айди: %s, Имя: %s, Баланс: %.2f руб.\n", ps.Users["333"].ID, ps.Users["333"].Name, ps.Users["333"].Balance)
}
