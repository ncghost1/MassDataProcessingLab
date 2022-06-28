package SearchUniqueNumbersLab

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

// GenerateBigFileForTest used in lab testing
func GenerateBigFileForTest(Row uint64, UniqueNum []uint32) {
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
	for i := uint64(0); i < Row; i += 2 {
		rand.Seed(time.Now().UnixNano())
		num := rand.Uint32()
		str := strconv.FormatUint(uint64(num), 10)
		_, err := f.WriteString(str + "\n")
		if err != nil {
			panic(err)
		}
		_, err = f.WriteString(str + "\n")
		if err != nil {
			panic(err)
		}
	}
	for i := 0; i < len(UniqueNum); i++ {
		str := strconv.FormatUint(uint64(UniqueNum[i]), 10)
		_, err := f.WriteString(str + "\n")
		if err != nil {
			panic(err)
		}
	}
}

type TwoBitmap struct {
	bitmap []byte
}

// Task 1: Complete the functions of TwoBitmap(grow, SetBit, GetBit)
// 任务 1: 完成 TwoBitmap 的 grow, SetBit, GetBit 函数

// grow is used for grow TwoBitmaps' size
// grow 用于扩展 TwoBitmap 的空间
func (b *TwoBitmap) grow(size uint64) {

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

}

// Set2Bit Set the value of bitPos*2 bit and bitPos*2+1 bit
// Please see the comments of Get2Bit for how to set
// Set2Bit 设置第 bitPos*2 位和第 bitPos*2+1 位的值
// 如何设置请看 Get2Bit 函数的注释
func (b *TwoBitmap) Set2Bit(bitPos uint64) {

}

// Task 2: Complete the ReadData function
// 任务 2： 完善 ReadData 函数

// ReadData read the data in big file,and store the data in TwoBitmap
// ReadData 读取大文件中的数据，并将这些数据存储在 TwoBitmap 中
func ReadData() {
	f, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Read data and store them in TwoBitmap
		// 读取数据并将它们存储在 TwoBitmap 中

	}
}

// Task 3: Complete the SearchUniqueNumbers function
// 任务 3： 完善 SearchUniqueNumbers 函数

// SearchUniqueNumbers return a slice containing all unique integers(exists,but only one)
// SearchUniqueNumbers 返回一个包含所有只出现过一次的整数的切片
func SearchUniqueNumbers() []uint32 {

}
