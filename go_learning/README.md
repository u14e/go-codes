## 安装

下载并安装: https://golang.google.cn/dl/
GOROOT: 默认 C:\Go\
GOPATH: 默认 %USERPROFILE%\go\
设置 GOPATH: setx GOPATH D:\Go (关闭终端，开启新终端即可)

> 参考链接: 
  - https://github.com/golang/go/wiki/SettingGOPATH
  - https://golang.google.cn/doc/install

## 安装 GoLand

下载并安装: https://www.jetbrains.com/go

激活: 
- https://github.com/xmge/gonote/blob/master/goland使用指南/2.安装.md
- http://idea.lanyus.com/

设置: 主题、字体(Hack)、go 语言 indent 为 2

## 应用程序入口(ch1)

1. 必须是 main 包: package main (目录名可以不是 main)
2. 必须是 main 方法: func main()
3. 文件名不一定是 main.go

## 退出返回值(ch1)

1. Go 中的 main 函数不支持任何返回值
2. 通过 os.Exit 来返回状态
3. 在程序中直接通过 os.Args 获取命令行参数

## 编写测试程序

1. 源码文件以 _test 结尾: xxx_test.go
2. 测试方法以 Test 开头: func TestXXX(t *testing.T) {...}

## 变量声明和赋值(ch2)

```go
// 变量
a := 1
// 常量
const (
    Monday = iota + 1
    Tuesday
    Wednesday
)
```

## 基本数据类型

bool
string
int int8 int16 int32 int64
unit unit8 unit16 unit32 unit64 uintptr
byte  // alias for unit8
rune  // alias for unit32, represents a Unicode code point
float32 float64
complex64 complex128

## 类型转化(ch3)

1. Go 语言不允许隐式类型转换
2. 别名和原有类型也不能进行隐式类型转换

## 指针类型(ch3)

1. 不支持指针运算
2. string 是值类型，其默认的初始值为空字符串，而不是 nil

## 运算符(ch4-按位清零&^)

## 用 == 比较数组(ch4)

1. 相同维度且含有相同个数元素的数组才可以比较
2. 每个元素都相同才相等

## 循环(ch5)

## if 条件(ch5)

1. condition 表达式结果必须为布尔值
2. 支持变量赋值

```go
if var declaration; condition {
  // code here
}
```

## case语句(ch5)

1. 条件表达式不不限制为常量量或者整数；
2. 单个 case 中，可以出现多个结果选项, 使⽤用逗号分隔；
3. 与 C 语⾔言等规则相反，Go 语⾔言不不需要⽤用break来明确退出⼀一个 case；
4. 可以不不设定 switch 之后的条件表达式，在此种情况下，整个 switch 结
构与多个 if…else… 的逻辑作⽤用等同

## 数组的声明

```go
// 声明并初始化为默认零值
var a [3]int
a[0] = 1

// 声明同时初始化
b := [3]int{1, 2, 3}

// 多维数组
c := [2][2]int{{1, 2}, {3, 4}}
```

## 数组 vs 切片(ch6)

数组是一个容器, 包含一段连续的存储空间

切片是一个共享的存储结构

切片实际上是一个结构体, 包含三个元素, 
第一个是一个指针， 指向一片连续的存储空间(也就是数组)
第二个是切片长度, 表示可访问元素个数
第三个是切片容量, 表示指针指向的内部数组的空间长度

Q: 为什么是 s = append(s, 1)
A: 因为结构体指向的连续存储空间地址发生了变化, 并不总是在连续存储空间后面添加值
当存储空间扩展时, 使用新的存储空间, 并把原来的值拷贝进来

1. 数组容量不可伸缩，切片可以
2. 数组可以比较，切片不可以

## Map 声明(ch7)

```go
m := map[string]int{"one": 1, "two": 2, "three": 3}
m1 := map[string]int{}
m1["one"] = 1

m2 := make(map[string]int, 10 /* Initial Capacity */)
```

## 实现 Set(ch7)

Go 的内置集合中没有 Set 实现，可以 map[type]bool

1. 元素的唯一性
2. 基本操作
    - 添加元素
    - 判断元素是否存在
    - 删除元素
    - 元素个数

## 字符串(ch8)

1. string 是数据类型，不是引用或指针类型
2. string 是只读的 byte slice, len 函数返回它所包含的 byte 数
3. string 的 byte 数组可以存放任何数据

## Unicode 和 UTF8(ch8)

1. Unicode 是一种字符集 (code point)
2. UTF8 是 Unicode 的存储实现 (转换为字节序列的规则)

## 编码和存储

| 字符 | "中" |
| --- | --- |
| Unicode | 0x4E2D |
| UTF-8 | 0xE4B8AD |
| string/[]byte | [0xE4, 0xB8, 0xAD] |

## 函数是一等公民(ch9)

1. 可以有多个返回值
2. 所有参数都是值传递: slice, map, channel 会有传引用的错觉
3. 函数可以作为变量的值
4. 函数可以作为参数和返回值

## 面向对象(ch10-ch12)

## 空接口与断言(ch12)

1. 空接口可以表示任何类型
2. 通过断言来将空接口转换为指定类型

```go
v, ok := p.(int)  // ok = true 时为转换成功
```

## Go 的错误机制(ch13)

1. 没有异常机制

2. error 类型实现了 error 接口

    ```go
    type error interface {
      Error() string
    }
    ```

3. 可以通过 errors.New 来快速创建错误实例

    ```go
    errors.New("This is a error.")
    ```

4. 最佳实践: 及早失败，避免嵌套

## panic(ch13)

1. panic 用于不可以恢复的错误
2. panic 退出前会执行 defer 指定的内容

## os.Exit

1. os.Exit 退出时不会调用 defer 指定的函数
2. os.Exit 退出时不输出当前调用栈信息

## 最常见的 "错误恢复"(ch13)

```go
defer func() {
  if err := recover(); err != nil {
    log.Error("recovered panic", err)
  }
}()
```

## package(ch14)

1. 基本复用模块单元(以首字母大写来表明可被包外代码访问)
2. 代码的 package 可以和所在的目录不一致
3. 同一目录里的 Go 代码的 package 要保持一致

## init 方法(ch14)

1. 在 main 被执行前，所有依赖的 package 的 init 方法都会被执行
2. 不同包的 init 函数按照包导入的依赖关系决定执行顺序
3. 每个包可以有多个 init 函数
4. 包的每个源文件也可以有多个 init 函数，这点比较特殊

## 使用远程包(ch14)

```bash
go get -u github.com/easierway/concurrent_map
```

## vendor 路径

查找依赖包路径的解决方案:
1. 当前包下的 vendor 目录
2. 向上级目录查找，直到找到 src 下的 vendor 目录
3. 在 GOPATH 下面查找依赖包
4. 在 GOROOT 目录下查找

## [glide](https://glide.readthedocs.io/en/latest/)

```bash
# 安装
go get -u github.com/Masterminds/glide

glide init
glide install
glide get ...
```

> 参考链接: 
> - https://studygolang.com/topics/4905
> - https://github.com/Masterminds/glide/issues/873#issuecomment-325311755

## 协程(ch15)
## 共享内存(ch16)
## CSP 并发机制(ch17)
## 多路选择实现超时控制(ch18)

## channel 的关闭(ch19)

- 向关闭的 channel 发送数据，会导致 panic
- v, ok <-ch; ok 为 bool 值，true 表示正常接受, false 表示通道关闭
- 所有的 channel 接收者都会在 channel 关闭时，立刻从阻塞等待中返回且上述 ok 值为 false。这个广播机制常被利用，进行多个订阅者同时发送信号(如：退出信号)

## Context(ch19)

- 根 Context: 通过 context.Background() 创建
- 子 Context: context.WithCancel(parentContext) 创建
    - ctx, cancel := context.WithCancel(context.Background())
- 当前 Context 被取消时, 基于他的子 context 都会被取消
- 接受取消通知: <-ctx.Done()

## 文档

- 本地: `godoc --http :8888`
- 中文在线: https://studygolang.com/pkgdoc