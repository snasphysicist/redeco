package redeco

import (
	"log"
	"regexp"
)

// groups runs a regex match and returns
// the capture group matches as key value pairs
func groups(regex string, s string) map[string]string {
	re := regexp.MustCompile(regex)
	match := re.FindStringSubmatch(s)
	return pairGroupNamesWithContent(re.SubexpNames(), match)
}

// pairGroupNamesWithContent pairs the names of the capture groups
// with the content that was found for each group
func pairGroupNamesWithContent(groupNames []string, matches []string) map[string]string {
	if len(groupNames) == 0 {
		return make(map[string]string)
	}
	if len(matches) == 0 {
		return make(map[string]string)
	}
	pm := make(map[string]string)
	captureGroupNames := groupNames[1:]
	matches = matches[1:]
	for i, name := range captureGroupNames {
		log.Printf("%d %s %s", i, name, matches[i])
		if i <= len(matches) {
			pm[name] = matches[i]
		}
	}
	return pm
}
