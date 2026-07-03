# MediaGo with Desktop GUI

> 🎯 MediaGo 的桌面 GUI 版本 - 双击即用，支持 92 个中文平台

[![License](https://img.shields.io/badge/license-Unlicense-green.svg)](LICENSE)
[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![Fyne](https://img.shields.io/badge/GUI-Fyne-orange.svg)](https://fyne.io)

**双模式运行** - 双击启动桌面 GUI，命令行使用 CLI 模式

![Desktop GUI](https://img.shields.io/badge/Desktop_GUI-✅-success)
![CLI Mode](https://img.shields.io/badge/CLI_Mode-✅-success)

---

## ✨ 特性

### 🎯 双模式操作

```bash
# 双击 mediago.exe → 打开桌面 GUI 窗口
mediago.exe

# 命令行使用 → CLI 模式（保留所有原功能）
mediago.exe https://www.bilibili.com/video/BV1xxx
```

### 🖥️ 桌面 GUI 模式

- ✅ **真正的桌面应用** - 使用 Fyne 原生 GUI
- ✅ **无需浏览器** - 独立桌面窗口
- ✅ **简单易用** - 可视化表单输入
- ✅ **实时反馈** - 下载进度实时显示
- ✅ **完整功能** - 支持格式选择、Cookies、代理

### 💻 CLI 模式（保留）

- ✅ **所有原功能** - 完全兼容原 MediaGo
- ✅ **批量下载** - 多 URL 支持
- ✅ **脚本自动化** - 适合批处理
- ✅ **远程服务器** - SSH 环境可用

---

## 📥 安装

### 下载预编译版本（推荐）

从 [Releases](https://github.com/obihs2501/mediago-webui/releases) 下载 Windows 版本：

- **Windows**: `mediago-windows-amd64.zip`

解压后双击即用！

> 注意：目前仅提供 Windows 版本。macOS 和 Linux 用户请使用原 [MediaGo](https://github.com/Sophomoresty/mediago) CLI 版本。

### 必需依赖

- **ffmpeg** - 用于视频合并和格式转换
  - Windows: [下载 ffmpeg](https://www.gyan.dev/ffmpeg/builds/)
  - macOS: `brew install ffmpeg`
  - Linux: `sudo apt install ffmpeg`

---

## 🎯 使用方法

### GUI 模式（双击启动）

1. **双击** `mediago.exe`
2. **填写表单**：
   - 视频链接
   - 画质选择（可选）
   - Cookies 文件（付费内容需要）
   - 代理地址（可选）
3. **点击下载**
4. **查看结果** - 实时显示在窗口中

### CLI 模式（命令行）

```bash
# 下载视频
mediago https://www.bilibili.com/video/BV1GJ411x7h7

# 使用 Cookies（付费内容）
mediago --cookies cookies.txt URL

# 选择画质
mediago -f 1080p URL

# 列出可用格式
mediago -F URL

# 下载整个课程/播放列表
mediago --yes-playlist URL

# 使用代理
mediago --proxy socks5://127.0.0.1:1080 URL
```

完整命令参数见原 [MediaGo 文档](https://github.com/Sophomoresty/mediago)

---

## 🛠️ 技术栈

- **核心**: Go 1.21+
- **GUI**: [Fyne v2.5](https://fyne.io) - 跨平台原生 GUI
- **CLI**: [Cobra](https://github.com/spf13/cobra) - 命令行框架
- **下载引擎**: 原 MediaGo 引擎（92 个平台支持）

### 为什么选择 Fyne？

| 特性 | Fyne | 
|------|------|
| 编译复杂度 | ⭐⭐⭐⭐⭐ |
| Windows 体积 | 15-20MB |
| 原生外观 | ✅ |
| 跨平台 | ✅ (目前仅 Windows) |
| 纯 Go | ✅ |

> **为什么只有 Windows 版本？**
> 
> Fyne 需要 CGO 编译，在 GitHub Actions 上跨平台编译较复杂。目前专注于 Windows 用户体验。
> 
> macOS/Linux 用户推荐使用原 [MediaGo CLI](https://github.com/Sophomoresty/mediago)。

---

## 🚀 从源码构建

### 前置要求

- Go 1.21+
- Git

### 构建步骤

```bash
# 1. 克隆仓库
git clone https://github.com/obihs2501/mediago-webui.git
cd mediago-webui

# 2. 下载依赖
go mod download

# 3. 编译
go build -o mediago ./cmd/mediago

# 4. 运行
./mediago          # GUI 模式
./mediago [URL]    # CLI 模式
```

### 编译选项

```bash
# 最小化体积（去除调试信息）
go build -ldflags="-s -w" -o mediago ./cmd/mediago

# Windows GUI 模式（无控制台窗口）
go build -ldflags="-s -w -H windowsgui" -o mediago.exe ./cmd/mediago
```

---

## 📦 项目结构

```
mediago-webui/
├── cmd/
│   └── mediago/
│       ├── main.go          # 主程序（GUI + CLI）
│       ├── extractors.go    # 提取器列表
│       └── output.go        # 输出格式化
├── internal/
│   ├── extractor/           # 92 个平台提取器
│   ├── download/            # 下载引擎
│   ├── cookie/              # Cookie 处理
│   └── util/                # 工具函数
├── go.mod                   # Go 依赖
└── README.md                # 本文件
```

---

## 🌐 支持的平台

支持 **92 个中文视频/课程平台**，包括：

- **视频平台**: Bilibili, Douyin, 抖音
- **在线课程**: iCourse163, Xuetang, Chaoxing, Zhihuishu
- **职业培训**: Huatu, Gaodun, Fenbi, Med66
- **企业培训**: DingTalk, Feishu, ClassIn
- **更多...** 完整列表见 `mediago --list-extractors`

---

## 🆚 对比

### vs. 原 MediaGo

| 特性 | 原 MediaGo | 本项目 |
|------|-----------|--------|
| CLI 模式 | ✅ | ✅ |
| GUI 模式 | ❌ | ✅ (Fyne 桌面窗口) |
| 使用难度 | 需要命令行知识 | 双击即用 |
| 适合人群 | 开发者 | 所有用户 |
| 功能 | 完整 | 完整 |

### vs. 其他 WebUI 方案

| 特性 | 本项目 | 浏览器 WebUI |
|------|--------|-------------|
| 启动方式 | 双击 | 双击 + 等待浏览器 |
| 窗口类型 | 原生桌面窗口 | 浏览器标签页 |
| 外观 | 系统原生 | Web 样式 |
| 资源占用 | 低 | 中（浏览器占用） |

---

## 💡 常见问题

### Q: 为什么双击没反应？

**A**: 
1. 检查是否已安装 ffmpeg
2. 在命令行运行查看错误信息：`./mediago.exe`

### Q: 提示找不到 ffmpeg？

**A**: 
- Windows: 下载 ffmpeg.exe 并放在与 mediago.exe 同目录
- macOS/Linux: 安装 ffmpeg 到系统 PATH

### Q: GUI 和 CLI 有什么区别？

**A**: 
- GUI: 双击启动，图形界面，适合日常使用
- CLI: 命令行，功能更强大，适合批量处理和自动化

### Q: 付费内容如何下载？

**A**: 需要提供 Cookies 文件：
1. 浏览器登录并购买课程
2. 使用插件导出 cookies.txt
3. GUI 模式：在表单中填写 cookies 文件路径
4. CLI 模式：`mediago --cookies cookies.txt URL`

### Q: 支持哪些平台？

**A**: 支持 92 个中文平台，运行 `mediago --list-extractors` 查看完整列表

---

## 🤝 贡献

欢迎贡献！

- 🐛 报告 Bug
- 💡 提出新功能
- 📝 改进文档
- 🔧 提交代码

---

## 📄 许可证

与原 MediaGo 项目相同，使用 [The Unlicense](LICENSE) 发布到公有领域。

---

## 🙏 致谢

- [MediaGo](https://github.com/Sophomoresty/mediago) - 核心下载引擎
- [Fyne](https://fyne.io) - 跨平台 GUI 框架
- [FFmpeg](https://ffmpeg.org/) - 视频处理工具

---

## 📞 支持

- **问题反馈**: [GitHub Issues](https://github.com/obihs2501/mediago-webui/issues)
- **原项目**: [MediaGo](https://github.com/Sophomoresty/mediago)

---

**立即下载，开始使用！** 🚀

[📥 下载最新版本](https://github.com/obihs2501/mediago-webui/releases)
