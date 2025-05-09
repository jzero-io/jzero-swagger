【已废弃⚠️】

# jzero-swagger

## 下载

```shell
go install github.com/jzero-io/jzero-swagger@latest
```

## 使用

### jzero 中使用

```shell
jzero gen swagger
```

### go-zero 中使用

```shell
goctl api plugin -plugin jzero-swagger="swagger -filename swagger.json --schemes http,https" -api desc/api/version.api -dir desc/swagger
```

## 优化以下内容

* [request/response 支持基础类型](https://github.com/jzero-io/jzero-swagger/commit/575a3af18e90424e7506efe6d0315734212a0e40)
* [支持 jzero route2code 特性, 自动填充 swagger](https://github.com/jzero-io/jzero-swagger/commit/deee751c16426debca7858943738246921d2d74b)
* [支持 post 请求带有 query param 参数](https://github.com/jzero-io/jzero-swagger/commit/304e7e4d705ee6c0e1a94d3fe5305b31cbd4f424)
