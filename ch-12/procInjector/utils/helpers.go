package utils

import (
	"syscall"
	"unicode/utf16"
)

// FullPath是一个辅助函数
func FullPath(name string) (path string, err error) {
	p, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return "", err
	}
	n := uint32(100)
	for {
		buf := make([]uint16, n)
		n, err = syscall.GetFullPathName(p, uint32(len(buf)), &buf[0], nil)
		if err != nil {
			return "", err
		}
		if n <= uint32(len(buf)) {
			return syscall.UTF16ToString(buf[:n]), nil
		}
	}
}

// StringToCharPtr将Go字符串转换为指向以null结尾的c字符串的指针。
// 这假设go字符串已经是ANSI编码的。
func StringToCharPtr(str string) *uint8 {
	chars := append([]byte(str), 0) // 以null结尾
	return &chars[0]
}

// StringToUTF16Ptr将Go字符串转换为指向以null结尾的UTF-16宽字符串的指针。
// 这假设str是UTF-8兼容编码,因此可以重新编码为UTF-16。
func StringToUTF16Ptr(str string) *uint16 {
	wchars := utf16.Encode([]rune(str + "\x00"))
	return &wchars[0]
}
