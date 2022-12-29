# Catty 实现使用GPT开放接口增强各类手机厂商自带的 Assistant

Catty build with GO.

### 功能
- [x] 接入苹果🍎的WorkFlow -> [点我接入WorkFlow](https://www.icloud.com/shortcuts/121f16f23dd04b3a8a02ce4b5245cbc7)
- [ ] 接入华为小艺
- [ ] 接入HomeAssistant
- [ ] Web

### 目录结构：
```bash
.
├── README.md
├── config  # 配置文件
│   ├── config.go
│   ├── config.yaml
│   └── config_test.go
├── controllers # 控制器
│   └── index.go
├── dao # 数据库相关
│   └── initDB.go
├── go.mod
├── go.sum
├── main.go # 主入口
├── models  # 接口模型与数据库模型
│   └── prompt.go
├── openai  # openai 请求接口
│   └── chat.go
└── routers # 路由
    └── router.go
```

### 使用方式
GET访问格式：`YOUR_IP/?prompt=What's Catty?`

POST访问格式：打算要不要做。

Docker部署：开发中～～～

### Q/A
Q：为什么要连接数据库？

A：收集问的问题，自己训练模型？，如果不要数据库，将`dao.Initdatabase()`注释掉

Q：后面会有什么更新？

A：这个项目仅仅是学习练手，有好玩的玩法或新功能会不定期更新。

Q：能否pr或fork？

A：欢迎pr，fork。这是个后端版本，希望能有人配合着写个扁平化样式的前端