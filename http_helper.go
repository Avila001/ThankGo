package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// начало решения

// Handy предоставляет удобный интерфейс
// для выполнения HTTP-запросов
type Handy struct {
	url      string
	headers  map[string]string
	client   *http.Client
	response *http.Response
}

// NewHandy создает новый экземпляр Handy
func NewHandy() *Handy {
	handy := Handy{}
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	return &handy
}

// URL устанавливает URL, на который пойдет запрос
func (h *Handy) URL(uri string) *Handy {
	h.url = uri
	return h
}

// Client устанавливает HTTP-клиента
// вместо умолчательного http.DefaultClient
func (h *Handy) Client(client *http.Client) *Handy {
	h.client = client
	return h
}

// Header устанавливает значение заголовка
func (h *Handy) Header(key, value string) *Handy {
	req.Header.Set(key, val)
	return nil
}

// Param устанавливает значение URL-параметра
func (h *Handy) Param(key, value string) *Handy {
	return nil
}

// Form устанавливает данные, которые будут закодированы
// как application/x-www-form-urlencoded и отправлены в теле запроса
// с соответствующим content-type
func (h *Handy) Form(form map[string]string) *Handy {
	return nil
}

// JSON устанавливает данные, которые будут закодированы
// как application/json и отправлены в теле запроса
// с соответствующим content-type
func (h *Handy) JSON(v any) *Handy {
	return nil
}

// Get выполняет GET-запрос с настроенными ранее параметрами
func (h *Handy) Get() *HandyResponse {
	req, _ := http.NewRequest(http.MethodGet, h.url, nil)
	for key, value := range h.headers {
		req.Header.Add(key, value)
	}
	resp, err := h.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return &HandyResponse{}
}

// Post выполняет POST-запрос с настроенными ранее параметрами
func (h *Handy) Post() *HandyResponse {
	return &HandyResponse{}
}

// HandyResponse представляет ответ на HTTP-запрос
type HandyResponse struct {
	StatusCode int
	Body       io.ReadCloser
	// ...
}

// OK возвращает true, если во время выполнения запроса
// не произошло ошибок, а код HTTP-статуса ответа равен 200
func (r *HandyResponse) OK() bool {
	return false
}

// Bytes возвращает тело ответа как срез байт
func (r *HandyResponse) Bytes() []byte {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bodyBytes
}

// String возвращает тело ответа как строку
func (r *HandyResponse) String() string {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	return bodyString
}

// JSON декодирует тело ответа из JSON и сохраняет
// результат по адресу, на который указывает v
func (r *HandyResponse) JSON(v any) {
	// работает аналогично json.Unmarshal()
	// если при декодировании произошла ошибка,
	// она должна быть доступна через r.Err()
	body, _ := io.ReadAll(r.Body)
	var data map[string]any
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}
}

// Err возвращает ошибку, которая возникла при выполнении запроса
// или обработке ответа
func (r *HandyResponse) Err() error {
	return nil
}

// конец решения

func main() {
	{
		// примеры запросов

		// GET-запрос с параметрами
		NewHandy().URL("https://httpbingo.org/get").Param("id", "42").Get()

		// HTTP-заголовки
		NewHandy().
			URL("https://httpbingo.org/get").
			Header("Accept", "text/html").
			Header("Authorization", "Bearer 1234567890").
			Get()

		// POST формы
		params := map[string]string{
			"brand":    "lg",
			"category": "tv",
		}
		NewHandy().URL("https://httpbingo.org/post").Form(params).Post()

		// POST JSON-документа
		NewHandy().URL("https://httpbingo.org/post").JSON(params).Post()
	}

	{
		// пример обработки ответа

		// отправляем GET-запрос с параметрами
		resp := NewHandy().URL("https://httpbingo.org/get").Param("id", "42").Get()
		if !resp.OK() {
			panic(resp.String())
		}

		// декодируем ответ в JSON
		var data map[string]any
		resp.JSON(&data)

		fmt.Println(data["url"])
		// "https://httpbingo.org/get"
		fmt.Println(data["args"])
		// map[id:[42]]
	}
}
