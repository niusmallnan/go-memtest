# go-memtest

在Go 1.12之前，运行时会在未使用的内存上发送MADV_DONTNEED信号，并且操作系统会立即回收未使用的内存页。 从Go 1.12开始，该信号已更改为MADV_FREE，这告诉操作系统它可以根据需要回收一些未使用的内存页面，这意味着除非系统承受着来自不同进程的内存压力，否则它不会真正去回收这部分内存。

值得注意的是，如果在Kubernetes中运行Go程序，它占用的内存可能会随着时间推移越来越大，而且如果你的系统负载不高，这个内存并不会被回收。这很可能对监控和调度产生一些困扰。

本程序辅助测试Golang GC策略，测试使用golang 1.14 。

[![](http://img.youtube.com/vi/cWj_fyWufdk/0.jpg)](http://www.youtube.com/watch?v=cWj_fyWufdk "Golang GC Test")

### 默认MADV_FREE策略

启动容器

```
docker run -idt niusmallnan/go-memtest
```

使用 `docker stats` 监控内存情况

进入容器，申请内存：

```
docker exec -it <id> bash
# allocate_memory.sh
```

内存占用会出现激增，但达到一定水准后，保持稳定。

手动触发回收：

```
docker exec -it <id> bash
# gops gc 1
```

内存没有被系统回收。

### 切换MADV_DONTNEED策略

启动容器

```
docker run -e GODEBUG=madvdontneed=1 -idt niusmallnan/go-memtest
```

使用 `docker stats` 监控内存情况

进入容器，申请内存：

```
docker exec -it <id> bash
# allocate_memory.sh
```

内存占用会出现激增，但达到一定水准后，保持稳定。

手动触发回收：

```
docker exec -it <id> bash
# gops gc 1
```

内存会被系统回收，同时，即使不去手动触发，GOGC也会在一段时间后自动回收。


### 内置API

查看memstats

```
curl --unix-socket /tmp/log.sock http://127.0.0.1/stats

docker logs -f <id>
```

Call GC

```
curl --unix-socket /tmp/log.sock http://127.0.0.1/freegc
```

Call FreeOSMemory

```
curl --unix-socket /tmp/log.sock http://127.0.0.1/freeosm
```

分配内存

```
curl --unix-socket /tmp/log.sock http://127.0.0.1/alloc
curl --unix-socket /tmp/log.sock http://127.0.0.1/alloc?size=10000000
```
