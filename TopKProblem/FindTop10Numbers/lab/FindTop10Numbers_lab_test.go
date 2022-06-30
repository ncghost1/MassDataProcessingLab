package FindTop10NumbersLab

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFindTop10Numbers(t *testing.T) {
	fmt.Println("---Lab Test: Find Top 10 Numbers---")
	Row := uint64(100000)
	Top10Numbers := make([]uint64, 10)
	MinInTop10 := uint64(maxLimit)

	// Generate Top 10 numbers
	// 生成 top 10
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		val := rand.Uint64() % MinInTop10
		if val < 10 {
			val += 10
		}
		flag := true
		for j := 0; j < i; j++ {
			if Top10Numbers[j] == val {
				flag = false
				break
			}
		}
		if flag == false {
			i-- // Regenerate Top10Numbers[i]
			continue
		}
		if val < MinInTop10 {
			MinInTop10 = val
			Top10Numbers[i] = val
		}
	}

	GenerateBigFileForTest(Row, MinInTop10, Top10Numbers)
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
	result := GetTop10()
	fmt.Println("Process: GetTop10 is completed.")
	fmt.Println("Top10 Numbers:")
	for i := 0; i < len(result); i++ {
		if result[i] != Top10Numbers[i] {
			t.Errorf("expected result: %v ,actual: %v", Top10Numbers[i], result[i])
		}
		fmt.Print("NO.", i+1, result[i], "\n")
	}

	fmt.Println("---Congratulations: your answer is correct!---")
}
