package provider

// EnglishInternetProvider provides data for english language for emails, URLs, username, etc.
type EnglishInternetProvider struct {
	lang    string
	country string
}

// NewEnglishInternetProvider builds an EnglishInternetProvider and returns it
func NewEnglishInternetProvider() EnglishInternetProvider {
	return EnglishInternetProvider{
		lang:    "en",
		country: "us",
	}
}

func (i EnglishInternetProvider) SafeEmailDomains() []string {
	return []string{"example.com", "example.org", "example.net"}
}

func (i EnglishInternetProvider) FreeEmailDomains() []string {
	return []string{"gmail.com", "yahoo.com", "hotmail.com"}
}

func (i EnglishInternetProvider) Tld() []string {
	return []string{"com", "com", "com", "com", "com", "com", "biz", "info", "net", "org"}
}

func (i EnglishInternetProvider) EmailFormats() []string {
	return []string{
		"{{userName}}@{{domainName}}",
		"{{userName}}@{{freeEmailDomain}}",
		"{{userName}}@{{safeEmailDomain}}",
	}
}

func (i EnglishInternetProvider) UserNameFormats() []string {
	return []string{
		"{{lastName}}.{{firstName}}",
		"{{firstName}}.{{lastName}}",
		"{{firstName}}",
		"{{lastName}}",
	}
}

func (i EnglishInternetProvider) UrlFormats() []string {
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

func (i EnglishInternetProvider) Words() []string {
	return []string{
		"alias", "consequatur", "aut", "perferendis", "sit", "voluptatem",
		"accusantium", "doloremque", "aperiam", "eaque", "ipsa", "quae", "ab",
		"illo", "inventore", "veritatis", "et", "quasi", "architecto",
		"beatae", "vitae", "dicta", "sunt", "explicabo", "aspernatur", "aut",
		"odit", "aut", "fugit", "sed", "quia", "consequuntur", "magni",
		"dolores", "eos", "qui", "ratione", "voluptatem", "sequi", "nesciunt",
		"neque", "dolorem", "ipsum", "quia", "dolor", "sit", "amet",
		"consectetur", "adipisci", "velit", "sed", "quia", "non", "numquam",
		"eius", "modi", "tempora", "incidunt", "ut", "labore", "et", "dolore",
		"magnam", "aliquam", "quaerat", "voluptatem", "ut", "enim", "ad",
		"minima", "veniam", "quis", "nostrum", "exercitationem", "ullam",
		"corporis", "nemo", "enim", "ipsam", "voluptatem", "quia", "voluptas",
		"sit", "suscipit", "laboriosam", "nisi", "ut", "aliquid", "ex", "ea",
		"commodi", "consequatur", "quis", "autem", "vel", "eum", "iure",
		"reprehenderit", "qui", "in", "ea", "voluptate", "velit", "esse",
		"quam", "nihil", "molestiae", "et", "iusto", "odio", "dignissimos",
		"ducimus", "qui", "blanditiis", "praesentium", "laudantium", "totam",
		"rem", "voluptatum", "deleniti", "atque", "corrupti", "quos",
		"dolores", "et", "quas", "molestias", "excepturi", "sint",
		"occaecati", "cupiditate", "non", "provident", "sed", "ut",
		"perspiciatis", "unde", "omnis", "iste", "natus", "error",
		"similique", "sunt", "in", "culpa", "qui", "officia", "deserunt",
		"mollitia", "animi", "id", "est", "laborum", "et", "dolorum", "fuga",
		"et", "harum", "quidem", "rerum", "facilis", "est", "et", "expedita",
		"distinctio", "nam", "libero", "tempore", "cum", "soluta", "nobis",
		"est", "eligendi", "optio", "cumque", "nihil", "impedit", "quo",
		"porro", "quisquam", "est", "qui", "minus", "id", "quod", "maxime",
		"placeat", "facere", "possimus", "omnis", "voluptas", "assumenda",
		"est", "omnis", "dolor", "repellendus", "temporibus", "autem",
		"quibusdam", "et", "aut", "consequatur", "vel", "illum", "qui",
		"dolorem", "eum", "fugiat", "quo", "voluptas", "nulla", "pariatur",
		"at", "vero", "eos", "et", "accusamus", "officiis", "debitis", "aut",
		"rerum", "necessitatibus", "saepe", "eveniet", "ut", "et",
		"voluptates", "repudiandae", "sint", "et", "molestiae", "non",
		"recusandae", "itaque", "earum", "rerum", "hic", "tenetur", "a",
		"sapiente", "delectus", "ut", "aut", "reiciendis", "voluptatibus",
		"maiores", "doloribus", "asperiores", "repellat",
	}
}
