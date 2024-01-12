package logging

import (
	"fmt"
	"time"
)

var (
	portNumber string = ":3000"
	debug      bool   = false
)

func Debug(b bool) {
	debug = b
}
func Logger(statement string) {
	if !debug {
		return
	}
	fmt.Printf("%s %s\n", time.Now().Format(time.RFC3339), statement)
}
