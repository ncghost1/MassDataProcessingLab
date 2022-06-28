package MaxCountIPLab

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMaxCountIP(t *testing.T) {
	fmt.Println("---Lab Test: Get the IP with the maximum count---")
	MaxIP := "192.168.0.1"
	MaxCount := 100
	GenerateBigFileForTest(MaxIP, MaxCount)
	fmt.Println("Process: GenerateBigFile is completed.")

	SplitBigFile(NumPartFile)
	fmt.Println("Process: SplitBigFile is completed.")

	GetPartMax()
	fmt.Println("Process: GetPartMax is completed.")

	result := GetMax()
	fmt.Println("Process: GetMax is completed.")
	fmt.Println("Result IP: " + result.IP)
	fmt.Println("Result count: " + strconv.FormatUint(result.count, 10))
	RemoveAndClosePartFile()

	if result.IP != MaxIP {
		t.Errorf("expected IP: %v ,actual:%v", MaxIP, result.IP)
	}
	if result.count != uint64(MaxCount) {
		t.Errorf("expected count: %v ,actual:%v", MaxCount, result.count)
	}
	fmt.Println("---Congratulations: your answer is correct!---")
}
