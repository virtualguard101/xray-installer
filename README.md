# xray-installer

在 VPS 本机执行的 Go 命令行工具。

启动后会交互式询问：
- 域名
- FlClash 节点名称（默认 `[自建 1] 美国家宽-Reality`）

然后自动继续安装。

```bash
xray-installer
```

它会自动：
- 检查运行环境（root / Linux / systemd / Debian/Ubuntu/Arch Linux）
- 校验域名 A 记录是否已经指向当前 VPS
- 安装 Xray
- 生成 VLESS + Reality + Vision 服务端配置
- 启动并启用 `xray` systemd 服务
- 生成文章同款长版 FlClash 配置 `proxy.yaml`
- 支持 `uninstall`
- 支持 `--dry-run`

## 当前边界

- 只支持 **Debian / Ubuntu / Arch Linux + systemd**
- 需要在 **VPS 本机**运行
- 域名 DNS 需要提前配好，并且能解析到当前 VPS
- 如果用了 Cloudflare，需要保证是 **DNS Only**

## 构建

```bash
go build ./cmd/xray-installer
```

## 运行

```bash
sudo ./xray-installer
```

## 帮助

```bash
./xray-installer --help
```

## 卸载

```bash
sudo ./xray-installer uninstall
```

## Dry run

```bash
./xray-installer --dry-run
```

生成文件：
- Xray 配置：`/usr/local/etc/xray/config.json`
- FlClash 配置：`/usr/local/etc/xray-installer/proxy.yaml`
- 安装记录：`/usr/local/etc/xray-installer/install-result.json`

## 验证

```bash
go test ./...
go build ./cmd/xray-installer
```

## 参考

- 文章：<https://manateelazycat.github.io/2026/04/08/best-proxy/>
- Xray 安装脚本：<https://github.com/XTLS/Xray-install>
- Xray 命令文档：<https://xtls.github.io/document/command>
