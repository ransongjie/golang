package str

import (
	"errors"
	"math"
	"math/rand"

	. "com.xcrj/golang/common"
)

//最长回文串
/**
 * Manacher 求回文串的最大长度
 * [..., i', maxC, i, ...]
 * 1. i在maxC+maxR内：扩展基础 r(i)=min{(maxC+maxR-i), maxC-(i-maxC)}
 * r=i关于maxC的对称点i'的回文半径r(i')
 * - r(i')在maxR内，r(i)=r(i')=rs[maxC-(i-maxC)], 外扩将失败
 * - r(i')在maxR外，r(i)=(maxC+maxR-i), 外扩将失败 (若可以外扩则r(i')不在maxR外)
 * - r(i')刚好压线，r(i)=(maxC+maxR-i)=rs[maxC-(i-maxC)], 外扩可能成功
 * 2. i不在maxC+maxR内(外或压线)：扩展基础 r(i)=1
 * 步骤
 * 1. 转换字符串
 * 2. 历史最大回文半径，历史最大回文半径的回文中心，每个字符的回文半径
 * 3. i的扩展基础，继续扩展i
 * 4. 更新历史最大回文半径，历史最大回文半径的回文中心
 * - i在maxC+maxR内，扩展基础 r(i)=min{(maxC+maxR-i), maxC-(i-maxC)}
 * - i不在maxC+maxR内，扩展基础 r(i)=1
 */
func RunManacher() {
	times := 10000
	maxLen := 100
	for i := 0; i < times; i++ {
		s := getString(maxLen)
		a := manacher(s)
		b := manacher2(s)
		if a != b {
			panic(errors.New("error a!=b."))
		}
	}
}

func getString(maxLen int) string {
	rand.Seed(10)
	len := rand.Intn(maxLen) + 1
	cs := make([]rune, len)
	for i := 0; i < len; i++ {
		c := rand.Intn(126 - 20 + 1)
		cs[i] = rune(c)
	}
	return string(cs)
}

func manacher(s string) int {
	if len(s) == 0 {
		return 0
	}

	t := convert(s)
	// 历史最大回文半径，历史最大回文半径的回文中心
	maxR, maxC, n := 0, 0, len(t)
	// 每个字符最大回文半径
	rs := make([]int, n)

	// 遍历
	for i := 0; i < n; i++ {

		if i >= maxC+maxR {
			rs[i] = 1
		} else {
			rs[i] = Min[int](maxC+maxR-i, rs[maxC-(i-maxC)])
		}

		for i+rs[i] < n && i-rs[i] > -1 && t[i+rs[i]] == t[i-rs[i]] {
			rs[i]++
		}

		if rs[i] > maxR {
			maxR = rs[i]
			maxC = i
		}
	}

	// s >#> t
	return maxR - 1
}

// #a#b#c#
func convert(s string) string {
	ans := []rune{}
	ans = append(ans, '#')
	for _, c := range s {
		ans = append(ans, c, '#')
	}
	return string(ans)
}

func manacher2(s string) int {
	if len(s) == 0 {
		return 0
	}

	t := convert2(s)
	// 历史最大回文半径的回文中心，C+历史回文半径，最大回文半径
	C, L, maxR, n := -1, -1, math.MinInt, len(t)
	// 每个字符的最大回文半径
	rs := make([]int, n)
	for i := 0; i < n; i++ {
		if i >= L {
			rs[i] = 1
		} else {
			rs[i] = Min[int](L-i, rs[C-(i-C)])
		}
		for i-rs[i] > -1 && i+rs[i] < n && t[i-rs[i]] == t[i+rs[i]] {
			rs[i]++
		}
		if i+rs[i] > L {
			L = i + rs[i]
			C = i
			maxR = Max[int](maxR, rs[i])
		}
	}
	return maxR - 1
}

func convert2(s string) string {
	n := len(s)*2 + 1
	ans := make([]rune, n)
	for i, j := 0, 0; i < n; i++ {
		if (i & 1) == 0 {
			ans[i] = '#'
		} else {
			ans[i] = rune(s[j])
			j++
		}
	}
	return string(ans)
}
