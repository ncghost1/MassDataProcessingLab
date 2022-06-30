package FindTheKthNumberAnswer

import (
	"fmt"
	"os"
	"testing"
)

// This is not actually a test, but a reference to the implementation steps
// If you want to test the answer code, you can copy the code to lab code and run lab test
// 这并不是测试，而是实现步骤的参考
// 若你要测试 answer 代码，你可以将 answer 代码复制到 lab 代码中，并运行 lab 的测试
func TestFindTheKthNumber(t *testing.T) {
	fmt.Println("---Find the Kth number---")
	Row := int64(1000000)
	kth := Row / 2
	bitPos := int64(64)
	memoryLimit := UINT64SIZE * 2000
	left, right := uint64(0), UINT64MAX
	GenerateBigFile(Row)
	fmt.Println("Process: GenerateBigFile is completed.")
	srcFile, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}

	result := SplitFileByBitThenGetKth(srcFile, Row, kth, bitPos, memoryLimit)
	fmt.Println("Process: SplitFileByBitThenGetKth is completed.")
	fmt.Printf("SplitFileByBitThenGetKth: The %vth number is %v.\n", kth, result)

	result = SplitFileByPivotThenGetKth(srcFile, Row, kth, left, right, memoryLimit)
	fmt.Println("Process: SplitFileByPivotThenGetKth is completed.")
	fmt.Printf("SplitFileByPivotThenGetKth: The %vth number is %v.\n", kth, result)
}
