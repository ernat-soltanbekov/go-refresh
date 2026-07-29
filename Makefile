.PHONY: all build run test clean

# По умолчанию собираем проект
all: build

# Компиляция бинарного файла
build:
	go build -o go-refresh cmd/go-refresh/main.go

# Запуск (пример с дефолтными файлами, если они нужны)
run:
	go run cmd/go-refresh/main.go sample.txt result.txt

# Запуск всех тестов в проекте
test:
	go test ./... -v

# Очистка скомпилированных файлов
clean:
	rm -f go-refresh result.txt