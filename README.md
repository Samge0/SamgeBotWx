## 微信机器人简易助手
基于[openwechat](https://github.com/eatmoreapple/openwechat)+[openai](https://openai.com)的api+[docker](https://www.docker.com/)，实现基本的消息（好友、群组）接收与自定义回复、公众号消息监听。
方便在微信中使用openai体验[chatGPT](https://chat.openai.com/chat)~

【温馨提示】非官方产品有被封号风险，请使用小号作为Bot尝试。

### 基本步骤

- 1、注册openai并获取其token；
- 2、运行docker，首次复制控制台的链接访问扫描登录（后续自动推消息到手机登录）；

### 运行本项目代码进行消息处理
 - `-v xxx:/app/tmp/wxBotCache` =》这里的xxx填写自己的映射路径，存放日志，例如下面的`/home/samge/docker_data/samge_wx_bot`
 - `OPEN_AI_TOKEN=`填写openai的token值，
 - `GROUP_IDS=`填写群的白名，多个用,分隔（群名，群需要保存到通讯录）
 - `FRIEND_IDS=`填写好友的白名单，多个用,分隔（好友昵称）
 - `MINE_NICKNAME=`填写当前机器人bot的昵称，用于判断群中是否@自己

```shell
docker run -d \
-p 8888:8080 \
--name samge_wx_bot \
-v /home/samge/docker_data/samge_wx_bot:/app/tmp/wxBotCache \
--pull=always \
--restart always \
-e LANG=C.UTF-8 \
-e TZ="Asia/Shanghai" \
-e OPEN_AI_TOKEN= \
-e GROUP_IDS= \
-e FRIEND_IDS= \
-e MINE_NICKNAME= \
samge/samge_wx_bot:v1
```

### 如果需要调试

- 配置：`cmd/wxBot/botConfig/botConfig.go`
- 安装依赖：`go mod tidy`
- 运行：`go run cmd/wxBot/main.go`

### 技术交流
[Join Discord >>](https://discord.com/invite/eRuSqve8CE)

### 免责声明
该程序仅供技术交流，使用者所有行为与本项目作者无关
