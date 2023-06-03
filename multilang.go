package multilang

import (
	"fmt"
	"strings"
)

type (
	LanguageDictionary map[string]string
	Variables          map[string]interface{}
)

type Dictionary struct {
	dict map[string]*LanguageDictionary
}

func (dict *Dictionary) Set(lang string, key string, value string) {
	if l, ok := dict.dict[lang]; ok {
		(*l)[key] = value
		return
	}

	l := &LanguageDictionary{}
	(*l)[key] = value
	dict.dict[lang] = l
}

func (dict *Dictionary) Get(lang string, key string, variables *Variables) (string, bool) {
	if _, ok := dict.dict[lang]; !ok {
		return "", false
	}

	l := dict.dict[lang]
	val, ok := (*l)[key]
	if ok {
		val = dict.replaceVariables(val, variables)
	}
	return val, ok
}

func (dict *Dictionary) SetLangDict(lang string, landDict *LanguageDictionary) {
	dict.dict[lang] = landDict
}

func (dict *Dictionary) GetLangDict(lang string) (*LanguageDictionary, bool) {
	l, ok := dict.dict[lang]
	return l, ok
}

func (dict *Dictionary) SetDict(d map[string]*LanguageDictionary) {
	dict.dict = d
}

func (dict *Dictionary) GetDict() map[string]*LanguageDictionary {
	return dict.dict
}

func (dict *Dictionary) replaceVariables(str string, variables *Variables) string {
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
			replacementString := fmt.Sprint(replacement)
			builder.WriteString(replacementString)
		} else {
			builder.WriteString(placeholder)
		}

		start = placeholderEnd
	}

	return builder.String()
}
