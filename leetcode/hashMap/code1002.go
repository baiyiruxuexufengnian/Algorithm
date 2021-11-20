package hashMap

func commonChars(words []string) []string {
	if len(words) == 0 {
		return words
	}

	mp := make(map[byte]int)
	ret := make([]string, 0)
	for _, c := range []byte(words[0]) {
		mp[c] ++
	}

	for _, str := range words {
		tmp := make(map[byte]int)
		for _, c := range []byte(str) {
			tmp[c] ++
		}
		for k, _ := range mp {
			if _, ok := tmp[k]; ok {
				mp[k] = min(mp[k], tmp[k])
			} else {
				delete(mp, k)
			}
		}
	}
	for k, _ := range mp {
		for mp[k] > 0 {
			ret = append(ret, string(k))
			mp[k] --
		}
	}
	return ret
}

func commonChars2(words []string) []string {
	length:=len(words)
	fre:=make([][]int,0)//统计每个字符串的词频
	res:=make([]string,0)
	//统计词频
	for i:=0;i<length;i++{
		var row [26]int//存放该字符串的词频
		for j:=0;j<len(words[i]);j++{
			row[words[i][j]-97]++
		}
		fre=append(fre,row[:])
	}
	//查找一列的最小值
	for j:=0;j<len(fre[0]);j++{
		pre:=fre[0][j]
		for i:=0;i<len(fre);i++{
			pre=min(pre,fre[i][j])
		}
		//将该字符添加到结果集（按照次数）
		tmpString:=string(j+97)
		for i:=0;i<pre;i++{
			res=append(res,tmpString)
		}
	}
	return res
}

func min(a,b int)int{
	if a>b{
		return b
	}
	return a
}