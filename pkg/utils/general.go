/*
 * MIT License
 *
 * Copyright (c) 2024 Bamboo
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 */

package utils

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/go-ping/ping"
)

func ConvertToIntList(stringList []string) ([]int, error) {
	intList := make([]int, 0, len(stringList))
	for _, idStr := range stringList {
		id, err := strconv.Atoi(strings.TrimSpace(idStr)) // 去除空白并转换为整数
		if err != nil {
			return nil, fmt.Errorf("无法解析 leafNodeId: '%s' 为整数", idStr)
		}
		intList = append(intList, id)
	}

	return intList, nil
}

// IsType 判断两个值是否是相同类型
func IsType(value1, value2 interface{}) bool {
	return reflect.TypeOf(value1) == reflect.TypeOf(value2)
}

// GetDefaultValue 返回值的默认值（零值）
func GetDefaultValue(value interface{}) interface{} {
	v := reflect.ValueOf(value)
	if !v.IsValid() {
		return nil
	}

	// 创建零值的副本并返回
	return reflect.Zero(v.Type()).Interface()
}

// GetMax 返回两个数值中的最大值，支持 int, float64 等常见类型
func GetMax(value1, value2 interface{}) (interface{}, error) {
	switch v1 := value1.(type) {
	case int:
		v2, ok := value2.(int)
		if !ok {
			return nil, errors.New("类型不匹配")
		}
		if v1 > v2 {
			return v1, nil
		}
		return v2, nil
	case float64:
		v2, ok := value2.(float64)
		if !ok {
			return nil, errors.New("类型不匹配")
		}
		return math.Max(v1, v2), nil
	default:
		return nil, errors.New("不支持的类型")
	}
}

// GetMin 返回两个数值中的最小值，支持 int, float64 等常见类型
func GetMin(value1, value2 interface{}) (interface{}, error) {
	switch v1 := value1.(type) {
	case int:
		v2, ok := value2.(int)
		if !ok {
			return nil, errors.New("类型不匹配")
		}
		if v1 < v2 {
			return v1, nil
		}
		return v2, nil
	case float64:
		v2, ok := value2.(float64)
		if !ok {
			return nil, errors.New("类型不匹配")
		}
		return math.Min(v1, v2), nil
	default:
		return nil, errors.New("不支持的类型")
	}
}

// ToUpperCase 将字符串转换为大写
func ToUpperCase(str string) string {
	return strings.ToUpper(str)
}

// ToLowerCase 将字符串转换为小写
func ToLowerCase(str string) string {
	return strings.ToLower(str)
}

// TrimSpaces 去掉字符串的前后空格
func TrimSpaces(str string) string {
	return strings.TrimSpace(str)
}

// IsSameDay 判断两个日期是否为同一天
func IsSameDay(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

// DaysBetween 计算两个日期之间的天数
func DaysBetween(t1, t2 time.Time) int {
	days := t2.Sub(t1).Hours() / 24
	return int(math.Abs(days))
}

// IsValidEmail 简单检查一个字符串是否是有效的电子邮件格式
func IsValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}

// MapToStringSlice 将 map 转换为 []string，要求偶数个元素，key和值依次排列
func MapToStringSlice(inputMap map[string]string) ([]string, error) {
	if inputMap == nil {
		return []string{}, nil
	}

	var result []string
	for key, value := range inputMap {
		result = append(result, key, value)
	}

	// 确保结果长度为偶数
	if len(result)%2 != 0 {
		return nil, fmt.Errorf("转换后的字符串切片长度为奇数，不符合键值对要求")
	}

	return result, nil
}

// StringSliceToMap 将 []string 转换为 map[string]string，要求输入长度为偶数，奇数索引为 key，偶数索引为 value
func StringSliceToMap(inputSlice []string) (map[string]string, error) {
	if len(inputSlice)%2 != 0 {
		return nil, fmt.Errorf("输入的字符串切片长度必须为偶数，实际长度为 %d", len(inputSlice))
	}

	result := make(map[string]string)
	for i := 0; i < len(inputSlice); i += 2 {
		key := inputSlice[i]
		value := inputSlice[i+1]
		result[key] = value
	}

	return result, nil
}

// Ping 检查指定的 IP 地址是否可达
func Ping(ipAddr string) bool {
	pinger, err := ping.NewPinger(ipAddr)
	if err != nil {
		fmt.Printf("创建 pinger 失败: %v\n", err)
		return false
	}

	// 设置使用 IPv4，如果需要 IPv6，可以设置为 false
	pinger.SetPrivileged(true) // 需要 root/管理员权限

	// 设置超时时间
	pinger.Timeout = time.Second * 3

	// 发送 1 个包
	pinger.Count = 1

	err = pinger.Run() // 开始 ping
	if err != nil {
		return false
	}

	stats := pinger.Statistics()
	return stats.PacketsRecv > 0
}