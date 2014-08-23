package main

import (
	"fmt"
	"github.com/koyachi/go-romankana"
	"log"
	"os"
	"os/exec"
)

func main() {
	s := romankana.KanaRoman(os.Args[1])
	fmt.Println(s)
	cmd := exec.Command("say", s)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}
