package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Проверяем, существует ли .env
	if _, err := os.Stat(".env"); err == nil {
		fmt.Println("Файл .env уже существует. Удалите его, если хотите пересоздать.")
		return
	}

	// Инициализируем данные по умолчанию
	defaults := map[string]string{
		"host":     "localhost",
		"port":     "5432",
		"sslmode":  "disable",
		"username": "debt_book_owner",
		"dbname":   "debt_book_db",
	}

	// Читаем ввод пользователя
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите параметры подключения к PostgreSQL (или нажмите Enter для использования значений по умолчанию):")

	input := make(map[string]string)
	for key, value := range defaults {
		fmt.Printf("%s (по умолчанию %s): ", key, value)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		if userInput == "" {
			input[key] = value
		} else {
			input[key] = userInput
		}
	}

	// Генерируем строку подключения
	password := input["password"]
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		input["username"], password,
		input["host"], input["port"],
		input["dbname"], input["sslmode"],
	)

	// Сохраняем в .env
	envContent := fmt.Sprintf("POSTGRES_URL=\"%s\"\n", connStr)
	err := os.WriteFile(".env", []byte(envContent), 0644)
	if err != nil {
		fmt.Printf("Ошибка при записи файла .env: %v\n", err)
		return
	}

	fmt.Println("✅ Файл .env успешно создан!")
}
