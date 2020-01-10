package prompt

import (
	"strings"
)

type PlaceholderValues struct {
	Org     string
	Space   string
	API     string
	User    string
	Version string
}

func Format(format string, values PlaceholderValues) string {
	// strings.NewReplacer takes a slice of strings
	// to replace, where index i is replaced by index i+1.
	// There must be an even number
	replacerPairs := []string{
		"%s", values.Space,
		"%o", values.Org,
		"%a", values.API,
		"%u", values.User,
		"%v", values.Version,
	}

	replacer := strings.NewReplacer(replacerPairs...)
	return replacer.Replace(format)
}
