/* ######################################################################
# Author: (zfly1207@126.com)
# Created Time: 2021-08-31 15:25:27
# File Name: exp_calc.go
# Description:
####################################################################### */

package exp_calc

import (
	"container/list"
	"fmt"
	"strings"
)

// TODO, 逐步补充常用运算符，等于、大于、小于、包含等
var registry = map[string]func(p interface{}, entry *entry) (bool, error){
	"turn": func(p interface{}, entry *entry) (bool, error) { return entry.args.(bool), nil },
}

func Register(name string, fn func(p interface{}, entry *entry) (bool, error)) {
	registry[name] = fn
}

type entry struct {
	isOperator bool
	name       string
	args       interface{}
	operator   string
}

type Calc struct {
	entries []*entry
}

// exp, format see function "exp2Entry":
//   appid:eq:[]
func New(exp string) *Calc {
	o := &Calc{}
	o.entries = o.infixExp2PostfixExp(o.parseExp(exp))
	return o
}

func (this *Calc) Calculate(p interface{}) (r bool, err error) {
	stack := list.New()
	for _, one := range this.entries {
		if _, ok := registry[one.operator]; one.isOperator == false && !ok {
			err = fmt.Errorf("operator#%s not support", one.operator)
			return
		}

		switch one.operator {
		case "&":
			var r1, r2 bool
			e1 := stack.Remove(stack.Back()).(*entry)
			if r1, err = registry[e1.operator](p, e1); err != nil {
				return
			}
			e2 := stack.Remove(stack.Back()).(*entry)
			if r2, err = registry[e2.operator](p, e2); err != nil {
				return
			}
			stack.PushBack(&entry{isOperator: false, operator: "turn", args: r1 && r2})
		case "|":
			var r1, r2 bool
			e1 := stack.Remove(stack.Back()).(*entry)
			if r1, err = registry[e1.operator](p, e1); err != nil {
				return
			}
			e2 := stack.Remove(stack.Back()).(*entry)
			if r2, err = registry[e2.operator](p, e2); err != nil {
				return
			}
			stack.PushBack(&entry{isOperator: false, operator: "turn", args: r1 || r2})
		// case "+":
		//	stack.PushBack(Do(stack.Remove(stack.Back())) - Do(stack.Remove(stack.Back())))
		default:
			stack.PushBack(one)
		}
	}

	e := stack.Remove(stack.Back()).(*entry)
	r, err = registry[e.operator](p, e)
	return
}

// 表达式转计算因子
func (this *Calc) exp2Entry(exp string, isOperator bool) (r *entry) {
	r = &entry{isOperator: isOperator}

	if r.isOperator == true {
		r.operator = exp
		return
	}

	// 将args进行复杂语句解析，如 util.SerializeValue
	t := strings.Split(exp, ":")
	r.name, r.operator, r.args = t[0], t[1], t[2]
	return
}

// 拆解表达式为中缀因子表
func (this *Calc) parseExp(exp string) (r []*entry) {
	r = make([]*entry, 0, 10)

	for first, second := 0, 0; ; {
		if len(exp) <= second {
			if t := strings.TrimSpace(exp[first:second]); len(t) > 0 {
				r = append(r, this.exp2Entry(t, false))
			}
			break
		}
		// space key
		if exp[second] == 32 {
			second++
			continue
		}
		// (、)、*、+、&、| key
		//if exp[second] != 40 && exp[second] != 41 && exp[second] != 42 && exp[second] != 43 {
		if exp[second] != 40 && exp[second] != 41 && exp[second] != 38 && exp[second] != 124 {
			second++
			continue
		}

		if t := strings.TrimSpace(exp[first:second]); len(t) > 0 {
			r = append(r, this.exp2Entry(t, false))
		}
		r = append(r, this.exp2Entry(fmt.Sprintf("%c", exp[second]), true))

		second++
		first = second
	}
	return
}

// 中缀因子表转后缀因子表
func (this *Calc) infixExp2PostfixExp(entries []*entry) (r []*entry) {
	r = make([]*entry, 0, len(entries))

	stack := list.New()
	for _, one := range entries {
		switch one.operator {
		case "(": // 如左括号则入栈
			stack.PushBack(one)
		case ")": // 如右括号则将元素弹出且写入中缀因子表，直到遇到左括号
			for stack.Len() > 0 {
				t := stack.Back()
				if t.Value.(*entry).operator == "(" {
					stack.Remove(t)
					break
				}
				r = append(r, stack.Remove(t).(*entry))
			}
		case "&", "|": // 如操作符，遇到高优先级运算符则将元素弹出且写入中缀因子表，直到遇到低优先级运算符元素
			for stack.Len() > 0 {
				t := stack.Back()
				if t.Value.(*entry).operator == "(" {
					break
				}
				// if (one.operator == "*" || one.operator == "/") && (t.Value.(*entry).operator == "+" || t.Value.(*entry).operator == "-") {
				//  break
				// }
				r = append(r, stack.Remove(t).(*entry))
			}
			stack.PushBack(one)
		default: // 如计算因子，直接写入中缀因子表
			r = append(r, one)
		}
	}

	for stack.Len() > 0 { // 将栈中所有元素输出
		r = append(r, stack.Remove(stack.Back()).(*entry))
	}
	return
}
