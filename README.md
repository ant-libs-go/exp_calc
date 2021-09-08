# exp\_calc
表达式计算器

# 功能
 - 参考 xorm、go-playground/validator.v8 等框架，用于实现复杂、灵活的逻辑判断，可用于业务中复杂的条件判断。如广告定向条件、流量分层策略等


# 基本使用

```golang
// 注册操作符
Register("in", func(p interface{}, entry *entry) (bool, error) {
	// 解析p中的参数，与entry中的args进行对比判断
	code ...
})

// 创建calc实例
calc := New("appid:in:[1001] & (age:in:{"range": [10,30]} | sex:in:[1,2])")

// 定义参数
p := &struct {
	user int32
	Age  int32
	sex  int32
}{3, 20, 2}

fmt.Println(o.Calculate(p))
```
