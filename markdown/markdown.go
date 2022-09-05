package markdown

import (
	"fmt"
	"html/template"
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
	if strings.HasPrefix(str, "#") {
		splitStr := strings.Split(str, " ")
		prefix := splitStr[0]
		prefixCount := len(prefix)
		str = fmt.Sprintf("<h%d>%s</h%d>", prefixCount, strings.Join(splitStr[1:], " "), prefixCount)
	} else if strings.HasPrefix(str, "```") && !openCodeBlock {
		if strings.HasSuffix(str, "```") {
			str = strings.ReplaceAll(str, "```", "")
			str = fmt.Sprintf("<code>%s</code>", str)
		} else {
			str = strings.ReplaceAll(str, "```", "")
			str = fmt.Sprintf("<code>%s<br>", str)
			openCodeBlock = true
		}
	} else if strings.HasSuffix(str, "```") {
		str = strings.ReplaceAll(str, "```", "</code>")
		openCodeBlock = false
	} else {
		str = fmt.Sprintf("<p>%s</p>", str)
	}

	return str, openCodeBlock
}
