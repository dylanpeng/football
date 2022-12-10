package reg

import "regexp"

func FindAllString(content, pattern string) (result []string, err error) {
	regItem, err := regexp.Compile(pattern)

	if err != nil {
		return
	}

	result = regItem.FindAllString(content, -1)

	return
}

func FindStringSubMatch(content, pattern string) (result []string, err error) {
	regItem, err := regexp.Compile(pattern)

	if err != nil {
		return
	}

	result = regItem.FindStringSubmatch(content)

	return
}
