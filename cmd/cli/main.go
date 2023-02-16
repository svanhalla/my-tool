package main

import (
	"fmt"
	"github.com/svanhalla/my-tool/pkg"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	if err := Run(); err != nil {
		println("Error: ", err.Error())
	}
}

func Run() error {
	app := cli.App{
		Commands: []*cli.Command{
			{
				Name:  "slug",
				Usage: "slugs argument",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "output",
						Usage: "full path to file where to put the slugged result, create file or append to existing file",
					},
				},
				Action: SlugCommand,
			},
			{
				Name:  "random",
				Usage: "writes a random string to output",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "number",
						Usage: "string of digits 0-9",
					},
					&cli.IntFlag{
						Name:  "length",
						Usage: "the length of the random string",
						Value: 20,
					},
				},
				Action: RandomCommand,
			},
		},
	}
	return app.Run(os.Args)
}

func SlugCommand(cliCtx *cli.Context) error {
	output := os.Stdout
	defer func() {
		err := output.Close()
		if err != nil {
			fmt.Printf("error; %s\n", err.Error())
		}
	}()

	if o := cliCtx.String("output"); len(o) > 0 {
		f, err := pkg.CreateOrOpenFile(o)
		if err != nil {
			return err
		}
		output = f
	}

	for _, s := range cliCtx.Args().Slice() {
		sluggedStr, err := pkg.Slugify(s)
		if err != nil {
			return err
		}
		_, _ = fmt.Fprintf(output, "%s\n", sluggedStr)
	}

	return nil
}

func RandomCommand(cliCtx *cli.Context) error {
	number := cliCtx.Bool("number")
	numberOfCharacters := cliCtx.Int("length")

	var wantedString string
	if number {
		wantedString = pkg.RandomNumber(numberOfCharacters)
	} else {
		wantedString = pkg.RandomString(numberOfCharacters)
	}

	_, _ = fmt.Fprintf(os.Stdout, "%s\n", wantedString)
	return nil
}
