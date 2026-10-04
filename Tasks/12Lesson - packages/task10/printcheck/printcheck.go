package printcheck

import "github.com/fatih/color"

func PrintCheck(name string, passed bool) {
	if passed {
		color.Green("OK: %s\n", name)
	} else {
		color.Red("FAIL: %s\n", name)
	}
}
