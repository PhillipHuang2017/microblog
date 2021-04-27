## 功能
用户注册、登陆、获取自己的信息

## 启动
需要安装docker。   
```bash
cd scripts
./create_project_env.sh
cd ..
docker compose -f docker-compose.yml up -d
```
启动后的服务状态：
```
$ docker compose ps
NAME                SERVICE             STATUS              PORTS
etcd                etcd                running             2379/tcp, 2380/tcp
mysql               mysql               running             3306/tcp, 33060/tcp
redis               redis               running             6379/tcp
user-api            user-api            running             0.0.0.0:8080->8080/tcp
user-rpc            user-rpc            running
```
停止项目：   
```
docker compose down
```

## 测试
```bash
# 注册
curl -X POST -H "Content-Type: application/json" -d '{"username": "phillip", "password": "123", "phone": "", "email": ""}' localhost:8080/user/register

# 登陆
curl -X POST -H "Content-Type: application/json" -d '{"username": "phillip", "password": "123", "phone": "", "email": ""}' localhost:8080/user/login
# {"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTk1MjY1NTYsImlhdCI6MTYxOTUyNDc1NiwidWlkIjoiMSJ9.fuArMjJtrs-jGLzFbYvQlMwRXCuqIuAZNb6RS0sG11U"}

# 获取自己的信息
curl -X GET -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTk1MjY1NTYsImlhdCI6MTYxOTUyNDc1NiwidWlkIjoiMSJ9.fuArMjJtrs-jGLzFbYvQlMwRXCuqIuAZNb6RS0sG11U" localhost:8080/user/info
# {"Id":"1","username":"phillip","gender":"unknown","birthday":"1900-01-01"}
```