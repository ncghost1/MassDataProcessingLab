package MaxCountIPLab

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMaxCountIP(t *testing.T) {
	fmt.Println("---Lab Test: Get the IP with the maximum count---")
	MaxIP := GenerateIP()
	// Only 10000 IPs are generated for fast testing
	// 为了快速测试，所以只生成 10000 条 IP，当然你也可以自己修改
	MaxCount := 10000
	GenerateBigFileForTest(MaxIP, MaxCount)
	fmt.Println("Process: GenerateBigFile is completed.")

	defer RemoveAndClosePartFile()
	SplitBigFile(NumPartFile)
	fmt.Println("Process: SplitBigFile is completed.")

	GetPartMax()
	fmt.Println("Process: GetPartMax is completed.")

	result := GetMax()
	fmt.Println("Process: GetMax is completed.")
	fmt.Println("Result IP: " + result.IP)
	fmt.Println("Result count: " + strconv.FormatUint(result.count, 10))

	if result.IP != MaxIP {
		t.Errorf("expected IP: %v ,actual:%v", MaxIP, result.IP)
	}
	if result.count != uint64(MaxCount) {
		t.Errorf("expected count: %v ,actual:%v", MaxCount, result.count)
	}
	fmt.Println("---Congratulations: your answer is correct!---")
}
