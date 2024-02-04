## Docker：微信机器人

### docker版本记录：
- 20240204：samge/samge_wx_bot:v2
- 20230210：samge/samge_wx_bot/wxBotCache:v1


### 构建应用docker镜像
```shell
docker build -t samge/samge_wx_bot:v2 . -f docker/Dockerfile
```

### 上传应用docker镜像
```shell
docker push samge/samge_wx_bot:v2
```

### run docker:

- 方式1：配置文件方式：
```shell
docker run -d \
-p 8888:8080 \
--name samge_wx_bot \
-v xxx/docker_data/samge_wx_bot/wxBotCache:/app/tmp/wxBotCache \
-v xxx/docker_data/samge_wx_bot/config.json:/app/config.json \
--restart always \
samge/samge_wx_bot:v2
```

- 方式2：使用环境变量方式运行：
```shell
docker run -d \
-p 8888:8080 \
--name samge_wx_bot \
-v xxx/docker_data/samge_wx_bot/wxBotCache:/app/tmp/wxBotCache \
-v xxx/docker_data/samge_wx_bot/config.json:/app/config.json \
--restart always \
-e sg.samge_wx_bot.base_url= \
-e sg.samge_wx_bot.api_key= \
-e sg.samge_wx_bot.model= \
-e sg.samge_wx_bot.group_ids= \
-e sg.samge_wx_bot.friend_ids= \
-e sg.samge_wx_bot.manager_ids= \
-e sg.samge_wx_bot.mine_nickname= \
samge/samge_wx_bot:v2
```


# 配置文件说明

````
{
  "base_url": "https://api.openai.com/v1",
  "bot_desc": "",
  
  "api_key": "sk-xxxxxxxxxxxxxxxxxxxxxxxxx",
  "model": "gpt-3.5-turbo-0301",
  "max_tokens": 2048,
  "temperature": 0.9,
  "top_p": 1,
  "frequency_penalty": 0.0,
  "presence_penalty": 0.6,
  
  "group_ids": "",
  "friend_ids": "",
  "manager_ids": "",
  "mine_nickname": "",
  "enable_reply": true
}

base_url：openai的请求地址，需要携带v1版本号，默认是：https://api.openai.com/v1 ，可配置转发地址/本地部署的模型api地址
bot_desc：提示内容prompt，功能等同给与AI一个身份设定（功能风格），默认为空，可自定义，例如：以下是与AI助手的对话。助手乐于助人，富有创造力，聪明且非常友好。

api_key：openai/本地模型/其他模型定义的api_key
model: GPT选用模型，默认gpt-3.5-turbo-0301，具体选项参考官网模型列表
max_tokens: 上下文长度。
temperature: GPT热度，0到1，默认0.9。数字越大创造力越强，但更偏离训练事实，越低越接近训练事实
top_p: 使用温度采样的替代方法称为核心采样，其中模型考虑具有top_p概率质量的令牌的结果。因此，0.1 意味着只考虑包含前 10% 概率质量的代币。
frequency_penalty: -2.0到2.0之间的数字。正值根据它们在文本中的现有频率惩罚新标记，降低模型逐字重复同一行的可能性。
presence_penalty: 数字介于-2.0和2.0之间。正值根据新标记到目前为止是否出现在文本中来惩罚它们，从而增加模型谈论新主题的可能性。

group_ids：需要接收消息的群名称，多个用英文逗号隔开，置空则接收所有群消息，例如：Samge测试群,其他逗比群
friend_ids：需要接收消息好友微信昵称(不是微信号)，多个用英文逗号隔开，置空则接收所有好友消息，例如：Samge001,SamgeBot,SamgeApp
manager_ids：管理员微信昵称(不是微信号)，目前有两个指令：开启回复 | 关闭回复。多个用英文逗号隔开，例如：Samge001,SamgeApp
mine_nickname：自己的昵称(不是微信号)，例如：STE
enable_reply：是否允许回复消息
````