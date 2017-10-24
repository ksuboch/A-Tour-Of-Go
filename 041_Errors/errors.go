package main

import (
	"fmt"
	"time"
)

// Тип error - встроенный интерфейс, похожий на fmt.Stringer:
// type error interface {
//     Error() string
// }
// пакет fmt ищет интерфейс error interface для вывода сообщений

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"It didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
