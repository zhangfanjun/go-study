package structure

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func ReaderAndWriterTest() {
	ReaderTest()
}

func ReaderTest() {
	fmt.Println("字符串转为字节数组，然后用bytes将字节数组转为i/o reader流，然后转为缓存流")
	fmt.Println("敲入newReader会发现有很多创建io流的方式，bytes，strings，zip，bar，gzip，csv")
	fmt.Println("操作 Reader 对象的方法共有 11 个，分别是 Read()、ReadByte()、ReadBytes()、ReadLine()、ReadRune ()、ReadSlice()、ReadString()、UnreadByte()、UnreadRune()、Buffered()、Peek()")

	fmt.Println("------------------------------------------------------------------------------------")
	str := "这是输入内容,go的reader"
	data := []byte(str)
	bytesReader := bytes.NewReader(data)
	reader1 := bufio.NewReader(bytesReader)
	var result [128]byte
	n, err := reader1.Read(result[:])
	i, err := reader1.Read(result[n:])
	fmt.Println("这里采用Read()，读到字节数组中，返回当前的位移：", string(result[:n]), n, err)
	fmt.Println("第二次读，每次读返回各自读取的位数，而不会得到一个总的下边数据：", string(result[:i+n]), i, err)

	fmt.Println("------------------------------------------------------------------------------------")
	stringReader := strings.NewReader(str)
	reader2 := bufio.NewReader(stringReader)
	b, err := reader2.ReadByte()
	fmt.Println("这里采用ReadByte(),输出一个字节，但是中文是两个字节，所以输出是乱码的：", string(b), err)

	fmt.Println("------------------------------------------------------------------------------------")
	stringReader = strings.NewReader(str)
	reader2 = bufio.NewReader(stringReader)
	var delim byte = ',' //注意这里是英文的字符，如果是中文的字符是占用两个字节的
	readBytes, err := reader2.ReadBytes(delim)
	fmt.Println("ReadBytes()方法的功能是读取数据直到遇到第一个分隔符\"delim\"：", string(readBytes), err)

	fmt.Println("------------------------------------------------------------------------------------")
	stringReader = strings.NewReader(str)
	reader2 = bufio.NewReader(stringReader)
	line, isPrefix, err := reader2.ReadLine()
	fmt.Println("ReadLine() 是一个低级的用于读取一行数据的方法，大多数调用者应该使用 ReadBytes('') 或者 ReadString('')。\nReadLine 返回一行，不包括结尾的回车字符，\n如果一行太长（超过缓冲区长度），参数 isPrefix 会设置为 true 并且只返回前面的数据，剩余的数据会在以后的调用中返回。")
	fmt.Println("ReadLine()读取一行：", string(line), isPrefix, err)

	fmt.Println("------------------------------------------------------------------------------------")
	stringReader = strings.NewReader(str)
	reader := bufio.NewReader(stringReader)
	r, size, err := reader.ReadRune()
	fmt.Println("ReadRune() 方法的功能是读取一个 UTF-8 编码的字符，并返回其 Unicode 编码和字节数。如果编码错误，ReadRune 只读取一个字节并返回 unicode.ReplacementChar(U+FFFD) 和长度 1")
	fmt.Println("ReadRune()每次读取一个字符，不管是英文还是中文的字符", string(r), size, err)

}
