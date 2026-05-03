package main

import (
	"debug/pe"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("Telegram.exe") // 修改为二进制文件或改为接受参数
	check(err)
	pefile, err := pe.NewFile(f)
	check(err)
	defer f.Close()
	defer pefile.Close()

	dosHeader := make([]byte, 96)
	sizeOffset := make([]byte, 4)

	// 十进制到ASCII(搜索MZ)
	_, err = f.Read(dosHeader)
	check(err)
	fmt.Println("[-----DOS Header / Stub-----]")
	fmt.Printf("[+] Magic Value: %s%s\n", string(dosHeader[0]), string(dosHeader[1]))

	// 验证PE+0+0(有效的PE格式)
	pe_sig_offset := int64(binary.LittleEndian.Uint32(dosHeader[0x3c:]))
	f.ReadAt(sizeOffset[:], pe_sig_offset)
	fmt.Println("[-----Signature Header-----]")
	fmt.Printf("[+] LFANEW Value: %s\n", string(sizeOffset))

	// 创建reader并读取COFF Header
	sr := io.NewSectionReader(f, 0, 1<<63-1)
	_, err = sr.Seek(pe_sig_offset+4, os.SEEK_SET)
	check(err)
	binary.Read(sr, binary.LittleEndian, &pefile.FileHeader)

	// 获取OptionalHeader的大小
	var sizeofOptionalHeader32 = uint16(binary.Size(pe.OptionalHeader32{}))
	var sizeofOptionalHeader64 = uint16(binary.Size(pe.OptionalHeader64{}))
	var oh32 pe.OptionalHeader32
	var oh64 pe.OptionalHeader64

	// 读取OptionalHeader
	switch pefile.FileHeader.SizeOfOptionalHeader {
	case sizeofOptionalHeader32:
		binary.Read(sr, binary.LittleEndian, &oh32)
	case sizeofOptionalHeader64:
		binary.Read(sr, binary.LittleEndian, &oh64)
	}
	// 打印File Header
	fmt.Println("[-----COFF File Header-----]")
	fmt.Printf("[+] Machine Architecture: %#x\n", pefile.FileHeader.Machine)
	fmt.Printf("[+] Number of Sections: %#x\n", pefile.FileHeader.NumberOfSections)
	fmt.Printf("[+] Size of Optional Header: %#x\n", pefile.FileHeader.SizeOfOptionalHeader)
	// 打印节名称
	fmt.Println("[-----Section Offsets-----]")
	fmt.Printf("[+] Number of Sections Field Offset: %#x\n", pe_sig_offset+6)
	// 这是Signature header (0x7c) + coff (20bytes) + oh32 (224bytes)的结束位置
	fmt.Printf("[+] Section Table Offset: %#x\n", pe_sig_offset+0xF8)
	// 打印Optional Header
	fmt.Println("[-----Optional Header-----]")
	fmt.Printf("[+] Entry Point: %#x\n", oh32.AddressOfEntryPoint)
	fmt.Printf("[+] ImageBase: %#x\n", oh32.ImageBase)
	fmt.Printf("[+] Size of Image: %#x\n", oh32.SizeOfImage)
	fmt.Printf("[+] Sections Alignment: %#x\n", oh32.SectionAlignment)
	fmt.Printf("[+] File Alignment: %#x\n", oh32.FileAlignment)
	fmt.Printf("[+] Characteristics: %#x\n", pefile.FileHeader.Characteristics)
	fmt.Printf("[+] Size of Headers: %#x\n", oh32.SizeOfHeaders)
	fmt.Printf("[+] Checksum: %#x\n", oh32.CheckSum)
	fmt.Printf("[+] Machine: %#x\n", pefile.FileHeader.Machine)
	fmt.Printf("[+] Subsystem: %#x\n", oh32.Subsystem)
	fmt.Printf("[+] DLLCharacteristics: %#x\n", oh32.DllCharacteristics)
	// 打印Data Directory
	fmt.Println("[-----Data Directory-----]")
	var winnt_datadirs = []string{
		"IMAGE_DIRECTORY_ENTRY_EXPORT",
		"IMAGE_DIRECTORY_ENTRY_IMPORT",
		"IMAGE_DIRECTORY_ENTRY_RESOURCE",
		"IMAGE_DIRECTORY_ENTRY_EXCEPTION",
		"IMAGE_DIRECTORY_ENTRY_SECURITY",
		"IMAGE_DIRECTORY_ENTRY_BASERELOC",
		"IMAGE_DIRECTORY_ENTRY_DEBUG",
		"IMAGE_DIRECTORY_ENTRY_COPYRIGHT",
		"IMAGE_DIRECTORY_ENTRY_GLOBALPTR",
		"IMAGE_DIRECTORY_ENTRY_TLS",
		"IMAGE_DIRECTORY_ENTRY_LOAD_CONFIG",
		"IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT",
		"IMAGE_DIRECTORY_ENTRY_IAT",
		"IMAGE_DIRECTORY_ENTRY_DELAY_IMPORT",
		"IMAGE_DIRECTORY_ENTRY_COM_DESCRIPTOR",
		"IMAGE_DIRECTORY_ENTRY_NUMBEROF_DIRECTORY_ENTRIES",
	}
	for idx, directory := range oh32.DataDirectory {
		fmt.Printf("[!] Data Directory: %s\n", winnt_datadirs[idx])
		fmt.Printf("[+] Image Virtual Address: %#x\n", directory.VirtualAddress)
		fmt.Printf("[+] Image Size: %#x\n", directory.Size)
	}
	fmt.Println("[-----Section Table-----]")
	for _, section := range pefile.Sections {
		fmt.Println("[+] --------------------")
		fmt.Printf("[+] Section Name: %s\n", section.Name)
		fmt.Printf("[+] Section Characteristics: %#x\n", section.Characteristics)
		fmt.Printf("[+] Section Virtual Size: %#x\n", section.VirtualSize)
		fmt.Printf("[+] Section Virtual Offset: %#x\n", section.VirtualAddress)
		fmt.Printf("[+] Section Raw Size: %#x\n", section.Size)
		fmt.Printf("[+] Section Raw Offset to Data: %#x\n", section.Offset)
		fmt.Printf("[+] Section Append Offset (Next Section): %#x\n", section.Offset+section.Size)
	}

	// s := pefile.Section(".text")
	// fmt.Printf("%v", *s)

	// "Section Table Offset" + (40bytes * number of sections)

}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// 修改镜像时需要考虑的事项
// 1. 舍入最后一个节
// 2. 在PE Header中添加stub
// 3. 增加节的数量
// 4. 更改IMAGE_SIZE
// 5. 在内存代码空洞中添加字节
