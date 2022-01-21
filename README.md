# v-services

这是一个简单golang微服务脚手架

### 编译运行
先编译docker文件
```shell
docker build -t v-template-01 .\v-template-01\.
```
然后使用docker-compose运行
```
docker-compose up
```
### 访问
端口号：8888为注册中心，所有访问进入到注册中心转发

### 待完善内容
1. nacos接入
2. make文件编写
3. 轻量化哥服务
4. etcd客户端注册