package FindTop10NumbersLab

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxLimit = 100000000
)

var (
	bm             bitmap
	NumPartFile    = uint64(100)
	partFile       = make([]*os.File, NumPartFile)
	srcPath        = "SourceFile.txt"
	partPathPrefix = "partFile"
)

func GenerateBigFile(Row uint64) {
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
	for i := uint64(0); i < Row; i++ {
		rand.Seed(time.Now().UnixNano())

		// We control range to ensure that bitmap memory does not exceed the limit on most computers
		// 我们控制范围是为了保证 bitmap 占用的内存在大多数电脑下不会超出内存限制
		val := rand.Uint64() % maxLimit
		str := strconv.FormatUint(val, 10)
		_, err := f.WriteString(str + "\n")
		if err != nil {
			panic(err)
		}
	}
}

// GenerateBigFileForTest used in lab testing
func GenerateBigFileForTest(Row, MinInTop10 uint64, Top10Numbers []uint64) {
	err := os.Truncate(srcPath, 0)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(srcPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(Top10Numbers); i++ {
		str := strconv.FormatUint(Top10Numbers[i], 10)
		_, err := f.WriteString(str + "\n")
		if err != nil {
			panic(err)
		}
	}

	for i := uint64(0); i < Row; i++ {
		rand.Seed(time.Now().UnixNano())
		val := rand.Uint64() % MinInTop10
		str := strconv.FormatUint(val, 10)
		_, err := f.WriteString(str + "\n")
		if err != nil {
			panic(err)
		}
	}
}

type bitmap struct {
	bitmap []byte
}

// grow is used for grow bitmaps' size
// grow 用于扩展 bitmap 的空间
func (b *bitmap) grow(size uint64) {
	if size < uint64(cap(b.bitmap)) {
		// No need to grow capacity
		// 无需扩大容量
		return
	}
	old := b.bitmap
	New := make([]byte, size+1)
	copy(New, old)
	b.bitmap = New
}

// SetBit set the bitPos bit to 1
// SetBit 将第 bitPos 位设置为 1
func (b *bitmap) SetBit(bitPos uint64) {
	bytePos := bitPos / 8
	if bytePos < uint64(cap(b.bitmap)) {
		b.bitmap[bytePos] |= 1 << ((bitPos) % 8)
	} else {
		b.grow(bytePos)
		b.bitmap[bytePos] |= 1 << ((bitPos) % 8)
	}
}

// GetBit get the value at bitPos (0 or 1)
// GetBit 获取第 bitPos 位的值（0或1）
func (b *bitmap) GetBit(bitPos uint64) int {
	bytePos := bitPos / 8
	if bytePos >= uint64(cap(b.bitmap)) {
		return 0
	} else {
		bit := b.bitmap[bytePos] & (1 << ((bitPos) % 8))
		if bit != 0 {
			return 1
		}
		return 0
	}
}

// Task 1: complete the SplitBigFile function
// 任务 1： 完善 SplitBigFile 函数

// SplitBigFile split the source file into partition files
// SplitBigFile 将源文件分割到多个小文件中
func SplitBigFile(NumPartFile uint64) {
	srcFile, err := os.Open(srcPath)
	defer srcFile.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(srcFile)

	// Create partFile
	// 创建小文件
	for i := 0; i < len(partFile); i++ {
		file, err := os.OpenFile(partPathPrefix+strconv.Itoa(i), os.O_CREATE|os.O_RDWR, 0777)
		partFile[i] = file
		if err != nil {
			panic(err)
		}
	}

	// Read SourceFile
	// 读取源文件（大文件）
	for scanner.Scan() {
		// When duplicate numbers are read, you can use bitmap or hash table to remove duplicates
		// 当读取到重复数字时，你可以借助 bitmap 或者 哈希表来去重（已提供 bitmap 的基本功能函数）

	}
}

type Item struct {
	number uint64
}

type ItemHeap []Item

// The following is the implementation of heap.Interface
// 以下是对堆的接口的实现，已保证是大根堆（Pop 出的是最大值）

func (h ItemHeap) Len() int { return len(h) }

func (h ItemHeap) Less(i, j int) bool {
	return h[i].number > h[j].number
}
func (h ItemHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *ItemHeap) Push(val interface{}) {
	*h = append(*h, val.(Item))
}

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var TopHeap = &ItemHeap{}

// Task 2: complete the GetPartTop10 function
// 任务 2： 完善 GetPartTop10 函数

// GetPartTop10 get the top 10 numbers in each partition files and push them into TopHeap
// GetPartTop10 获取每个小文件中的 top 10，并把它们加入 TopHeap 中（你也可以使用别的方式来排序）
func GetPartTop10() {
	for i := 0; i < len(partFile); i++ {
		f := partFile[i]
		// Reset offset to 0 （重置文件指针偏移量为0，即回到起始位置）
		_, err := f.Seek(0, 0)
		if err != nil {
			panic(err)
		}

	}
}

// Task 3: complete the GetTop10 function
// 任务 3： 完善 GetTop10 函数

// GetTop10 get the real top 10 and return the slice containing them
// GetTop10 获取真正的 top 10，并返回包含它们的切片
func GetTop10() []uint64 {

}

func RemoveAndClosePartFile() {
	for i := 0; i < len(partFile); i++ {
		partFile[i].Close()
		err := os.Remove(partPathPrefix + strconv.Itoa(i))
		if err != nil {
			panic(err)
		}
	}
}
