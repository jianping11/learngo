package common_prefix

import (
	"bytes"
	"path/filepath"
	"strings"
)

//commonPrefix 负责提取共同字符串的
func CommonPrefix(texts []string) string {

	//初始化一个rune的二元数组，用来储存转换成rune后的text，长度容量设为传入的数组的长度
	components := make([][]rune, len(texts))
	for i, text := range texts {
		components[i] = []rune(text)
	}

	//如果数组的长度是0，则返回“”，
	if len(components) == 0 || len(components[0]) == 0 {
		return ""
	}
	//使用bytes.Buffer保存数据
	var common bytes.Buffer
FINISH:
	//取出第一行数据，遍历这一行的每一列
	for column := 0; column < len(components[0]); column++ {
		//取出第一行的第一个字符，用来比较
		char := components[0][column]
		//从第二行开始，遍历二元数组的每一行，取出每一行的第column列与第一行取出的char进行比较
		//当有一行的数据与char不相同，则跳出外围循环，否则，将char字符保留起来，
		for row := 1; row < len(components); row++ {
			if column >= len(components[row]) || components[row][column] != char {
				break FINISH
			}
		}
		common.WriteRune(char)
	}
	//返回数据
	return common.String()
}

func CommonPathPrefix(paths []string) string {
	const separator = string(filepath.Separator)
	components := make([][]string, len(paths))
	for i, path := range paths {
		components[i] = strings.Split(path, separator)
		if strings.HasPrefix(path, separator) {
			components[i] = append([]string{separator}, components[i]...)
		}
	}

	if len(components) == 0 || len(components[0]) == 0 {
		return ""
	}
	var common []string
FINISH:
	for column := range components[0] {
		part := components[0][column]
		for row := 1; row < len(components); row++ {
			if len(components[row]) == 0 || column >= len(components[row]) || components[row][column] != part {
				break FINISH
			}
		}
		common = append(common, part)
	}
	return filepath.Join(common...)
}
