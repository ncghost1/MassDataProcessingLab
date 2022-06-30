package MaxCountIPAnswer

import (
	"fmt"
	"strconv"
	"testing"
)

// This is not actually a test, but a reference to the implementation steps
// If you want to test the answer code, you can copy the code to lab code and run lab test
// 这并不是测试，而是实现步骤的参考
// 若你要测试 answer 代码，你可以将 answer 代码复制到 lab 代码中，并运行 lab 的测试
func TestMaxCountIP(t *testing.T) {
	fmt.Println("---Get the IP with the maximum count---")

	// This process could take a while, you can set RowSum smaller to speed up.
	// 生成大文件会花上一段时间，你可以调小 RowSum 使生成 IP 数量少一些来加速生成
	Row := 1000000 // We generate 1 million IPs for testing 我们生成100万个 IP 进行测试
	GenerateBigFile(Row)
	fmt.Println("Process: GenerateBigFile is completed.")

	// Step 1: split source file to each partition file
	// 第一步：分而治之
	defer RemoveAndClosePartFile()
	SplitBigFile(NumPartFile)
	fmt.Println("Process: SplitBigFile is completed.")

	// Step 2: get the IP with the maximum count in each partition file,
	// the IP and its count will be saved to the 'partMaxVal'.
	// 第二步：获取每个小文件中出现次数最多的 IP，该 IP 与计数保存到 partMaxVal 哈希表中
	GetPartMax()
	fmt.Println("Process: GetPartMax is completed.")

	// Step 3: use partMaxVal and heap sort to get the Item(IP,count) which has the maximum count.
	// Note that if there are multiple max values, we only get one of them
	// 第三步：在上一步我们获得了每个分区文件中出现次数最多的 IP 和它的次数，它们被保存在 partMaxVal 中，
	// 这一步我们使用 partMaxVal 和堆排序获取出现次数最多的 IP 与它的出现次数，并保存在 Item 二元组中返回
	// 注意如果有多个 IP 出现次数都为最大值，我们只返回其中一个即可
	result := GetMax()
	fmt.Println("Process: GetMax is completed.")
	fmt.Println("Result IP: " + result.IP)
	fmt.Println("Result count: " + strconv.FormatUint(result.count, 10))
}
