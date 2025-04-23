# ssh-manager
使用 golang 的 cobra 实现一个 轻松管理ssh别名配置的命令行工具



## 使用方法
### 构建
```shell
# 本地构建
make build

# Docker 构建
make docker-build
# Docker 运行
make docker-run
```



### 运行

```shell
./ssh-manager --help
./ssh-manager add --alias myalias --host 192.168.1.100 --user ubuntu --port 22

# 查看配置
cat ~/.ssh/config

Host myalias
    HostName 192.168.1.100
    User ubuntu
    Port 22


./ssh-manager list

./ssh-manager remove --alias myalias
```



