string不能修改。可以赋值
string底层是类型指针和长度。



https://gitee.com/tiger1103/gfast快速开发go的golang二次框架



一个查询语句的状态变化是这样的：

MySQL查询语句进入执行阶段后，先把状态设置成 Sending data。
然后，发送执行结果的列相关的信息（meta data) 给客户端。
再继续执行语句的流程。
执行完成后，把状态设置成空字符串。即“Sending data”并不一定是指“正在发送数据”，而可能是处于执行器过程中的任意阶段。比如，你可以构造一个锁等待场景，就能看到Sending data状态。


go使用redis
https://mp.weixin.qq.com/s/0gdqvy5xXaw00JDLwVA6PA



grpc  https://github.com/developer-learning/reading-go

使用，以及写好的grpc池


jukylin
需要好好看下
https://github.com/developer-learning/reading-go


raft实现
 https://github.com/developer-learning/reading-go


# go依赖更新

## go list -m all
查看所有依赖版本
## go mod tidy
自动更新依赖包为最新

## go mod vendor
将依赖包下载到vendor目录下
若依赖是本地的，则不会下载，需要手动添加

