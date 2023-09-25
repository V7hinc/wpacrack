
## 使用ubuntu环境
1. 安装aircrack-ng
```
sudo apt install aircrack-ng -y
```
2. 拉取密码字典
```
git clone https://github.com/conwnet/wpa-dictionary.git wpa-dictionary
```
3. 构建go代码
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./crack_amd64 ./cli/main.go
```
4. 将cap文件放在cap_file目录下
5. 测试wechat bot并运行字典爆破， 看是否能收到消息
```
./crack_amd64 dictCrack -c cap_file -p wpa-dictionary -k e34de5c8-1543-45f8-bc9d-c152e147d76b -b
```
或直接运行爆破
```
./crack_amd64 dictCrack -c cap_file -p wpa-dictionary -k e34de5c8-1543-45f8-bc9d-c152e147d76b
```
6. 测试wechat bot并运行所有密码爆破（时间超久）， 看是否能收到消息
```
./crack_amd64 allCrack -c cap_file -k e34de5c8-1543-45f8-bc9d-c152e147d76b -b
```
或直接运行爆破
```
./crack_amd64 allCrack -c cap_file -k e34de5c8-1543-45f8-bc9d-c152e147d76b
```

## docker部署模式
1. 构建docker镜像
```
docker build -t ghcr.io/v7hinc/wpacrack:latest .
或直接拉取构建好的镜像
docker pull ghcr.io/v7hinc/wpacrack:latest
```
2. 拉取密码字典
```
git clone https://github.com/conwnet/wpa-dictionary.git wpa-dictionary
```
3. 创建cap_file文件夹并放入cap文件
```
mkdir cap_file
```
4. 运行爆破，key为企业微信机器人key
```
docker run --rm -it -v `pwd`/cap_file:/app/cap_file -v `pwd`/wpa-dictionary:/app/wpa-dictionary -e key="e34de5c8-1543-45f8-bc9d-c152e147d76b" ghcr.io/v7hinc/wpacrack:latest
```

