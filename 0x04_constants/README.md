# Constants

## 定义

术语常量在Go中用于表示固定值，例如5,-89,"I love Go",67.89等。

## 赋值

`const variable = value`

常量是应在编译时知道的值。因此，不能分配给它函数调用返回的值，因为函数调用是在运行时进行的。

## 字符串常量

双引号之间的任何值都是Go中的字符串常量。例如，像“Hello World”或“Sam”这样的字符串都是Go中的常量。

字符串常量属于什么类型？答案是他们是无类型的。

答案是无类型常量具有与它们相关联的默认类型，并且当且仅当一行代码需要它时才提供它。

有没有办法创建一个类型常量？答案是肯定的。以下代码创建一个类型化常量。

```go
const typedhello string = "Hello World"
```

上面代码中的typedhello是string类型的常量。

Go是一种强类型语言。不允许在分配期间混合类型。让我们通过一个程序来看看这意味着什么。

```go
package main

func main() {  
        var defaultName = "Sam" //allowed
        type myString string
        var customName myString = "Sam" //allowed
        customName = defaultName //not allowed

}
```

## 数字表达式

数值常量可以在表达式中自由混合和匹配，只有在将它们分配给变量或在需要类型的代码中的任何位置使用时才需要类型。