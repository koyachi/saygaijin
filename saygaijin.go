package main

import (
	"fmt"
	"github.com/koyachi/go-romankana"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		sayIo(os.Stdin)
	} else {
		sayString(os.Args[1])
	}
}

func sayString(input string) {
	s := romankana.KanaRoman(input)
	fmt.Println(s)
	cmd := exec.Command("say", s)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}

func sayIo(r io.Reader) {
	say, err := NewSayCommand()
	if err != nil {
		log.Fatal(err)
	}
	err = say.Run(r)
	if err != nil {
		log.Fatal(err)
	}
}
