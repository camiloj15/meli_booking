package utils

import (
	"log"
	"os"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}
