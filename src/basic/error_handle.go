package main

import (
	"fmt"
)

// ------------自定义异常--------------
type FileError struct {
	Op   string
	Name string
	Path string
}

// 初始化函数
func NewFileError(op string, name string, path string) *FileError {
	return &FileError{Op: op, Name: name, Path: path} // 返回一个结构体
}

// 使用 FileError 结构体实现接口 error.Error() 方法， 此时结构体拥有Error()
func (f *FileError) Error() string {
	return fmt.Sprintf("路径为 %v 的文件 %v，在 %v 操作时出错", f.Path, f.Name, f.Op)
}

// ------------自定义异常--------------

func division(a int, b int) (float64, error) {
	if b == 0 {
		err := fmt.Errorf("Err: b = %d !", 0) // 创建一个异常
		return 0, err
	}
	return float64(a / b), nil
}

func main() {
	f, err := division(10, 0)
	if err != nil {
		fmt.Println(err.Error()) // 打印异常信息
	}
	fmt.Println(f)
	file_err := NewFileError("一些操作", "这个文件", "/home")
	fmt.Println(file_err.Error())
}
