/*
 * @Author: reber
 * @Mail: reber0ask@qq.com
 * @Date: 2022-04-28 10:26:09
 * @LastEditTime: 2022-04-28 10:26:11
 */
package utils

import (
	"sort"
	"strings"

	"github.com/syyongx/php2go"
)

// 反转 [][]string
func SliceReverse(s [][]string) [][]string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// slice 转为 string
func SliceToString(slc []string) string {
	return "[" + strings.Join(slc, ", ") + "]"
}

// slice 排序
func SortSlice(t []string) {
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
}

// 判断 needle 是否在 slice, array, map 中
func InSlice(needle interface{}, haystack interface{}) bool {
	return php2go.InArray(needle, haystack)
}

// slice(string类型)元素去重
func UniqStringSlice(slc []string) []string {
	result := make([]string, 0)
	tempMap := make(map[string]bool, len(slc))
	for _, e := range slc {
		if tempMap[e] == false {
			tempMap[e] = true
			result = append(result, e)
		}
	}
	return result
}

// slice(int类型)元素去重
func UniqIntSlice(slc []int) []int {
	result := make([]int, 0)
	tempMap := make(map[int]bool, len(slc))
	for _, e := range slc {
		if tempMap[e] == false {
			tempMap[e] = true
			result = append(result, e)
		}
	}
	return result
}
