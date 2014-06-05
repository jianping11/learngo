package palindrome_ans

import (
	"fmt"
	"unicode/utf8"
)

//非递归版本,
func IsPalindrome(word string) bool {
	//将word 转化成unicode码点，防止word为非asnii码
	runes := []rune(word)

	// 设置循环的次数为len(runes)/2,,这样可以通用的处理runes个数为奇数后者偶数，例如aboba需要循环比较两次就可以了，忽略中间的o
	for i := 0; i < len(runes)/2; i++ {
		fmt.Println("tow words: ", runes[i], " ", runes[len(runes)-1-i])
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

//判断是否十回文单词，如ROTOR，PULLUP等 (使用递归)
func IsPalindrome_old(word string) bool {
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

	return IsPalindrome_old(word[sizeOfFirst : len(word)-sizeOfLast])
}
