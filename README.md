<p align="center">
    <img src="./src/logo.png">
</p>
<p align="center">
    a terminal todo manager written in go.
</p>
<p align="center">
<img src="https://goreportcard.com/badge/github.com/jeasonlau/medum?longCache=true&style=for-the-badge">
<img src="https://img.shields.io/badge/license-MIT-orange.svg?longCache=true&style=for-the-badge">
<img src="https://img.shields.io/badge/version-v1.0.0-red.svg?longCache=true&style=for-the-badge">
</p>



## 介绍

medum是一款使用go语言开发的命令行待办事项管理器。

在我的构想中，它应该摈弃掉我认为的不必要的，冗余的功能，专注于deadline的管理，在此基础上尽量做到简单方便、易于使用；但同时，因为有一些个人主观性的取舍，该项目更倾向于个人项目，不大可能因为用户需求改变开发方向。

>   目标并未完全实现，为了简便处理做了较多的妥协（例如时间参数仅仅接受`month.day.hour.minute`格式），这将是接下来的开发方向。

## 初衷

+   最近在学习go语言，准备整个简单项目练练手。

    >   然而更多地用到了数据库知识和第三方package..

+   近期ddl过多，打算开发一个自己能用得到的小工具。

    >   结果为了开发这个工具差点错过ddl..

+   目前在用的[todo](https://github.com/foobuzz/todo)使用较为不便（比如对日期格式的要求过于复杂），同时支持了很多在我看来没有太大必要的功能（比如星级、分类、是否可见等等）。

    >   写完之后发现自己的也没有简单到哪里去（主要是因为不太熟悉golang`time`标准库的用法，踩了不少坑），可能唯一的优点就是运行速度比那一款快了吧..~~废话，如果golang写的还没python快不如去死算了！~~

## 安装

>   在输出中使用到了第三方的emoji package打印emoji表情，暂不确定windows平台的兼容性。

当前并没有打包好的二进制可执行程序，需要自己构建。

+   安装go与git。

+   克隆仓库进行构建：

    ```bash
    git clone https://github.com/jeasonlau/medum
    cd medum
    go build -v .
    ```

+   （可选）将`medum`移至环境变量：

    ```bash
    sudo mv ./medum /usr/bin/medum
    ```

## 用法

```bash
⟩ medum --help
usage: medum [<flags>]

Flags:
      --help                   Show context-sensitive help (also try --help-long
                               and --help-man).
  -b, --begin time=BEGIN TIME  event's begin time.
  -e, --end time=END TIME      event's end time.
  -n, --name=NAME              event's name.
  -d, --delete                 delete outdated events.
```

## 引用

在项目中使用到的第三方package:

+   `github.com/alecthomas/kingpin`：处理命令行参数
+   `github.com/fatih/color`：跨平台的彩色输出
+   `github.com/kyokomi/emoji`：终端输出emoji表情
+   `github.com/mattn/go-sqlite3`：sqlite3驱动
+   `github.com/mitchellh/go-homedir`:跨平台的用户目录获取