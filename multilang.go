package multilang

import (
	"strings"
)

type (
	LanguageDictionary map[string]string
	Dictionary         map[string]*LanguageDictionary
	Variables          map[string]string
)

var (
	dict = &Dictionary{}
)

func Set(lang string, key string, value string) {
	if l, ok := (*dict)[lang]; ok {
		(*l)[key] = value
		return
	}

	l := (*dict)[lang]
	l = &LanguageDictionary{}
	(*l)[key] = value
}

func Get(lang string, key string, variables *Variables) (string, bool) {
	if _, ok := (*dict)[lang]; !ok {
		return "", false
	}

	l := (*dict)[lang]
	val, ok := (*l)[key]
	if ok {
		val = replaceVariables(val, variables)
	}
	return val, ok
}

func SetLangDict(lang string, landDict *LanguageDictionary) {
	(*dict)[lang] = landDict
}

func GetLangDict(lang string, key string) (*LanguageDictionary, bool) {
	l, ok := (*dict)[lang]
	return l, ok
}

func SetDict(d *Dictionary) {
	dict = d
}

func GetDict() *Dictionary {
	return dict
}

func replaceVariables(str string, variables *Variables) string {
	var builder strings.Builder
	start := 0

	for {
		placeholderStart := strings.IndexByte(str[start:], '$')
		if placeholderStart == -1 {
			builder.WriteString(str[start:])
			break
		}

		builder.WriteString(str[start : start+placeholderStart])

		placeholderEnd := strings.IndexByte(str[start+placeholderStart:], ' ')
		if placeholderEnd == -1 {
			placeholderEnd = len(str)
		} else {
			placeholderEnd += start + placeholderStart
		}

		placeholder := str[start+placeholderStart : placeholderEnd]

		if replacement, exists := (*variables)[placeholder]; exists {
			builder.WriteString(replacement)
		} else {
			builder.WriteString(placeholder)
		}

		start = placeholderEnd
	}

	return builder.String()
}
