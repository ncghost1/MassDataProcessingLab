package SearchUniqueNumbersAnswer

import (
	"fmt"
	"testing"
)

// This is not actually a test, but a reference to the implementation steps
// If you want to test the answer code, you can copy the code to lab code and run lab test
// 这并不是测试，而是实现步骤的参考
// 若你要测试 answer 代码，你可以将 answer 代码复制到 lab 代码中，并运行 lab 的测试
func TestUniqueNumber(t *testing.T) {
	fmt.Println("---Search Unique Numbers---")
	Row := uint64(1000000)
	GenerateBigFile(Row)
	fmt.Println("Process: GenerateBigFile is completed.")
	// When we read the data in the big file, we store the data in TwoBitmap at the same time.
	// 当我们读取大文件中的数据时，同时将数据存储在 TwoBitmap 中
	ReadData()
	fmt.Println("Process: ReadData is completed.")

	// Use TwoBitmap to find the integers that only appears once
	// 利用在 ReadData 建立好的 TwoBitmap 来查找只出现过一次的整数
	res := SearchUniqueNumbers()
	fmt.Println("Process: SearchUniqueNumbers is completed.")

	fmt.Println("Unique numbers:")
	fmt.Println(res)
}
