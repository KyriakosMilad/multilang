package multilang

import "strings"

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
	for placeholder, value := range *variables {
		str = strings.Replace(str, placeholder, value, -1)
	}

	return str
}
