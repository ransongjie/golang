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

# data_type
防止溢出，类型分类，存储范围，字面量，默认值，类型转换

# operator
赋值运算符，算术运算符，比较运算符，逻辑运算符，位运算符，指针相关运算符
- 只有a++和a--

# control
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
泛型
- 使用时机：经常要分别为不同的类型写完全相同逻辑的代码，使用泛型将是最合适的选择
- 注意：任何泛型类型都必须传入类型实参 实例化后才可以使用

可用
- 泛型切片, map, chan
- 泛型函数
- 泛型结构体, 接口, 接收器, 没有泛型方法

# concurrent
- 任务, goroutine, MPG模型, chan, 锁, sync
- 主线程执行完毕协程退出

MPG: M(main 主线程) P(协程执行所需上下文) G(协程)
- M0上G0被阻塞, 新开M1或从线程池中拿M1执行M0上G1...

