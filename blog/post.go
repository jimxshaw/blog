package blog

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/Depado/bfchroma"
	blackfriday "github.com/russross/blackfriday/v2"
)

type Post struct {
	Title       string
	Description string
	Date        time.Time
	Tags        []string
	Url         string
	Body        template.HTML
}

const (
	titleParser       = "Title: "
	descriptionParser = "Description: "
	dateParser        = "Date: "
	tagsParser        = "Tags: "
)

func readBody(scanner *bufio.Scanner) []byte {
	scanner.Scan()
	buffer := bytes.Buffer{}

	for scanner.Scan() {
		fmt.Fprintln(&buffer, scanner.Text())
	}

	newBuffer := buffer.Bytes()

	// Blackfriday is a Markdown processor that can output HTML.
	content := bytes.TrimSpace(blackfriday.Run(newBuffer, blackfriday.WithRenderer(bfchroma.NewRenderer(
		bfchroma.Style("vs"),
	))))
	return content
}

func CreateUrl(title string) string {
	// Remove all spaces within the title and replace it with a dash.
	title = strings.ToLower(strings.Replace(title, " ", "-", -1))

	// Regex will mark anything NOT listed here.
	regex, err := regexp.Compile("[^a-z0-9\\-]+")
	if err != nil {
		log.Fatal(err)
	}

	url := regex.ReplaceAllString(title, "")
	url = strings.Replace(url, "---", "-", -1)

	return "/blog/" + url
}
