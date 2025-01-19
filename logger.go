package logs

import (
	"log"
	"os"
)

func InitLogger() *log.Logger {
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return log.New(file, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}
