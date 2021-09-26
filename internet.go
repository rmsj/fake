package faker

import (
	"math/rand"
	"strings"

	"github.com/rmsj/faker/random"
)

// InternetProvider must be implemented by types that wants to provide data source
// for emails, URLs, etc.
type InternetProvider interface {
	FreeEmailDomains() []string
	SafeEmailDomains() []string
	Tld() []string
	UserNameFormats() []string
	EmailFormats() []string
	UrlFormats() []string
	Words() []string
}

// internet provides random fake email, urls, etc.
type internet struct {
	Provider InternetProvider
	Person   person
}

// NewInternet constructs a internet type value and returns it
func NewInternet(provider InternetProvider, person person) internet {
	return internet{
		Provider: provider,
		Person:   person,
	}
}

//Username returns a random username
func (i internet) Username() string {
	userNameFormat := random.FromSliceOfString(i.Provider.UserNameFormats())

	var userName string
	if strings.Contains(userNameFormat, "firstName") {
		userName = strings.ReplaceAll(userNameFormat, "{{firstName}}", i.Person.FirstName())
	}
	if strings.Contains(userName, "lastName") {
		userName = strings.ReplaceAll(userName, "{{lastName}}", i.Person.LastName())
	}
	return userName
}

//Email returns a random email address
func (i internet) Email() string {
	emailFormat := random.FromSliceOfString(i.Provider.EmailFormats())

	var email string
	if strings.Contains(emailFormat, "userName") {
		email = strings.ReplaceAll(emailFormat, "{{userName}}", i.Username())
	}
	if strings.Contains(email, "domainName") {
		email = strings.ReplaceAll(email, "{{domainName}}", i.DomainName())
	}
	if strings.Contains(email, "FreeEmailDomains") {
		email = strings.ReplaceAll(email, "{{FreeEmailDomains}}", random.FromSliceOfString(i.Provider.FreeEmailDomains()))
	}
	if strings.Contains(email, "SafeEmailDomains") {
		email = strings.ReplaceAll(email, "{{SafeEmailDomains}}", random.FromSliceOfString(i.Provider.SafeEmailDomains()))
	}

	return strings.ToLower(email)
}

//SafeEmail returns a random email address from a safe domain
func (i internet) SafeEmail() string {
	return strings.ToLower(i.Username() + "@" + random.FromSliceOfString(i.Provider.SafeEmailDomains()))
}

//FreeEmail returns a random free email address
func (i internet) FreeEmail() string {
	return strings.ToLower(i.Username() + "@" + random.FromSliceOfString(i.Provider.FreeEmailDomains()))
}

//CompanyEmail returns a random "company" email address
func (i internet) CompanyEmail() string {
	return strings.ToLower(i.Username() + "@" + i.DomainName())
}

// DomainName provides a random domain name
func (i internet) DomainName() string {
	return i.domainWord() + "." + i.tld()
}

func (i internet) UrlFormats() []string {
	return []string{
		"http://www.{{domainName}}/",
		"http://{{domainName}}/",
		"http://www.{{domainName}}/{{slug}}",
		"http://www.{{domainName}}/{{slug}}",
		"https://www.{{domainName}}/{{slug}}",
		"http://www.{{domainName}}/{{slug}}.html",
		"http://{{domainName}}/{{slug}}",
		"http://{{domainName}}/{{slug}}",
		"http://{{domainName}}/{{slug}}.html",
		"https://{{domainName}}/{{slug}}.html",
	}
}

// Url provides a random URL
func (i internet) Url() string {
	urlFormat := random.FromSliceOfString(i.Provider.UrlFormats())

	var url string
	if strings.Contains(urlFormat, "domainName") {
		url = strings.ReplaceAll(urlFormat, "{{domainName}}", i.DomainName())
	}
	if strings.Contains(url, "slug") {
		url = strings.ReplaceAll(url, "{{slug}}", i.slug(rand.Intn(3)))
	}

	return url
}

func (i internet) slug(size int) string {
	slug := random.FromSliceOfString(i.Provider.Words())
	if size > 1 {
		for words := 0; words < size-1; words++ {
			slug = strings.Join([]string{slug, random.FromSliceOfString(i.Provider.Words())}, "-")
		}
	}

	return slug
}

func (i internet) domainWord() string {
	return strings.ToLower(i.toAscii(i.Person.LastName()))
}

func (i internet) tld() string {
	return random.FromSliceOfString(i.Provider.Tld())
}

func (i internet) toAscii(s string) string {
	transliteration := map[string]string{
		"Ĳ": "I", "Ö": "O", "Œ": "O", "Ü": "U", "ä": "a", "æ": "a",
		"ĳ": "i", "ö": "o", "œ": "o", "ü": "u", "ß": "s", "ſ": "s",
		"À": "A", "Á": "A", "Â": "A", "Ã": "A", "Ä": "A", "Å": "A",
		"Æ": "A", "Ā": "A", "Ą": "A", "Ă": "A", "Ç": "C", "Ć": "C",
		"Č": "C", "Ĉ": "C", "Ċ": "C", "Ď": "D", "Đ": "D", "È": "E",
		"É": "E", "Ê": "E", "Ë": "E", "Ē": "E", "Ę": "E", "Ě": "E",
		"Ĕ": "E", "Ė": "E", "Ĝ": "G", "Ğ": "G", "Ġ": "G", "Ģ": "G",
		"Ĥ": "H", "Ħ": "H", "Ì": "I", "Í": "I", "Î": "I", "Ï": "I",
		"Ī": "I", "Ĩ": "I", "Ĭ": "I", "Į": "I", "İ": "I", "Ĵ": "J",
		"Ķ": "K", "Ľ": "K", "Ĺ": "K", "Ļ": "K", "Ŀ": "K", "Ł": "L",
		"Ñ": "N", "Ń": "N", "Ň": "N", "Ņ": "N", "Ŋ": "N", "Ò": "O",
		"Ó": "O", "Ô": "O", "Õ": "O", "Ø": "O", "Ō": "O", "Ő": "O",
		"Ŏ": "O", "Ŕ": "R", "Ř": "R", "Ŗ": "R", "Ś": "S", "Ş": "S",
		"Ŝ": "S", "Ș": "S", "Š": "S", "Ť": "T", "Ţ": "T", "Ŧ": "T",
		"Ț": "T", "Ù": "U", "Ú": "U", "Û": "U", "Ū": "U", "Ů": "U",
		"Ű": "U", "Ŭ": "U", "Ũ": "U", "Ų": "U", "Ŵ": "W", "Ŷ": "Y",
		"Ÿ": "Y", "Ý": "Y", "Ź": "Z", "Ż": "Z", "Ž": "Z", "à": "a",
		"á": "a", "â": "a", "ã": "a", "ā": "a", "ą": "a", "ă": "a",
		"å": "a", "ç": "c", "ć": "c", "č": "c", "ĉ": "c", "ċ": "c",
		"ď": "d", "đ": "d", "è": "e", "é": "e", "ê": "e", "ë": "e",
		"ē": "e", "ę": "e", "ě": "e", "ĕ": "e", "ė": "e", "ƒ": "f",
		"ĝ": "g", "ğ": "g", "ġ": "g", "ģ": "g", "ĥ": "h", "ħ": "h",
		"ì": "i", "í": "i", "î": "i", "ï": "i", "ī": "i", "ĩ": "i",
		"ĭ": "i", "į": "i", "ı": "i", "ĵ": "j", "ķ": "k", "ĸ": "k",
		"ł": "l", "ľ": "l", "ĺ": "l", "ļ": "l", "ŀ": "l", "ñ": "n",
		"ń": "n", "ň": "n", "ņ": "n", "ŉ": "n", "ŋ": "n", "ò": "o",
		"ó": "o", "ô": "o", "õ": "o", "ø": "o", "ō": "o", "ő": "o",
		"ŏ": "o", "ŕ": "r", "ř": "r", "ŗ": "r", "ś": "s", "š": "s",
		"ť": "t", "ù": "u", "ú": "u", "û": "u", "ū": "u", "ů": "u",
		"ű": "u", "ŭ": "u", "ũ": "u", "ų": "u", "ŵ": "w", "ÿ": "y",
		"ý": "y", "ŷ": "y", "ż": "z", "ź": "z", "ž": "z", "Α": "A",
		"Ά": "A", "Ἀ": "A", "Ἁ": "A", "Ἂ": "A", "Ἃ": "A", "Ἄ": "A",
		"Ἅ": "A", "Ἆ": "A", "Ἇ": "A", "ᾈ": "A", "ᾉ": "A", "ᾊ": "A",
		"ᾋ": "A", "ᾌ": "A", "ᾍ": "A", "ᾎ": "A", "ᾏ": "A", "Ᾰ": "A",
		"Ᾱ": "A", "Ὰ": "A", "ᾼ": "A", "Β": "B", "Γ": "G", "Δ": "D",
		"Ε": "E", "Έ": "E", "Ἐ": "E", "Ἑ": "E", "Ἒ": "E", "Ἓ": "E",
		"Ἔ": "E", "Ἕ": "E", "Ὲ": "E", "Ζ": "Z", "Η": "I", "Ή": "I",
		"Ἠ": "I", "Ἡ": "I", "Ἢ": "I", "Ἣ": "I", "Ἤ": "I", "Ἥ": "I",
		"Ἦ": "I", "Ἧ": "I", "ᾘ": "I", "ᾙ": "I", "ᾚ": "I", "ᾛ": "I",
		"ᾜ": "I", "ᾝ": "I", "ᾞ": "I", "ᾟ": "I", "Ὴ": "I", "ῌ": "I",
		"Θ": "T", "Ι": "I", "Ί": "I", "Ϊ": "I", "Ἰ": "I", "Ἱ": "I",
		"Ἲ": "I", "Ἳ": "I", "Ἴ": "I", "Ἵ": "I", "Ἶ": "I", "Ἷ": "I",
		"Ῐ": "I", "Ῑ": "I", "Ὶ": "I", "Κ": "K", "Λ": "L", "Μ": "M",
		"Ν": "N", "Ξ": "K", "Ο": "O", "Ό": "O", "Ὀ": "O", "Ὁ": "O",
		"Ὂ": "O", "Ὃ": "O", "Ὄ": "O", "Ὅ": "O", "Ὸ": "O", "Π": "P",
		"Ρ": "R", "Ῥ": "R", "Σ": "S", "Τ": "T", "Υ": "Y", "Ύ": "Y",
		"Ϋ": "Y", "Ὑ": "Y", "Ὓ": "Y", "Ὕ": "Y", "Ὗ": "Y", "Ῠ": "Y",
		"Ῡ": "Y", "Ὺ": "Y", "Φ": "F", "Χ": "X", "Ψ": "P", "Ω": "O",
		"Ώ": "O", "Ὠ": "O", "Ὡ": "O", "Ὢ": "O", "Ὣ": "O", "Ὤ": "O",
		"Ὥ": "O", "Ὦ": "O", "Ὧ": "O", "ᾨ": "O", "ᾩ": "O", "ᾪ": "O",
		"ᾫ": "O", "ᾬ": "O", "ᾭ": "O", "ᾮ": "O", "ᾯ": "O", "Ὼ": "O",
		"ῼ": "O", "α": "a", "ά": "a", "ἀ": "a", "ἁ": "a", "ἂ": "a",
		"ἃ": "a", "ἄ": "a", "ἅ": "a", "ἆ": "a", "ἇ": "a", "ᾀ": "a",
		"ᾁ": "a", "ᾂ": "a", "ᾃ": "a", "ᾄ": "a", "ᾅ": "a", "ᾆ": "a",
		"ᾇ": "a", "ὰ": "a", "ᾰ": "a", "ᾱ": "a", "ᾲ": "a", "ᾳ": "a",
		"ᾴ": "a", "ᾶ": "a", "ᾷ": "a", "β": "b", "γ": "g", "δ": "d",
		"ε": "e", "έ": "e", "ἐ": "e", "ἑ": "e", "ἒ": "e", "ἓ": "e",
		"ἔ": "e", "ἕ": "e", "ὲ": "e", "ζ": "z", "η": "i", "ή": "i",
		"ἠ": "i", "ἡ": "i", "ἢ": "i", "ἣ": "i", "ἤ": "i", "ἥ": "i",
		"ἦ": "i", "ἧ": "i", "ᾐ": "i", "ᾑ": "i", "ᾒ": "i", "ᾓ": "i",
		"ᾔ": "i", "ᾕ": "i", "ᾖ": "i", "ᾗ": "i", "ὴ": "i", "ῂ": "i",
		"ῃ": "i", "ῄ": "i", "ῆ": "i", "ῇ": "i", "θ": "t", "ι": "i",
		"ί": "i", "ϊ": "i", "ΐ": "i", "ἰ": "i", "ἱ": "i", "ἲ": "i",
		"ἳ": "i", "ἴ": "i", "ἵ": "i", "ἶ": "i", "ἷ": "i", "ὶ": "i",
		"ῐ": "i", "ῑ": "i", "ῒ": "i", "ῖ": "i", "ῗ": "i", "κ": "k",
		"λ": "l", "μ": "m", "ν": "n", "ξ": "k", "ο": "o", "ό": "o",
		"ὀ": "o", "ὁ": "o", "ὂ": "o", "ὃ": "o", "ὄ": "o", "ὅ": "o",
		"ὸ": "o", "π": "p", "ρ": "r", "ῤ": "r", "ῥ": "r", "σ": "s",
		"ς": "s", "τ": "t", "υ": "y", "ύ": "y", "ϋ": "y", "ΰ": "y",
		"ὐ": "y", "ὑ": "y", "ὒ": "y", "ὓ": "y", "ὔ": "y", "ὕ": "y",
		"ὖ": "y", "ὗ": "y", "ὺ": "y", "ῠ": "y", "ῡ": "y", "ῢ": "y",
		"ῦ": "y", "ῧ": "y", "φ": "f", "χ": "x", "ψ": "p", "ω": "o",
		"ώ": "o", "ὠ": "o", "ὡ": "o", "ὢ": "o", "ὣ": "o", "ὤ": "o",
		"ὥ": "o", "ὦ": "o", "ὧ": "o", "ᾠ": "o", "ᾡ": "o", "ᾢ": "o",
		"ᾣ": "o", "ᾤ": "o", "ᾥ": "o", "ᾦ": "o", "ᾧ": "o", "ὼ": "o",
		"ῲ": "o", "ῳ": "o", "ῴ": "o", "ῶ": "o", "ῷ": "o", "А": "A",
		"Б": "B", "В": "V", "Г": "G", "Д": "D", "Е": "E", "Ё": "E",
		"Ж": "Z", "З": "Z", "И": "I", "Й": "I", "К": "K", "Л": "L",
		"М": "M", "Н": "N", "О": "O", "П": "P", "Р": "R", "С": "S",
		"Т": "T", "У": "U", "Ф": "F", "Х": "K", "Ц": "T", "Ч": "C",
		"Ш": "S", "Щ": "S", "Ы": "Y", "Э": "E", "Ю": "Y", "Я": "Y",
		"а": "A", "б": "B", "в": "V", "г": "G", "д": "D", "е": "E",
		"ё": "E", "ж": "Z", "з": "Z", "и": "I", "й": "I", "к": "K",
		"л": "L", "м": "M", "н": "N", "о": "O", "п": "P", "р": "R",
		"с": "S", "т": "T", "у": "U", "ф": "F", "х": "K", "ц": "T",
		"ч": "C", "ш": "S", "щ": "S", "ы": "Y", "э": "E", "ю": "Y",
		"я": "Y", "ð": "d", "Ð": "D", "þ": "t", "Þ": "T", "ა": "a",
		"ბ": "b", "გ": "g", "დ": "d", "ე": "e", "ვ": "v", "ზ": "z",
		"თ": "t", "ი": "i", "კ": "k", "ლ": "l", "მ": "m", "ნ": "n",
		"ო": "o", "პ": "p", "ჟ": "z", "რ": "r", "ს": "s", "ტ": "t",
		"უ": "u", "ფ": "p", "ქ": "k", "ღ": "g", "ყ": "q", "შ": "s",
		"ჩ": "c", "ც": "t", "ძ": "d", "წ": "t", "ჭ": "c", "ხ": "k",
		"ჯ": "j", "ჰ": "h", "ţ": "t", "ʼ": "\"", "̧": "", "ḩ": "h",
		"‘": "\"", "’": "\"", "ừ": "u", "/": "", "ế": "e", "ả": "a",
		"ị": "i", "ậ": "a", "ệ": "e", "ỉ": "i", "ồ": "o", "ề": "e",
		"ơ": "o", "ạ": "a", "ẵ": "a", "ư": "u", "ằ": "a", "ầ": "a",
		"ḑ": "d", "Ḩ": "H", "Ḑ": "D", "ș": "s", "ț": "t", "ộ": "o",
		"ắ": "a", "ş": "s", "\"": "", "ու": "u", "ա": "a", "բ": "b",
		"գ": "g", "դ": "d", "ե": "e", "զ": "z", "է": "e", "ը": "y",
		"թ": "t", "ժ": "zh", "ի": "i", "լ": "l", "խ": "kh", "ծ": "ts",
		"կ": "k", "հ": "h", "ձ": "dz", "ղ": "gh", "ճ": "ch", "մ": "m",
		"յ": "y", "ն": "n", "շ": "sh", "ո": "o", "չ": "ch", "պ": "p",
		"ջ": "j", "ռ": "r", "ս": "s", "վ": "v", "տ": "t", "ր": "r",
		"ց": "ts", "փ": "p", "ք": "q", "և": "ev", "օ": "o", "ֆ": "f",
	}

	st := s
	for k, v := range transliteration {
		st = strings.ReplaceAll(st, k, v)
	}

	return st
}
