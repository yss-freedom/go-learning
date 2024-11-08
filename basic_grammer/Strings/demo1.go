package main

import (
	"fmt"
	"strings"
)

/*
在Go语言中，strings包是一个非常重要且功能丰富的标准库，它提供了一系列用于操作字符串的函数。
从基本的字符串查找、替换、分割到高级的比较、格式化等，
strings包几乎涵盖了所有字符串处理的需求。
*/

//1,判断字符串是否包含子串
/*
strings.Contains(s, substr string) bool 函数用于判断字符串s是否包含子串substr，
如果包含则返回true，否则返回false。
*/

func main() {
	//str := "Hello, World!"
	//substr := "Hello"
	//fmt.Println(strings.Contains(str, substr)) // 输出: true

	//2. 判断某个字符串是否包含了多个字符串中的某一个
	//ContainsAny函数用于检查字符串s是否包含字符串chars中的任何Unicode码点。如果包含，则返回true；否则返回false。
	//只要包含一个就返回true
	//str := "yss,大侠"
	//
	//fmt.Println(strings.ContainsAny(str, "z大"))
	////都不包含返回false
	//fmt.Println(strings.ContainsAny(str, "z大"))

	//3. 字符串计数
	//Count(s, substr string) int 方法返回字符串s中包含子串substr的个数。
	s := "hello world, world"
	fmt.Println(strings.Count(s, "world")) // 输出: 2

	//4. 查找子串在字符串中第一次出现的位置
	//strings.Index(s, substr string) int函数用于查找子串substr在字符串s中首次出现的位置（索引从0开始），
	//如果未找到则返回-1。
	//str := "Hello, World! World"
	//substr := "World"
	//index := strings.Index(str, substr)
	//fmt.Println(index) // 输出: 7

	//5. 查找子串最后一次出现的位置
	//strings.LastIndex(s, substr string) int函数与Index类似，但它查找的是子串substr在字符串s中最后一次出现的位置。
	//查的是字符串最后一次首字母出现的位置
	//str := "Hello, Hello, World!"
	//substr := "Hello"
	//lastIndex := strings.LastIndex(str, substr)
	//fmt.Println(lastIndex) // 输出: 7   查的是字符串最后一次首字母出现的位置

	//6. 判断字符串是否以指定前缀或后缀开头/结尾
	//strings.HasPrefix(s, prefix string) bool和strings.HasSuffix(s, suffix string) bool函数
	//分别用于判断字符串s是否以指定的前缀prefix或后缀suffix开头/结尾。
	//str := "Hello, World!"
	//fmt.Println(strings.HasPrefix(str, "Hello"))  // 输出: true
	//fmt.Println(strings.HasSuffix(str, "World!")) // 输出: true

	//7. 字符串的替换
	/*
		替换字符串中的子串
		strings.Replace(s, old, new string, n int) string函数用于将字符串s中的old子串替换为new子串，n表示替换的次数。如果n为小于0，则表示替换所有匹配的子串。
		n必须写，没有默认值
	*/
	//str := "Hello, World!"
	//oldSubstr := "Hello"
	//newSubstr := "Hi"
	//newStr := strings.Replace(str, oldSubstr, newSubstr, 1)
	//fmt.Println(newStr) // 输出: Hi, World!
	//
	//// 替换所有匹配的子串
	//newStrAll := strings.Replace(str, "o", "a", -1)
	//fmt.Println(newStrAll) // 输出: Hella, Warld!

	//从Go 1.12版本开始，strings包引入了ReplaceAll函数，用于替换字符串s中所有的old子串为new。
	//func ReplaceAll(s, old, new string) string
	//str := "Hello, World! Hello, Go!"
	//newStr := strings.ReplaceAll(str, "Hello", "Hi")
	//fmt.Println(newStr) // 输出: Hi, World! Hi, Go!

	//8. 字符串的分割
	/*
		将字符串分割成切片
		strings.Split(s, sep string) []string函数用于将字符串s按照指定的分隔符sep进行分割，并返回一个字符串切片。
		如果sep为空字符串，或者一个找不到的分隔符，则会将字符串s切分成单个字符的切片。
	*/
	//str := "Hello, World!"
	//sep := ", "
	////以逗号空格作为切割符
	//
	//strs := strings.Split(str, sep)
	//fmt.Println(strs) // 输出: [Hello World!]
	//
	//// 使用空字符串作为分隔符
	//strsEmpty := strings.Split(str, "")
	//fmt.Println(strsEmpty) // 输出: [H e l l o ,   W o r l d !]

	//9. 字符串的连接
	/*
		将切片中的字符串连接起来
		strings.Join(a []string, sep string) string函数用于将字符串切片a中的字符串使用指定的分隔符sep连接起来，并返回连接后的字符串。
	*/
	strs := []string{"Hello", "World!"}
	sep := ", "
	joinedStr := strings.Join(strs, sep)
	fmt.Println(joinedStr) // 输出: Hello, World!

	//10. 字符串比较
	/*
		虽然strings包本身不直接提供字符串比较的函数（因为Go语言的==和!=操作符已经足够用于比较字符串），但了解如何比较字符串并理解其背后的机制是很重要的。
		特别是当涉及到比较时区分大小写或不区分大小写时。
		区分大小写比较：直接使用==和!=。
		不区分大小写比较：可以使用strings.EqualFold方法。
	*/
	s1 := "Hello, Go!"
	s2 := "hello, go!"
	s3 := "Hello, Go!"

	// 区分大小写比较
	fmt.Println(s1 == s2) // 输出: false
	fmt.Println(s1 == s3) // 输出: true

	// 不区分大小写比较
	fmt.Println(strings.EqualFold(s1, s2)) // 输出: true

	//11. 字符串的大小写转换
	//strings.ToUpper(s string) string和strings.ToLower(s string) string函数分别用于将字符串s中的所有字符转换为大写或小写。
	//str := "Hello, World!"
	//upperStr := strings.ToUpper(str)
	//lowerStr := strings.ToLower(str)
	//fmt.Println(upperStr) // 输出: HELLO, WORLD!
	//fmt.Println(lowerStr) // 输出: hello, world!

	//12. 去除字符串首尾的空白字符
	//strings.TrimSpace(s string) string函数用于去除字符串s开头和结尾的空白字符（如空格、换行符等）。
	//str := " Hello, World! "
	//trimmedStr := strings.TrimSpace(str)
	//fmt.Println(trimmedStr) // 输出: Hello, World!

	//13. 去除字符串首尾指定的字符
	//strings.Trim(s, cutset string) string函数用于去除字符串s开头和结尾由cutset指定的字符。
	//str := "!!!Hello, World!!!"
	//trimmedStr := strings.Trim(str, "! ")
	//fmt.Println(trimmedStr) // 输出: Hello, World

	//14. 字符串的重复
	//Repeat(s string, count int) string方法将字符串s重复count次，并返回结果字符串。
	s = "Go!"
	repeated := strings.Repeat(s, 4)
	fmt.Println(repeated)

	//15. 字符串提取
	//可以通过字符串下标提取想要的字符串范围，取单个得到的是uint8的ASIIC吗数字，取多个得到的是字符串
	str := "Hello, World! Hello, Go!"
	fmt.Println(str[0:5])        //Hello
	fmt.Printf("%T\n", str[0:5]) //string
	fmt.Printf("%T\n", str[0])   //uint8
	//单个字符可以通过string转化为字符串
	fmt.Println(string(str[0])) //H

}
