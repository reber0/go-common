/*
 * @Author: reber
 * @Mail: reber0ask@qq.com
 * @Date: 2021-11-10 09:48:35
 * @LastEditTime: 2022-04-25 17:19:04
 */

package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
	"github.com/syyongx/php2go"
)

// 获取两个 string 的相似度
func GetRatio(first string, second string) (ratio float64) {
	_ = php2go.SimilarText(first, second, &ratio)
	return ratio / 100
}

// HandleError 用于处理 error
func HandleError(action string, err error) {
	if err != nil {
		_ = fmt.Errorf(fmt.Sprintf("%s => %s\n", action, err))
	}
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

// 反转 [][]string
func ReverseSlice(s [][]string) [][]string {
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

// 随机生成 X-Forwarded-For
func RandomUserAgent() string {
	userAgent := []string{
		"Mozilla/5.0 (iPhone; CPU iPhone OS 15_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/100.0.4896.77 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 15_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) EdgiOS/100.0.1185.50 Version/15.0 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 10_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 OPT/3.2.9",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 12_3_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.3 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 12_6_1 like Mac OS X) AppleWebKit/612.4.9 (KHTML, like Gecko) Mobile/19D52 QHBrowser/2 QihooBrowser/5.2.4",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 15_3_1 like Mac OS X; zh-cn) AppleWebKit/601.1.46 (KHTML, like Gecko) Mobile/19D52 Quark/5.6.5.1336 Mobile",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 15_3_1 like Mac OS X; zh-CN) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/19D52 UCBrowser/13.8.9.1722 Mobile  AliApp(TUnionSDK/0.1.20.4)",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML,  like Gecko) Version/6.0 Mobile/10A403 Safari/8536.25",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:91.0) Gecko/20100101 Firefox/91.0",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.12 Safari/537.36 OPR/86.0.4363.23 (Edition B2)",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.3 Safari/605.1.15",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_16_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.69 Safari/537.36 QIHU 360EE",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 12_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.23 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36 Edg/100.0.1185.50",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:99.0) Gecko/20100101 Firefox/99.0",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 OPR/86.0.4363.23",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; Touch; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.15 Safari/537.36 QIHU 360SE",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.16 Safari/537.36",
		"Mozilla/5.0 (X11; Ubuntu; Linux x86 64; rv:79.0) Gecko/20100101 Firefox/79.0",
		"Mozilla/5.0 (Linux; Ubuntu 16.04) AppleWebKit/537.36 Chromium/57.0.2987.110 Safari/537.36",
	}
	return userAgent[RandomInt(0, len(userAgent)-1)]
}

// 随机生成 X-Forwarded-For
func RandomXFF() string {
	int1 := RandomInt(1, 255)
	int2 := RandomInt(1, 255)
	int3 := RandomInt(1, 255)
	int4 := RandomInt(1, 255)
	xff := fmt.Sprintf("%d.%d.%d.%d", int1, int2, int3, int4)
	return xff
}

// 获取区间中的随机整数，返回数字范围 [mai, max]
func RandomInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

// 获取指定长度的随机字符串
func RandomStr(length int) string {
	b_str := []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		s := b_str[rand.Intn(len(b_str))]
		result = append(result, s)
	}
	return string(result)
}

// 获取文件内容
func FileGetContents(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return string(content)
}

// 按行读取文件内容
func FileEachLineRead(filename string) []string {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0664)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var datas []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		datas = append(datas, sc.Text())
	}
	return datas
}

// 判定文件是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}

// 时间戳转时间字符串 => 2006-01-02 15:04:05
func UnixToStr(timestamp interface{}) string {
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
