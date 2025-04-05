## 修改docker容器端口
1.重启一个容器
2.在之前容器基础上commit，在开端口。
3.在运行中添加，繁琐。修改容器配置文件，重启docker服务
找到容器id，打开 hostconfig.json 配置文件
vim /var/lib/docker/containers/{hash_of_the_container}/hostconfig.json
vim hostconfig.json
修改"PortBindings":{} 这个配置项


