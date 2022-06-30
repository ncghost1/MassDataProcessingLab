# Mass Data Processing Lab 海量数据处理 Lab
海量数据处理是很经典且常见的一类面试题目，网络上也有很多这类题目的情景与对应方案，但大多都是在文字说明而没有人真的用代码去实现。我对海量数据处理这一类的题目很感兴趣，所以在这里我会实现一些海量数据处理的典型题目的代码 Demo 供大家参考，也欢迎大家发现问题后提 issue 或 pr 指正，当然也支持你提出一个新场景并贡献对应的 Go 代码实现！，要带上 answer、 lab 以及对应的 test 哦。😃😃<br>
### 2022/6/28：<br>
今天做了一个决定，就是把“海量数据处理Demo”更名为“海量数据处理lab”，因为我不希望大家只是进来看看我实现的方式后就过了，因为这类问题还有更多的实现方法，就算方法相同，不同的人写出来的代码细节上也可能不一样。或许你也想尝试亲手做一做呢？所以我将把这个仓库改造成 lab，每一类问题分为 lab 和 answer 文件夹，lab 即是你要去完善并使用配套的 test 来验证是否通过的代码，answer 即是我原本实现的 demo，可以作为答案参考。欢迎大家踊跃参与亲身实践！😀😀

## 题目列表：
### 一、 海量数据中的最值问题：<br>
| 题目 | lab 链接 | answer 链接 |
| - | :-: | :-: |
| 1. 海量日志数据，提取出访问次数最多的IP： | [MaxCountIP_lab](https://github.com/ncghost1/MassDataProcessingLab/tree/main/MaxValueProblem/MaxCountIP/lab) | [MaxCountIP_answer](https://github.com/ncghost1/MassDataProcessingLab/tree/main/MaxValueProblem/MaxCountIP/answer) |

### 二、海量数据的某个数据是否存在或重复存在的问题：<br>
| 题目 | lab 链接 | answer 链接 |
| - | :-: | :-: |
| 1. 给出海量的不重复整数，之后指定一个数，快速判断指定数字是否存在 | [IsExistsOrNot_lab](https://github.com/ncghost1/MassDataProcessingLab/tree/main/ExistsOrDuplicateProblem/IsExistsOrNot/lab) | [IsExistsOrNot_answer](https://github.com/ncghost1/MassDataProcessingLab/tree/main/ExistsOrDuplicateProblem/IsExistsOrNot/answer) |
| 2. 海量整数中找出不重复的整数 | [SearchUniqueNumbers_lab](https://github.com/ncghost1/MassDataProcessingLab/tree/main/ExistsOrDuplicateProblem/SearchUniqueNumbers/lab) | [SearchUniqueNumbers_answer](https://github.com/ncghost1/MassDataProcessingLab/tree/main/ExistsOrDuplicateProblem/SearchUniqueNumbers/answer) |

### 三、海量数据的 top K 问题：<br>
| 题目 | lab 链接 | answer 链接 |
| - | :-: | :-: |
| 1. 如何在海量数据中找出最大的 10 个数： | [FindTop10Numbers_lab](https://github.com/ncghost1/MassDataProcessingLab/tree/main/TopKProblem/FindTop10Numbers/lab) | [FindTop10Numbers_answer](https://github.com/ncghost1/MassDataProcessingLab/tree/main/TopKProblem/FindTop10Numbers/answer) |

### 四、海量数据的 No.K 问题：<br>
| 题目 | lab 链接 | answer 链接 |
| - | :-: | :-: |
| 1. 在海量数据中找出从小到大第 k 个数： | [FindTheKthNumber_lab](https://github.com/ncghost1/MassDataProcessingLab/tree/main/NoKProblem/FindTheKthNumber/lab)  | [FindTheKthNumber_answer](https://github.com/ncghost1/MassDataProcessingLab/tree/main/NoKProblem/FindTheKthNumber/answer)  |
