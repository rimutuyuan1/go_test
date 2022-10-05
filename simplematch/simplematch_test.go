package simplematch

import (
	"fmt"
	"testing"
)

func TestSimplematch(t *testing.T) {
	var tests = []struct {
		A string
		B string
	}{
		{"西风吹老洞庭波，一夜湘君白发多。", "洞庭湖"},
		{"醉后不知天在水，满船清梦压星河。", "满船清梦"},
	}

	for _, str := range tests {
		testsname := fmt.Sprintf("测试内容是否匹配：“%s” vs “%s”", str.A, str.B)
		t.Run(testsname, func(t *testing.T) {
			ret := Simplematch(str.A, str.B)
			fmt.Println("内容匹配", ret)
		})
	}
}
