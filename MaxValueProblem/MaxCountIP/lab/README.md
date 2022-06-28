## lab 题目：海量日志数据，提取出访问次数最多的IP
### 任务说明：
**Task 1: 完善 SplitBigFile 函数**<br>
SplitBigFile 函数实现将源文件（大文件）分割成多个小文件，也就是我们要做的第一步：分而治之。<br>
**Task 2: 完善 GetPartMax 函数**<br>
GetPartMax 函数获取每个小文件中出现次数最多的 IP，将该 IP 与计数值保存到代码中提供的 partMaxVal 哈希表中。<br>
**Task 3: 完善 GetMax 函数**<br>
在 GetMax 函数中，我们使用 GetPartMax 得来的 partMaxVal 哈希表与堆排序的方法，返回一个保含出现次数最多的 IP 与它的出现次数的二元组 Item（见代码）。<br>
**测试说明**<br>
当你完成了以上的函数之后，在该 lab 文件夹下执行`go test`即可进行 lab 测试！祝贺您一次通过哦！另外，鼓励大家使用不一样的方法实现 lab。
