# MediaGo WebUI

> 🎯 为 [MediaGo](https://github.com/Sophomoresty/mediago) 下载器构建的现代化 Web 用户界面

[![Build Status](https://github.com/obihs2501/mediago-webui/workflows/Build%20and%20Release/badge.svg)](https://github.com/obihs2501/mediago-webui/actions)
[![License](https://img.shields.io/badge/license-Unlicense-green.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org/dl/)

通过浏览器轻松管理 MediaGo 下载任务，支持 92 个中文视频/课程平台。

## ✨ 功能特性

- 🎯 **可视化操作** - 通过 Web 界面提交和管理下载任务
- 🔄 **实时更新** - WebSocket 实时推送任务状态
- ⚙️ **高级选项** - 支持格式选择、Cookie、代理等配置
- 📊 **任务管理** - 查看任务历史、取消进行中的任务
- 🌐 **平台支持** - 查看所有 92 个支持的中文平台

## 📥 安装运行

### 方式一：下载预编译版本（推荐）

从 [Releases](https://github.com/obihs2501/mediago-webui/releases) 页面下载对应系统的二进制文件：

- **Windows**: `mediago-webui-windows-amd64.exe`
- **Linux**: `mediago-webui-linux-amd64`
- **macOS**: `mediago-webui-darwin-amd64`

下载后直接运行：

```bash
# Windows
mediago-webui-windows-amd64.exe

# Linux/macOS
chmod +x mediago-webui-*
./mediago-webui-linux-amd64
```

然后访问：**http://localhost:8080**

### 方式二：从源码构建

**前置要求**: Go 1.21+

```bash
# 1. 克隆仓库
git clone https://github.com/obihs2501/mediago-webui.git
cd mediago-webui

# 2. 安装依赖
go mod download

# 3. 构建
go build -o mediago-webui.exe server/main.go

# 4. 运行
./mediago-webui.exe
```

### 方式三：使用 Docker

```bash
docker-compose up -d
```

### ⚠️ 重要：MediaGo 二进制文件

确保 `mediago.exe` 在以下位置之一：
- WebUI 同级目录
- WebUI 目录内
- 系统 PATH 中

下载地址：https://github.com/Sophomoresty/mediago/releases

## 使用说明

### 基本下载

1. 在"视频链接"输入框中粘贴视频 URL
2. 点击"开始下载"按钮
3. 任务会出现在下方的任务列表中

### 高级选项

点击"高级选项"展开更多配置：

- **格式选择** - 选择视频画质 (1080p/720p/480p/best/worst)
- **输出模板** - 自定义输出文件名和路径
- **Cookies 文件** - 用于需要登录的付费内容
- **从浏览器读取** - 自动从 Chrome/Edge/Firefox 提取 Cookies
- **代理** - 配置 HTTP/SOCKS5 代理
- **下载整个播放列表** - 下载课程/播放列表中的所有内容

### 任务管理

- 实时查看任务状态（等待中/下载中/已完成/失败）
- 取消正在进行的下载任务
- 查看任务错误信息
- 点击"刷新"手动更新任务列表

### 查看支持的平台

点击"支持的平台"查看所有 92 个受支持的中文视频/课程平台列表。

## API 接口

### POST `/api/download`

提交下载任务

**请求体:**
```json
{
  "url": "https://www.bilibili.com/video/BV1GJ411x7h7",
  "format": "1080p",
  "output": "downloads/%(title)s.%(ext)s",
  "cookies": "cookies.txt",
  "cookies_browser": "chrome",
  "proxy": "socks5://127.0.0.1:1080",
  "yes_playlist": false
}
```

### GET `/api/tasks`

获取所有任务列表

### GET `/api/tasks/{id}`

获取单个任务详情

### POST `/api/tasks/{id}/cancel`

取消指定任务

### GET `/api/extractors`

获取支持的平台列表

### WebSocket `/ws`

实时任务更新推送

## 项目结构

```
mediago-webui/
├── server/
│   └── main.go           # Go 后端服务
├── web/
│   ├── templates/
│   │   └── index.html    # 主页面
│   └── static/
│       ├── css/
│       │   └── style.css # 样式
│       └── js/
│           └── app.js    # 前端逻辑
├── downloads/            # 默认下载目录（自动创建）
├── go.mod
└── README.md
```

## 配置

### 环境变量

- `PORT` - 服务端口（默认: 8080）

示例：
```bash
PORT=3000 go run server/main.go
```

## 技术栈

- **后端**: Go + gorilla/mux + gorilla/websocket
- **前端**: 原生 HTML/CSS/JavaScript
- **实时通信**: WebSocket

## 注意事项

1. **MediaGo 路径**: 服务会按以下顺序查找 `mediago` 二进制：
   - 当前目录
   - 上级目录
   - 系统 PATH

2. **下载目录**: 默认下载到 `downloads/` 目录（可在高级选项中自定义）

3. **Cookies**: 下载付费内容时需要提供登录 Cookies

4. **代理**: 某些平台可能需要代理才能正常访问

## 许可证

与 MediaGo 项目相同，本项目同样使用 [The Unlicense](../LICENSE) 发布到公有领域。

## 致谢

- [MediaGo](https://github.com/Sophomoresty/mediago) - 核心下载引擎
- [gorilla/mux](https://github.com/gorilla/mux) - HTTP 路由
- [gorilla/websocket](https://github.com/gorilla/websocket) - WebSocket 支持
