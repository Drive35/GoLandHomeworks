package main

import (
	"fmt"
	"github.com/urfave/cli"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	pathDir = ""
	url     = ""
)

func mainAction3(c *cli.Context) error {
	if len(os.Args) != 3 {
		fmt.Println("usage: download url filename")
		os.Exit(1)
	}
	url := os.Args[1]
	filename := os.Args[2]

	err := downloadImage(url, filename)
	if err != nil {
		panic(err)
	}

}
func downloadImage() error {
	out, err := os.Create(pathDir)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err + io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "App for downloading image from URL"
	app.Description = "Just set addresses of directory and URL"
	app.Usage = "Download image from URL to the specified directory"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path,p",
			Value:       "./Documents",
			Destination: &pathDir,
		},
		cli.StringFlag{
			Name:        "url,i",
			Value:       "https://fainaidea.com/wp-content/uploads/2018/01/1480439569193232018.jpg",
			Destination: &url,
		},
	}
	app.Action = mainAction3
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
