# golang
|Package|Description|
|---|---|
|data_type|Data type|
|operator|Operator|
|condition|Condition|
|funcs|function|
|oop|struct and interface|
|err|error handling|
|generic|generic paradigm|
|concurrent|concurrent|
|reflect|reflect|
|random|random|
|io|io|

# data_type
防止溢出，类型分类，存储范围，字面量，默认值，类型转换

# operator
赋值运算符，算术运算符，比较运算符，逻辑运算符，位运算符，指针相关运算符
- 只有a++和a--

# condition
条件，循环
- 默认无case穿透，需要穿透要加上fallthrough

# funcs
文件函数，结构体函数，闭包-匿名函数，多值返回，函数作入参，访问控制
- 文件函数都是包下唯一，首字母大写函数可以在其它包中调用

# oop
结构体，接口
- 接口由 方法的聚合体 》类型的聚合体

# err
定义异常，抛出异常，捕获异常，处理异常defer

# generic
- 类型不同但操作相同
- 形参，实参，约束，实例化
- 泛型 slice, map, chan，func，struct，interface{}，receiver（成员方法）

interface与泛型
- 只含有方法的接口
- 只含有类型的接口
- 含有方法和类型的接口

# concurrent
- 任务, goroutine, MPG模型, chan，select，context, 锁, sync, `sync/atomic`
- MPG: M(main 主线程) P(协程执行所需上下文) G(协程)
- `sync/atomic`：变量进行原子性操作

# reflect
- interface{}包含类型指针`_type *_type`，数据指针`data unsafe.Pointer`
- reflect.Type 反射类型
- reflect.Value 反射值

# random
- 未设定种子数而产生的随机数是固定数

# io