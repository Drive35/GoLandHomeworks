package main

import (
	"github.com/kkdai/youtube"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	pathDir = ""
	url     = ""
	file    = ""
)

func downloadFromYoutube(url string) error {
	yt := youtube.NewYoutube(false, true)
	err := yt.DecodeURL(url)
	if err != nil {
		return err
	}
	fileName := yt.Author + "_" + yt.Title + ".mp4"
	err = yt.StartDownload(pathDir, fileName, "high", 0)
	if err != nil {
		return err
	}
	return nil
}

func mainAction4() error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	dataStr := string(data)
	dataStrArr := strings.Split(dataStr, "\r\n")
	for _, v := range dataStrArr {
		err = downloadFromYoutube(v)
		if err != nil {
			return err
		}
	}
	return nil
}

func downloadAction(c *cli.Context) error {
	if url != "" {
		err := downloadFromYoutube(url)
		if err != nil {
			return err
		}
	} else if file != "" {
		return mainAction4()
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "App for downloading video"
	app.Description = "Just set the url and video files will bee downloaded"
	app.Usage = "Set parameter url to link"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "path,p",
			Value:       "./Documents",
			Destination: &pathDir,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "download",
			Aliases: []string{"d"},
			Usage:   "for download video",
			Action:  downloadAction,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "file,f",
					Usage:       "file string",
					Destination: &file,
					Value:       "",
				},
				cli.StringFlag{
					Name:        "url,u",
					Usage:       "url string",
					Destination: &url,
					Value:       "",
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
