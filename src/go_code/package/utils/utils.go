package utils

// 要想被外面的文件调用,方法名的首个字母必须大写
func Calculate(a float32,b float32) float32 {
	c := a +b
	return c
}   