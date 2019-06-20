package main

import (
	"log"
	"os"

	"github.com/gopherdojo/dojo5/kadai3-2//pkg/multithreadDownloader"

	"github.com/urfave/cli"
)

func defineApp(app *cli.App) {
	app.Name = "goDownloader"
	app.Usage = "file downloader made of Go."
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "pallarel, p",
			Value: 0,
			Usage: "pallarel number to download file",
		},
		cli.StringFlag{
			Name:  "url, u",
			Usage: "file url to download",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.String("url") == "" {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}
		dc := &multithreadDownloader.DownlodeClient{c.String("url"), 0, false, 0, false}
		err := dc.SetResponceHeader()
		if err != nil {
			log.Println(err)
		}

		dc.Download(c.Int("parallel"))
		return nil
	}
}

func main() {
	app := cli.NewApp()
	defineApp(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
