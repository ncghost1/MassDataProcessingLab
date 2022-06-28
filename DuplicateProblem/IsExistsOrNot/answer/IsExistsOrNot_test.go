package IsExistsOrNotAnswer

import (
	"fmt"
	"math/rand"
	"testing"
)

// This is not actually a test, but a reference to the implementation steps
// 这并不是测试，而是实现步骤的参考
func TestIsExistsOrNot(t *testing.T) {
	fmt.Println("---Lab Test: Is the number exists or Not---")
	Row := uint64(1000000)
	GenerateBigFile(Row)
	fmt.Println("Process: GenerateBigFile is completed.")

	// When we read the data in the big file, we store the data in bitmap at the same time.
	// 当我们读取大文件中的数据时，同时将数据存储在 bitmap 中
	ReadData()
	fmt.Println("Process: ReadData is completed.")

	// Query whether the number exists
	// 查询数字是否存在
	Num := uint64(rand.Uint64() % Row)
	if ExistsQuery(Num) == true {
		fmt.Printf("Number: %v is exists.\n", Num)
	} else {
		fmt.Printf("Number: %v is not exists.\n", Num)
	}
}
