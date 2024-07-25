// stringutil 包含有用于處理字符串的工具函數。
package tools

import (
	"fmt"
	"net/http"

	"github.com/gobuffalo/packr/v2"
)

// 獲得URL的GET參數
func GetUrlArg(r *http.Request, name string) string {
	var arg string
	values := r.URL.Query()
	arg = values.Get(name)
	return arg
}

// Reverse 將其实参字符串以符文為单位左右反轉。
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Reverse2 將其实参字符串以符文為单位左右反轉。
func Reverse2(s string) string {
	r := []rune(s)
	left := 0
	right := len(r) - 1
	for left < right {
		r[left], r[right] = r[right], r[left]
		left++
		right--
	}
	return string(r)
}

// 獲得文件内容，可以打包到二進制
func FileGetContent(file string) string {
	str := ""
	box := packr.New("tmpl", "../static")
	content, err := box.FindString(file)
	if err != nil {
		return str
	}
	return content
}

func ShowStringByte(str string) {
	s := []byte(str)
	for i, c := range s {
		fmt.Println(i, c)
	}
}
func NilChannel() {
	var ch chan int
	ch <- 1
}
func Int2Str(i interface{}) string {
	return fmt.Sprintf("%v", i)
}
