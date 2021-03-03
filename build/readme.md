
# go build

## mac
```shell script
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

## linux
```shell script
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

## windows
```shell script
        SET CGO_ENABLED=0
        SET GOOS=darwin
        SET GOARCH=amd64
        go build main.go
         
        SET CGO_ENABLED=0
        SET GOOS=linux
        SET GOARCH=amd64
        go build main.go
```

## 扩展
```shell script
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w --extldflags "-static -fpic"' main.go
```
-s -w 去掉调试信息，可以减小构建后文件体积，
--extldflags "-static -fpic" 完全静态编译[2]，这样编译生成的文件就可以任意放到指定平台下运行，而不需要运行环境配置。