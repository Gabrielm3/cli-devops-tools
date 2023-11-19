package utils

import "github.com/likexian/whois"

func GetWhois(host string) string {
	result, _ := whois.Whois(host)
	return result
}