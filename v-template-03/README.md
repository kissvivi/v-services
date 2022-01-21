# v-template-03

#### 介绍
这是一个简单的业务项目脚手架

后期可扩展微服务

修改setting.yaml文件后运行
```
go run main.go server
```

#### 依赖注入工具 wire
#####安装wire
```
go get github.com/google/wire/cmd/wire
```
安装后会在$GOPATH/bin中生成一个可执行程序wire，
这就是代码生成器。
建议将$GOPATH/bin加入系统环境变量$PATH中，
可直接在命令行中执行wire命令。
##### 使用wire命令
wire.go目录下直接指向` wire `命令则自动生成wire_gen.go文件


### 待完善内容
