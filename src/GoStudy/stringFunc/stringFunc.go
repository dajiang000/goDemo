package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := "124将"
	s := []rune(str1)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}
	num, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println("错误原因", err)
	} else {
		fmt.Println(num)
	}
	i := 1122
	str3 := strconv.Itoa(i)
	fmt.Println(str3)
	byte1 := []byte(str1)

	fmt.Println(byte1)
	byte2 := []byte{49, 51, 52, 53, 54}
	str4 := string(byte2)
	fmt.Println(str4)
	str5 := "dsssss"
	fmt.Println(strings.Count(str5, "sss"))
	num2 := 1242
	str6 := strconv.FormatInt(int64(num2), 2)
	fmt.Println(str6)

	str7 := "jiangda"
	b := strings.Contains(str7, "jiang")
	fmt.Println(b)
	str8 := "dsssd"
	num3 := strings.Count(str8, "ss")
	fmt.Println(num3)

	str9 := "dSSSd"
	b1 := strings.EqualFold(str9, str8)
	fmt.Println(b1)

	str10 := "ddadf"
	loc := strings.Index(str10, "dd")
	loc1 := strings.LastIndex(str10, "dd")
	fmt.Println(loc, loc1)
	str11 := "do do jiang"
	res1 := strings.Replace(str11, "do", "shuai", 1)
	fmt.Println(res1)

	str12 := "1,12,14,121"
	res2 := strings.Split(str12, ",1")
	fmt.Println(res2)

	str13 := "dA JIang"
	str14 := strings.ToLower(str13)
	str15 := strings.ToUpper(str13)
	fmt.Println(str14, str15)

	str16 := "diassssssd ia"
	str1 = strings.Trim(str16, "ad")
	fmt.Println(str1)
	str17 := " diassssssd ia "
	str1 = strings.TrimSpace(str17)
	str18 := "diassssssd iaa"
	b2 := strings.HasPrefix(str18, "di")
	b3 := strings.HasSuffix(str18, "aa")
	fmt.Println(b2, b3)
}
