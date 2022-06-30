package FindTheKthNumberLab

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	UINT64SIZE = uint64(8)
	SAFESIZE   = uint64(64)
	UINT64MAX  = uint64(1<<64 - 1)
)

var (
	srcPath = "SourceFile"
)

func GenerateBigFile(Row int64) {
	// First: clear file (if it exists)
	// 首先清空文件内容（如果文件存在的话）
	err := os.Truncate(srcPath, 0)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(srcPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	// Generate the big file
	// 生成大文件
	for i := int64(0); i < Row; i++ {
		rand.Seed(time.Now().UnixNano())
		val := rand.Uint64() % UINT64MAX
		str := strconv.FormatUint(val, 10)
		_, err := f.WriteString(str + "\n")
		if err != nil {
			panic(err)
		}
	}
}

// GenerateBigFileForTest used in lab testing
func GenerateBigFileForTest(Row, kth int64, number uint64) {
	if Row < 0 || uint64(Row) >= UINT64MAX {
		panic("Row can not greater than " + strconv.FormatUint(UINT64MAX, 10))
	}
	// First: clear file (if it exists)
	// 首先清空文件内容（如果文件存在的话）
	err := os.Truncate(srcPath, 0)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(srcPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	// Generate the big file
	// 生成大文件
	str := strconv.FormatUint(number, 10)
	_, err = f.WriteString(str + "\n")
	if err != nil {
		panic(err)
	}
	for i := int64(0); i < kth-1; i++ {
		rand.Seed(time.Now().UnixNano())
		val := rand.Uint64() % number
		str = strconv.FormatUint(val, 10)
		_, err := f.WriteString(str + "\n")
		if err != nil {
			panic(err)
		}
	}
	for i := int64(0); i < Row-kth; i++ {
		rand.Seed(time.Now().UnixNano())
		val := rand.Uint64()%(UINT64MAX-number) + number
		str = strconv.FormatUint(val, 10)
		_, err := f.WriteString(str + "\n")
		if err != nil {
			panic(err)
		}
	}
}

// QuickSelect is used to quickly find the kth number
// Example(find 5th number): QuickSelect(s,0,int64(len(s)-1),5)
// Challenge: try to implement QuickSelect yourself
// QuickSelect 用来快速寻找到从小到大的第 k 个数
// 示例（找到第5个数）：QuickSelect(s,0,int64(len(s)-1),5)
// 挑战：尝试自己实现 QuickSelect
func QuickSelect(s []uint64, start, end, k int64) uint64 {
	p := s[end]
	l := start
	for i := start; i < end; i++ {
		if s[i] <= p {
			s[l], s[i] = s[i], s[l]
			l++
		}
	}
	s[l], s[end] = s[end], s[l]

	if l == k-1 {
		return s[l]
	} else if l < k-1 {
		return QuickSelect(s, l+1, end, k)
	} else {
		return QuickSelect(s, start, l-1, k)
	}
}

// Task 1: complete the SplitFileByBitThenGetKth function
// 任务 1：完善 SplitFileByBitThenGetKth 函数

// SplitFileByBitThenGetKth recursively get the value of each bit from the highest bit to the lowest bit,
// and then divides the numbers into different files according to the bit value (0 or 1)
// When the size of all numbers in the current file is lower than the memoryLimit,
// you can load all the numbers into memory, find the k-th number and return it
// Parameter:
// f: pointer to current file
// count: the total number of numbers in the current file
// k: the rank of the k-th number in the current file in descending order
// bitPos: current bit to check
// memoryLimit: memory limit size
// SplitFileByBitThenGetKth 递归地从二进制最高位到最低位获取每一位的值，之后按比特值（0或1）将数字分到不同的文件中。
// 当前文件中所有数字的大小总和低于内存限制时，便可以把所有数字加载进内存，查找第 k 个数并返回它
// 参数说明：
// f：指向当前文件的指针
// count：当前文件内的数字总数
// k： 我们要找的第 k 个数在当前文件中的排位（从小到大顺序）
// bitPos： 函数当前要检查的比特位
// memoryLimit： 内存限制大小
func SplitFileByBitThenGetKth(f *os.File, count, k, bitPos int64, memoryLimit uint64) uint64 {
	defer func() {
		fileName := f.Name()
		if fileName != srcPath {
			os.Remove(f.Name())
		}
	}()
	defer func() {
		fileName := f.Name()
		if fileName != srcPath {
			f.Close()
		}
	}()
	if bitPos < 0 || count < 0 {
		// When you get here, please check whether your code is correct
		// or whether the memoryLimit is too small
		// 当你的代码到了这里时，请检查你的代码是否正确，或测试使用的 memoryLimit 是否设置过小
		panic("unexpected error")
	}

	f.Seek(0, 0)
	scanner := bufio.NewScanner(f)

	// We checked that count is non-negative before, so it can be converted to uint64
	// 我们在之前已经对 count 进行了非负数的检查，所以可以放心转换为 uint64
	if uint64(count)*UINT64SIZE+SAFESIZE <= memoryLimit {

	} else {

	}
}

// Task 2: complete the SplitFileByPivotThenGetKth function
// 任务 2：完善 SplitFileByPivotThenGetKth 函数

// SplitFileByPivotThenGetKth compares the number in the file with the value as pivot,
// and divides the number into different partition files according to the comparison results.
// When the size of all numbers in the current file is lower than the memory limit,
// you can load all the numbers into memory, find the k-th number and return it
// We use a bisection algorithm to implement this function
// Parameter:
// f: pointer to current file
// count: the total number of numbers in the current file
// k: the rank of the k-th number in the current file in descending order
// left: minimum value of the current pivot range (left endpoint of bisection algorithm)
// right: maximum value of the current pivot range (right endpoint of bisection algorithm)
// memoryLimit: memory limit size
// SplitFileByPivotThenGetKth 将文件中的数字与作为轴点的值进行比较，根据比较结果的不同将数字分到不同的分区文件中
// 当前文件中所有数字的大小总和低于内存限制时，便可以把所有数字加载进内存，查找第 k 个数并返回它
// 我们使用二分法来实现该函数
// 参数说明：
// f：指向当前文件的指针
// count：当前文件内的数字总数
// k： 我们要找的第 k 个数在当前文件中的排位（从小到大顺序）
// left： 当前轴点值范围的最小值（二分法中的左端点）
// right： 当前轴点值范围的最大值（二分法中的右端点）
// memoryLimit： 内存限制大小
func SplitFileByPivotThenGetKth(f *os.File, count, k int64, left, right, memoryLimit uint64) uint64 {
	defer func() {
		fileName := f.Name()
		if fileName != srcPath {
			os.Remove(f.Name())
		}
	}()
	defer func() {
		fileName := f.Name()
		if fileName != srcPath {
			f.Close()
		}
	}()

	if left >= right || count < 0 {
		panic("unexpected error")
	}

	f.Seek(0, 0)
	scanner := bufio.NewScanner(f)

	// We checked that count is non-negative before, so it can be converted to uint64
	// 我们在之前已经对 count 进行了非负数的检查，所以可以放心转换为 uint64
	if uint64(count)*UINT64SIZE+SAFESIZE <= memoryLimit {

	} else {

	}
}
