## lab 题目：在海量数据中找出从小到大的第 k 个数
### 思路解析：
假设前提是内存不足以存储所有的数据，所以我们依然是选择将大文件分成小文件来处理，关于分法，本次 lab 需要完成如下两种：<br>
Method 1: 二进制位比较。从最高位到最低位依次比较每一位的值，值为0放一个文件中，值为1放另一个文件中，这样分割到越低位时，文件越小。<br>
Method 2: 轴点值比较。这里我们可以使用二分算法，将中间值作为轴点值，令当前文件中的每一个数字与轴点值进行比较，小于等于轴点值的可以放同一个文件，大于轴点值的放到另一个文件。<br>
那么这两个方法都如何往下递归呢？我们这里借助快速排序的思想，以二进制位比较为例：假如原本有 10 亿无符号整数，要找第5亿个数，我们首先考虑最高位，将最高位为0的分到 file_0，为1的分到 file_1，一边分一边给两个文件分到的数字计数。分完之后假设 file_0 的计数值是6亿，file_1则是剩下4亿，由于是无符号，最高位为0的 file_0 文件中的数全都比 file_1 要小，则第5亿个数就在 file_0 中，且是 file_0 中从小到大的第5亿个数，则我们向 file_0 递归。如果这里是 file_0 有4亿而 file_1 有6亿的话，那么第5亿个数在 file_1 中，我们需要向 file_1 递归，但是注意！递归到 file_1 时第5亿个数的从小到大的排位变化了，第5亿个数在 file_1 中的排位应是从小到大的第1亿个数。就这样一直分割文件和递归下去，直到文件大小低于内存限制时，我们便可以使用排序或者快速选择将答案数字找出来。<br>

### 任务说明：
**Task 1: 完善 SplitFileByBitThenGetKth 函数**<br>
SplitFileByBitThenGetKth 函数使用二进制位比较的方式分割文件，当文件中所有整数的大小加上64B的安全预留大小不超过内存限制时，即可以使用排序或快速选择（已提供）等方法返回第 k 个数。<br>
**Task 2: 完善 SplitFileByPivotThenGetKth 函数**<br>
SplitFileByPivotThenGetKth 函数使用轴点值比较的方式分割文件，当文件中所有整数的大小加上64B的安全预留大小不超过内存限制时，即可以使用排序或快速选择（已提供）等方法返回第 k 个数。<br>

**测试说明**<br>
当你完成了以上的函数之后，在该 lab 文件夹下执行`go test`即可进行 lab 测试！祝贺您一次通过哦！另外，鼓励大家使用不一样的方法实现 lab。
