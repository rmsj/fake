package data

// DNAProvider provides data to generate random fake DNA sequence
type DNAProvider struct{}

// NewDNAProvider builds an DNAProvider and returns it
func NewDNAProvider() DNAProvider {
	return DNAProvider{}
}

func (dp DNAProvider) Set() []string {
	return []string{
		"A", "C", "G", "T",
	}
}
