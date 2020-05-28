<p align="center">
    <img src="./src/logo.png">
</p>
<p align="center">
    a terminal todo manager written in go.
</p>
<p align="center">
<img src="https://goreportcard.com/badge/github.com/amtoaer/medum?longCache=true&style=for-the-badge">
<img src="https://img.shields.io/badge/license-MIT-orange.svg?longCache=true&style=for-the-badge">
<img src="https://img.shields.io/badge/version-v1.1.4-red.svg?longCache=true&style=for-the-badge">
</p>




## 介绍

medum是一款使用go语言开发的命令行待办事项管理器。

在我的构想中，它应该摈弃掉我认为的不必要的，冗余的功能，专注于deadline的管理，在此基础上尽量做到简单方便、易于使用；但同时，因为有一些个人主观性的取舍，该项目更倾向于个人项目，不大可能因为用户需求改变开发方向。

>   目标并未完全实现，为了简便处理做了较多的妥协（例如时间参数仅仅接受`month.day.hour.minute`格式），这将是近期的开发目标。

## 初衷

+   最近在学习go语言，准备整个简单项目练练手。

    >   然而更多地用到了数据库知识和第三方package..

+   近期ddl过多，打算开发一个自己能用得到的小工具。

    >   结果为了开发这个工具差点错过ddl..

+   目前在用的[todo](https://github.com/foobuzz/todo)使用较为不便（比如对日期格式的要求过于复杂），同时支持了很多在我看来没有太大必要的功能（比如星级、分类、是否可见等等）。

    >   写完之后发现自己的也没有简单到哪里去（主要是因为不太熟悉golang`time`标准库的用法，踩了不少坑），可能唯一的优点就是运行速度比那一款快了吧..~~废话，如果golang写的还没python快不如去死算了！~~


## 安装

>   注：windows用户请使用微软全新的命令行工具`windows terminal`，否则无法正常运行。
>
>   感谢[@Ray-Keiyaku](https://github.com/Ray-Keiyaku)同学的测试。
>
>   `powershell`:
>
>   ![powershell](https://i.loli.net/2020/05/28/l8DrgIikwUKvFn1.png)
>
>   `windows terminal`:
>
>   ![windows terminal](https://i.loli.net/2020/05/28/NRwEbDW64GVPorT.png)

1.  （linux/windows）下载[release](https://github.com/amtoaer/medum/releases)页面中对应的的二进制程序，参照[#用法](#用法)运行。

    >   注：每次更新内容在该页面都有详细说明，可自己酌情选择版本（推荐最新）。

2.  手动构建。

+   安装go与git。

+   克隆仓库进行构建：

    ```bash
    git clone https://github.com/amtoaer/medum
    cd medum
    go build -v .
    ```

+   将`medum`移至环境变量（以linux为例）：

    ```bash
    sudo mv ./medum /usr/bin/medum
    ```

## 用法

```bash
⟩ medum --help
NAME:
   medum - a terminal todo manager

USAGE:
   medum [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --begin value, -b value  event's begin time
   --end value, -e value    event's end time
   --name value, -n value   event's name
   --remove, -r             remove outdated events (default: false)
   --done value, -d value   the id of your finished event (default: 0)
   --help, -h               show help (default: false)
```

常用操作演示：

```bash
# 当前时间为5.12日晚

# 插入新事件，deadline为5/15，12:00（缺省-b，起始时刻默认为当前时刻）
⟩ medum -e 5.15.12.00 -n 测试1
successfully insert event

# 插入新事件，自己指定起始时间
⟩ medum -e 5.15.12.00 -b 5.14.12.00 -n 测试2
successfully insert event

# 插入过期事件
⟩ medum -e 5.11.12.00 -b 5.10.12.00 -n 测试3
successfully insert event

# 打印事件
⟩ medum
3 | 🍺 测试3 | ⌛ no time remaining
1 | 🍺 测试1 | ⌛ 2 days remaining
2 | 🍺 测试2 | ⌛ 1 days beginning

# 删除过期事件
⟩ medum -r
successfully delete outdated events

# 再次打印事件，发现过期事件已经被删除
⟩ medum
1 | 🍺 测试1 | ⌛ 2 days remaining
2 | 🍺 测试2 | ⌛ 1 days beginning

# 完成序号为1的测试1，将其删除
⟩ medum -d 1
successfully delete event with specific ID

# 再次打印事件，测试1已经被删除
⟩ medum
2 | 🍺 测试2 | ⌛ 1 days beginning
```

## 截图

以下为我在`linux`上使用的截图：

![截图录屏_选择区域_20200513005252](https://i.loli.net/2020/05/13/zycFTamlK8uM79A.png)

## 配置

该项目支持自定义输出颜色，配置文件位于`~/.medum/config.json`，内容如下：

```json
{
    "NumberColor": "red",
    "EventColor": "blue",
    "TimeColor": "yellow"
}
```

具体来说，指的是：

```bash
8 | 🍺 数电作业 | ⌛ 13 hours remaining
↑        ↑               ↑
red     blue           yellow
```

允许的颜色共有十六种，分别为：

```go
var funcs = map[string]interface{}{
	"red":       color.New(color.FgRed),
	"blue":      color.New(color.FgBlue),
	"cyan":      color.New(color.FgCyan),
	"green":     color.New(color.FgGreen),
	"yellow":    color.New(color.FgYellow),
	"magenta":   color.New(color.FgMagenta),
	"white":     color.New(color.FgWhite),
	"black":     color.New(color.FgBlack),
	"hired":     color.New(color.FgHiRed),
	"hiblue":    color.New(color.FgHiBlue),
	"hicyan":    color.New(color.FgHiCyan),
	"higreen":   color.New(color.FgHiGreen),
	"hiyellow":  color.New(color.FgHiYellow),
	"himagenta": color.New(color.FgHiMagenta),
	"hiwhite":   color.New(color.FgHiWhite),
	"hiblack":   color.New(color.FgHiBlack),
}
```

## 引用

在项目中使用到的第三方package:

+   `github.com/urfave/cli/v2`：处理命令行参数
+   `github.com/fatih/color`：跨平台的彩色输出
+   `github.com/mattn/go-sqlite3`：sqlite3驱动
+   `github.com/mitchellh/go-homedir`:跨平台的用户目录获取