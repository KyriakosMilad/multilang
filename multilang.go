package multilang

type (
	LanguageDictionary map[string]string
	Dictionary         map[string]*LanguageDictionary
)

var (
	dict = map[string]*LanguageDictionary{}
)

func Set(lang string, key string, value string) {
	if l, ok := dict[lang]; ok {
		(*l)[key] = value
		return
	}

	l := dict[lang]
	l = &LanguageDictionary{}
	(*l)[key] = value
}

func Get(lang string, key string) (string, bool) {
	if _, ok := dict[lang]; !ok {
		return "", false
	}

	l := dict[lang]
	val, ok := (*l)[key]
	return val, ok
}

func SetLangDict(lang string, landDict *LanguageDictionary) {
	dict[lang] = landDict
}

func GetLangDict(lang string, key string) (*LanguageDictionary, bool) {
	l, ok := dict[lang]
	return l, ok
}
