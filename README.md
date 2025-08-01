# match

后端go代码

# 环境依赖

> go 1.20+

> MySQL 5.7

> MySQL 创建一个数据库，然后导入 `*.sql` 文件

# 下载-打包

```shell
# 打包
$ go build

# 运行，**这里这个命令可以用 supervisor 这个进程守护工具守护一下**
$ nohup ./xx -c config/config-example.yml &

```

# 配置文件

> `config/config-example.yml` 这个配置文件可以随便放在某个目录下

> 运行时加上 -c 命令后面跟文件路径


* * *

- 配置文件介绍

```yaml
app:
  port: 10009 #程序运行端口

mysql:
  host: 127.0.0.1 #MySQL连接地址
  port: 3306 #MySQL端口
  db: xxx #数据库名
  user_name: root #用户名
  password: root #密码
  max_idle_conn: 10
  max_open_conn: 256
  conn_max_life_time: 600
```

#Nginx的配置文件

```shell
server {
    listen        80;
    server_name  vaiwan.com;
    root   "你的前端静态目录";
    location / {
    }
    
    # go程序的代理
    location /api/ {
        proxy_pass http://127.0.0.1:10001/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}

```
