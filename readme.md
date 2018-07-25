

## 一个服务程序的基本控制:

* 安装服务:

```shell
$ ./gdaemon -a install
```

* 启动服务:

```shell
$ systemctl start gdaemon
```

* 删除服务:

```shell
$ ./gdaemon -a remove
```

* 重启错误的服务:
```shell
$ systemctl reset-failed
$ systemctl daemon-reload 
```
###### 注：删除服务之前需要先stop,否则会造成残留信息

### build.sh 封装完成了常见的打包操作:

* 默认打包:

```shell
$ sh build.sh
```

* 打出一个mini包

```shell
$ sh build.sh mini
```

* 打出一个rpm包

```shell
$ sh build.sh rpm
```
###### 注：依赖于gdaemon.spec ，可以对rpm包的具体细节做修改

