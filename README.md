# mai_pb_updater

将 maimai 乐曲数据从 Diving Fish API 同步到 PocketBase 数据库的工具。

## 功能

- 从 Diving Fish API 获取所有 maimai 乐曲数据
- 将数据批量创建/更新到 PocketBase
- 支持 ETag 缓存，减少不必要的网络请求
- 增量更新：自动判断创建和更新操作

## 配置

配置文件：`config.yml`

```yaml
divingfish:
    base_url: https://www.diving-fish.com/api/maimaidxprober
    etag: ""  # 自动更新

pocketbase:
    base_url: https://pb.example.com
    collection_name: maimai_music_data
    identity: example@example.com
    password: example
    max_batch_size: 50
```

配置优先级：环境变量 `CONFIG_PATH` > 命令行参数 `-config` > 默认值 `./config.yml`

## 使用

```bash
go build -o mai_pb_updater .
./mai_pb_updater -config ./config.yml
```

## 项目结构

```
.
├── main.go              # 主程序入口
├── config/              # 配置管理
│   ├── config.go        # 配置结构体
│   └── viper.go         # Viper 初始化
├── divingfish/          # Diving Fish API 客户端
│   ├── client.go        # API 请求
│   └── model.go         # 数据模型
└── pocketbase/          # PocketBase 客户端
    ├── client.go        # 数据库操作
    └── model.go         # 数据模型
```

## 数据同步流程

1. 从 PocketBase 登录获取认证 Token
2. 调用 Diving Fish API 获取最新乐曲数据
3. 从 PocketBase 获取已存在的乐曲数据
4. 对比两份数据，分离出需要创建和更新的记录
5. 批量创建新记录，批量更新已有记录
