### 环境

- go > 1.15
- mysql
- redis

### 配置

- 根据 `config/config.default.yaml` 配置 `config/config.yaml` 

### 初始化

- 初始化数据库表 `go run main.go -m true`  

### 启动

- `go run main.go` 或指定配置文件 `go run main.go -c config/config.yaml`
