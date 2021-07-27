## go-zero使用程序

### 功能试用
+ [ ] 支持中间件，方便扩展
+ [ ] 面向故障编程，弹性设计
+ [ ] 内建服务发现、负载均衡
+ [ ] 内建限流、熔断、降载，且自动触发，自动恢复
+ [ ] API 参数自动校验
+ [ ] 超时级联控制
+ [ ] 自动缓存控制
+ [ ] 链路跟踪、统计报警等
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