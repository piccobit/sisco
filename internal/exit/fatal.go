package exit

import (
	"fmt"
	"os"
)

func fatal(code int, arg string) {
	fmt.Fprintf(os.Stderr, arg)

	os.Exit(code)
}

func Fatal(code int, arg any) {
	fatal(code, fmt.Sprint(arg))
}

func Fatalln(code int, arg any) {
	fatal(code, fmt.Sprintln(arg))
}

func Fatalf(code int, format string, args ...any) {
	fatal(code, fmt.Sprintf(format, args))
}
