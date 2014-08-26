package main

import (
	"bufio"
	"fmt"
	"github.com/koyachi/go-romankana"
	"log"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) == 1 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			say(scanner.Text())
		}
	} else {
		say(os.Args[1])
	}
}

func say(input string) {
	s := romankana.KanaRoman(input)
	fmt.Println(s)
	cmd := exec.Command("say", s)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
}
