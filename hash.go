package tools

//RSHash 根据字符串，计算出一个hash值
func RSHash(key string) int {
	b := 378551
	a := 63689
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = hash*a + int(key[i])
		a = a * b
	}
	return (hash & 0x7FFFFFFF)
}
