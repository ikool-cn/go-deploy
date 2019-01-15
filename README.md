# go-deploy
Deploy code to cluster servers based on svn, onekey update onekey rollback

# Screenshot
![](https://github.com/ikool-cn/go-deploy/blob/master/Screenshot.png)

### 简介
- 支持多项目、多站点、多环境、多机器的批量更新和回滚操作
- 相比瓦力无需担心php执行权、免密登陆目标机器造成的安全威胁
- 无需svn账号密码
- 增量更新速度更快
- server和client采用tcp通讯，client支持断线自动重连
- 支持befor_deploy、after_deploy 更新前、更新后的执行命令 如：sudo service php-fpm reload

### 部署流程

 - 所有的目标机器使用svn预先部署好环境，假设web运行的账号为www

    ```
    su www
    svn co svn://1.2.3.4/app /data/wwwroot/app
    svn up 请确保这一步可以免密登陆```

 - 请确保www拥有shell权限
    ```
    cat /etc/passwd|grep www 
    www:x:501:501::/home/www:/bin/bash
    ```

 - 配置config.json

 - 使用www账号启动server，把config.json和server在相同的目录。
    ```
    su www
    ./pathto/server
    ```
 - 打开浏览器请确保http://x.x.x.:http_port可以正常访问

 - client端部署，同样使用web账号www运行
      ```
        su www
        ./pathto/client -s serverIP:serverTcpPort
      ```
ok! 至此就部署完毕了，打开浏览器测试一下把。

### web访问安全问题
 1. 可以使用nginx+auth_basic 来保护页面的访问。
 2. 每次用完关闭server，下次开启server时client会自动重连。