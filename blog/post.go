package blog

import (
	"html/template"
	"log"
	"regexp"
	"strings"
	"time"
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
