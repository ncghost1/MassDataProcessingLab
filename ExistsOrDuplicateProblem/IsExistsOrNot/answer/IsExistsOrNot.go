package IsExistsOrNotAnswer

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	bm      bitmap
	srcPath = "SourceFile.txt"
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
	rand.Seed(time.Now().UnixMilli())
	for i := uint64(0); i < Row; i++ {
		str := strconv.FormatUint(i, 10)
		_, err := f.WriteString(str + "\n")
		if err != nil {
			panic(err)
		}
		i += rand.Uint64() % 2
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
	if bytePos >= uint64(cap(b.bitmap)) || bitPos < 0 {
		return 0
	} else {
		bit := b.bitmap[bytePos] & (1 << ((bitPos) % 8))
		if bit != 0 {
			return 1
		}
		return 0
	}
}

// ReadData read the data in big file,and store the data in bitmap
// ReadData 读取大文件中的数据，并将这些数据存储在 bitmap 中
func ReadData() {
	f, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		raw := scanner.Text()
		data, err := strconv.ParseUint(raw, 10, 64)
		if err != nil {
			panic(err)
		}
		bm.SetBit(data)
	}
}

func ExistsQuery(num uint64) bool {
	if bm.GetBit(num) == 0 {
		return false
	}
	return true
}
