package FindTheKthNumberAnswer

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

// QuickSelect is used to quickly find the kth number
// Example(find 5th number): QuickSelect(s,0,int64(len(s)-1),5)
// QuickSelect 用来快速寻找到从小到大的第 k 个数
// 示例（找到第5个数）：QuickSelect(s,0,int64(len(s)-1),5)
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

// SplitFileByBitThenGetKth recursively get the value of each bit from the highest bit to the lowest bit,
// and then divides the numbers into different files according to the bit value (0 or 1).
// SplitFileByBitThenGetKth 递归地从二进制最高位到最低位获取每一位的值，之后按比特值（0或1）将数字分到不同的文件中。
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
		var arr []uint64
		for scanner.Scan() {
			str := scanner.Text()
			val, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				panic(err)
			}
			arr = append(arr, val)
		}
		return QuickSelect(arr, 0, int64(len(arr)-1), k)
	} else {
		cnt0 := int64(0)
		cnt1 := int64(0)

		f0, err := os.OpenFile(f.Name()+"_0", os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			panic(err)
		}
		f1, err := os.OpenFile(f.Name()+"_1", os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			panic(err)
		}

		for scanner.Scan() {
			str := scanner.Text()
			val, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				panic(err)
			}
			if val&(1<<bitPos) == 0 {
				f0.WriteString(str + "\n")
				cnt0++
			} else {
				f1.WriteString(str + "\n")
				cnt1++
			}
		}

		if cnt0 >= k {
			fileName := f1.Name()
			f1.Close()
			os.Remove(fileName)
			return SplitFileByBitThenGetKth(f0, cnt0, k, bitPos-1, memoryLimit)
		} else {
			fileName := f0.Name()
			f0.Close()
			os.Remove(fileName)
			return SplitFileByBitThenGetKth(f1, cnt1, k-cnt0, bitPos-1, memoryLimit)
		}
	}
}

// SplitFileByPivotThenGetKth compares the number in the file with the value as pivot,
// and divides the number into different partition files according to the comparison results.
// Tip: you can use the dichotomy method
// SplitFileByPivotThenGetKth 将文件中的数字与作为轴点的值进行比较，根据比较结果的不同将数字分到不同的分区文件中。
// 提示：你可以使用二分法
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
		var arr []uint64
		for scanner.Scan() {
			str := scanner.Text()
			val, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				panic(err)
			}
			arr = append(arr, val)
		}
		return QuickSelect(arr, 0, int64(len(arr)-1), k)
	} else {
		cnt0 := int64(0)
		cnt1 := int64(0)
		pivot := left + (right-left)/2
		f0, err := os.OpenFile(f.Name()+"_0", os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			panic(err)
		}
		f1, err := os.OpenFile(f.Name()+"_1", os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			panic(err)
		}

		for scanner.Scan() {
			str := scanner.Text()
			val, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				panic(err)
			}
			if val <= pivot {
				f0.WriteString(str + "\n")
				cnt0++
			} else {
				f1.WriteString(str + "\n")
				cnt1++
			}
		}

		if cnt0 >= k {
			fileName := f1.Name()
			f1.Close()
			os.Remove(fileName)
			right = pivot
			return SplitFileByPivotThenGetKth(f0, cnt0, k, left, right, memoryLimit)
		} else {
			fileName := f0.Name()
			f0.Close()
			os.Remove(fileName)
			left = pivot + 1
			return SplitFileByPivotThenGetKth(f1, cnt1, k-cnt0, left, right, memoryLimit)
		}
	}
}
