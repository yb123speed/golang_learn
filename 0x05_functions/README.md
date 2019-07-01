# Functions

## 定义

函数是执行特定任务的代码块。函数接受输入，对输入执行一些计算并生成输出。

## 声明

在go中声明函数的一般语法是

```go
func functionname(parametername type) returntype {  
 //function body
}
```

函数声明以func关键字开头，后跟functionname。这些参数之间指定(和)随后returntype的功能。指定参数的语法是参数名称，后跟类型。可以指定任意数量的参数(parameter1 type, parameter2 type)。然后在它之间有一段代码{，}它是函数体。

参数和返回类型在函数中是可选的。因此，以下语法也是有效的函数声明。

```go
func functionname() {  
}
```

如果连续参数属于同一类型，我们可以避免每次都写入类型，并且在结束时写入一次就足够了。即 `price int, no int` 可以写成 `price, no int` 。

## 调用

调用函数的语法是`functionname(parameters)`。可以使用代码调用上述函数。

## 多个返回值

可以从函数返回多个值。

如果函数返回多个返回值，则应在`(`和`)`之间指定它们。

## 命名的返回值

可以从函数返回命名值。如果命名了返回值，则可以将其视为在函数的第一行中声明为变量。

```go
func rectProps(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return //no explicit return value
}
```

area和perimeter是上述函数中的命名返回值。请注意，函数中的return语句不会显式返回任何值。由于area和perimeter在函数声明中指定为返回值，因此遇到return语句时会自动从函数返回它们。

## 空白标识符

_被称为Go中的空白标识符。它可以用来代替任何类型的任何值。

_标识符用于丢弃参数。

