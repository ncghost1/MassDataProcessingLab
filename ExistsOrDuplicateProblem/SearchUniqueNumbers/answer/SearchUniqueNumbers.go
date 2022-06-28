package SearchUniqueNumbersAnswer

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	bm      TwoBitmap
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
	for i := uint64(0); i < Row; i++ {
		rand.Seed(time.Now().UnixNano())
		num := rand.Uint32()
		str := strconv.FormatUint(uint64(num), 10)
		_, err := f.WriteString(str + "\n")
		if err != nil {
			panic(err)
		}
	}
}

type TwoBitmap struct {
	bitmap []byte
}

// grow is used for grow TwoBitmaps' size
// grow 用于扩展 TwoBitmap 的空间
func (b *TwoBitmap) grow(size uint64) {
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

// Get2Bit get the value of the two bit binary number composed of bitPos*2 (low bit) and bitPos*2+1 (high bit)
// the Binary values represent meaning:
// 00: not exist
// 01: exist, but only one
// 11: exist, but more than one
// 10: Meaningless
// Get2Bit 获取第 bitPos*2（低位） 与 bitPos*2+1（高位） 组成的两位二进制数的值
// 该二进制值表示的意义：
// 00: 不存在
// 01: 存在，但只有一个
// 11: 存在，但多于一个
// 10: 无意义
func (b *TwoBitmap) Get2Bit(bitPos uint64) int {
	bitPos *= 2
	bytePos := bitPos / 8
	if bytePos >= uint64(cap(b.bitmap)) {
		return 0
	} else {
		result := 0
		bit1 := b.bitmap[bytePos] & (1 << ((bitPos) % 8))
		if bit1 > 0 {
			bit1 = 1
		}
		bit2 := b.bitmap[bytePos] & (1 << ((bitPos + 1) % 8))
		if bit2 > 0 {
			bit2 = 1 << 1
		}
		result |= int(1 & bit1)
		result |= int(1 << 1 & bit2)
		return result
	}
}

// Set2Bit Set the value of bitPos*2 bit and bitPos*2+1 bit
// Please see the comments of Get2Bit for how to set
// Set2Bit 设置第 bitPos*2 位和第 bitPos*2+1 位的值
// 如何设置请看 Get2Bit 函数的注释
func (b *TwoBitmap) Set2Bit(bitPos uint64) {
	bitPos *= 2
	bytePos := bitPos / 8
	if bytePos < uint64(cap(b.bitmap)) {
		val := b.Get2Bit(bitPos / 2)
		if val == 0 {
			b.bitmap[bytePos] |= 1 << ((bitPos) % 8) // 00 -> 01
		} else if val == 1 {
			b.bitmap[bytePos] |= 1 << ((bitPos + 1) % 8) // 01 -> 11
		}
	} else {
		b.grow(bytePos)
		b.bitmap[bytePos] |= 1 << ((bitPos) % 8) // 00 -> 01
	}
}

// ReadData read the data in big file,and store the data in TwoBitmap
// ReadData 读取大文件中的数据，并将这些数据存储在 TwoBitmap 中
func ReadData() {
	f, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Read data and store them in bitmap
		// 读取数据并将它们存储在 bitmap 中
		raw := scanner.Text()
		data, err := strconv.ParseUint(raw, 10, 64)
		if err != nil {
			panic(err)
		}
		bm.Set2Bit(data)
	}
}

// SearchUniqueNumbers return a slice containing all unique integers(exists,but only one)
// SearchUniqueNumbers 返回一个包含所有只出现过一次的整数的切片
func SearchUniqueNumbers() []uint32 {
	var result []uint32
	f, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Read data and store them in TwoBitmap
		// 读取数据并将它们存储在 TwoBitmap 中
		raw := scanner.Text()
		data, err := strconv.ParseUint(raw, 10, 64)
		if err != nil {
			panic(err)
		}
		val := bm.Get2Bit(data)
		if val == 1 {
			// Here we can safely convert uint64 to uint32
			// because we ensure that all integers are uint32 when generating source files
			// 这里我们可以放心将 uint64 强制转换为 uint32，是因为我们生成源文件时保证了整数都是 uint32 类型的
			result = append(result, uint32(data))
		}
	}
	return result
}
