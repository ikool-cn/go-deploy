# go-deploy
web集群一键上线工具，支持svn、git。增量更新、回滚，100+服务器节点秒级完成部署。

# Screenshot
![](https://github.com/ikool-cn/go-deploy/blob/master/Screenshot.png)

### 特性
- 支持svn和git
- 支持多项目、多节点、多环境、分布式集群环境批量更新和回滚操作
- golang的高性能并发加持，上百台节点秒级完成部署
- 无需svn或git账号密码、无需免密登录目标机 降低安全风险
- 增量更新、回滚，速度更快
- server和client采用tcp通讯+心跳保活 节点在线状态实时监控
- 支持befor_deploy、after_deploy 部署前和部署后的hook命令，清理缓存、执行重启等操作。 如：sudo service php-fpm reload
- 控制面板、简单易用、无需复杂配置。
- server端配置hook命令，避免web端配置hook带来的风险

### 部署流程
 - 可以自行编译或者直接下载bin下的二进制文件。
 
 - 所有的client节点标机器必须先使用svn部署好环境，假设web运行的系统账号为www，后续的所有操作均在www账号下进行

    ```
    su www
    svn co svn://1.2.3.4/app /data/wwwroot/app
    cd /data/wwwroot/app
    svn up
    #确保下次执行更新 不需要输入账号密码
    ```
    并配置好你的web运行环境，如nginx。

 - 请确保www账号拥有shell权限
    ```
    cat /etc/passwd|grep www 
    www:x:501:501::/home/www:/bin/bash
    ```
 - 每个项目需要在server节点拉取一份代码，并将其目录配置在config.json对应的项目的fetchlogpath，控制面板读取的commit log将在这个目录下读取，切忌请不要将任何web指向这个目录。

 - 配置config.json，添加应用和node节点

 - 使用www账号启动server
    ```
    su www
    chmod +x /pathto/server
    /pathto/server -c /pathto/config.json
    ```
 - 打开浏览器查看web管理界面 http://ip:port 是否可以正常访问

 - 所有的节点机器启动client，同样使用web账号www运行
    ```
   su www
   chmod +x /pathto/client
   /pathto/client -l :8081 #config.json里的node节点端口
    ```
    至此就部署完毕了。

### web管理访问安全问题
 - 第一种方案：使用nginx+auth_basic 来保护页面的访问。
    通过htpasswd命令生成用户名及对应密码数据库文件
    ```
    htpasswd -c /usr/local/openresty/nginx/conf/vhost/passwd.db yourusername
    ```
    配置nginx
    ```
    server {
             listen 80;
             server_name yourdomain;
             auth_basic "User Authentication";
             auth_basic_user_file /usr/local/openresty/nginx/conf/vhost/passwd.db;
             location / {
                 proxy_pass http://127.0.0.1:http_port;
             }
         }
    ```
 - 第二种方案：每次上线后关闭server。
