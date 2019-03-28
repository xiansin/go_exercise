# GO 练习

### 运算符优先级
优先级|运算符
---|---
7|^ !
6|* / % << >> & &^
5|+ - \| ^
4|== != < <= >= >
3|<-
2|&&
1|\|\|

### 格式化说明符
    格式化输出
    %d      格式化整数
    %0d     输出定长的整数，其中开头数字 0 是必须的
	%x      和 %X 用于格式化 16 进制表示的数字
	%g      用于格式化浮点型
	%f      输出浮点数
	%e      输出科学计数法
	%v      表示复数 如果只需要表示复数的一部分可以使用 %f
	%n.mg   用于表示数字 n 并精确到小数点后 m 位，除了使用 g 之外，还可以使用 e 和 f,eg：%5.2e
    %c      表示字符
    %b      表示补码
    %p      指针格式符

> 当和字符串配合使用时 `%v` 和 `%d` 会输出用于表示该字符的整数；`%U` 输出格式为`U+hhhh`的字符串 
### string包

    // 是否存在某一个前缀
    strings.HasPerFix(str,s) boolean
    // 是否存在某一个后缀
    strings.HasSufFix(str,s) boolean
    // 是否存在某个字符
    string.Contains(str,s) boolean
    // 某个字符首次出现的位置
    strings.Index(str,s) int
    // 某个字符最后一次出现的位置
    strings.LastIndex(str,s) int
    // 查询非 ASCII 编码的字符在父字符串中的位置
    strings.IndexRune(str,ASCII) int
    // 计算某个字符出现非重叠次数
    strings.Count(str,s) int
    // 重复字符串次数
    strings.Repeat(str,int) string
    // 替换字符串 n为替换次数 n为 -1 全部替换
    strings.Replace(str,old,new,n) string
    // 大小写转换
    strings.ToLower() string
    strings.ToUpper() string
    // 去除指定字符
    strings.TrimSpace(str) string
    strings.Trim(str,s) string
    strings.TrimLeft(str,s) string
    strings.TrimRight(str,s) sting
    // 分隔字符串
    strings.Fields(str) array 以空白符进行分隔
    strings.Split(str,s) array 以指定字符进行分隔
    // 链接字符串
    strings.Join([],s) string
    // 读取字符串 并指向字符串指针
    strings.Reader()
    strings.Read()
    strings.ReadByte()
    strings.ReadRune()

### strconv 包
    
    strconv
    
### 补码计算
    负数的二进制
    -60
    先求出60的
        原码:0011 1100
        反码:1100 0011 
        补码:1100 0100 // 反码+1 
    补码即为：-60 的二进制
