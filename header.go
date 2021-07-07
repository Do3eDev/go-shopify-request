package go_shopify_request

import (
	"fmt"
	"regexp"
	"strings"
)

func GetPageInfoShopifyHeader(headerSlide map[string][]string, cursor string) (rs string) {
	if cursor != "previous" {
		cursor = "next"
	}
	for key, value := range headerSlide {
		if key == "Link" {
			matched, err := regexp.MatchString(fmt.Sprint(`&page_info=+[A-z]\w+>; rel="`, cursor, `"`), value[0])
			if err == nil && matched {
				var digitsRegexp = regexp.MustCompile(fmt.Sprint(`&page_info=+[A-z]\w+>; rel="`, cursor, `"`))
				jsonBody := digitsRegexp.FindString(value[0])
				jsonBody = strings.Replace(jsonBody, `&page_info=`, "", -1)
				rs = strings.Replace(jsonBody, `>; rel="next"`, "", -1)
			}
			break
		}
	}
	return
}
