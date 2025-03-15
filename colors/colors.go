package colors

import "fmt"

const (
	RESET_COLOR  = "\x1b[0m"
	RED_COLOR    = "\x1b[31m"
	GREEN_COLOR  = "\x1b[32m"
	YELLOW_COLOR = "\x1b[33m"
	BLUE_COLOR   = "\x1b[34m"
	BOLD         = "\x1b[1m"
	UNDERLINE    = "\x1b[4m"
	CLEAR_SCREEN = "\x1b[2J\x1b[H"
)

func PrintLine(c rune, length int) {
	for i := 0; i < length; i++ {
		fmt.Print(string(c))
	}
	fmt.Println()
}

func PrintColored(text string, color string) {
	fmt.Printf("%v%v%v", color, text, RESET_COLOR)
}
