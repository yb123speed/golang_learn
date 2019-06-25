## 声明一个变量

`var name type` 是声明单个变量的语法。

## 声明具有初始值的变量

`var name type = initialvalue` 是声明具有初始值的变量的语法。

## 类型推断

如果使用语法 `var name = initialvalue` 声明变量，Go将自动从初始值推断出该变量的类型。

## 多变量声明

`var name1，name2 type = initialvalue1，initialvalue2` 是多变量声明的语法。


## 在某些情况下，我们可能希望在单个语句中声明属于不同类型的变量。这样做的语法是:

```go
var (  
      name1 = initialvalue1
      name2 = initialvalue2
)
```

## 简短的声明

`name := initialvalue` 是声明变量的简写语法。

`:=` 左侧没有新的变量,它会抛出错误no new variables on left side of :=。即 `:=` 左侧至少有一个新变量。

由于Go是强类型的，因此声明属于一种类型的变量不能赋予另一种类型的值。
