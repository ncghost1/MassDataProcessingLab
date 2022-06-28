package SearchUniqueNumbersLab

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestUniqueNumber(t *testing.T) {
	fmt.Println("---Lab Test: Search Unique Numbers---")
	Row := uint64(100)
	UniqueNumbersCount := 3
	UniqueNumbers := make([]uint32, UniqueNumbersCount)
	rand.Seed(time.Now().Unix())
	num := rand.Uint32()
	UniqueNumbers[0] = num
	UniqueNumbers[1] = num + 1
	UniqueNumbers[2] = num + 2
	GenerateBigFileForTest(Row, UniqueNumbers)
	fmt.Println("Process: GenerateBigFile is completed.")

	// When we read the data in the big file, we store the data in TwoBitmap at the same time.
	// 当我们读取大文件中的数据时，同时将数据存储在 TwoBitmap 中
	ReadData()
	fmt.Println("Process: ReadData is completed.")

	// Use TwoBitmap to find the integers that only appears once
	// 利用在 ReadData 建立好的 TwoBitmap 来查找只出现过一次的整数
	actual := SearchUniqueNumbers()
	fmt.Println("Process: SearchUniqueNumbers is completed.")

	fmt.Println("Unique numbers:")
	fmt.Println(actual)
	if len(actual) != len(UniqueNumbers) {
		t.Errorf("length of actual is not equal to UniqueNumbers,expected: %v,actual: %v", len(UniqueNumbers), len(actual))
	}

	set := make([]bool, UniqueNumbersCount)
	for i := 0; i < len(actual); i++ {
		flag := false
		for j := 0; j < len(UniqueNumbers); j++ {
			if set[j] == false && actual[i] == UniqueNumbers[j] {
				flag = true
				set[j] = true
				break
			}
		}
		if flag == false {
			t.Errorf("Wrong value: %v", actual[i])
		}
	}
	fmt.Println("---Congratulations: your answer is correct!---")
}
