package inputprompt

import (
	"bufio"
	"fmt"
	"os"
)

func Wait(message string) {
	fmt.Println(message)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

type Callback[T any] func() (T, error)

func Command[T any](message string, fn Callback[T]) (T, error) {
	fmt.Print(message)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	return fn()
}
