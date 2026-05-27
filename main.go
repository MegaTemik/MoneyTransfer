package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID      string
	Name    string
	Balance float64
}

func (u *User) Deposit(amount float64) error {
	if amount <= 0.01 {
		return errors.New("the amount of deposit must be greater than 0!")
	}
	u.Balance += amount
	return nil
}

func (u *User) WithDraw(amount float64) error {
	if u.Balance >= amount && amount > 0 {
		u.Balance -= amount
		return nil
	}
	return errors.New("not enough money to withdraw from the user ID: " + u.ID)
}

type Transaction struct {
	FromID string
	ToID   string
	Amount float64
}

type PaymentSystem struct {
	Users        map[string]*User
	Transactions []Transaction
}

func NewPaymentSystem() *PaymentSystem {
	return &PaymentSystem{
		Users:        make(map[string]*User),
		Transactions: make([]Transaction, 0),
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
	ps.Transactions = append(ps.Transactions, Transaction{
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

	ps.Users[tr.FromID].WithDraw(tr.Amount)
	ps.Users[tr.ToID].Deposit(tr.Amount)

	return nil
}

func main() {

	ps := NewPaymentSystem()

	ps.AddUser("111", "Григорий", 50000)
	ps.AddUser("222", "Петр", 50000)
	ps.AddUser("333", "Василий", 15000)

	ps.AddTransaction("111", "222", 10000)
	ps.AddTransaction("333", "222", 5000)
	ps.AddTransaction("222", "333", 100000)

	err := ps.Users["111"].Deposit(100)
	if err != nil {
		fmt.Println(err)
	}
	err = ps.Users["222"].Deposit(500)
	if err != nil {
		fmt.Println(err)
	}
	err = ps.Users["333"].Deposit(1000)
	if err != nil {
		fmt.Println(err)
	}

	err = ps.Users["111"].WithDraw(10000)
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

	for _, v := range ps.Transactions {
		err := ps.ProcessingTransactions(v)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Printf("Айди: %s, Имя: %s, Баланс: %.2f руб.\n", ps.Users["111"].ID, ps.Users["111"].Name, ps.Users["111"].Balance)
	fmt.Printf("Айди: %s, Имя: %s, Баланс: %.2f руб.\n", ps.Users["222"].ID, ps.Users["222"].Name, ps.Users["222"].Balance)
	fmt.Printf("Айди: %s, Имя: %s, Баланс: %.2f руб.\n", ps.Users["333"].ID, ps.Users["333"].Name, ps.Users["333"].Balance)
}
