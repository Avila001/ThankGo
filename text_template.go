package main

/*
import (
"bytes"
"fmt"
"testing"
"text/template"
)

// начало решения

var templateText = `Алиса, добрый день! Ваш баланс - {{if gt(.,100}} {{.}}₽. Все в порядке. {{ - else if . > 0 && . < 100 - }} {{.}}₽. Пора пополнить. {{ - else if . == 0 - }} {{.}}₽. Доступ заблокирован.`

// конец решения

type User struct {
	Name    string
	Balance int
}

// renderToString рендерит данные по шаблону в строку
func renderToString(tpl *template.Template, data any) string {
	var buf bytes.Buffer
	tpl.Execute(&buf, data)
	return buf.String()
}

func main() {
	tpl := template.New("message")
	tpl = template.Must(tpl.Parse(templateText))

	user := User{"Алиса", 500}
	got := renderToString(tpl, user)

	const want = "Алиса, добрый день! Ваш баланс - 500₽. Все в порядке."
	if got != want {
		fmt.Printf("%v: got '%v'", user, got)
	}
}

func Test(t *testing.T) {
	tpl := template.New("message")
	tpl = template.Must(tpl.Parse(templateText))

	user := User{"Алиса", 500}
	got := renderToString(tpl, user)

	const want = "Алиса, добрый день! Ваш баланс - 500₽. Все в порядке."
	if got != want {
		t.Errorf("%v: got '%v'", user, got)
	}
}
*/
