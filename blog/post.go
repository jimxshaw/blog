package blog

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"io"
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
	Tags        map[string]struct{}
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

func newPost(file io.Reader) (Post, error) {
	scanner := bufio.NewScanner(file)

	readLines := func(parser string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), parser)
	}

	title := readLines(titleParser)
	description := readLines(descriptionParser)
	date := readLines(dateParser)
	tagsArray := strings.Split(readLines(tagsParser), ", ")
	url := CreateUrl(title)
	body := template.HTML(readBody(scanner))

	const dateFormat = "2000-Jan-01"

	parsedDate, err := time.Parse(dateFormat, date)
	if err != nil {
		return Post{}, nil
	}

	tagsMap := make(map[string]struct{})
	for _, tag := range tagsArray {
		tagsMap[tag] = struct{}{}
	}

	return Post{
		Title:       title,
		Description: description,
		Date:        parsedDate,
		Tags:        tagsMap,
		Url:         url,
		Body:        body,
	}, nil
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
