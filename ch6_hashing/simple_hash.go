package ch6_hashing

import "fmt"

//简单设计哈希算法

//加法哈希
//对输入的每个字符的 ASCII 码进行相加，将得到的总和作为哈希值。
func addHash(key string) int {

	var hash int64
	var modulus int64

	modulus = 1000000007
	for _, b := range []byte(key) {
		hash = (hash + int64(b)) % modulus
	}
	return int(hash)

}

//乘法哈希
func mulHash(key string) int {
	var hash int64
	var modulus int64

	modulus = 1000000007
	for _, b := range []byte(key) {
		hash = (31*hash + int64(b)) % modulus
	}
	return int(hash)
}

//亦或哈希
func xorHash(key string) int {
	hash := 0
	modulus := 1000000007
	for _, b := range []byte(key) {
		fmt.Println(int(b))
		hash ^= int(b)
		hash = (31*hash + int(b)) % modulus
	}
	return hash & modulus
}

//旋转哈希
//将每个字符的 ASCII 码累积到一个哈希值中，每次累积之前都会对哈希值进行旋转操作。
func rotHash(key string) int {
	var hash int64
	var modulus int64

	modulus = 1000000007
	for _, b := range []byte(key) {
		hash = ((hash << 4) ^ (hash >> 28) ^ int64(b)) % modulus
	}
	return int(hash)
}
