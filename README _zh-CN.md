# kvm-agent

<p align="center">
    <a href="./README.md"><b>English</b></a> •
    <a href="./README_zh-CN.md"><b>中文</b></a>
</p>

用于监控和报告虚拟机/主机指标的服务。

## 构建

在安装了 `go` 的本地平台中构建。
```bash
make build
```

编译 rpm/deb。
```bash
make build-rpm
make build-deb
```

## 运行

### 在 arm64 平台上安装 `kvm-agent

1. 下载本地文件中的 `kvm-agent_0.1.5_linux_arm64.rpm` 软件包。

2. 安装 `kvm-agent_0.1.5_linux_arm64.rpm`。
```bash
sudo rpm -i --force kvm-agent-1.0.0.aarch64.rpm
```

3. 修改 `/etc/kvm-agent/kvm-agent.conf` 以设置 `kvm-agent` 配置。
```bash
app
  version: 1.0.0
  debug: true # 调试模式
  log_file： /var/log/kvm-agent # 日志模式

agent：
  uuid: 6e8c2828-5460-49a2-99a9-6ed809ae4d1d # 代理的uuid
  period： 10 # 每 10 秒收集一次数据
  gzip: true # 压缩报告数据

redis：
  IP: 127.0.0.1 # 主机 redis IP
  port： 6379 # 主机 Redis 端口
  password: utyCyszHbhR9w5XX # 主机 redis 密码
  db: 0 # host redis db

dm：
  IP: 127.0.0.1 # 主机 dm IP
  port： 5236 # 主机 dm 端口
  username: SYSDBA # 主机 dm 用户名
  password: SYSDBA # 主机 dm 密码
```

4. 启动 `kvm-agent` 服务。
```bash
sudo systemctl start kvm-agent
```

5. 检查 `kvm-agent` 服务状态。
```bash
sudo systemctl status kvm-agent
```

6. 检查 `kvm-agent` 日志。
```bash
sudo journalctl -u kvm-agent
```