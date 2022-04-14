package Utils

import (
	"fmt"
	"io"
	"io/fs"
	"os"
)

func CopyFile(dstName, srcName string) (err error) {
	file, err := os.Open(dstName)
	filepath, err := os.Create(srcName)
	if err != nil {
		return err
	}
	//写入二进制数据
	_, err = io.Copy(filepath, file)
	if err != nil {
		return err
	}
	// 结束打开的文件释放占用状态（内存）
	file.Close()
	// 结束创建的新文件的占用状态）（硬盘）
	filepath.Close()
	return nil
}

// PathExists 存在返回true
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// AutoCreatePath 自动创建目录，创建成功返回true 否返回false
func AutoCreatePath(path string, FileMode fs.FileMode) bool {
	if !PathExists(path) {
		err := os.Mkdir(path, FileMode)
		if err != nil {
			fmt.Println(fmt.Sprintf("创建目录'%s'失败", path))
			return false
		}
	}
	return true
}
