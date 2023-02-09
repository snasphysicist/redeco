package redeco

import "regexp"

// groups runs a regex match and returns
// the capture group matches as key value pairs
func groups(regex string, s string) map[string]string {
	compRegEx := regexp.MustCompile(regex)
	match := compRegEx.FindStringSubmatch(s)
	paramsMap := make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
