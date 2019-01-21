# go-deploy
基于SVN的web集群一键部署工具，一键更新、回滚所有集群节点。

# Screenshot
![](https://github.com/ikool-cn/go-deploy/blob/master/Screenshot.png)

### 特性
- 支持多项目、多节点、多环境、分布式环境批量更新和回滚操作
- 基于goroute的并行操作，上百台节点秒级完成
- 相比瓦力无需担心免密登陆目标机器造成的安全威胁
- 无需svn账号密码
- 增量更新速度更快，不需要全量检出
- server和client采用tcp通讯+心跳保活 节点在线状态实时监控
- 支持befor_deploy、after_deploy 前置后置行命令 如：sudo service php-fpm reload

### 部署流程

 - 所有的节点标机器必须先使用svn部署好环境，假如web运行的账号为www，以下所有操作均在www账号下进行

    ```
    su www
    svn co svn://1.2.3.4/app /data/wwwroot/app
    cd /data/wwwroot/app
    svn up
    #请确保这一步可以免密更新
    ```
    并配置好你的web运行环境，如nginx。

 - 请确保www账号拥有shell权限
    ```
    cat /etc/passwd|grep www 
    www:x:501:501::/home/www:/bin/bash
    ```

 - 配置config.json，添加应用和client节点

 - 使用www账号启动server
    ```
    su www
    /pathto/server -c /pathto/config.json
    ```
 - 打开浏览器查看web管理界面 http://serverip:http_port 是否可以正常访问

 - 所有的节点机器启动client，同样使用web账号www运行
    ```
   su www
   ./pathto/client -l :8081 #自定义端口
    ```
ok! 至此就部署完毕了。

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
