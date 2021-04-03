package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "up",
		Usage: "FILE...",
		Action: func(c *cli.Context) error {
			fileName := c.Args().Get(0)
			fmt.Println("Uploading", fileName)

			file, err := os.Open(fileName)
			payload := io.MultiReader(file)
			req, err := http.NewRequest("POST", "https://filebin.net", payload)
			check(err)
			resp, err := http.DefaultClient.Do(req)
			check(err)
			defer resp.Body.Close()
			fmt.Println(resp)

			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
