package FindTheKthNumberLab

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestFindTheKthNumber(t *testing.T) {
	fmt.Println("---Lab Test: Find the Kth number---")
	bitPos := int64(64)
	memoryLimit := UINT64SIZE * 1024
	left, right := uint64(0), UINT64MAX
	Row := int64(100000)

	fmt.Printf("Lab Info: The size of all numbers is %v bytes, The memory limit is %v bytes，good luck!\n", uint64(Row)*UINT64SIZE, memoryLimit)

	kth := Row / 2
	expected := uint64(rand.Uint64() % UINT64MAX / 2)

	fmt.Printf("---Test1: Find the %vth(median) number---\n", kth)

	GenerateBigFileForTest(Row, kth, expected)
	fmt.Println("Process: GenerateBigFile is completed.")

	srcFile, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}

	result := SplitFileByBitThenGetKth(srcFile, Row, kth, bitPos, memoryLimit)
	fmt.Println("Process: SplitFileByBitThenGetKth is completed.")
	if result != expected {
		t.Errorf("expected: %v, actual: %v\n", expected, result)
	}
	fmt.Printf("SplitFileByBitThenGetKth: The %vth number is %v\n", kth, result)

	result = SplitFileByPivotThenGetKth(srcFile, Row, kth, left, right, memoryLimit)
	fmt.Println("Process: SplitFileByPivotThenGetKth is completed.")
	if result != expected {
		t.Errorf("expected: %v, actual: %v\n", expected, result)
	}
	fmt.Printf("SplitFileByPivotThenGetKth: The %vth number is %v\n", kth, result)

	kth = 1
	expected = uint64(rand.Uint64() % UINT64MAX / 2)

	fmt.Printf("---Test2: Find the %vth(minimum) number---\n", kth)

	GenerateBigFileForTest(Row, kth, expected)
	fmt.Println("Process: GenerateBigFile is completed.")

	srcFile, err = os.Open(srcPath)
	if err != nil {
		panic(err)
	}

	result = SplitFileByBitThenGetKth(srcFile, Row, kth, bitPos, memoryLimit)
	fmt.Println("Process: SplitFileByBitThenGetKth is completed.")
	if result != expected {
		t.Errorf("expected: %v, actual: %v\n", expected, result)
	}
	fmt.Printf("SplitFileByBitThenGetKth: The %vth number is %v\n", kth, result)

	result = SplitFileByPivotThenGetKth(srcFile, Row, kth, left, right, memoryLimit)
	fmt.Println("Process: SplitFileByPivotThenGetKth is completed.")
	if result != expected {
		t.Errorf("expected: %v, actual: %v\n", expected, result)
	}
	fmt.Printf("SplitFileByPivotThenGetKth: The %vth number is %v\n", kth, result)

	kth = Row
	expected = uint64(rand.Uint64() % UINT64MAX)

	fmt.Printf("---Test3: Find the %vth(maximum) number---\n", kth)

	GenerateBigFileForTest(Row, kth, expected)
	fmt.Println("Process: GenerateBigFile is completed.")

	srcFile, err = os.Open(srcPath)
	if err != nil {
		panic(err)
	}

	result = SplitFileByBitThenGetKth(srcFile, Row, kth, bitPos, memoryLimit)
	fmt.Println("Process: SplitFileByBitThenGetKth is completed.")
	if result != expected {
		t.Errorf("expected: %v, actual: %v\n", expected, result)
	}
	fmt.Printf("SplitFileByBitThenGetKth: The %vth number is %v\n", kth, result)

	result = SplitFileByPivotThenGetKth(srcFile, Row, kth, left, right, memoryLimit)
	fmt.Println("Process: SplitFileByPivotThenGetKth is completed.")
	if result != expected {
		t.Errorf("expected: %v, actual: %v\n", expected, result)
	}
	fmt.Printf("SplitFileByPivotThenGetKth: The %vth number is %v\n", kth, result)

	fmt.Println("---Congratulations: You passed all the tests! ╰(*°▽°*)╯---")
}
