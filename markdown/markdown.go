package markdown

import (
	"fmt"
	"html/template"
	"log"
	"strconv"
	"strings"
)

func MdToHtml(markdown string) (result template.HTML) {
	var resultLines []string
	var resultLine string
	openCodeBlock := false
	lines := strings.Split(markdown, "\n")
	for _, line := range lines {
		resultLine, openCodeBlock = parse(line, openCodeBlock)
		resultLines = append(resultLines, resultLine)
	}
	return template.HTML(strings.Join(resultLines, ""))
}

func parse(str string, openCodeBlock bool) (string, bool) {
	str = strings.TrimSpace(str)

	if openCodeBlock {
		if strings.HasSuffix(str, "```") {
			str = strings.ReplaceAll(str, "```", "</code>")
			openCodeBlock = false
		} else {
			str = fmt.Sprintf("<p>%s</p>", str)
		}
	} else {
		if strings.HasPrefix(str, "```") {
			if strings.HasSuffix(str, "```") {
				str = strings.ReplaceAll(str, "```", "")
				str = fmt.Sprintf("<code>%s</code>", str)
			} else {
				str = strings.ReplaceAll(str, "```", "")
				str = fmt.Sprintf("<code>%s<br>", str)
				openCodeBlock = true
			}
		} else if strings.HasPrefix(str, "#") {
			splitStr := strings.Split(str, " ")
			prefix := splitStr[0]
			prefixCount := len(prefix)
			str = fmt.Sprintf("<h%d>%s</h%d>", prefixCount, strings.Join(splitStr[1:], " "), prefixCount)
		} else if strings.HasPrefix(str, "``a") {
			splitStr := strings.Split(str, " ")
			link := splitStr[1]
			text := strings.Join(splitStr[2:], " ")
			str = fmt.Sprintf("<a href='%s' target='_blank' rel='noopener noreferrer'>%s</a>", link, text)
		} else if strings.HasPrefix(str, "``img") {
			splitStr := strings.Split(str, " ")
			src := splitStr[1]
			options := splitStr[2:]
			w, h := parseWidthHeight(options)
			str = fmt.Sprintf("<img class='content-img' src='%s'", src)
			if w > 0 {
				str = fmt.Sprintf("%s width='%d'", str, w)
			}
			if h > 0 {
				str = fmt.Sprintf("%s height='%d'", str, h)
			}
			str = str + ">"
		} else {
			str = fmt.Sprintf("<p>%s</p>", str)
		}
	}

	return str, openCodeBlock
}

func parseWidthHeight(options []string) (w int, h int) {
	for _, option := range options {
		parts := strings.Split(option, "=")
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		if strings.Contains(key, "w") {
			width, err := strconv.Atoi(value)
			if err != nil {
				log.Printf("width parse error : %s \n", err.Error())
			} else {
				w = width
			}
		} else if strings.Contains(key, "h") {
			height, err := strconv.Atoi(value)
			if err != nil {
				log.Printf("height parse error : %s \n", err.Error())
			} else {
				h = height
			}
		}
	}
	return
}
