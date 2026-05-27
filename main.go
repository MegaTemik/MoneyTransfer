package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID      string
	Name    string
	Balance int64
}

func (u *User) Deposit(amount int64) error {
	if amount <= 1 {
		return errors.New("the amount must be greater than 0!")
	}
	u.Balance += amount
	return nil
}

func (u *User) WithDraw(amount int64) error {
	if u.Balance >= amount {
		u.Balance -= amount
		return nil
	}
	return errors.New("not enough money on the balance!")
}

func main() {
	u1 := &User{
		ID:      "111",
		Name:    "Григорий",
		Balance: 5000021,
	}

	u2 := &User{
		ID:      "222",
		Name:    "Пётр",
		Balance: 10101000,
	}

	u3 := &User{
		ID:      "333",
		Name:    "Василий",
		Balance: 1500000,
	}

	u1.Deposit(-10_000 * 100)
	u1.WithDraw(60_000 * 100)

	u2.WithDraw(100_000 * 100)
	u2.Deposit(50_000 * 100)

	u3.Deposit(5_000 * 100)
	u3.WithDraw(1_000 * 100)

	fmt.Printf("Айди: %s, Имя: %s, Баланс: %.2f руб.\n", u1.ID, u1.Name, float64(u1.Balance)/100)
	fmt.Printf("Айди: %s, Имя: %s, Баланс: %.2f руб.\n", u2.ID, u2.Name, float64(u2.Balance)/100)
	fmt.Printf("Айди: %s, Имя: %s, Баланс: %.2f руб.\n", u3.ID, u3.Name, float64(u3.Balance)/100)
}
