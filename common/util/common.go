package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/wumansgy/goEncrypt"
	"io"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

//MDF5 256位加密
func EncryptSha256(str string) string {
	return goEncrypt.Sha256Hex([]byte(str))
}

//MDF5 512位加密
func DecryptSha512(str string) string {
	return goEncrypt.Sha512Hex([]byte(str))
}

func Md5V1(str string) string {
	w := md5.New()
	_, _ = io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func GetToken(ags ...string) string {
	bs := strconv.FormatInt(time.Now().UnixNano(), 10)
	for _, item := range ags {
		bs += item
	}
	return fmt.Sprintf("%x", sha1.Sum([]byte(bs)))
}

// 随机生成指定位数的大写字母和数字的组合
func GetRandomString(l int) string {
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 随机生成指定位数的数字
func GetRandomInt(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 16位
func GetOrderID() int64 {
	atoi, _ := strconv.Atoi(time.Now().Format("20060102150405") + GetRandomInt(2))
	return int64(atoi)
}

// Contains 判断切片包含某元素
func Contains(ary interface{}, value interface{}) bool {
	switch reflect.TypeOf(ary).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(ary)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(value, s.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}

// ToMapList 转 map[string]interface{}
func ToMapListByTagJson(data interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	v := reflect.ValueOf(data)
	t := reflect.TypeOf(data)
	for i := 0; i < v.NumField(); i++ {
		m[t.Field(i).Tag.Get("json")] = v.Field(i).Interface()
	}
	return m
}

func ToJsonMap(data interface{}) map[string]interface{} {
	var result map[string]interface{}
	marshal, _ := json.Marshal(data)
	_ = json.Unmarshal(marshal, &result)
	return result
}
