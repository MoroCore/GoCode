package main

import (
	"fmt"
	// "bufio"
	// "os"
	"strings"
)
// 时间复杂的 n4   4重循环
func main() {

	// counts := make(map[string]int)
	// input := bufio.NewScanner(os.Stdin)	
	
	// for input.Scan(){
	// 	counts[input.Text()]++
	// }
	// for _ , _ := range counts {

	// 	max := getMaxLengthChildStringByForce("aaaaaa")
	// 	fmt.Println("最长子串的长度为",max)
	// }
	// max := getMaxLengthChildStringByForce("aaaaaa")

	max := lengthOfLongestSubstring("abadfcagagvagv")
	fmt.Println("最长子串的长度为",max)
}

var (
	chlidString string
	length int
)

// 方法1  暴力破解
// 思路  获取该字符串的所有子串 比较长度 
func getMaxLengthChildStringByForce( s string) int {

	for i := 0 ; i < len(s) ; i++ {
		for j := i+1 ; j < len(s) + 1 ; j++ {

			chlidString = s[i:j]
			right := decideStringIsRightByMap(chlidString)
			if right == true {
				if len(chlidString) > length {
					length = len(chlidString)
				}
			}
		}
	}
	return length
}

func decideStringIsRightByForce(s string) bool {


	for i := 0 ; i< len(s) - 1 ;i++{
		for j := i+1 ; j < len(s) ;j++{
			if s[i] == s[j]{
				return false
			}
		}
	}
	return true
}

//方案2  获取所有的子字符串  但是判断子字符串是否有重复字符时
func decideStringIsRightByMap(s string) bool {

	counts := make(map[string]int)
	for i := 0 ; i < len(s) ; i++{
		str  := s[i:i+1]
		counts[str]++
		if(counts[str] != 1){
			return false
		}
	}
	return true
}

//滑动窗口算法
func lengthOfLongestSubstring(s string) int {

	max := 0
	start := 0
	end := 0

	for i := 0 ; i < len(s) ; i++ {

		//返回的是下标
		cIndex := strings.LastIndex(s[start:end],string(s[i]))
		fmt.Println("cIndex=",cIndex)
		if cIndex > -1 {
			if max < (end - start){
				max = end - start
			}
			if cIndex == 0 {
				start++
			}else {
				start += cIndex + 1
			}

		}
		end++
		if (i == len(s) - 1 && max < end - start){
			max = end - start
		}
	}
	return max
}

func  test(s string) int {

	max := 0
	start := 0
	end := 0

	for i := 0 ; i < len(s) ; i++ {

		index := strings.LastIndex(s[start:end],string(s[i]))

		if index > -1 {
			if max < (end - start) {
				max = end - start
			}
			if index == 0 {
				start++
			}else {
				start += index + 1
			}

		}
		end++
		if i == len(s) -1  && max < (end - start) {
			max = end - start
		}
	}
	return max
}
