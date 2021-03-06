## lab 题目：海量日志数据，提取出访问次数最多的IP
### 思路解析：
假设前提是内存不足以存储所有的 IP 地址。<br>
Step 1：首先将存储了海量 IP 地址的大文件分成多个小文件（在代码中分成了100份），通过哈希映射的方式将各个 IP 地址映射到各文件中。<br>
ps: 实际上该分成多少份应从内存进行考虑，应做到即使一个小文件内的所有 IP 地址都不相同，在内存中建立存储这些 IP 的哈希表也不会超过内存限制。<br>
Step 2：当大文件转化成了多份小文件，那么我们便可以采用哈希表来进行频率统计,然后获取每个小文件的最值。<br>
Step 3：将每个小文件的最值进行排序，最终得到访问次数最多的 IP 地址。

### 任务说明：
**Task 1: 完善 SplitBigFile 函数**<br>
SplitBigFile 函数实现将源文件（大文件）分割成多个小文件，也就是我们要做的第一步：分而治之。<br>
**Task 2: 完善 GetPartMax 函数**<br>
GetPartMax 函数获取每个小文件中出现次数最多的 IP，将该 IP 与计数值保存到代码中提供的 partMaxVal 哈希表中。<br>
**Task 3: 完善 GetMax 函数**<br>
在 GetMax 函数中，我们使用 GetPartMax 得来的 partMaxVal 哈希表与堆排序的方法，返回一个保含出现次数最多的 IP 与它的出现次数的二元组 Item（见代码）。<br>

**测试说明**<br>
当你完成了以上的函数之后，在该 lab 文件夹下执行`go test`即可进行 lab 测试！祝贺您一次通过哦！另外，鼓励大家使用不一样的方法实现 lab。
