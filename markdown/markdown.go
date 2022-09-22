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
	trimStr := strings.TrimSpace(str)

	if openCodeBlock {
		if strings.HasSuffix(trimStr, "```") {
			str = strings.ReplaceAll(str, "```", "</pre>")
			openCodeBlock = false
		} else {
			str = fmt.Sprintf("<p>%s</p>", str)
		}
	} else {
		if strings.HasPrefix(trimStr, "```") {
			if strings.HasSuffix(trimStr, "```") && len(trimStr) > len("```") {
				str = strings.ReplaceAll(str, "```", "")
				str = fmt.Sprintf("<pre>%s</pre>", str)
			} else {
				str = strings.ReplaceAll(str, "```", "")
				str = fmt.Sprintf("<pre>%s<br>", str)
				openCodeBlock = true
			}
		} else if strings.HasPrefix(trimStr, "#") {
			splitStr := strings.Split(trimStr, " ")
			prefix := splitStr[0]
			prefixCount := len(prefix)
			str = fmt.Sprintf("<h%d>%s</h%d>", prefixCount, strings.Join(splitStr[1:], " "), prefixCount)
		} else if strings.HasPrefix(trimStr, "``a") {
			splitStr := strings.Split(str, " ")
			link := splitStr[1]
			text := strings.Join(splitStr[2:], " ")
			str = fmt.Sprintf("<a href='%s' target='_blank' rel='noopener noreferrer'>%s</a>", link, text)
		} else if strings.HasPrefix(trimStr, "``img") {
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
			if len(str) == 0 {
				str = "<br>"
			}
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
