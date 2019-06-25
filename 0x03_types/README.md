# Types

## 以下是go中可用的基本类型

- bool
- Numeric Types
    - int8, int16, int32, int64, int
    - uint8, uint16, uint32, uint64, uint
    - float32, float64
    - complex64, complex128
    - byte
    - rune
- string

## 复数的类型

complex64:具有float32实部和虚部的复数

complex128:具有float64实部和虚部的复数

内置函数`complex`用于构造具有实部和虚部的复数。
 
复杂函数具有以下定义

```go
func complex(r, i FloatType) ComplexType  
```

它将实部和虚部作为参数并返回复杂类型。实部和虚部都应该是相同的类型。即float32或float64。如果实部和虚部都是float32，则此函数返回complex64类型的复数值。如果实部和虚部都是float64类型，则此函数返回complex128类型的复数值。

## 其他数字类型

byte是uint8 的别名

rune是int32的别名

## 类型转换

Go对显式输入非常严格。没有自动类型提升或转换。

`T(v)` 是将值v转换为类型T的语法。