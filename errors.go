package main

// не меняйте импорты, они нужны для проверки
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// bankAccount представляет счет
type accountErrors struct {
	balance   int
	overdraft int
}

func main_errors() {
	var acc accountErrors
	var trans []int

	acc, trans, err := parseInputErrors()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("-> ")
	fmt.Println(acc, trans)
}

// parseInputErrors считывает счет и список транзакций из os.Stdin.
func parseInputErrors() (accountErrors, []int, error) {
	accSrc, transSrc := readInputErrors()
	acc, err := parseAccountErrors(accSrc)
	if err != nil {
		return accountErrors{}, []int{}, err
	}
	trans, err := parseTransactionsErrors(transSrc)
	if err != nil {
		return acc, trans, err
	}
	return acc, trans, err
}

// readInputErrors возвращает строку, которая описывает счет
// и срез строк, который описывает список транзакций.
// эту функцию можно не менять
func readInputErrors() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	accSrc := scanner.Text()
	var transSrc []string
	for scanner.Scan() {
		transSrc = append(transSrc, scanner.Text())
	}
	return accSrc, transSrc
}

// parseAccountErrors парсит счет из строки
// в формате balance/overdraft.
func parseAccountErrors(src string) (accountErrors, error) {
	var acc accountErrors
	parts := strings.Split(src, "/")
	balance, err := strconv.Atoi(parts[0])
	if err != nil {
		return acc, err
	}
	overdraft, err := strconv.Atoi(parts[1])
	if err != nil {
		return accountErrors{}, err
	}
	if overdraft < 0 {
		return acc, errors.New("expect overdraft >= 0")
		//panic("expect overdraft >= 0")
	}
	if balance < -overdraft {
		return acc, errors.New("balance cannot exceed overdraft")
		//panic("balance cannot exceed overdraft")
	}
	return accountErrors{balance, overdraft}, nil
}

// parseTransactionsErrors парсит список транзакций из строки
// в формате [t1 t2 t3 ... tn].
func parseTransactionsErrors(src []string) ([]int, error) {
	trans := make([]int, len(src))
	for idx, s := range src {
		t, err := strconv.Atoi(s)
		if err != nil {
			return trans, err
		}
		trans[idx] = t
	}
	return trans, nil
}
