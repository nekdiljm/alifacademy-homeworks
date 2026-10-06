package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("id не найдено")

func findItem(id int) error {
	if id != 10 {
		return ErrNotFound
	}

	return nil
}

func loadItem(id int) error {
	if err := findItem(id); err != nil {
		return fmt.Errorf("не удалось загрузить item %d: %w\n", id, err)
	}

	return nil
}

func main() {
	if err := loadItem(11); err != nil {
		errNotFound := errors.Is(err, ErrNotFound)
		fmt.Println("Ошибка", err, errNotFound)
	}
}
