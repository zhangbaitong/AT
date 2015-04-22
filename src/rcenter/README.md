#资源注册中心

##实现功能

1. 注册资源:res/reg
2. 停用资源:res/stop
3. 请求资源:res/get
4. 更新资源:res/update

##代码说明

启动
```
go run rcenter/main.go
```
测试功能
```
go run rcenter/test.go
```
测试提交请求
```
go run rcenter/httptest.go
```


##改进计划

1. 使用新的路由实现（github.com\julienschmidt\httprouter）
2. 调试HTTPS问题
3. 联合调试