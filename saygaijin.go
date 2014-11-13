package main

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "saygaijin"
	app.Usage = "say japanese with gaijin voice."
	app.Flags = []cli.Flag{
		/*
			cli.StringFlag{
				Name:  "input-file, f",
				Value: "-",
				Usage: "Specify a file to be spoken.",
			},
		*/
		cli.StringFlag{
			Name:  "output-file, o",
			Value: "",
			Usage: "Specify the path for an audio file to be written.",
		},
	}
	app.Action = func(c *cli.Context) {
		str := ""
		for _, s := range c.Args() {
			str += " " + s
		}

		say, err := NewSayCommand()
		if err != nil {
			log.Fatal(err)
		}

		if len(str) > 0 {
			err = say.RunString(str)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			say.OutputFile = c.String("output-file")
			err = say.Run(os.Stdin)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	app.Run(os.Args)
}
