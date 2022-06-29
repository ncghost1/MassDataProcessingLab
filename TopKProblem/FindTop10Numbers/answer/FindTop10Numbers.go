package FindTop10NumbersAnswer

import (
	"bufio"
	"container/heap"
	"github.com/dchest/siphash"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxLimit = 100000000
	baseKey  = "MassDataProcess4" // For sipHash （用于sipHash）
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

// SplitBigFile split the source file into partition files
// SplitBigFile 将源文件分割到多个小文件中
func SplitBigFile(NumPartFile uint64) {
	srcFile, err := os.Open(srcPath)
	defer srcFile.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(srcFile)
	// We use sipHash as our hash algorithm
	// 我们使用 sipHash 作为我们要用的 hash 算法
	h := siphash.New([]byte(baseKey))

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
		number := scanner.Text()
		num, err := strconv.ParseUint(number, 10, 64)
		if err != nil {
			panic(err)
		}

		// If the bit has been set,we skip it.
		// 如果 num 对应的 bit 位已经被设置则跳过（去重操作）
		if bm.GetBit(num) == 1 {
			continue
		}

		// set the bit to 1
		// 将 num 对应的 bit 置 1
		bm.SetBit(num)
		// Use number as hash key
		// 将 number 作为哈希用的 key
		_, err = h.Write([]byte(number))
		if err != nil {
			panic(err)
		}
		// get hash
		// 获取读到的 number 对应的哈希值
		hash := h.Sum64() % NumPartFile
		h.Reset() // Reset hash key（重置 key）

		// Append number to the partFile corresponding to the hash
		// 将 number 追加写入到哈希值所对应的小文件上
		_, err = partFile[hash].WriteString(number + "\n")
		if err != nil {
			panic(err)
		}
	}
}

type Item struct {
	number uint64
}

type ItemHeap []Item

// The following is the implementation of heap.Interface
// 以下是对堆的接口的实现

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

// GetPartTop10 get the top 10 numbers in each partition files and push them into TopHeap
// GetPartTop10 获取每个小文件中的 top 10，并把它们加入 TopHeap 中
func GetPartTop10() {
	for i := 0; i < len(partFile); i++ {
		f := partFile[i]
		h := &ItemHeap{}
		// Reset offset to 0 （重置文件指针偏移量为0，即回到起始位置）
		_, err := f.Seek(0, 0)
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(f)

		// Scan the current partition file to get the top 10 number
		// 扫描当前小文件，获取该文件中最大的10个数字
		for scanner.Scan() {
			number := scanner.Text()
			val, err := strconv.ParseUint(number, 10, 64)
			if err != nil {
				panic(err)
			}
			heap.Push(h, Item{number: val})
		}

		for i := 0; i < 10; i++ {
			if h.Len() == 0 {
				break
			}
			val := heap.Pop(h)
			heap.Push(TopHeap, val)
		}
	}
}

// GetTop10 get the real top 10 and return the slice containing them
// GetTop10 获取真正的 top 10，并返回包含它们的切片
func GetTop10() []uint64 {
	var result []uint64

	// Use TopHeap to pop up the top 10 numbers
	// 使用 TopHeap 弹出 top 10
	for i := 0; i < 10; i++ {
		if TopHeap.Len() == 0 {
			break
		}
		raw := heap.Pop(TopHeap)
		item, ok := raw.(Item)
		if !ok {
			panic("type error")
		}
		val := item.number
		result = append(result, val)
	}
	return result
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
