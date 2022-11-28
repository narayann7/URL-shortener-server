package services

import (
	"net/url"
	"regexp"
)

func IsVaildUrl(str string) bool {

	if str[0:8] == "https://" || str[0:7] == "http://" {
		u, err := url.Parse(str)
		if err != nil {
			return false
		}
		return CheckForDomain(u.Hostname())
	} else {
		return CheckForDomain(str)
	}

}

func CheckForDomain(domain string) bool {
	RegExp := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z
		]{2,3})$`)
	return RegExp.MatchString(domain)
}
