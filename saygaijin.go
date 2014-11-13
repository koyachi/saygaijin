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
		cli.StringFlag{
			Name:  "voice, vo",
			Value: "",
			Usage: "Specify the voice to be used.",
		},
		cli.IntFlag{
			Name:  "rate, r",
			Value: 0,
			Usage: "Speech rate to be used, in words per minute.",
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

		say.OutputFile = c.String("output-file")
		say.Voice = c.String("voice")
		rate := c.Int("rate")
		if rate != 0 {
			say.Rate = rate
		}
		if len(str) > 0 {
			err = say.RunString(str)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err = say.Run(os.Stdin)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	app.Run(os.Args)
}
