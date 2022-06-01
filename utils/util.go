/*
 * @Author: reber
 * @Mail: reber0ask@qq.com
 * @Date: 2021-11-10 09:48:35
 * @LastEditTime: 2022-06-01 23:15:33
 */

package utils

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/nsf/termbox-go"
)

// 获取区间中的一个随机整数，返回数字范围 [mai, max]
func RandomInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

// 获取指定长度的随机字符串
func RandomString(length int) string {
	b_str := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		s := b_str[rand.Intn(len(b_str))]
		result = append(result, s)
	}
	return string(result)
}

// 时间戳转时间字符串 => 2006-01-02 15:04:05
func Unix2String(timestamp interface{}) string {
	// 通过 i.(type) 来判断是什么类型,下面的 case 分支匹配到了则执行相关的分支
	switch timestamp.(type) {
	case int:
		t := int64(timestamp.(int)) // interface 转为 int 再转为 int64
		return time.Unix(t, 0).Format("2006-01-02 15:04:05")
	case int64:
		return time.Unix(timestamp.(int64), 0).Format("2006-01-02 15:04:05")
	case string:
		t, _ := strconv.ParseInt(timestamp.(string), 10, 64) // interface 转为 string 再转为 int64
		return time.Unix(t, 0).Format("2006-01-02 15:04:05")
	}
	return ""
}

// 获取终端宽度
func GetTermWidth() int {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	width, _ := termbox.Size()
	termbox.Close()

	return width
}

// 获取两个 string 的相似度
func GetRatio(first string, second string) (percent float64) {
	// https://github.com/syyongx/php2go/blob/master/php.go#L870

	var similarText func(string, string, int, int) int
	similarText = func(str1, str2 string, len1, len2 int) int {
		var sum, max int
		pos1, pos2 := 0, 0

		// Find the longest segment of the same section in two strings
		for i := 0; i < len1; i++ {
			for j := 0; j < len2; j++ {
				for l := 0; (i+l < len1) && (j+l < len2) && (str1[i+l] == str2[j+l]); l++ {
					if l+1 > max {
						max = l + 1
						pos1 = i
						pos2 = j
					}
				}
			}
		}

		if sum = max; sum > 0 {
			if pos1 > 0 && pos2 > 0 {
				sum += similarText(str1, str2, pos1, pos2)
			}
			if (pos1+max < len1) && (pos2+max < len2) {
				s1 := []byte(str1)
				s2 := []byte(str2)
				sum += similarText(string(s1[pos1+max:]), string(s2[pos2+max:]), len1-pos1-max, len2-pos2-max)
			}
		}

		return sum
	}

	l1, l2 := len(first), len(second)
	if l1+l2 == 0 {
		return 0
	}
	sim := similarText(first, second, l1, l2)
	percent = float64(sim*200) / float64(l1+l2)

	return percent / 100
}
