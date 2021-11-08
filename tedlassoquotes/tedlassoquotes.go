package tedlassoquotes

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func Run(args []string) int {
	var app tedlassoApp
	err := app.fromArgs(args)
	if err != nil {
		return 2
	}
	if err = app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v\n", err)
		return 1
	}
	return 0
}

type tedlassoApp struct {
	client   http.Client
	tag      string
	dumpJSON bool
}

func (app *tedlassoApp) fromArgs(args []string) error {
	// Shallow copy of default client
	app.client = *http.DefaultClient
	fl := flag.NewFlagSet("tedlasso-grab", flag.ContinueOnError)
	fl.StringVar(
		&app.tag, "c", "", "Character to fetch (default to random: valid choices: Ted|Roy|Beard)",
	)
	app.client.Timeout = 2 * time.Second
	outputType := fl.String(
		"o", "text", "Print output in format: text/json",
	)
	if err := fl.Parse(args); err != nil {
		return err
	}
	if *outputType != "text" && *outputType != "json" {
		fmt.Fprintf(os.Stderr, "got bad output type: %q\n", *outputType)
		fl.Usage()
		return flag.ErrHelp
	}
	app.dumpJSON = *outputType == "json"
	return nil
}

func (app *tedlassoApp) run() error {
	u := buildURL(app.tag)

	var resp APIResponse
	if err := app.fetch(u, &resp); err != nil {
		return err
	}
	if app.dumpJSON {
		return printJSON(resp)
	}

	return printOut(resp)
}

// fetch the Ted Lasso quote from the API
func (app *tedlassoApp) fetch(url string, data interface{}) error {
	resp, err := app.client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(data)
}

// Quote is the JSON output of this app
type Quote struct {
	Quote      string `json:"quote"`
	Author     string `json:"author"`
	ProfileImg string `json:"profile_img"`
	Tag        string `json:"tag"`
}

// printJSON to stdout with indentations
func printJSON(ar APIResponse) error {
	o := Quote{
		Quote:      ar.Quote,
		Author:     ar.Author,
		Tag:        ar.Tag,
		ProfileImg: ar.ProfileImg,
	}
	b, err := json.MarshalIndent(&o, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

// printOut to stdout with formatting
func printOut(ar APIResponse) error {
	_, err := fmt.Printf(
		"\n%s\n— %s\n\n",
		ar.Quote,
		ar.Author,
	)
	return err
}
