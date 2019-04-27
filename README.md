# postgres 数据库操作 API
作者 XinyaoTian

## 基于 Docker 或 Kubernetes 运行

该应用的新版本为容器化运行提供便利。使用 postgresManager/ 目录下的 Dockerfile 文件即可将该工程的所有源码封装进 Image 中。
而本应用的各项数据库配置信息通过在容器启动时，传入的环境变量进行读取。

目前可供配置的信息如下：

- POSTGRES_PROTOCOL 数据库使用的协议。
- POSTGRES_IPADDR 数据库服务所在 IP 地址。
- POSTGRES_PORT 同上，网络位置中所在的端口号。
- POSTGRES_USERNAME 数据库用户名。
- POSTGRES_PASSWORD 数据库密码。
- POSTGRES_DBNAME 数据库名。
- POSTGRES_SSLMODE SSL模式。

在 docker 容器首次启动时，这些配置信息将会输出到标准输出，请留意。

配置示例：

    # for db connection testing
    POSTGRES_SSLMODE=disable
    POSTGRES_PROTOCOL=postgres
    POSTGRES_IPADDR=184.157.54.15
    POSTGRES_PORT=5432
    POSTGRES_USERNAME=dbuser
    POSTGRES_PASSWORD=password
    POSTGRES_DBNAME=testdb
    
    export POSTGRES_SSLMODE
    export POSTGRES_PROTOCOL
    export POSTGRES_IPADDR
    export POSTGRES_PORT
    export POSTGRES_USERNAME
    export POSTGRES_PASSWORD
    export POSTGRES_DBNAME

## 用法

1. 更改 postgresManager/config 中的配置信息，以符合您的数据库接入要求
    
2. 编译并运行 main.go ；此 API 服务默认启动在 9090 端口

3. 对相应的 url 发起 http 请求即可使用

## 目前可用的 URL 规则

1. 返回所有数据库的 post 信息

        /all-queries
    
2. 基于用户名返回 post 信息

        /queries/:username
        
3. 插入新的 post(注意 recordId 必须为不重复的 hash 字段)
        
        /insert-post/:recordId
        # curl 命令示例: $ curl -X POST -d 'username=SAM&post=Hello&postTime=2012-02-05' http://127.0.0.1:9090/insert-post/00ac

## To Be Continue...


