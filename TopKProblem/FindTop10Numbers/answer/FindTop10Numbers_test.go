package FindTop10NumbersAnswer

import (
	"fmt"
	"testing"
)

// This is not actually a test, but a reference to the implementation steps
// 这并不是测试，而是实现步骤的参考
func TestFindTop10Numbers(t *testing.T) {
	fmt.Println("---Find Top 10 Numbers---")

	Row := uint64(1000000)
	GenerateBigFile(Row)
	fmt.Println("Process: GenerateBigFile is completed.")

	// Step 1: split source file to each partition file
	// 第一步：分而治之
	defer RemoveAndClosePartFile()
	SplitBigFile(NumPartFile)
	fmt.Println("Process: SplitBigFile is completed.")

	// Step 2: get the top 10 numbers in each partition file,
	// the numbers will be saved to the 'TopHeap'.
	// 第二步：获取每个小文件中最大的十个数字，这些数字保存到 TopHeap 堆（优先队列）中
	GetPartTop10()
	fmt.Println("Process: GetPartTop10 is completed.")

	// Step 3: use TopHeap to get the Top 10 numbers
	// 第三步：最后使用 TopHeap 获取最大的十个数字
	res := GetTop10()
	fmt.Println("Process: GetTop10 is completed.")
	fmt.Println("Top10 Numbers:")
	for i := 0; i < len(res); i++ {
		fmt.Print("NO.", i+1, res[i], "\n")
	}
}
