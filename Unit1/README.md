<h2>****Redis第一章****</h2>
初识Redis
- 常用的几种类型
**字符串、哈希、列表、集合、有序集合、位图、超日志、地理空间索引和流**
- 切换数据库
  - **select index** (Redis单机服务器有16个数据库(0~15) )
  - **flushdb**(删除所有key)
  - **getrange key 3 7**(类似切片获取部分内容)
  - **getset name lisi**(获取旧的值并设置为新的值)
  - **psetex key expire value**(为key设置有效时间)
  - **set user object**(对象缓存)
- 对于接口限流的操作，限制请求次数200次，超过200次就等待浏览量统计、阅读量统计
  - **incr key**键自增1(计数器的方式，但是到达一定程度删除这个键)
  - **incrby key count**(计数器自己指定步长,适用场景是分布式自增)
- 分布式锁
  - 多线程通过锁来限制资源抢占控制
    - 秒杀伪代码：
    <code>setnx seckill true #抢到锁设置成功返回</code><br>
    <code>setnx seckkill true #返回0没有抢到锁</code> <br>
    <code>del seckkill #执行完任务删除</code> <br>
    <code>expire seckkill 20000</code> <br>
    <code>ttl seckkill #查看key还有多久过期</code> <br>
    <code>expire seckkill #查看key过期时间</code> <br>
  
