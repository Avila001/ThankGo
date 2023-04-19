package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// errInsufficientFunds сигнализирует,
// что на счете недостаточно денег,
// чтобы выполнить списание
var errInsufficientFunds error = errors.New("insufficient funds")

// errInvalidAmount сигнализирует,
// что указана некорректная сумма транзакции
var errInvalidAmount error = errors.New("invalid transaction amount")

// bankAccount представляет счет
type bankAccount struct {
	balance   int
	overdraft int
}

// deposit зачисляет деньги на счет.
func (acc *bankAccount) deposit(amount int) error {
	// ...
	if amount <= 0 {
		return errInvalidAmount
	}
	acc.balance += amount
	return nil
}

// withdraw списывает деньги со счета.
func (acc *bankAccount) withdraw(amount int) error {
	// ...
	if amount <= 0 {
		return errInvalidAmount
	}
	if amount > acc.balance+acc.overdraft {
		return errInsufficientFunds
	}
	acc.balance -= amount
	return nil
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

type test struct {
	acc   bankAccount
	trans []int
}

var testsbankAccount = map[string]test{
	"{100 10} [10 -50 20]":   {bankAccount{100, 10}, []int{10, -50, 20}},
	"{30 0} [-20 -10]":       {bankAccount{30, 0}, []int{-20, -10}},
	"{30 0}, [-20 -10 -10]":  {bankAccount{30, 0}, []int{-20, -10, -10}},
	"{30 0}, [-100]":         {bankAccount{30, 0}, []int{-100}},
	"{0 0}, [10 20 30]":      {bankAccount{0, 0}, []int{10, 20, 30}},
	"{0 0}, [10 -10 20 -20]": {bankAccount{0, 0}, []int{10, -10, 20, -20}},
	"{20 10}, [-20 -10]":     {bankAccount{20, 10}, []int{-20, -10}},
	"{20 10}, [-20 -10 -10]": {bankAccount{20, 10}, []int{-20, -10, -10}},
	"{0 100}, [-20 -10]":     {bankAccount{0, 100}, []int{-20, -10}},
	"{0 30}, [-20 -10]":      {bankAccount{0, 30}, []int{-20, -10}},
	"{0 30}, [-20 -10 -10]":  {bankAccount{0, 30}, []int{-20, -10, -10}},
	"{70 30}, [-100 100]":    {bankAccount{70, 30}, []int{-100, 100}},
	"{100 10}, [10 0 20]":    {bankAccount{100, 10}, []int{10, 0, 20}},
}

func mainBankAccount() {
	var err error
	name, err := readString()
	if err != nil {
		log.Fatal(err)
	}
	testCase, ok := testsbankAccount[name]
	if !ok {
		log.Fatalf("Test case '%v' not found", name)
	}
	for _, t := range testCase.trans {
		if t >= 0 {
			err = testCase.acc.deposit(t)
		} else {
			err = testCase.acc.withdraw(-t)
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	if err == nil {
		fmt.Println(testCase.acc)
	}
}

// readString считывает и возвращает строку из os.Stdin
func readString() (string, error) {
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return strings.TrimSpace(str), nil
}
