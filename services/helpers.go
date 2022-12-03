package services

import (
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/narayann7/gourl/database"
)

func IsVaildUrl(str string) bool {
	if len(str) >= 8 && (str[0:8] == "https://" || str[0:7] == "http://") {
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

func EnforceHttp(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	} else {
		return url
	}
}

func CheckForDominError(url string) bool {

	if url == os.Getenv("DOMAIN") {
		return false
	}
	newUrl := strings.Replace(url, "https://", "", 1)
	newUrl = strings.Replace(url, "http://", "", 1)
	newUrl = strings.Replace(url, "www.", "", 1)
	newUrl = strings.Split(newUrl, "/")[0]
	return newUrl != os.Getenv("DOMAIN")

}
func CheckForVaildExpiry(expiry *time.Duration) bool {

	//max limit of url is 1 day
	//which means 24 * 60 = 1440 minutes
	min := time.Minute
	min = min * *expiry

	if min > 0 && min.Minutes() <= 1440 {
		*expiry = min / time.Minute
		return true
	} else {
		return false
	}

}
func CheckForVaildCustomUrl(customUrl string) bool {
	//length should = 8
	if len(customUrl) != 8 {
		return false
	}
	var isVaild bool = false
	for i := range customUrl {
		c := customUrl[i]
		//a vaild custom url should contain only Alphabet and Digits
		if (c >= 48 && c <= 57) || (c >= 65 && c <= 90) || (c >= 90 && c <= 122) {
			isVaild = true
		} else {
			isVaild = false
			break
		}
	}
	return isVaild
}

func SendErrorHtml() string {
	return `<html>
	<head></head>
	<body
	  style="
		 display: flex;
		 height: 100vh;
		 width: 100vw;
		 padding: 0%;
		 margin: 0%;
		 align-items: center;
		 justify-content: center;
		 flex-direction: column;
	  "
	>
	  <h3>Url is not vaild. Please visit</h3>
	  <a href="https://google.com"></a>
	</body>
 </html>
 `

}
func GetNewHashes(size int) []string {

	dataList := []string{}

	for i := 0; i < size; i++ {
		newHash := database.GetNewHash()
		dataList = append(dataList, newHash)
	}

	return dataList
}
