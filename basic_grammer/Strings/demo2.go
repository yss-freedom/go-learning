package main

import (
	"fmt"
	"strconv"
)

/*
在Go语言（Golang）的编程实践中，strconv包是一个非常重要的标准库，它提供了在基本数据类型（如整型、浮点型、布尔型）和字符串之间的转换功能。
无论是在处理用户输入、文件读写、网络传输还是其他需要数据格式转换的场景中，strconv包都扮演着至关重要的角色。

strconv是两个单词的缩写。string convert = strconv
strconv包位于Go语言的标准库中，它提供了一系列用于字符串和基本数据类型之间转换的函数。这些函数主要包括以下几种类型：
将基本数据类型转换为字符串的函数，如Itoa、FormatInt、FormatFloat、FormatBool等。
将字符串解析为基本数据类型的函数，如Atoi、ParseInt、ParseFloat、ParseBool等。
附加到已存在字节数组的函数，如AppendInt、AppendFloat、AppendBool等。
其他辅助函数，如IsPrint、IsGraphic、Quote、Unquote等。
将字符串转换为其他基础类型的过程叫做解析parse，其他基础类型转换为字符串的过程成为格式化format
*/

//2.1 整数转字符串
//在Go语言中，将整数转换为字符串是一个常见的需求。
//strconv包中的Itoa函数和FormatInt函数都可以实现这一功能。

// 整形转字符串
func main() {
	intNum := 42
	//整形转字符串方法strconv.Itoa
	strFromInt := strconv.Itoa(intNum)
	fmt.Printf("整型转换为字符串: %s\n", strFromInt)
	fmt.Printf("整型转换为字符串数据类型: %T\n", strFromInt)

	/*
		对于大整数或需要指定进制的无符号整数场景，可以使用FormatInt函数。
		func FormatInt(i int64, base int) string
		i 表示需要转换的整数，base 表示源数字要转换后的进制数，最终返回的是 i 的字符串形式
	*/
	bigIntValue := int64(9223372036854775807) // MaxInt64
	bigIntStr := strconv.FormatInt(bigIntValue, 10)
	fmt.Println("大整数转字符串:", bigIntStr)

	// 转换为二进制字符串
	binaryStr := strconv.FormatInt(bigIntValue, 2)
	fmt.Println("大整数转二进制字符串:", binaryStr)

	//2.2 字符串转整数
	//将字符串转换为整数，可以使用Atoi函数或ParseInt函数。
	/*
		示例1：使用Atoi函数
		注意，字符串转整形，需要进行错误捕获
		func Atoi(s string) (int, error)
	*/
	//Atoi函数将字符串转换为int类型。注意，如果字符串不能被解析为整数，Atoi会返回错误。
	//str := "123"
	//intFromStr, err := strconv.Atoi(str)
	//if err != nil {
	//	fmt.Println("转换错误:", err)
	//	return
	//}
	//fmt.Printf("字符串转换为整型: %d\n", intFromStr)

	/*
		示例2：使用ParseInt函数
		ParseInt函数比Atoi更灵活，允许指定基数（进制）和位大小。
		func ParseInt(s string, base int, bitSize int) (i int64, err error)
		s string表示需要转换的字符串
		base int 表示需要转换的数据的进制
		bitSize int 表示转换后的整数为多少位int类型
		返回一个int64整形和一个error
	*/
	hexStr := "FF"
	hexValue, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("转换错误:", err)
		return
	}
	fmt.Printf("十六进制字符串转整数: %d\n", hexValue)

	//3. 浮点数与字符串的转换
	/*
		3.1 浮点数转字符串
		将浮点数转换为字符串，可以使用FormatFloat函数。
		func FormatFloat(f float64, fmt byte, prec, bitSize int) string
		f：要格式化的浮点数
		fmt：格式标记，'b’表示二进制，'e’表示科学计数法，'f’表示十进制无指数，'g’表示最少计数法表示
		prec：表示精度，对于’f’和’g’格式，代表除小数点以外的位数。也可以是-1，表示根据浮点数的小数位数自动确定保留的位数。
		bitsize：浮点数类型，32表示float32，64表示float64
	*/
	floatValue := 3.14159
	floatStr := strconv.FormatFloat(floatValue, 'f', -1, 64)
	fmt.Println("浮点数转字符串:", floatStr)

	// 保留两位小数
	preciseFloatStr := strconv.FormatFloat(floatValue, 'f', 2, 64)
	fmt.Println("精确到两位小数的浮点数转字符串:", preciseFloatStr)

	//3.2 字符串转浮点数
	/*
		将字符串转换为浮点数，可以使用ParseFloat函数。
		func ParseFloat(s string, bitSize int) (float64, error)

		将字符串解析为浮点数，如果 s 符合语法规则，会返回一个最为接近 s 表示值得浮点数（IEEE754 规范舍入）
		bitSize 指定解析结果的浮点数类型，32 是 float32，64 是 float64
		返回值 err 是 *NumErr 类型的，语法有误时，err.Error = ErrSyntax；结果超出表示范围时，返回值 f 为 ±Inf，err.Error= ErrRange
	*/
	//str := "3.14"
	//floatValue, err = strconv.ParseFloat(str, 64)
	//if err != nil {
	//	fmt.Println("转换错误:", err)
	//	return
	//}
	//fmt.Println("字符串转浮点数:", floatValue)

	//4. 布尔值与字符串的转换
	/*
		4.1 布尔值转字符串
		将布尔值转换为字符串时，可以使用strconv.FormatBool函数。该函数将true转换为字符串"true"，将false转换为字符串"false"。
	*/
	boolValue := true
	boolStr := strconv.FormatBool(boolValue)
	fmt.Println("布尔值转字符串:", boolStr)
	fmt.Printf("布尔值转字符串数据类型:%T\n", boolStr)

	//4.2 字符串转布尔值
	/*
		将字符串转换为布尔值时，可以使用strconv.ParseBool函数。
		该函数会尝试将字符串解析为布尔值，并返回解析后的布尔值和可能发生的错误。有效的布尔值字符串为"true"和"false"（不区分大小写）。
	*/

	str := "true"
	boolValue, err = strconv.ParseBool(str)
	if err != nil {
		panic(err)
	}
	fmt.Println("字符串转布尔值:", boolValue)
	fmt.Printf("字符串转布尔值数据类型:%T\n", boolValue)

	//5. strconv包的其他功能
	/*
		5.1 Append系列函数
		strconv包还提供了Append…系列函数，这些函数可以将基本数据类型追加到已存在的字节数组中，而无需重新分配整个数组。这在处理大量数据时非常有用，可以显著提高性能。
	*/
	// 追加整数到字节数组
	num1 := 123
	byteSlice := []byte("Number: ")
	byteSlice = strconv.AppendInt(byteSlice, int64(num1), 10)
	fmt.Println("追加整数:", string(byteSlice))

	// 追加布尔值到字节数组
	boolVal := true
	byteSlice = []byte("Bool: ")
	byteSlice = strconv.AppendBool(byteSlice, boolVal)
	fmt.Println("追加布尔值:", string(byteSlice))

	// 追加浮点数到字节数组
	floatVal := 3.14
	byteSlice = []byte("Float: ")
	byteSlice = strconv.AppendFloat(byteSlice, floatVal, 'f', -1, 64)
	fmt.Println("追加浮点数:", string(byteSlice))

	/*
		5.2 辅助函数
		strconv包还包含一些辅助函数，如IsPrint、IsGraphic、Quote和Unquote等，它们提供了额外的功能，如检查字符的可打印性、将字符串转换为带引号和转义字符的字符串字面值等。
		strconv.Quote 函数用于将字符串转换为双引号括起来的、转义了特殊字符的Go字符串字面量。而 strconv.Unquote 函数则执行相反的操作，即解析一个被双引号括起来、可能包含转义字符的字符串。
	*/
	chars := []rune{'H', 'e', 'l', '\n', '♥', 127}
	for _, char := range chars {
		fmt.Printf("Character: %c, IsPrint: %v\n", char, strconv.IsPrint(char))
		fmt.Printf("Character: %c, IsGraphic: %v\n", char, strconv.IsGraphic(char))
	}

	str = `路多辛的,"所思所想"!。`
	quoted := strconv.Quote(str)
	fmt.Println("Quoted: ", quoted)

	unquoted, err := strconv.Unquote(quoted)
	if err != nil {
		panic(err)
	}
	fmt.Println("Unquoted: ", unquoted)
	/*
		注意事项
		在使用Parse…系列函数时，一定要检查返回的错误值，以确保转换成功。
		对于大整数或无符号整数，应使用FormatInt和ParseInt函数，并指定适当的基数和位数。
		在处理浮点数时，注意精度和格式的控制。
		strconv包中的函数主要用于基本类型和字符串之间的转换，对于复杂数据结构的序列化和反序列化，应考虑使用encoding/json、encoding/xml等包。
	*/
}
