package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	// 绝对路径
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		// 终止程序
		panic(err)
	}
	return &DirEntry{absDir}
}

// Go结构体不需要显示实现接口，只要方法匹配即可
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	// 拼接目录和class文件名
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
