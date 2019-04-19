# postgres 数据库操作 API
作者 XinyaoTian

## 用法

1. 更改 postgresManager/routeHandler 中的 3 个 Init 函数的数据库信息，以符合您的数据库
    
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


