## 微信机器人简易助手
基于[openwechat](https://github.com/eatmoreapple/openwechat)+[openai](https://openai.com)的api+[docker](https://www.docker.com/)，实现基本的消息（好友、群组）接收与自定义回复、公众号消息监听。
方便在微信中使用openai[chatGPT](https://chat.openai.com/chat) 或者 使用自己部署的开源大模型（如`Qwen`、`GLM`、`llama2`、`baichuan2`……）<br>

因为本项目是基于[openwechat](https://github.com/eatmoreapple/openwechat) ，微信Bot详细文档请到这里查看：[openwechat](https://github.com/eatmoreapple/openwechat) <br>

【温馨提示】非官方产品有被封号风险，请使用小号作为Bot尝试。<br>

### 基本步骤

- 1、【可选】注册openai并获取其token（或使用自己部署的开源模型）；
- 2、运行docker，首次复制控制台的链接访问扫描登录（后续自动推消息到手机登录）；

### 使用docker运行
[点击查看docker运行时的参数说明>>](docker/README.md)

```shell
mv config.dev.json xxx/docker_data/samge_wx_bot/config.json
```

```shell
docker run -d \
-p 8888:8080 \
--name samge_wx_bot \
-v xxx/docker_data/samge_wx_bot/wxBotCache:/app/tmp/wxBotCache \
-v xxx/docker_data/samge_wx_bot/config.json:/app/config.json \
--restart always \
samge/samge_wx_bot:v2
```

### 如果需要调试

- 配置：`cmd/wxBot/botConfig/botConfig.go`
- 安装依赖：`go mod tidy`
- 运行：`go run cmd/wxBot/main.go`


### 技术交流
- [Join Discord >>](https://discord.com/invite/eRuSqve8CE)
- WeChat：`SamgeApp`


### 免责声明
该程序仅供技术交流，使用者所有行为与本项目作者无关
