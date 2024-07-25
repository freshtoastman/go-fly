### 郑重提示
禁止將本項目用于含病毒、木马、色情、赌博、诈骗、违禁用品、假冒产品、虚假信息、數字货币、金融等违法违规业務

當前項目仅供个人学习测试，禁止一切線上商用行為，禁止一切违法使用！！！




### 項目简介

Golang语言开源客服系统，主要使用了gin + jwt-go + websocket + go.uuid + gorm + cobra + VueJS + ElementUI + MySQL等技术


### 安装使用


* 先安装和運行mysql數據庫 ，版本>=5.5 ，創建數據庫
 
```
 create database gofly charset utf8mb4;
```
   
*  配置數據庫链接信息，在config目錄mysql.json中
```php
{
	"Server":"127.0.0.1",
	"Port":"3306",
	"Database":"gofly",
	"Username":"go-fly",
	"Password":"go-fly"
}
```
* 安装配置Golang運行环境，請参照下面的命令去执行
```php
wget https://studygolang.com/dl/golang/go1.20.2.linux-amd64.tar.gz
tar -C /usr/local -xvf go1.20.2.linux-amd64.tar.gz
mv go1.20.2.linux-amd64.tar.gz /tmp
echo "PATH=\$PATH:/usr/local/go/bin" >> /etc/profile
echo "PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
source /etc/profile
go version
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```
* 下载代碼

    在任意目錄 git clone https://github.com/taoshihan1991/go-fly.git
    
    進入go-fly 目錄
   
* 导入數據庫 go run go-fly.go install

* 源碼運行 go run go-fly.go server

* 源碼打包 go build -o kefu  會生成kefu可以执行文件

* 二進制文件運行

   linux:   ./kefu server [可選 -p 8082 -d]
   
   windows: kefu.exe server [可選 -p 8082 -d]
   
* 關閉程序

   killall kefu


程序正常運行后，监听端口8081，可以直接ip+端口8081訪問

也可以配置域名訪問，反向代理到8081端口，就能隐藏端口號
### 客服對接
聊天链接

http://127.0.0.1:8081/chatIndex?kefu_id=kefu2

弹窗使用

```
    (function(a, b, c, d) {
        let h = b.getElementsByTagName('head')[0];let s = b.createElement('script');
        s.type = 'text/javascript';s.src = c+"/static/js/kefu-front.js";s.onload = s.onreadystatechange = function () {
            if (!this.readyState || this.readyState === "loaded" || this.readyState === "complete") d(c);
        };h.appendChild(s);
    })(window, document,"http://127.0.0.1:8081",function(u){
        KEFU.init({
            KEFU_URL:u,
            KEFU_KEFU_ID: "kefu2",
        })
    });

```
### 版權聲明

當前項目是完整功能代碼 , 但是仍然仅支持个人演示测试 , 不包含線上使用 ，禁止一切商用行為。
使用本软件時,請遵守当地法律法规,任何违法用途一切后果請自行承担.