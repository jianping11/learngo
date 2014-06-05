package main

//［section1］------使用映射来创造分支

// var FunctionForSuffix = map[string]func(string) ([]stirng, error){
// 	"。zip": ZipFileList,
// 	".tar": TarFileList,
// }
//
// func ArchiveFileListMap(file string) ([]string, error) {
// 	if function, ok := FunctionForSuffix[Suffix(file)]; ok {
// 		return function(file)
// 	}
//
// 	return nil, errors.New("unrecognized archive")
// }
