package json

import (
	"regexp"
	"strings"
)

var arrayIndexRegex = regexp.MustCompile(`(?m)\[[0-9]+]`)

func SplitKey(key string) []string {
	results := make([]string, 0)
	for _, splitted := range strings.Split(key, ".") {
		arraykeys := make([]string, 0)
		arraykeys = append(arraykeys, arrayIndexRegex.FindAllString(splitted, -1)...)
		if idx := strings.LastIndex(splitted, strings.Join(arraykeys, "")); idx >= 0 {
			splitted = splitted[:idx]
		}
		results = append(results, splitted)
		results = append(results, arraykeys...)
	}
	return results
}
