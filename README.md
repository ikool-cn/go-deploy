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

### 编译安装
 - 可以自行编译或者直接下载bin目录下的二进制文件(编译环境centos 7.9)。
 
### Server端配置
```
|---server
|---server.json
```
 
1. 上传server二进制文件到服务器
     
2. 创建或修改server.json，添加项目以及相应的节点。
3. 获取commit log
   如果项目类型为svn则fetchlogpath设置为空，切换到应用程序执行账户下，这里假设为www账户，手动获取一次日志，目的是让工具拉日志的时候可以免密。
```bash
su www
svn log --limit 10 svn://x.x.x.x/project/
```
   如果项目类型为git则fetchlogpath需要配置一个目录，并将项目拉取一份在这个目录下。请不要将任何web站点指向这个目录，这个目录的作用仅仅是为了获取git提交日志用。

4. 使用www账号启动server
```bash
su www
chmod +x /pathto/server
/pathto/server -c /pathto/server.json
```
     
5. 打开浏览器查看web管理界面 http://ip:port 是否可以正常访问 (port为server.json里配置的listen_http 监听端口)

### Client节点配置
```
 |---client
```

1. 上传client二进制文件 到所有的节点服务器，Node节点只用部署一个单文件(以下操作每个Node节点服务器都需要执行同样的操作)

2. 首选拉取一份项目代码，假设www为你的web运行账户(使用ps aux|grep php-fpm查看)
   查看www账号是否有shell权限
```bash
cat /etc/passwd|grep www 
www:x:501:501::/home/www:/bin/bash
```
   如果没有shell权限则需要执行下面命令
```bash
usermod -s /bin/bash www #www是用户名
```
   拉取项目代码，确保下次执行可以免密更新
```bash
su www
svn co svn://1.2.3.4/app /data/wwwroot/app
cd /data/wwwroot/app
svn up
```
3. 启动client
```bash
su www
chmod +x /pathto/client
/pathto/client -l :8081 #8081为client监听端口
```

### web管理访问安全问题
   部署完后web管理界面直接暴露给所有人的，可以加一层nginx反向代理，设置一个auth认证
```bash
htpasswd -c /usr/local/openresty/nginx/conf/vhost/passwd.db yourusername
```
   配置nginx
```bash
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
