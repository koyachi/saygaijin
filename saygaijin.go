package main

import (
	"io"
	"log"
	"os"
)

func main() {
	say, err := NewSayCommand()
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) == 1 {
		err = say.Run(r)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = say.RunString(input)
		if err != nil {
			log.Fatal(err)
		}
	}
}
