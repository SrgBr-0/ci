// main.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Структура для хранения данных формы
type PageData struct {
	Result string
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Получаем значения чисел из формы
		r.ParseForm()
		num1, err1 := strconv.Atoi(r.FormValue("num1"))
		num2, err2 := strconv.Atoi(r.FormValue("num2"))

		// Проверяем на ошибки
		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid numbers", http.StatusBadRequest)
			return
		}

		// Считаем результат
		result := num1 + num2

		// Передаем результат в структуру для отображения на странице
		data := PageData{
			Result: fmt.Sprintf("%d + %d = %d", num1, num2, result),
		}

		// Используем HTML-шаблон для отображения страницы
		tmpl, err := template.New("calculator").Parse(`
			<html>
				<head><title>Plus Calculator</title></head>
				<body>
					<h1>Simple Calculator</h1>
					<form method="POST">
						<label for="num1">Number 1:</label>
						<input type="text" name="num1" id="num1" required>
						<br>
						<label for="num2">Number 2:</label>
						<input type="text" name="num2" id="num2" required>
						<br>
						<button type="submit">Add</button>
					</form>
					<h2>Result: {{.Result}}</h2>
				</body>
			</html>
		`)
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		// Отправляем данные на страницу
		tmpl.Execute(w, data)
	} else {
		// Для GET-запроса просто отрисовываем форму
		tmpl, err := template.New("calculator").Parse(`
			<html>
				<head><title>Simple Calculator</title></head>
				<body>
					<h1>Plus Calculator</h1>
					<form method="POST">
						<label for="num1">Number 1:</label>
						<input type="text" name="num1" id="num1" required>
						<br>
						<label for="num2">Number 2:</label>
						<input type="text" name="num2" id="num2" required>
						<br>
						<button type="submit">Add</button>
					</form>
					<h2>Result: </h2>
				</body>
			</html>
		`)
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}
		// Отправляем пустую страницу с формой
		tmpl.Execute(w, nil)
	}
}

func main() {
	http.HandleFunc("/", addHandler)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
