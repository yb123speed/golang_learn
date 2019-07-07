# Loops

循环语句用于重复执行代码块。

for是Go中唯一可用的循环。Go没有while或do while循环，这些循环存在于其他语言如C.

## for循环语法

```go
for initialisation; condition; post {  
}
```

初始化语句只执行一次。初始化循环后，将检查条件。如果条件的计算结果为true，则{ }执行循环体，然后执行post语句。post语句将在每次成功迭代循环后执行。执行post语句后，将重新检查该条件。如果是，则循环将继续执行，否则for循环终止。

在Go中，所有三个组件即初始化，条件和发布都是可选的。

## break
该break语句用于在完成正常执行之前突然终止for循环，并将控件移动到for循环之后的代码行。

## continue

该continue语句用于跳过for循环的当前迭代。在continue语句之后的for循环中出现的所有代码都不会在当前迭代中执行。循环将继续下一次迭代。

## Labels

Labels可用于从内部for循环内部打破外部循环。

## infinite loop

创建无限循环的语法是，

```go
for {  
}
```