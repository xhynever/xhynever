1 for 循环      for应该是使用最高频的一个命令了，如果需要处理一些简单的东西 for 循环就足够了   for i in `cat file`;do echo $i ;done 
2 awk 按某一列相同变量求和 求和：awk 'BEGIN{sum=0}{sum+=$1}END{print sum}' data.txt变量：awk -v var='bianliang' 'BEGIN{print var}'
3 某一列相加awk '{a[$1]+=$2}END{for(i in a){print i,a[i]}}'

4 统计个数按列统计个数，统计第二列各个阈值内个数的分布。awk '$2>2000{a+=1}$2>=1000&&$2<2000{b+=1}$2>=500&&$2<1000{c+=1}$2<500&&$2>200{d+=1}$2<200{e+=1}END{print ">2000 "a"\n1000-2000 "b"\n500-1000 "c"\n<200-500 "d"\n<200 "e}'

5 两个文件 按 列合并
awk 'NR==FNR{a[i]=$0;i++}NR>FNR{print a[j]","$0;j++}' 50host.cp 50hosts

a文件	b文件	合并
1	a	1，a
2	b	2，b
3	c	3，c

6 两个文件求交集
grep -v -f file1 file2 && grep -v -f file2 file1
但是这个命令比较慢 文件小的时候可以使用

comm 这个命令比较快 需要提前排序
comm -3 file1 file2
7 扫盘
fuser -km /dev/sda3 杀死所有进程
把sda3先umount下来
然后fsck -y /dev/sda3

8 日志清理
find -mtime +3 -name "lock.log.*" -exec rm -rf {} \;
-mtime +3 3天以上 -3 3天以内
{}是find的结果集合， \;是结束命令
9 sed 修改第几个变量
sed -i "/^ip:.*/{x;s/^/./;/^\.\{1\}$/{x;s/.*/ip:$metaserver_ip_0/;x};x;}
10 shell 中 获取当前所在目录

basepath=$(cd `dirname $0`; pwd)

11 查看线程数
pstree -p `ps -e | grep master | awk '{print $1}'` |wc -l


12 source

source是子脚本里面的变量能传递给父脚本;一般情况下需要把脚本里面的变量反馈处理就需要source命令，也是在日常工作非常高频的命令之一
写一个sub.sh 里面test=“a”
14 rsync    一开始的工作的时候线上部署还是最原始的状态的状态需要写脚本去同步，就是用的是rsync 这里面的好处是包括了批量同步，已经增量同步rsync -avtP  ./ok dec:/tmp/ rsync还有一个妙用就是批量删除大量文件第一步先创建一个空文件  mkdir  /tmp/blank 第二步就是把空文件同步进去要删除的文件里面  rsync --delete-before -d /tmp/blank/  ./dest/

15 lsof 这个命令是查看端口被占用最常见的命令，比如你一个服务启动不起来发现端口被占用了就需要看一下是哪个进程占用的就可以看一下

losf -i:8080
16 ps aux

有的人喜欢用 ps -ef 这两个参数其实区别不大大家按照自己的情况来使用即可
这里可以复习一个pstree root 查看某个用户下所有服务
pstree -p 700 查看这个进程下的所有线程数量
17 tail -f 看日志滚动的时候会经常用到这里有个问题大家可以想一下查看日志应该如何避免日志被打爆的情况

接下来的几个命令会跟网络定位有关系

### 根据pid，找到相关信息
作者：咸鱼Linux运维
链接：https://zhuanlan.zhihu.com/p/612782776
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

#! /bin/bash

read -p "请输入要查询的PID: " P

n=`ps -aux| awk '$2~/^'${P}'$/{print $0}'|wc -l`

if [ $n -eq 0 ];then
 echo "该PID不存在！！"
 exit
fi
echo -e "\e[32m--------------------------------\e[0m"
echo "进程PID: ${P}"
echo "进程命令：$(ps -aux| awk '$2~/^'$P'$/{for (i=11;i<=NF;i++) printf("%s ",$i)}')"
echo "进程所属用户: $(ps -aux| awk '$2~/^'$P'$/{print $1}')"
echo "CPU占用率：$(ps -aux| awk '$2~/^'$P'$/{print $3}')%"
echo "内存占用率：$(ps -aux| awk '$2~/^'$P'$/{print $4}')%"
echo "进程开始运行的时间：$(ps -aux| awk '$2~/^'$P'$/{print $9}')"
echo "进程运行的时间：$(ps -aux| awk '$2~/^'$P'$/{print $10}')"
echo "进程状态：$(ps -aux| awk '$2~/^'$P'$/{print $8}')"
echo "进程虚拟内存：$(ps -aux| awk '$2~/^'$P'$/{print $5}')"
echo "进程共享内存：$(ps -aux| awk '$2~/^'$P'$/{print $6}')"
echo -e "\e[32m--------------------------------\e[0m"

### 查看tcp连接情况
作者：咸鱼Linux运维
链接：https://zhuanlan.zhihu.com/p/612782776
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

#! /bin/bash

#统计不同状态 tcp 连接（除了 LISTEN ）
all_status_tcp=$(netstat -nt | awk 'NR>2 {print $6}' | sort | uniq -c)

#打印各状态 tcp 连接以及连接数
all_tcp=$(netstat -na | awk '/^tcp/ {++S[$NF]};END {for(a in S) print a, S[a]}')


#统计有哪些 IP 地址连接到了本地 80 端口（ipv4）
connect_80_ip=$(netstat -ant| grep -v 'tcp6' | awk '/:80/{split($5,ip,":");++S[ip[1]]}END{for (a in S) print S[a],a}' |sort -n)


#输出前十个连接到了本地 80 端口的 IP 地址（ipv4）
top10_connect_80_ip=$(netstat -ant| grep -v 'tcp6' | awk '/:80/{split($5,ip,":");++S[ip[1]]}END{for (a in S) print S[a],a}' |sort -rn|head -n 10)


echo -e "\e[31m不同状态(除了LISTEN) tcp 连接及连接数为：\e[0m\n${all_status_tcp}"
echo -e "\e[31m各个状态 tcp 连接以及连接数为：\e[0m\n${all_tcp}"
echo -e "\e[31m连接到本地80端口的 IP 地址及连接数为：\e[0m\n${connect_80_ip}"
echo -e "\e[31m前十个连接到本地80端口的 IP 地址及连接数为：\e[0m\n${top10_connect_80_ip}"

### 系统性能

作者：咸鱼Linux运维
链接：https://zhuanlan.zhihu.com/p/612782776
来源：知乎
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

#!/bin/bash

#物理内存使用量
mem_used=$(free -m | grep Mem | awk '{print$3}')

#物理内存总量
mem_total=$(free -m | grep Mem | awk '{print$2}')

#cpu核数
cpu_num=$(lscpu  | grep 'CPU(s)' | awk 'NR==1 {print$2}')

#平均负载
load_average=$(uptime  | awk -F : '{print$5}')

#用户态的CPU使用率
cpu_us=$(top -d 1 -n 1 | grep Cpu | awk -F',' '{print $1}' | awk '{print $(NF-1)}')

#内核态的CPU使用率
cpu_sys=$(top -d 1 -n 1 | grep Cpu | awk -F',' '{print $2}' | awk '{print $(NF-1)}')

#等待I/O的CPU使用率
cpu_wa=$(top -d 1 -n 1 | grep Cpu | awk -F',' '{print $5}' | awk '{print $(NF-1)}')

#处理硬中断的CPU使用率
cpu_hi=$(top -d 1 -n 1 | grep Cpu | awk -F',' '{print $6}' | awk '{print $(NF-1)}')

#处理软中断的CPU使用率
cpu_si=$(top -d 1 -n 1 | grep Cpu | awk -F',' '{print $7}'| awk '{print $(NF-1)}')

echo -e "物理内存使用量(M)为：${mem_used}"
echo -e "物理内存总量(M)为：${mem_total}"
echo -e "cpu核数为：${cpu_num}"
echo -e "平均负载为：${load_average}"
echo -e "用户态的CPU使用率为：${cpu_us}"
echo -e "内核态的CPU使用率为：${cpu_sys}"
echo -e "等待I/O的CPU使用率为：${cpu_wa}"
echo -e "处理硬中断的CPU使用率为：${cpu_hi}"
echo -e "处理软中断的CPU使用率为：${cpu_si}"



nvm使用
下载指定版本node
 NVM_NODEJS_ORG_MIRROR=http://nodejs.org/dist nvm ls-remote

 NVM_NODEJS_ORG_MIRROR=http://nodejs.org/dist nvm install 18.20.3

处理github拉取依赖，443报错问题。  挂梯子，然后清理dns
 ipconfig /flushdns
