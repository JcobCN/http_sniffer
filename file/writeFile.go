package file

import (
	"fmt"
	"os"
)

func WriteWithOs(name, content string){
	data := []byte(content)
	fl, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		fmt.Printf("open file error %#v\n", err, os.SyscallError{})
		return
	}
	defer fl.Close()

	fmt.Println(name)
	n, err := fl.Write(data)
	if err != nil{
		fmt.Println("出错")
	} else if err == nil && n < len(data) {
		fmt.Println("写入缺少")
	} else {
		fmt.Println("写入成功")
	}

}