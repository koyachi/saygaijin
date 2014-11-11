// TODO: Makefile or cat ~/tmp/20141107/kana.txt | go run *.go
package main

import (
	"bufio"
	"fmt"
	"github.com/koyachi/go-romankana"
	"io"
	"log"
	"os/exec"
)

type SayCommand struct {
	cmd        *exec.Cmd
	stdin      io.WriteCloser
	chCmdStart chan bool
	chCmdEnd   chan bool
}

func NewSayCommand() (*SayCommand, error) {
	s := &SayCommand{}
	err := s.init()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *SayCommand) init() error {
	var err error
	s.chCmdStart = make(chan bool)
	s.chCmdEnd = make(chan bool)

	s.cmd = exec.Command("say")
	s.stdin, err = s.cmd.StdinPipe()
	if err != nil {
		return err
	}

	return nil
}

func (s *SayCommand) Run(r io.Reader) error {
	scanner := bufio.NewScanner(r)
	go s.startCommand()

	banged := false
	for scanner.Scan() {
		t := romankana.KanaRoman(scanner.Text() + " ")
		fmt.Println(t)
		_, err := io.WriteString(s.stdin, t)
		if err != nil {
			return err
		}
		if !banged {
			s.chCmdStart <- true
			banged = true
		}
	}
	s.stdin.Close()
	return nil
}

func (s *SayCommand) startCommand() {
	<-s.chCmdStart
	//fmt.Println("start say command")
	if err := s.cmd.Start(); err != nil {
		log.Fatal(err)
	}
	//fmt.Println("wait say command...")
	if err := s.cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	//fmt.Println("end say command")
	s.chCmdEnd <- true
}
