# MediaGo WebUI

> 🎯 为 [MediaGo](https://github.com/Sophomoresty/mediago) 下载器构建的桌面应用

[![Build Status](https://github.com/obihs2501/mediago-webui/workflows/Build%20Wails%20App/badge.svg)](https://github.com/obihs2501/mediago-webui/actions)
[![License](https://img.shields.io/badge/license-Unlicense-green.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org/dl/)

**真正的桌面应用** - 使用 Wails 构建，双击即可打开 GUI 窗口，无需浏览器！

![MediaGo WebUI](https://img.shields.io/badge/Platform-Windows-blue)

## ✨ 功能特性

- 🖥️ **原生桌面窗口** - 不再需要打开浏览器访问 localhost
- 🎯 **一键启动** - 双击 `.exe` 文件即可使用
- 🔄 **实时反馈** - 下载进度即时显示
- ⚙️ **完整配置** - 支持格式选择、Cookies、代理等
- 🌐 **92个平台** - 支持所有 MediaGo 支持的中文视频/课程平台
- 🎨 **精美界面** - 温暖陶土色主题，现代化设计

## 📥 安装使用

### 下载预编译版本（推荐）

从 [Releases](https://github.com/obihs2501/mediago-webui/releases) 页面下载：

**Windows**: `mediago-webui.exe`

### 使用步骤

1. **下载** `mediago-webui.exe`
2. **确保** `mediago.exe` 在同一目录或 PATH 中
3. **双击运行** `mediago-webui.exe`
4. **GUI 窗口自动打开**，开始下载！

### ⚠️ 重要：需要 MediaGo 核心程序

确保 `mediago.exe` 在以下位置之一：
- 与 `mediago-webui.exe` 同一目录 ⭐（推荐）
- 系统 PATH 中

下载地址：https://github.com/Sophomoresty/mediago/releases

## 🎯 使用方法

1. **启动应用** - 双击 `mediago-webui.exe`
2. **输入链接** - 粘贴视频 URL
3. **配置选项**（可选）：
   - 选择视频画质
   - 设置 Cookies（付费内容需要）
   - 配置代理
4. **点击下载** - 即时显示下载结果
5. **文件位置** - 默认保存在 `downloads/` 目录

## 🛠️ 技术栈

- **框架**: [Wails v2](https://wails.io/) - Go + Web 混合桌面应用
- **后端**: Go 1.21+
- **前端**: 原生 HTML/CSS/JavaScript
- **窗口**: WebView2（Windows 原生）
- **设计**: 温暖陶土色调 + 现代布局

## 🔧 从源码构建

### 前置要求

- Go 1.21+
- Wails CLI: `go install github.com/wailsapp/wails/v2/cmd/wails@latest`
- Windows: WebView2 Runtime（通常系统已安装）

### 构建步骤

```bash
# 1. 克隆仓库
git clone https://github.com/obihs2501/mediago-webui.git
cd mediago-webui

# 2. 构建
wails build

# 3. 运行
./build/bin/mediago-webui.exe
```

### 开发模式

```bash
wails dev
```

## 📦 项目结构

```
mediago-webui/
├── main.go              # Wails 应用入口
├── mediago.go           # MediaGo 调用逻辑
├── frontend/
│   └── dist/
│       └── index.html   # 内嵌的 Web 界面
├── wails.json           # Wails 配置
├── go.mod               # Go 依赖
└── .github/
    └── workflows/
        └── build.yml    # 自动构建配置
```

## 🚀 GitHub Actions 自动构建

每次推送代码或创建标签时，GitHub Actions 会自动：
1. 安装 Wails 工具链
2. 编译 Windows 桌面应用
3. 上传编译产物

创建 Release：
```bash
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

## 🎨 界面预览

- 温暖的奶油色背景（#F7F4EF）
- 陶土色强调元素（#C4612F）
- Georgia 衬线标题 + Inter 无衬线正文
- 圆角卡片 + 柔和阴影
- 响应式布局

## ❓ 常见问题

**Q: 双击没反应？**  
A: 检查是否已安装 WebView2 Runtime（Windows 10/11 通常已内置）

**Q: 提示找不到 mediago？**  
A: 确保 `mediago.exe` 与 `mediago-webui.exe` 在同一目录

**Q: 下载失败？**  
A: 
- 检查链接是否正确
- 付费内容需要提供 Cookies
- 部分平台需要代理访问

**Q: 能在 Linux/macOS 上用吗？**  
A: 当前仅构建 Windows 版本，理论上可以构建其他平台

## 🆚 与原 MediaGo 的区别

| 特性 | MediaGo | MediaGo WebUI |
|------|---------|---------------|
| 界面 | 命令行 | 桌面 GUI |
| 使用方式 | 输入命令 | 点击按钮 |
| 配置 | 命令参数 | 可视化表单 |
| 反馈 | 终端输出 | 窗口显示 |
| 适合人群 | 开发者 | 所有用户 |

## 📄 许可证

与 MediaGo 项目相同，本项目使用 [The Unlicense](LICENSE) 发布到公有领域。

## 🙏 致谢

- [MediaGo](https://github.com/Sophomoresty/mediago) - 核心下载引擎
- [Wails](https://wails.io/) - Go 桌面应用框架
- [WebView2](https://developer.microsoft.com/microsoft-edge/webview2/) - Windows 原生渲染

---

**立即下载，开始使用！** 🎉

访问 [Releases](https://github.com/obihs2501/mediago-webui/releases) 页面下载最新版本。
