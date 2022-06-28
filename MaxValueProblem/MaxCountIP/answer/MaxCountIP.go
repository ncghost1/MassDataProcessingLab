package MaxCountIPAnswer

import (
	"bufio"
	"container/heap"
	"github.com/dchest/siphash"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	IPMod   = 256
	baseKey = "MassDataProcess1" // For sipHash （用于sipHash）
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
		IP := scanner.Text()

		// Use IP as hash key
		// 将 IP 作为哈希用的 key
		_, err = h.Write([]byte(IP))
		if err != nil {
			panic(err)
		}
		// get hash
		// 获取读到的 IP 对应的哈希值
		hash := h.Sum64() % NumPartFile
		h.Reset() // Reset hash key（重置 key）

		// Append IP to the partFile corresponding to the hash
		// 将 IP 追加写入到哈希值所对应的小文件上
		_, err = partFile[hash].WriteString(IP + "\n")
		if err != nil {
			panic(err)
		}
	}
}

func GetPartMax() {
	for i := 0; i < len(partFile); i++ {
		tempMap := make(map[string]uint64)
		f := partFile[i]
		// Reset offset to 0 （重置文件指针偏移量为0，即回到起始位置）
		_, err := f.Seek(0, 0)
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(f)
		var MaxIP string
		var MaxCount uint64

		// Scan the current partition file to get the MaxIP and MaxCount
		// 扫描当前小文件，获取该文件中出现次数最多的 IP 和 出现次数
		for scanner.Scan() {
			IP := scanner.Text()
			tempMap[IP]++
			if tempMap[IP] > MaxCount {
				MaxIP = IP
				MaxCount = tempMap[IP]
			}
		}
		partMaxVal[MaxIP] = MaxCount
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

func GetMax() Item {
	h := &ItemHeap{}

	// heap sort （我们使用堆排序来找出最大值）
	for i, v := range partMaxVal {
		heap.Push(h, Item{
			IP:    i,
			count: v,
		})
	}

	// get the Item(IP,count) which has the maximum count
	// 获取保存了出现次数最多的 IP 与它的出现次数的二元组 Item
	res, ok := heap.Pop(h).(Item)
	if !ok {
		panic("error: Type Error")
	}
	return res
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
