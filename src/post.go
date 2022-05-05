package blog

import (
	"html/template"
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
