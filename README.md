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