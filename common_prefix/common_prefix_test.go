package common_prefix

import (
	"fmt"
	"path/filepath"
	"testing"
)

var testData = [][]string{
	{"/home/user/goeg", "/home/user/goeg/prefix",
		"/home/user/goeg/prefix/extra"},
	{"/home/user/goeg", "/home/user/goeg/prefix",
		"/home/user/prefix/extra"},
	{"/pecan/π/goeg", "/pecan/π/goeg/prefix",
		"/pecan/π/prefix/extra"},
	{"/pecan/π/circle", "/pecan/π/circle/prefix",
		"/pecan/π/circle/prefix/extra"},
	{"/home/user/goeg", "/home/users/goeg",
		"/home/userspace/goeg"},
	{"/home/user/goeg", "/tmp/user", "/var/log"},
	{"/home/mark/goeg", "/home/user/goeg"},
	{"home/user/goeg", "/tmp/user", "/var/log"},
}

func TestCommonPrefix(t *testing.T) {
	words := []string{"ab", "abbbb", "abab_123"}
	fmt.Println("	共同前缀是：", CommonPrefix(words))
	fmt.Println("系统的separptor是 ", string(filepath.Separator))

}

func TestCommonPathPrefix(t *testing.T) {
	for _, data := range testData {
		fmt.Println("前缀是：", CommonPathPrefix(data))
	}

}
