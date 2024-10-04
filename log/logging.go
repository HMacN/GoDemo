package log

import (
	"fmt"
	"log"
)

func Write(format string, a ...any) {
	log.Println(fmt.Sprintf(format, a...))
}
