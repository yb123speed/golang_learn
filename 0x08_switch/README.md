# Switch

switch是一个条件语句，它计算表达式并将其与可能的匹配列表进行比较，并根据匹配执行代码块。它可以被认为是编写多个`if else`子句的惯用方式。

**不允许具有相同常量值的重复case。**

## Default case

当其他情况都不匹配时，将执行默认case。

**默认值不一定是switch语句中的最后一种情况。它可以出现在交换机的任何位置。**

## 多表达式

switch中的表达式是可选的，可以省略。如果省略表达式，则认为switch true并且评估每个case表达式是否为真，并且执行相应的代码块。

## Fallthrough（下通）

fallthrough语句用来控制执行当前case后穿透到下一个case。

Switch和case表达式不一定只是常量。它们也可以在运行时进行评估。

fallthrough应该是case最后一个语句。如果它出现在中间某处，编译器将抛出错误fallthrough statement out of place