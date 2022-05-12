/*
 * @Author: reber
 * @Mail: reber0ask@qq.com
 * @Date: 2022-04-28 10:26:09
 * @LastEditTime: 2022-05-12 15:15:37
 */
package utils

import (
	"reflect"
	"sort"
	"strings"
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
	// https://github.com/syyongx/php2go/blob/master/php.go#L1265

	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		panic("haystack: haystack type muset be slice, array or map")
	}

	return false
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
