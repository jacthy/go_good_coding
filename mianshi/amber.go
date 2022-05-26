package main

// amber 三面笔试题

// 获取所有子串（可重复），注意要是找子集则不可重复
// 思路：利用递归，将字符串分成str[0],str[1:]，获取str[1:]应展出的子串，递归下去
// 流程 如:abc 第一步：分为a,bc,第二步分为：b,c，最后的c直接返回，然后b和[c]组成子字符串是[b,bc,c]
// 再向上递归就出现了a和[b,bc,c]的组合，这时a和数组的组合是[ab,abc,ac],再将原来的数据a和[b,bc,c]合并进去
func getSet(str string) []string {
	if len(str) == 1 {
		return []string{str}
	}
	nextStrArr := getSet(str[1:])
	resultArr := []string{string(str[0])}
	for _, b := range nextStrArr {
		resultArr = append(resultArr, string(str[0])+b)
	}
	resultArr = append(resultArr, nextStrArr...)
	return resultArr
}
