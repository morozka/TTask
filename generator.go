package main

import "fmt"

var charSet = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var nbit = len(charSet)

//GenerateShortURL - converts the unique number of URL to n-bit number system
func GenerateShortURL(id int) (string, error) {
	if id < 0 {
		return "", fmt.Errorf("%v", "Incorrect ID, ID must be greater than zero")
	}
	if id == 0 {
		return "0", nil
	}
	var shortURL []rune
	var nextRune int
	for id > 0 {
		nextRune = id % nbit
		shortURL = append(shortURL, charSet[nextRune])
		id = id / nbit
	}
	return string(shortURL), nil
}
