

## 一个服务程序的基本控制:

* 安装服务:

```shell
$ ./sservice install
```

* 启动服务:

```shell
$ systemctl start sservice
```

* 删除服务:

```shell
$ ./sservice remove
```

#### 注：删除服务之前需要先stop,否则会造成残留信息
