# MediaGo WebUI

> 🎯 为 [MediaGo](https://github.com/Sophomoresty/mediago) 下载器构建的桌面应用

[![Build Status](https://github.com/obihs2501/mediago-webui/workflows/Build%20Wails%20App/badge.svg)](https://github.com/obihs2501/mediago-webui/actions)
[![License](https://img.shields.io/badge/license-Unlicense-green.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org/dl/)

**真正的桌面应用** - 使用 Wails 构建，双击即可打开 GUI 窗口，无需浏览器！

![MediaGo WebUI](https://img.shields.io/badge/Platform-Windows-blue)

## ✨ 功能特性

- 🖥️ **自动打开浏览器** - 双击启动，自动打开网页界面
- 🎯 **一键启动** - 双击 `.exe` 文件即可使用
- 🔄 **实时反馈** - 下载进度即时显示
- ⚙️ **完整配置** - 支持格式选择、Cookies、代理等
- 🌐 **92个平台** - 支持所有 MediaGo 支持的中文视频/课程平台
- 🎨 **精美界面** - 温暖陶土色主题，现代化设计

## 📥 安装使用

### 下载预编译版本（推荐）

从 [Releases](https://github.com/obihs2501/mediago-webui/releases) 页面下载：

**Windows**: `mediago-webui.exe`

### ⚠️ 必需组件

**MediaGo WebUI 需要三个文件才能工作**：

1. **mediago-webui.exe** - 本项目的桌面应用
2. **mediago.exe** - 核心下载器 ([下载](https://github.com/Sophomoresty/mediago/releases))
3. **ffmpeg.exe** - 视频处理工具 ([下载](https://www.gyan.dev/ffmpeg/builds/)) ⭐ **必需**

**为什么需要 FFmpeg?**  
Bilibili 等平台使用 DASH 流格式，必须用 FFmpeg 合并音视频，否则会报错：`ffmpeg required for DASH streams`

### 推荐的文件结构

```
你的文件夹/
├── mediago-webui.exe    # 桌面应用（本项目）
├── mediago.exe          # 核心下载器
├── ffmpeg.exe           # 视频处理（必需！）
└── downloads/           # 下载目录（自动创建）
```

### 使用步骤

1. **下载三个 exe 文件**（见上方链接）
2. **放在同一目录**
3. **双击运行** `mediago-webui.exe`
4. **GUI 窗口自动打开**，开始下载！

📖 **详细安装指南**: 查看 [INSTALLATION.md](INSTALLATION.md)

## 🎯 使用方法

1. **启动应用** - 双击 `mediago-webui.exe`
2. **浏览器自动打开** - 显示网页界面
3. **配置选项**（可选）：
   - 选择视频画质
   - 设置 Cookies（付费内容需要）
   - 配置代理
4. **点击下载** - 即时显示下载结果
5. **文件位置** - 默认保存在 `downloads/` 目录

## 🛠️ 技术栈

- **后端**: Go 1.21+ (标准库 net/http)
- **前端**: 原生 HTML/CSS/JavaScript
- **嵌入**: embed.FS - 前端打包在 `.exe` 中
- **浏览器**: 自动启动系统默认浏览器
- **设计**: 温暖陶土色调 + 现代布局

## 🔧 从源码构建

### 前置要求

- Go 1.21+

### 构建步骤

```bash
# 1. 克隆仓库
git clone https://github.com/obihs2501/mediago-webui.git
cd mediago-webui

# 2. 构建
go build -o mediago-webui.exe .

# 3. 运行
./mediago-webui.exe
```

### 开发模式

```bash
go run .
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

**Q: 启动后没有打开浏览器？**  
A: 手动访问 http://localhost:8080/web/

**Q: 提示找不到 mediago？**  
A: 确保 `mediago.exe` 与 `mediago-webui.exe` 在同一目录

**Q: 提示找不到 ffmpeg？**  
A: 下载 FFmpeg 并放在与 `mediago-webui.exe` 同一目录。详见 [INSTALLATION.md](INSTALLATION.md)

**Q: 下载失败？**  
A: 
- 检查链接是否正确
- 付费内容需要提供 Cookies
- 部分平台需要代理访问
- 确保已安装 ffmpeg

## 🆚 与原 MediaGo 的区别

| 特性 | MediaGo | MediaGo WebUI |
|------|---------|---------------|
| 界面 | 命令行 | 浏览器 GUI |
| 使用方式 | 输入命令 | 点击按钮 |
| 配置 | 命令参数 | 可视化表单 |
| 反馈 | 终端输出 | 网页显示 |
| 启动 | 手动运行 | 自动打开浏览器 |
| 适合人群 | 开发者 | 所有用户 |

## 📄 许可证

与 MediaGo 项目相同，本项目使用 [The Unlicense](LICENSE) 发布到公有领域。

## 🙏 致谢

- [MediaGo](https://github.com/Sophomoresty/mediago) - 核心下载引擎
- [FFmpeg](https://ffmpeg.org/) - 视频处理工具

---

**立即下载，开始使用！** 🎉

访问 [Releases](https://github.com/obihs2501/mediago-webui/releases) 页面下载最新版本。
