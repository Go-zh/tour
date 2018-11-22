# Go 语言之旅

《Go 语言之旅》是官方 Go Tour 的中文翻译版。
请访问 https://tour.go-zh.org/ 开始学习。

## 下载/安装

要从源码安装本教程，首先请[设置一个工作空间](https://go-zh.org/doc/code.html)并执行

	$ go get -u github.com/Go-zh/tour/gotour

这会在你工作空间的 `bin` 目录中创建一个可离线执行的 `gotour` 文件。

（如果安装过程中出现 `package` 或 `import` 字样的错误提示，那么说明依赖库的导入路径又挂了。这时请猛戳 @OlingCat 并督促其解决= =||）

## 贡献方式

贡献方式应遵循与 Go 项目相同的流程：http://go-zh.org/doc/contribute.html

要在本地运行 tour 服务，请执行

```sh
dev_appserver.py app.yaml
```

然后在你的浏览器中浏览 [http://localhost:8080/](http://localhost:8080)。

## 问题报告/发送补丁

本教程中文版直接托管在 Github 上，提交更改请直接发送 PR。

问题报告请在 github.com/Go-zh/tour/issues 上发起。

## 授权许可

除特别声明外，go-tour 源码文件均采用 BSD 风格的授权许可分发，许可内容见 `LICENSE` 文件。
