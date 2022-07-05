package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/flosch/pongo2/v6"
	"github.com/go-resty/resty/v2"
	"github.com/linbuxiao/rss-push-helper/feedly"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

type output struct {
	Title string
	Href  string
}

type formatType int

const (
	formatError formatType = iota
	formatHTML
	formatJSON
)

func main() {
	app := &cli.App{
		Name: "feedly-helper",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "token",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
			},
		},
		Action: func(ctx *cli.Context) error {
			format := ctx.String("format")
			if format == "" {
				format = "html"
			}
			f := parseFormat(format)
			if f == formatError {
				return errors.New("error format")
			}
			token := ctx.String("token")
			res, err := GetFeeds(token, f)
			if err != nil {
				return err
			}
			fmt.Println(res)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func parseFormat(text string) formatType {
	switch text {
	case "html":
		return formatHTML
	case "json":
		return formatJSON
	default:
		return formatError
	}
}

func GetFeeds(token string, format formatType) (string, error) {
	client := resty.New()
	client.SetBaseURL("https://cloud.feedly.com")
	client.SetAuthToken(token)
	var collections []*feedly.Collection
	_, err := client.R().SetResult(&collections).Get("/v3/collections")
	if err != nil {
		return "", err
	}
	var o []*output
	for _, v := range collections {
		var data feedly.GetStreamIDsReponse
		_, err := client.R().SetResult(&data).Get("/v3/streams/ids?streamId=" + v.Id)
		if err != nil {
			return "", err
		}
		var entryData []*feedly.GetEntriesResponse
		_, err = client.R().SetBody(data.Ids).SetResult(&entryData).Post("/v3/entries/.mget")

		for _, x := range entryData {
			yesterday := time.Now().AddDate(0, 0, -1)
			updated := time.UnixMilli(x.Updated)
			if updated.Before(yesterday) {
				continue
			}
			for _, y := range x.Alternate {
				o = append(o, &output{
					Title: x.Title,
					Href:  y.Href,
				})
			}
		}
	}
	// render
	switch format {
	case formatJSON:
		b, err := json.Marshal(o)
		return string(b), err
	case formatHTML:
		t := pongo2.Must(pongo2.FromString("{% for v in arr %}<li><a href=\"{{v.Href}}\">{{v.Title}}</a></li>\n{% endfor %}"))
		return t.Execute(pongo2.Context{
			"arr": o,
		})
	default:
		return "", nil
	}
}
