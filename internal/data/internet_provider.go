package data

// InternetProvider provides data for english language for emails, URLs, username, etc.
type InternetProvider struct{}

// NewInternetProvider builds an InternetProvider and returns it
func NewInternetProvider() InternetProvider {
	return InternetProvider{}
}

// SafeEmailDomains returns safe email domain options
func (i InternetProvider) SafeEmailDomains() []string {
	return []string{"example.com", "example.org", "example.net"}
}

// FreeEmailDomains returns free domain emails options
func (i InternetProvider) FreeEmailDomains() []string {
	return []string{"gmail.com", "yahoo.com", "hotmail.com"}
}

// Tld returns top level domain options
func (i InternetProvider) Tld() []string {
	return []string{"com", "com", "com", "com", "co.nz", "com.au", "biz", "info", "net", "org"}
}

// EmailFormats returns options for email format
func (i InternetProvider) EmailFormats() []string {
	return []string{
		"{{userName}}@{{domainName}}",
		"{{userName}}@{{freeEmailDomain}}",
		"{{userName}}@{{safeEmailDomain}}",
	}
}

// UserNameFormats returns options for user name formats
func (i InternetProvider) UserNameFormats() []string {
	return []string{
		"{{lastName}}.{{firstName}}",
		"{{firstName}}.{{lastName}}",
		"{{firstName}}",
		"{{lastName}}",
	}
}

// UrlFormats returns options for URL formats
func (i InternetProvider) UrlFormats() []string {
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
