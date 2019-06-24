# 运行go程序

运行go程序有几种不同的方法。让我们逐一看看它们。

1. 使用go run命令 
    - 在命令提示符下键入 
    ```
    go run workspacepath/src/helloworld/helloworld.go
    ```
    - 上面命令中的workspacepath应该替换为工作空间的路径（在windows中C:/Users/YourName/go，在linux或Mac中$HOME/go）

    - 您应该Hello World在控制台中看到输出。

2. 使用go install命令 
    - 运行go install hello命令，然后workspacepath/bin/hello运行程序。
    - 上面命令中的workspacepath应该替换为工作空间的路径（在windows中C:/Users/YourName/go，在linux或Mac中$HOME/go）
    - 当你输入go install hello时，go工具会在工作区内搜索hello包（hello被称为包，我们将在后面详细介绍包中）。然后它在工作空间的bin目录中创建一个名为hello（hello.exe在windows的情况下）的二进制文件。运行go install hello后，目录结构如下所示
```
       go
         bin  
           hello
         src
           hello
               helloworld.go
```