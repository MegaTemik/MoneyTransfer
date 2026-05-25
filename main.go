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

func (u *User) Deposit(amount float64) {
	u.Balance += amount
}

func (u *User) WithDraw(amount float64) error {
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
	Users        map[string]*User
	Transactions []Transaction
}

func NewPaymentSystem() *PaymentSystem {
	return &PaymentSystem{
		Users:        make(map[string]*User),
		Transactions: make([]Transaction, 0),
	}
}

func (ps *PaymentSystem) AddUser(Id string, user *User) {
	ps.Users[Id] = user
}

func (ps *PaymentSystem) AddTransaction(transaction Transaction) {
	ps.Transactions = append(ps.Transactions, transaction)
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
	u1 := &User{
		ID:      "111",
		Name:    "Григорий",
		Balance: 50000,
	}

	u2 := &User{
		ID:      "222",
		Name:    "Пётр",
		Balance: 101010,
	}

	u3 := &User{
		ID:      "333",
		Name:    "Василий",
		Balance: 15000,
	}

	u1.Deposit(100)
	err := u1.WithDraw(60000000000)
	if err != nil {
		fmt.Println(err)
	}

	err = u2.WithDraw(1000)
	if err != nil {
		fmt.Println(err)
	}
	u2.Deposit(5000)

	u3.Deposit(50)
	err = u3.WithDraw(1)
	if err != nil {
		fmt.Println(err)
	}

	tr1 := Transaction{
		FromID: u1.ID,
		ToID:   u2.ID,
		Amount: 10000,
	}

	tr2 := Transaction{
		FromID: u3.ID,
		ToID:   u2.ID,
		Amount: 5000,
	}

	tr3 := Transaction{
		FromID: u2.ID,
		ToID:   u3.ID,
		Amount: 100000,
	}

	ps := NewPaymentSystem()
	ps.Transactions = append(ps.Transactions, tr1, tr2, tr3)

	ps.Users[u1.ID] = u1
	ps.Users[u2.ID] = u2
	ps.Users[u3.ID] = u3

	for _, v := range ps.Transactions {
		err := ps.ProcessingTransactions(v)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Printf("Айди: %s, Имя: %s, Баланс: %.2f руб.\n", u1.ID, u1.Name, u1.Balance)
	fmt.Printf("Айди: %s, Имя: %s, Баланс: %.2f руб.\n", u2.ID, u2.Name, u2.Balance)
	fmt.Printf("Айди: %s, Имя: %s, Баланс: %.2f руб.\n", u3.ID, u3.Name, u3.Balance)
}
