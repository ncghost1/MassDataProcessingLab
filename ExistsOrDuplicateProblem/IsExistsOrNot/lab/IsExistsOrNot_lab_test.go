package IsExistsOrNotLab

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestIsExistsOrNot(t *testing.T) {
	fmt.Println("---Lab Test: Is the number exists or Not---")
	Row := uint64(10000)
	NotExistsNum := rand.Uint64() % Row
	GenerateBigFileForTest(Row, NotExistsNum)
	fmt.Println("Process: GenerateBigFile is completed.")

	// When we read the data in the big file, we store the data in bitmap at the same time.
	// 当我们读取大文件中的数据时，同时将数据存储在 bitmap 中
	ReadData()
	fmt.Println("Process: ReadData is completed.")

	// Query whether the number exists
	// 查询数字是否存在
	// First test number does not exist
	// 首先测试数字不存在的情况
	Num := NotExistsNum
	actual := ExistsQuery(Num)
	if actual == true {
		t.Errorf("expected result: %v ,actual:%v", false, actual)
	}

	// Test number exists
	// 测试数字存在的情况
	if Num-1 >= 0 {
		Num = Num - 1
		actual = ExistsQuery(Num)
		if actual == false {
			t.Errorf("expected result: %v ,actual:%v", true, actual)
		}
	} else if Num+1 < Row {
		Num = Num + 1
		actual = ExistsQuery(Num)
		if actual == false {
			t.Errorf("expected result: %v ,actual:%v", true, actual)
		}
	}

	fmt.Println("---Congratulations: your answer is correct!---")
}
