## 博客后台接口 (www.littlebug.vip)

<a href="https://travis-ci.org/wanchaochao/blog_api"><img src="https://travis-ci.org/wanchaochao/blog_api.svg" alt="Build Status"></a>
[![stable](https://img.shields.io/badge/gin--gonic%2Fgin-1.3.0-green.svg)](https://github.com/gin-gonic/gin)
[![stable](https://img.shields.io/badge/ilibs-open-brightgreen.svg)](https://github.com/ilibs)
[![stable](https://img.shields.io/badge/gosql-v1.1.1-brightgreen.svg)](https://github.com/ilibs/gosql)

<p align="center">
	<a href="https:www.littlebug.vip">
		<img src="http://littlebug.oss-cn-beijing.aliyuncs.com/www.littlebug.vip/favicon.ico" width="75">
	</a>
</p>


<p align="center">
	在线接口: api.littlebug.vip  (登录接口 : api.littlebug.vip/login)
</p>

## install

```
// 为了方便部署，go mod vendor已经将vendor目录加入了项目文件

git clone git@github.com:Wanchaochao/blog_api.git

// 配置文件
cd blog_api
go get
go mod vendor
cp -rf ./config/demo ./config/pro
cp -rf ./config/demo ./config/local

// captcha配置为腾讯免费滑块验证码配置
[腾讯滑块验证码](https://007.qq.com/captcha)

go

// 本机为mac，以mac为例，添加app_env

vim ~/.bash_profile

export APP_ENV="local"

make

./app http -addr=:8083

```


<p align="center">
	<a href="https:www.littlebug.vip">
		<img src="http://littlebug.oss-cn-beijing.aliyuncs.com/test/6E86E115-5DBF-4DB9-A095-EB0DD0F693A7.png" width="500">
	</a>
</p>


<p align="center">
	看到这里本地的golang http服务已经成功启动了
</p>

## 服务器部署

```
make start

// 我的服务器为阿里云香港服务器,centos7
cd /etc/nginx/conf.d

vim api.littlebug.vip.conf

// 加入如下代码

server {
    server_name  api.littlebug.vip;

    listen 443 ssl; # managed by Certbot
    ssl_certificate /etc/letsencrypt/live/api.littlebug.vip/fullchain.pem; # managed by Certbot
    ssl_certificate_key /etc/letsencrypt/live/api.littlebug.vip/privkey.pem; # managed by Certbot
    # include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot
    ssl_session_timeout 5m;
    ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:ECDHE:ECDH:AES:HIGH:!NULL:!aNULL:!MD5:!ADH:!RC4;
    ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
    ssl_prefer_server_ciphers on;

    charset utf-8;
    #如果是css,js|fonts|png|svg|html|txt 资源文件 nginx 直接处理，不提交到后台让go处理。
    # nginx 会在root 对应的目录下 去找这些资源文件
    location / {
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header Host $http_host;
        proxy_redirect off;
        proxy_pass http://localhost:8083;
        if ($request_method = 'OPTIONS') {
            add_header 'Access-Control-Allow-Origin' '*';
            add_header 'Access-Control-Allow-Methods' 'GET, POST, OPTIONS';
            add_header 'Access-Control-Allow-Headers' 'DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,multipart/form-data, application/json,Access-token';
            return 204;
        }

        add_header 'Access-Control-Allow-Origin' '*';
        add_header 'Access-Control-Allow-Methods' 'GET, POST, OPTIONS';
        add_header 'Access-Control-Allow-Headers' 'DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,multipart/form-data, application/json,Access-token';
    }

    access_log  /var/log/nginx/api.littlebug.log  access;
}
```


## 接口测试



```
// 修改本地配置文件
// config/local/common.json

// 将这里改为true,不然会被中间件拦截下来
"http_test" : true

// 代码位置 router/middleware/token.go

if config.App.Common.HttpTest {
    c.Next()
    return nil
}

```
**这里千万要注意线上一定要设置APP_ENV环境变量,否则使用的将是local里的配置,如果你的http_test配置为true,则接口没有任何安全性可言了,pro里的http_test一定要关闭**

### 运行接口测试

**别忘了开启http服务**
<img src="http://littlebug.oss-cn-beijing.aliyuncs.com/www.littlebug.vip/run_http_test01.png">
<img src="http://littlebug.oss-cn-beijing.aliyuncs.com/www.littlebug.vip/run_http_test02.png">

### 如何使用 http client
<img src="http://littlebug.oss-cn-beijing.aliyuncs.com/www.littlebug.vip/http_client_01.png">
<img src="http://littlebug.oss-cn-beijing.aliyuncs.com/www.littlebug.vip/http_client_02.png">
<img src="http://littlebug.oss-cn-beijing.aliyuncs.com/www.littlebug.vip/http_client_03.png">


<p align="center">
    <b>你一定注意到 "managed by Certbot",这里是使用Certbot配置的https证书,非常方便快捷</b>
    <br/>
	<a href="https://certbot.eff.org/">
		<img src="https://certbot.eff.org/images/certbot-logo-1A.svg" width="150">
	</a>
</p>

<p align="center">
	中文教程链接:
	    <a href="https://laravel-china.org/articles/5883/give-your-website-a-https-certificate-per-second">
	        让你的网站秒配https
	    </a>
</p>

<p align="center">
    事实证明我还是太年轻了,从阿里云的证书到CertBot再到<a href="https://boris1993.github.io/tools/caddy/set-up-caddy-server.html">Caddy</a>...结果大陆访问香港的https网站会被移动联通拦截...😭
    最后回归原始全部回到http
</p>




