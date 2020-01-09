# Golang基础学习-redis使用
## 1.开发环境

```
OS: MAC
docker: 19.03.5
Golang: go version go1.13.4 darwin/amd64
redis: redis5.0
```

## 2.安装 redis
> docker和golang 已经提前安装；

* docker pull redis:latest
* docker run --name redisserver -d -p 6379:6379  -v /Users/Keil/data:/data redis redis-server --appendonly yes

>  /Users/Keil/data：存储数据的地方，实际可根据自己开发环境测试

## 3. 引入redis的package

* go get -u  github.com/go-redis/redis

 
 ## 4. redis常用方法
 
 ```.env
package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var (
	RedisIp    = "127.0.0.1"
	RedisPort  = "6379"
	expireTime = 600
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     RedisIp + ":" + RedisPort,
		Password: "",
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		panic("redis连接错误")
		return
	}

	// 判断 key是否存在

	a, err := rdb.Exists("ming").Result()
	if err != nil {
		fmt.Println("判断key存在失败")
		return
	}
	if a == 1 {
		fmt.Println("key存在")
	}

	//存储key
	err = rdb.Set("ming", "ming", time.Duration(expireTime)*time.Second).Err()
	if err != nil {
		fmt.Println("设置key失败")
		return
	}

	//获取key
	value, err := rdb.Get("ming").Result()
	if err != nil {
		fmt.Println("设置key失败")
		return
	}
	fmt.Println(value)

	//设置过期时间
	err = rdb.Expire("ming", time.Duration(300)*time.Second).Err()
	if err != nil {
		fmt.Println("设置过期时间")
		return
	}

	/*
		设置hash
	*/
	status, err := rdb.HSet("ming1", "id", "12313").Result()
	if err != nil {
		fmt.Println("err:redis服务异常")
		return
	}
	if true == status {
		fmt.Println("值已存在")
	} else {
		fmt.Println("设置成功")
	}

	// 获取key
	key, err := rdb.HGet("ming1", "id").Result()
	if err != nil {
		fmt.Println("获取不到值")
		return
	}
	fmt.Println(key)

	// 判断hash中,值是否存在
	status, _ = rdb.HExists("ming1", "id").Result()
	if true == status {
		fmt.Println("值已存在")
	} else {
		fmt.Println("设置成功")
	}

	// 删除hash中的值

	statusDel, err := rdb.HDel("ming1", "id").Result()
	if err != nil {
		fmt.Println("err:redis服务异常")
		return
	}
	if 1 == statusDel {
		fmt.Println("删除hash值：id成功")
	}

	//删除key
	statusDel, err = rdb.Del("ming").Result()
	if 1 == statusDel {
		fmt.Println("删除值成功")
	}
}

```

## 5 参考
 * [go-redis](https://godoc.org/github.com/go-redis/redis)
 * [dockerhub-redis](https://hub.docker.com/_/redis?tab=description)