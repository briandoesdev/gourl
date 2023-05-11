package main

type Url struct {
	url         string
	protocol    string
	credentials string
	host        Host
	port        int
	path        string
	query       string
	fragment    string
}

type Host struct {
	subdomain string
	domain    string
	tld       string
}

// ParseStringToUrl takes a string and returns a Url struct
func ParseStringToUrl(url string) *Url {
	return nil
}

// ParseProtocol returns the protocol of the Url
func ParseProtocol(url string) string {
	return ""
}
