package main

import (
	"errors"
	"fmt"
)

func readConfig() error {
	return errors.New("read config error")
}

func startServer() error {
	if err := readConfig(); err != nil {
		return fmt.Errorf("%w, start server err", err)
	}

	return nil
}

func run() error {
	if err := startServer(); err != nil {
		return fmt.Errorf("%w, run server error", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)

		for err != nil {
			fmt.Println(err)
			err = errors.Unwrap(err)
		}
	}
}
