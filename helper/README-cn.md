[English](https://github.com/niubir/user-service/blob/main/helper/README-en.md) | 简体中文

# 什么是区域服务

区域服务是行政区域代码服务.

基于[高德地图](https://console.amap.com)实现.

# 快速开始

## 配置

[config.yml](https://github.com/niubir/user-service/blob/main/config/config.yml)
```
mode: TEST
debug: true
amapKey: <set you amap key>
autoFreshTime: 00:00:00
http:
  enable: true
  port: 10011
grpc:
  enable: true
  port: 10012
```

## Golang运行

```sh
git clone https://github.com/niubir/user-service && cd user-service
CONFIG_DIR='you config.yml path' go run main.go
```

## Docker运行

```sh
docker pull niubir/user-service:latest
docker run -p 10011:10011 -p 10012:10012 -v <you config.yml path>/config:/config -d niubir/user-service:latest
```

# 使用方式

## http

```
swgger http://localhost:10011/swagger/index.html
```

## grpc

```
proto https://github.com/niubir/user-service/blob/main/grpc/user.proto
```

## cli

```
git clone https://github.com/niubir/user-service-client
```

# 配置选项([config.yml](https://github.com/niubir/user-service/blob/main/config/config.yml))

| Option | Default | Description |
| - | - | - |
| mode | PRODUCTION | system mode(PRODUCTION,TEST,DEVELOPMENT) |
| logLevel | false | log level(DEBUG,INFO,WARN,ERROR) |
| amapKey | - | [Amap App](https://console.amap.com/dev/key/app) |
| autoFreshTime | - | at what time does the automatic fresh of user every day, for example:00:00:00 |
| http.enable | true | enable http |
| http.debug | false | enable http debug |
| http.port | 10011 | http port |
| grpc.enable | true | enable grpc |
| grpc.debug | false | enable grpc debug |
| grpc.port | 10012 | grpc port |

# 环境变量选项
| Option | Default | Description |
| - | - | - |
| CONFIG_DIR | /config | config dir |
| CONFIG_FILENAME | config.yml | config filename |
| LOG_DIR | /logs | log dir |
| AREA_DIR | /data | user dir |
