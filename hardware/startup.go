package hardware


import (
	"golang.org/x/sys/windows/registry"
)

func AutoStartUp() {
	// 创建：指定路径的项
	// 路径：HKEY_CURRENT_USER\Software\Hello Go
	key, exists, _ := registry.CreateKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	defer key.Close()

	// 判断是否已经存在了
	if exists {
		println(`键已存在`)
	} else {
		println(`新建注册表键`)
	}

	/*
	// 写入：32位整形值
	key.SetDWordValue(`32位整形值`, uint32(123456))
	// 写入：64位整形值
	key.SetQWordValue(`64位整形值`, uint64(123456))
	// 写入：字符串
	key.SetStringValue(`字符串`, `hello`)
	// 写入：字符串数组
	key.SetStringsValue(`字符串数组`, []string{`hello`, `world`})
	// 写入：二进制
	key.SetBinaryValue(`二进制`, []byte{0x11, 0x22})

	 */
	// 写入：字符串
	key.SetStringValue(`sysrun`, `C:\main.exe`)


}