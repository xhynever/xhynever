在分布式的业务中 , 如果有的共享资源需要安全的被访问和处理 , 那就需要分布式锁

分布式锁的几个原则;

1.「锁的互斥性」：在分布式集群应用中，共享资源的锁在同一时间只能被一个对象获取。

2. 「可重入」：为了避免死锁，这把锁是可以重入的，并且可以设置超时。

3. 「高效的加锁和解锁」：能够高效的加锁和解锁，获取锁和释放锁的性能也好。

4. 「阻塞、公平」：可以根据业务的需要，考虑是使用阻塞、还是非阻塞，公平还是非公平的锁。

redis实现分布式锁主要靠setnx命令

1. 当key存在时失败 , 保证互斥性

2.设置了超时 , 避免死锁

3.利用mutex保证当前程序不存在并发冲突问题
package redis

import (
    "context"
    "github.com/go-redis/redis/v8"
    "github.com/taoshihan1991/miaosha/setting"
    "log"
    "sync"
    "time"
)

var rdb *redis.Client
var ctx = context.Background()
var mutex sync.Mutex

func NewRedis() {
    rdb = redis.NewClient(&redis.Options{
        Addr:     setting.Redis.Ip + ":" + setting.Redis.Port,
        Password: "", // no password set
        DB:       0,  // use default DB
    })
}
func Lock(key string) bool {
    mutex.Lock()
    defer mutex.Unlock()
    bool, err := rdb.SetNX(ctx, key, 1, 10*time.Second).Result()
    if err != nil {
        log.Println(err.Error())
    }
    return bool
}
func UnLock(key string) int64 {
    nums, err := rdb.Del(ctx, key).Result()
    if err != nil {
        log.Println(err.Error())
        return 0
    }
    return nums
}