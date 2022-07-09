
讲工作期间go的使用技巧写下。
主要的方向为golang的运维开发
讲工作期间go的使用技巧写下。但不涉及公司源码。
重点是使用go语言一年半了，回头再重新看go语言基础，颇有收获。

每日一库计划：
flag：
常用到的套路，是对命令行的解析。一个程序启动后可以执行[start|stop|restart|check|monitor|reload]
常用的方法
cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")
	check := flag.Bool("check", false, "check collector")

	flag.Parse()
用-c="cfg.json"传入选择的cfg文件。
指定配置文件的路径，如redis-server ./redis.conf以当前目录下的配置文件redis.conf启动 Redis 服务器；

sssssgaidon 

