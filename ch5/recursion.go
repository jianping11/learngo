package main

import (
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {

	if isPalindrome := IsPalindrome(os.Args[1]); isPalindrome {
		fmt.Printf("%s 是回文单词 \n", os.Args[1])
	} else {
		fmt.Printf("%s 不是回文单词 \n", os.Args[1])
	}
}

//判断是否十回文单词，如ROTOR，PULLUP等
func IsPalindrome(word string) bool {
	//退出条件,当只有一个字符时，退出循环
	if utf8.RuneCountInString(word) <= 1 {
		return true
	}

	//进行判断
	first, sizeOfFirst := utf8.DecodeRuneInString(word)
	last, sizeOfLast := utf8.DecodeLastRuneInString(word)
	if first != last {
		return false
	}

	return IsPalindrome(word[sizeOfFirst : len(word)-sizeOfLast])
}
