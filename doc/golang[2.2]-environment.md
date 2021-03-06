# golang快速入门[2.2]-go语言开发环境配置-macOS
## macos安装Go语言开发包
* 配置go语言的开发环境的第一步是要在[go官网下载页面](https://golang.google.cn/dl/)下载开发包
* macOS需要下载.pkg后缀文件
* macOS和iOS操作系统Apple软件包使用.pkg扩展名,在内部使用Xar格式

![image](../image/17.png)
* 双击下载的安装包即可开始安装,一路点击“继续”即可

![image](../image/18.png)
* 安装包会默认安装在 /usr/local 目录下，如下所示。

![image](../image/19.png)
* 这个目录的结构遵守 GOPATH 规则，后面的章节详细介绍。目录中各个文件夹的含义如下表所示。

| 目录名 | 说明                                                                  |
|--------|-----------------------------------------------------------------------|
| api    | 每个版本的 api 变更差异                                               |
| bin    | go 源码包编译出的编译器（go）、文档工具（godoc）、格式化工具（gofmt） |
| doc    | 英文版的 Go 文档                                                      |
| lib    | 引用的一些库文件                                                      |
| misc   | 杂项用途的文件，例如 Android 平台的编译、git 的提交钩子等             |
| pkg    | linux 平台编译好的中间文件                                          |
| src    | 标准库的源码                                                          |
| test   | 测试用例                                                              |

* 安装完成之后，在终端运行 go version，如果显示类似下面的信息，表明安装成功。
```
» go version                                                                                                                                                                          jackson@192
go version go1.13.6 darwin/amd64
```

## 设置 GOPATH 环境变量
* 开始写 go 项目代码之前，需要我们先配置好GOPATH环境变量，GOPATH是go中的一个重要概念，我们在之后的文章中会详细介绍
* 编辑 ~/.bash_profile（在终端中运行 `vim ~/.bash_profile` ）来添加下面这行代码（如果文件 .bash_profile不存在，则需要创建一个）
* 注意gopath的路径也是可以自己指定的，这里是$HOME/go
* 提示：$HOME 是macos下指代用户主目录的环境变量，可以通过在终端运行 `echo $HOME` 查看具体路径

```
export GOPATH=$HOME/go
```
* 保存然后退出你的编辑器。然后在终端中运行下面命令，代表让上面的命令立即生效
```
source ~/.bash_profile
```
* GOROOT 也就是 Go 开发包的安装目录路径，默认环境变量会自动配置在 /usr/local/go。可以在 bash_profile 文件中设置。
```
export GOROOT=/usr/local/go
```
然后保存并退出编辑器，运行 `source ~/.bash_profile` 命令即可。

* 在任意目录下使用终端执行 go env 命令，输出如下结果说明Go语言开发包已经安装成功

```
» go env                                                                                                                                                                              jackson@192
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/jackson/Library/Caches/go-build"
GOENV="/Users/jackson/Library/Application Support/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="/Users/jackson/go"
GOPRIVATE=""
GOPROXY="direct"
GOROOT="/usr/local/go"
```
## 参考资料
* [在Mac OS上安装Go语言开发包](http://c.biancheng.net/view/3994.html)
