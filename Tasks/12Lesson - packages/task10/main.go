package main

import (
	_ "task10/init"
	"task10/printcheck"
)

func main() {
	printcheck.PrintCheck("Nekdil", true)
	printcheck.PrintCheck("fatal error", false)
}
