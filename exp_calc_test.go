/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2021-08-31 15:25:27
# File Name: exp_calc_test.go
# Description:
####################################################################### */

package exp_calc

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Register("in", func(p interface{}, entry *Entry) (bool, error) {
		return true, nil
	})
	os.Exit(m.Run())
}

func TestBasic(t *testing.T) {
	p := &struct {
		user int32
		Age  int32
		sex  int32
	}{3, 20, 2}

	o := New("appid:in:[1,2,3] & (age:in:[10,20,30] | sex:in:[1,2])")
	fmt.Println(o.Calculate(p))
}

/*
func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = aa()
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bb()
	}
}

func BenchmarkC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = cc()
	}
}

func BenchmarkD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = dd()
	}
}

func BenchmarkE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ee()
	}
}

func BenchmarkF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculate("true&(false|true)")
	}
}

var str = "true&(false|true)"

func aa() []string {
	exp := str
	exp = strings.ReplaceAll(exp, " ", "")

	// 拆解表达式为中缀因子表
	infixEntries := make([]string, 0, 10)
	for i, v := 0, strings.Split(exp, "("); i < len(v); i++ {
		for i, v := 0, strings.Split(v[i], ")"); i < len(v); i++ {
			for i, v := 0, strings.Split(v[i], "*"); i < len(v); i++ {
				for i, v := 0, strings.Split(v[i], "+"); i < len(v); i++ {
					if len(v[i]) == 0 {
						continue
					}
					infixEntries = append(infixEntries, v[i])
					util.IfDo(i < len(v)-1, func() { infixEntries = append(infixEntries, "+") })
				}
				util.IfDo(i < len(v)-1, func() { infixEntries = append(infixEntries, "*") })
			}
			util.IfDo(i < len(v)-1, func() { infixEntries = append(infixEntries, ")") })
		}
		util.IfDo(i < len(v)-1, func() { infixEntries = append(infixEntries, "(") })
	}

	return infixEntries
}

func bb() []string {
	exp := str
	infixEntries := make([]string, 0, 10)
	exp = strings.ReplaceAll(exp, "(", " ( ")
	exp = strings.ReplaceAll(exp, ")", " ) ")
	exp = strings.ReplaceAll(exp, "*", " * ")
	exp = strings.ReplaceAll(exp, "+", " + ")

	for _, v := range strings.Split(exp, " ") {
		if len(v) == 0 {
			continue
		}
		infixEntries = append(infixEntries, v)
	}

	return infixEntries
}

func cc() []string {
	exp := strings.ReplaceAll(str, " ", "")
	infixEntries := make([]string, 0, 10)

	for first, second := 0, 0; ; {
		if len(exp) <= second {
			if entry := strings.TrimSpace(exp[first:second]); len(entry) > 0 {
				infixEntries = append(infixEntries, entry)
			}
			break
		}
		// space key
		if exp[second] == 32 {
			second++
			continue
		}
		// (、)、*、+ key
		if exp[second] != 40 && exp[second] != 41 && exp[second] != 42 && exp[second] != 43 {
			second++
			continue
		}

		if entry := strings.TrimSpace(exp[first:second]); len(entry) > 0 {
			infixEntries = append(infixEntries, entry)
		}
		infixEntries = append(infixEntries, fmt.Sprintf("%c", exp[second]))

		second++
		first = second
	}

	return infixEntries
}

func ee() []string {
	exp := []byte(str)
	infixEntries := make([]string, 0, 50)

	for first, second := 0, 0; ; {
		if len(exp) <= second {
			if entry := bytes.TrimSpace(exp[first:second]); len(entry) > 0 {
				infixEntries = append(infixEntries, string(entry))
			}
			break
		}
		// space key
		if exp[second] == 32 {
			second++
			continue
		}
		// (、)、*、+ key
		if exp[second] != 40 && exp[second] != 41 && exp[second] != 42 && exp[second] != 43 {
			second++
			continue
		}

		if entry := bytes.TrimSpace(exp[first:second]); len(entry) > 0 {
			infixEntries = append(infixEntries, string(entry))
		}
		infixEntries = append(infixEntries, fmt.Sprintf("%c", exp[second]))

		second++
		first = second
	}

	return infixEntries
}

func dd() []string {
	exp := []byte(str)

	for idx := 0; idx < len(exp); idx++ {
		if exp[idx] != 40 && exp[idx] != 41 && exp[idx] != 42 && exp[idx] != 43 {
			continue
		}
		t := []byte{}
		t = append(t, exp[:idx]...)
		t = append(t, 32, exp[idx], 32)
		t = append(t, exp[idx+1:]...)
		exp = t
		idx += 2
	}
	infixEntries := strings.Split(string(exp), " ")

	return infixEntries
}
*/

// vim: set noexpandtab ts=4 sts=4 sw=4 :
