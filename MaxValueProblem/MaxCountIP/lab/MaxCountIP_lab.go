package MaxCountIPLab

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	IPMod = 256
)

// NumPartFile: Number of partition file （小文件/分割文件数量）
// partFile: An array that holds pointers to partition files (存放指向小文件的文件指针的数组）
// partMaxVal: A map that stores the maximum count of IP in each partition file and its count （保存每个小文件中出现次数最多的 IP 和它的次数的 map）
// srcPath: Source file path 源文件（大文件）路径
// partPathPrefix: partition path prefix 小文件的路径前缀
var (
	NumPartFile    = uint64(100)
	partFile       = make([]*os.File, NumPartFile)
	partMaxVal     = make(map[string]uint64, NumPartFile)
	srcPath        = "SourceFile.txt"
	partPathPrefix = "partFile"
)

func GenerateIP() string {
	var build strings.Builder
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		num := rand.Intn(IPMod)
		field := strconv.Itoa(num)
		build.WriteString(field)
		build.WriteString(":")
	}
	IP := build.String()
	IP = strings.TrimRight(IP, ":")
	return IP
}

func GenerateBigFile(RowSum int) {
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
	for i := 0; i < RowSum; i++ {
		str := GenerateIP() + "\n"
		_, err := f.WriteString(str)
		if err != nil {
			panic(err)
		}
	}
}

func GenerateBigFileForTest(maxIP string, maxCount int) {
	// Use in lab test
	err := os.Truncate(srcPath, 0)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(srcPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	for i := 0; i < maxCount; i++ {
		str := maxIP + "\n"
		_, err := f.WriteString(str)
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < maxCount-1; i++ {
		str := GenerateIP() + "\n"
		_, err := f.WriteString(str)
		if err != nil {
			panic(err)
		}
	}
}

// Task 1: Complete the SplitBigFile function
// 任务 1： 完善 SplitBigFile 函数

// SplitBigFile Splitting source files into multiple partition files
// SplitBigFile 实现将源文件（大文件）分割成多个小文件
func SplitBigFile(numPartFile uint64) {
	srcFile, err := os.Open(srcPath)
	defer srcFile.Close()
	if err != nil {
		panic(err)
	}

	// Step1: Create partition files and store these file pointers in the partFile array.
	// Note that the path of partition files is the same as the path in RemoveAndClosePartFile
	// 第一步： 创建小文件，并把这些文件指针存放到 partFile 数组中
	// 注意这些文件路径的写法要和 RemoveAndClosePartFile 中的保持一致

	// Step2: Read SourceFile and split into multiple partition files , tip: using hash mapping.
	// 第二步： 读取源文件（大文件）并分割成多个小文件，提示：使用哈希映射

}

// Task 2: Complete the GetPartMax function
// 任务 2： 完善 GetPartMax 函数

// GetPartMax get the IP with the maximum count in each partition file,
// the IP and its count will be saved to the 'partMaxVal'.
// GetPartMax 获取每个小文件中出现次数最多的 IP，该 IP 与计数保存到 partMaxVal 哈希表中
func GetPartMax() {

	// Step1: Read each partition file
	// 第一步： 读取每个小文件
	for i := 0; i < len(partFile); i++ {
		tempMap := make(map[string]uint64)
		f := partFile[i]
		// Reset offset to 0 （重置文件指针偏移量为0，即回到起始位置）
		_, err := f.Seek(0, 0)
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(f)

		// Step2: Scan the current partition file to get the MaxIP and MaxCount
		// Step2：扫描当前小文件，获取该文件中出现次数最多的 IP 和 出现次数
		for scanner.Scan() {

		}
		// Step3: Store the IP with the maximum count and its count in this file into partMaxVal
		// 第三步： 将该文件计数值最大的 IP 与它的计数值保存到 partMaxVal 中

	}
}

type Item struct {
	IP    string
	count uint64
}

type ItemHeap []Item

// The following is the implementation of heap.Interface
// 以下是对堆的接口的实现

func (h ItemHeap) Len() int { return len(h) }

func (h ItemHeap) Less(i, j int) bool {
	if h[i].count != h[j].count {
		return h[i].count > h[j].count
	} else {
		return h[i].IP > h[j].IP
	}
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

// Task 3: Complete the GetMax function
// 任务 3： 完善 GetMax 函数

// GetMax Here we use the partMaxVal and heap sort to find the IP with the largest count.
func GetMax() Item {
	h := &ItemHeap{}

	// Step1: heap sort 
	// 第一步：使用堆排序找出最大值（上面已实现堆排序相关接口，当然你也可以使用别的方式）
	
	// Step2: get the Item(IP,count) which has the maximum count
	// 第二步： 获取保存了出现次数最多的 IP 与它的出现次数的二元组 Item

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
