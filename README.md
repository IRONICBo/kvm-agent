# kvm-agent

<p align="center">
    <a href="./README.md"><b>English</b></a> •
    <a href="./README_zh-CN.md"><b>中文</b></a>
</p>

A service to monitor and report VM/Host metrics.

## Build

Build in local platform with `go` installed.
```bash
make build
```

Build rpm/deb.
```bash
make build-rpm
make build-deb
```

## Run

### Install `kvm-agent` in arm64 platform

1. Download `kvm-agent_0.1.5_linux_arm64.rpm` packages in local files.

2. Install `kvm-agent_0.1.5_linux_arm64.rpm`.
```bash
sudo rpm -i --force kvm-agent-1.0.0.aarch64.rpm
```

3. Modify `/etc/kvm-agent/kvm-agent.conf` to set `kvm-agent` configuration.
```bash
app:
  version: 1.0.0
  debug: true # debug mode
  log_file: /var/log/kvm-agent # log mode

agent:
  uuid: 6e8c2828-5460-49a2-99a9-6ed809ae4d1d # agent uuid
  period: 10 # collect data every 5 seconds
  gzip: true # compress report data

redis:
  ip: 127.0.0.1 # host redis ip
  port: 6379  # host redis port
  password: utyCyszHbhR9w5XX # host redis password
  db: 0 # host redis db

dm:
  ip: 127.0.0.1 # host dm ip
  port: 5236 # host dm port
  username: SYSDBA # host dm username
  password: SYSDBA # host dm password
```

4. Start `kvm-agent` service.
```bash
sudo systemctl start kvm-agent
```

5. Check `kvm-agent` service status.
```bash
sudo systemctl status kvm-agent
```

6. Check `kvm-agent` log.
```bash
sudo journalctl -u kvm-agent
```