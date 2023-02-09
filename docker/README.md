## Docker：微信机器人

### 构建应用docker镜像
```shell
docker build -t samge/samge_wx_bot:v1 . -f docker/wxBot/Dockerfile
```

### 上传应用docker镜像
```shell
docker push samge/samge_wx_bot:v1
```


### run docker:
-v xxx:/app/tmp/qqBotCache =》这里的xxx填写自己的映射路径，存放日志，例如下面的`/home/samge/docker_data/samge_wx_bot`
OPEN_AI_TOKEN=填写openai的token值，
GROUP_IDS=填写群的白名，多个用,分隔（群名，群需要保存到通讯录）
FRIEND_IDS=填写好友的白名单，多个用,分隔（好友昵称）
MINE_NICKNAME=填写当前机器人bot的昵称，用于判断群中是否@自己
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