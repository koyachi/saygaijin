package main

import (
	"bufio"
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
	bang := make(chan bool)
	sayStdin, sayCmd := sayStdin()
	scanner := bufio.NewScanner(r)

	go func() {
		<-bang
		fmt.Println("start say command")
		if err := sayCmd.Start(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("wait say command")
		if err := sayCmd.Wait(); err != nil {
			log.Fatal(err)
		}
		fmt.Println("end say command")
	}()

	banged := false
	func() {
		for scanner.Scan() {
			t := romankana.KanaRoman(scanner.Text() + " ")
			fmt.Println(t)
			_, err := io.WriteString(sayStdin, t)
			if err != nil {
				log.Fatal()
			}
			if !banged {
				bang <- true
				banged = true
			}
		}
	}()
}

// TODO: rename function
func sayStdin() (io.WriteCloser, *exec.Cmd) {
	cmd := exec.Command("say")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	return stdin, cmd
}
