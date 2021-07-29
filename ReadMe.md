## go-zero使用程序

### 功能试用
+ [ ] 支持中间件，方便扩展
+ [ ] 面向故障编程，弹性设计
+ [ ] 内建服务发现、负载均衡
+ [ ] 内建限流、熔断、降载，且自动触发，自动恢复
+ [ ] API 参数自动校验
+ [ ] 超时级联控制
+ [x] 数据库缓存代码自动化生成，自动缓存控制
  ```shell
  # 内置工具通过sql文件自动化生成，也可以通过数据库表生成
  
  # 带redis缓存版本，操作MySQL版本
  goctl model mysql ddl -c -src user.sql -dir ./model
  
  # 不带redis缓存版本，只操作MySQL版本
  goctl model mysql ddl -src user.sql -dir ./model
  ```
+ [ ] 链路跟踪 
+ [ ] 服务监控
  服务监控目前没有支持服务对redis和mysql的调用监控。
  + promethus部署
  ```shell

  ```  
+ [ ] 日志收集  
+ [ ] 高并发支撑，高并发




### 运行
```shell
# 运行rpc服务
cd rpc
make run  

# 运行rest服务
cd api
make run 
```

### 创建用户
```shell
curl -v -X POST "http://localhost:8888/v1/user" \
  -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"password\": \"123456\", \"username\": \"frank\"}"
```

### 获取用户信息
```shell
curl -v "http://localhost:8888/v1/user/4"
```

### 参考资料
+ [官方文档](https://go-zero.dev/cn/)