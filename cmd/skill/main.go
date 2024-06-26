// пакеты исполняемых приложений должны называться main
package main

import (
	"log"
	"net/http"
)

// функция main вызывается автоматически при запуске приложения
func main() {
	ParseFlags()

	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

// функция run будет полезна при инициализации зависимостей сервера перед запуском
func run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/alice/", webhook)
	return http.ListenAndServe(FlagRunAddr, mux)
}

// функция webhook — обработчик HTTP-запроса
func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// разрешаем только Get-запросы
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// установим правильный заголовок для типа данных
	w.Header().Set("Content-Type", "application/json")
	// пока установим ответ-заглушку, без проверки ошибок
	w.Write([]byte(`
      {
        "response": {
          "text": "Извините, я пока ничего не умею"
        },
        "version": "1.0"
      }
    `))
}
