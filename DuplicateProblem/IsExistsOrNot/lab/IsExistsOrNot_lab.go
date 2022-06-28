package IsExistsOrNotLab

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

// GenerateBigFileForTest used in lab testing
func GenerateBigFileForTest(Row, NotExistsNum uint64) {
	err := os.Truncate(srcPath, 0)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(srcPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().UnixMilli())
	for i := uint64(0); i < Row; i++ {
		if i == NotExistsNum {
			continue
		}
		str := strconv.FormatUint(i, 10)
		_, err := f.WriteString(str + "\n")
		if err != nil {
			panic(err)
		}
	}
}

type bitmap struct {
	bitmap []byte
}

// Task 1: Complete the functions of bitmap(grow, SetBit, GetBit)
// 任务 1: 完成 bitmap 的 grow, SetBit, GetBit 函数

// grow is used for grow bitmaps' size
// grow 用于扩展 bitmap 的空间
func (b *bitmap) grow(size uint64) {

}

// SetBit set the bitPos bit to 1
// SetBit 将第 bitPos 位设置为 1
func (b *bitmap) SetBit(bitPos uint64) {

}

// GetBit get the value at bitPos (0 or 1)
// GetBit 获取第 bitPos 位的值（0或1）
func (b *bitmap) GetBit(bitPos uint64) int {

}

// Task 2: Complete the ReadData function
// 任务 2： 完善 ReadData 函数

// ReadData read the data in big file,and store the data in bitmap
// ReadData 读取大文件中的数据，并将这些数据存储在 bitmap 中
func ReadData() {
	f, err := os.Open(srcPath)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Read data and store them in bitmap
		// 读取数据并将它们存储在 bitmap 中

	}
}

func ExistsQuery(num uint64) bool {
	if bm.GetBit(num) == 0 {
		return false
	}
	return true
}
