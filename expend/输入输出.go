package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
输入输出
标准库的fmt包没有使用缓冲区（那就慢），所以考虑到代码的运行效率，我们应当避免直接使用fmt包。
而使用bufio包再封装出两个好用的输入输出函数，就一点毛病没有了：
*/
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer writer.Flush()

	var a, b int
	scanf("%d %d\n", &a, &b)
	printf("%d\n", a+b)
}
