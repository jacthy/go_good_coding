package main


// 顺丰区块链面试算法题
// 连续，没有重复字符的最长子串
// 思路：窗口的概念，一个临时窗口，遇到重复的时候就判断一下是不是比之前缓存的要长，是就替换，然后清空临时窗口
// 其中一个小优化点：遍历的指针通过回退的方式回退到重复的那个位置，如abcd 下一个是c，那就直接从c这个位置往前
func getSonStr(str string) string {
	var finalArr []byte
	var tmpArr []byte
	for i := 0; i < len(str); i++ {
		repeatIndex := ifHavRepeat(tmpArr, str[i])
		if repeatIndex < 0 { // 未出现重复
			tmpArr = append(tmpArr, str[i])
			if i == len(str)-1 {
				if len(tmpArr) > len(finalArr) {
					finalArr = tmpArr
				}
			}
		} else {
			//出现重复
			i = i - len(tmpArr) + repeatIndex //指针就要回退到
			if len(tmpArr) > len(finalArr) {
				finalArr = tmpArr
			}
			tmpArr = []byte{}
		}
	}
	return string(finalArr)
}


func ifHavRepeat(str []byte, val byte) int {
	for i, b := range str {
		if b == val {
			return i
		}
	}
	return -1
}