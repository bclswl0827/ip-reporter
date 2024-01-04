# ip-reporter

在 IP 随机变化的 WiFi 校园网中，定时向固定 IP 的服务器上报自身 IP 位址

相比于使用 DDNS，此方案不需注册域名，不需配置 DNS，客户端设备也不需接入网际互联网

# 服务端

服务端使用 Go + Gin 编写，可以部署在多种平台上，包括树莓派、路由器、云服务器等

## 部署

确保服务器上已经安装了 Go 环境，对于 Linux 系统，以下命令可以在 Bash 中运行；对于 Windows 系统，可以在 Git Bash 中运行

**若服务器与编译环境不在同一平台或使用不同架构，可以使用交叉编译，Go 语言如何交叉编译请自行 Google 搜索**

### 取得源码

```bash
$ git clone https://github.com/bclswl0827/ip-reporter --depth=1
$ cd ip-reporter/server
```

### 编译

```bash
$ go build -ldflags "-s -w" -o ./ip-reporter-server
```

### 运行

编辑 `config.json` 文件，下表列出了各个字段的含义

| 字段     | 类型     | 描述              |
| :------- | :------- | :---------------- |
| `listen` | `string` | 监听地址和端口    |
| `cors`   | `bool`   | 是否允许跨域请求  |
| `life`   | `int`    | IP 有效期，单位秒 |

配置完整示例如下

```json
{
    "life": 60,
    "cors": false,
    "listen": "0.0.0.0:8000"
}
```

配置文件编辑完成后，运行服务端，假设配置文件路径位于 `/path/to/config.json`

```bash
$ ./ip-reporter-server -config /path/to/config.json
```

服务器端运行后，可以访问 `http://ip:port` 查看当前已经上报的 IP 地址

## 注意事项

 - 请确保服务器上的防火墙已经放行了服务端监听的端口
 - 服务器端 IP 地址尽量不要变化，否则客户端将无法上报 IP 地址

# 客户端

## 部署

确保客户端上已经安装了 Go 环境，对于 Linux 系统，以下命令可以在 Bash 中运行；对于 Windows 系统，可以在 Git Bash 中运行

**若客户端与编译环境不在同一平台或使用不同架构，可以使用交叉编译，Go 语言如何交叉编译请自行 Google 搜索**

### 取得源码

```bash
$ git clone https://github.com/bclswl0827/ip-reporter --depth=1
$ cd ip-reporter/client
```

### 编译

```bash
$ go build -ldflags "-s -w" -o ./ip-reporter-client
```

### 运行

编辑 `config.json` 文件，下表列出了各个字段的含义

| 字段         | 类型     | 描述                                                           |
| :----------- | :------- | :------------------------------------------------------------- |
| `timeout`    | `int`    | 上报超时时间，单位秒，超过此时间未上报成功则重试               |
| `interval`   | `int`    | 上报间隔时间，单位秒，须小于服务器配置的 IP 有效期             |
| `report_uri` | `string` | 服务端地址，格式为 `http://ip:port/collect`                    |
| `device_tag` | `string` | 设备标识，用于区分不同的设备，例如 `raspberrypi@lab0308`       |
| `ip_prefix`  | `string` | IP 地址前缀，例如校园网 IP 以 `10.10` 开头，需根据实际情况配置 |

配置完整示例如下

```json
{
    "timeout": 5,
    "interval": 5,
    "ip_prefix": "10.160",
    "device_tag": "test-device",
    "report_uri": "http://1.2.3.4:8000/collect"
}
```

配置文件编辑完成后，运行客户端，假设配置文件路径位于 `/path/to/config.json`

```bash
$ ./ip-reporter-client -config /path/to/config.json
```
