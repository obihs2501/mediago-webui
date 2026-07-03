# MediaGo WebUI Integration

> 🎯 为 [MediaGo](https://github.com/Sophomoresty/mediago) 添加内置 Web 界面的集成方案

[![License](https://img.shields.io/badge/license-Unlicense-green.svg)](LICENSE)
[![MediaGo](https://img.shields.io/badge/MediaGo-v0.2.0+-blue.svg)](https://github.com/Sophomoresty/mediago)

**一个命令，两种模式** - 为 MediaGo 下载器添加可选的图形界面

![使用方式](https://img.shields.io/badge/命令行-✅-success)
![WebUI](https://img.shields.io/badge/WebUI-✅-success)

---

## ✨ 特性

### 🎯 设计理念

本项目提供一个**集成方案**，将 Web 界面直接嵌入 MediaGo 可执行文件中：

- ✅ **单个 exe 文件** - 无需额外文件
- ✅ **双模式运行** - CLI 和 WebUI 随意切换
- ✅ **零外部依赖** - Web 界面打包在二进制中
- ✅ **完全兼容** - 不影响任何原有功能

### 🖥️ 使用方式

```bash
# 命令行模式（原功能）
mediago https://www.bilibili.com/video/BV1xxx

# Web 界面模式（新功能）
mediago --webui
```

启动 WebUI 后：
- ✅ 自动启动 HTTP 服务器（localhost:8080）
- ✅ 自动打开系统默认浏览器
- ✅ 显示精美的可视化界面
- ✅ 支持所有配置选项

---

## 📦 集成方案

### 方式 1：应用补丁到源码（推荐）

#### 需要的文件

1. **修改后的 `main.go`** - 添加 WebUI 功能
2. **Web 界面文件** - `web/index.html`
3. **文档** - `WEBUI.md`

#### 集成步骤

```bash
# 1. Clone MediaGo 源码
git clone https://github.com/Sophomoresty/mediago.git
cd mediago

# 2. 下载本项目的集成文件
# （从本仓库的 integration/ 目录）

# 3. 复制文件到对应位置
cp integration/main.go cmd/mediago/main.go
cp -r integration/web cmd/mediago/
cp integration/WEBUI.md .

# 4. 编译
go build -o mediago.exe ./cmd/mediago
```

#### 修改内容总结

```diff
+ import "embed"
+ import "net/http"

+ //go:embed web/*
+ var webFiles embed.FS

+ --webui flag
+ --webui-port flag
+ runWebUI() function
+ handleDownload() API
+ handleGetSites() API
+ openBrowser() function
```

---

## 🎨 界面预览

### 主界面
- 🌐 温暖的陶土色设计
- 📝 清晰的表单输入
- ⚙️ 可展开的高级选项
- 📊 实时下载反馈

### 功能
- ✅ 视频链接输入
- ✅ 画质选择（自动/1080p/720p/480p）
- ✅ Cookies 文件配置
- ✅ 代理设置
- ✅ 下载结果显示

---

## 🔧 技术细节

### 架构

```
MediaGo (single binary)
├── CLI Mode (原功能)
│   └── cobra commands
└── WebUI Mode (新功能)
    ├── embed.FS (embedded web files)
    ├── HTTP Server (localhost:8080)
    ├── API Endpoints
    │   ├── POST /api/download
    │   └── GET /api/sites
    └── Auto Browser Launch
```

### 技术栈

- **后端**: Go 标准库 `net/http`
- **嵌入**: Go 1.16+ `embed` 包
- **前端**: 纯 HTML/CSS/JavaScript
- **设计**: 响应式布局 + 温暖配色

### 优势

1. **零额外体积** - HTML/CSS/JS 压缩后约 10KB
2. **完全离线** - 无需网络，本地运行
3. **跨平台** - Windows/Linux/macOS 通用
4. **易维护** - 单个 `index.html` 文件

---

## 📖 使用文档

### 命令行选项

```bash
# 启动 WebUI（默认端口 8080）
mediago --webui

# 自定义端口
mediago --webui --webui-port 9000

# 命令行模式（不受影响）
mediago [原有所有参数]
```

### API 端点

```
POST /api/download
  - url: 视频链接
  - format: 画质选择
  - cookies: Cookies 文件路径
  - proxy: 代理地址

GET /api/sites
  - 返回支持的 92 个平台列表
```

---

## 🚀 为什么选择集成方案？

### vs. 独立 WebUI 项目

| 特性 | 本方案（集成） | 独立项目 |
|------|---------------|----------|
| 文件数量 | 1 个 exe | 2 个 exe |
| 配置 | 无需配置 | 需要指定 mediago 路径 |
| 更新 | 同步更新 | 需分别更新 |
| 分发 | 单文件分发 | 多文件打包 |
| 用户体验 | ⭐⭐⭐⭐⭐ | ⭐⭐⭐ |

### vs. 完全 GUI 应用（Wails/Electron）

| 特性 | 本方案 | Wails/Electron |
|------|--------|----------------|
| 体积 | +10KB | +10-100MB |
| 编译复杂度 | 简单 | 复杂 |
| CLI 功能 | 完整保留 | 需额外实现 |
| 开发难度 | 低 | 中高 |

---

## 📂 项目结构

```
mediago-webui/
├── README.md                    # 本文件
├── integration/                 # 集成文件
│   ├── main.go                  # 修改后的主程序
│   ├── web/
│   │   └── index.html           # Web 界面
│   └── WEBUI.md                 # 使用文档
├── docs/
│   ├── INTEGRATION_GUIDE.md     # 详细集成指南
│   └── API.md                   # API 文档
└── LICENSE
```

---

## 🤝 贡献

### 提交 PR 到原项目

如果您希望将此功能合并到 MediaGo 主项目：

1. Fork [Sophomoresty/mediago](https://github.com/Sophomoresty/mediago)
2. 应用本项目的集成文件
3. 测试编译和功能
4. 创建 Pull Request

### 改进本集成方案

欢迎：
- 🎨 改进界面设计
- 🔧 优化代码实现
- 📖 完善文档说明
- 🐛 报告问题

---

## 📄 许可证

与 MediaGo 项目相同，本项目使用 [The Unlicense](LICENSE) 发布到公有领域。

---

## 🙏 致谢

- [MediaGo](https://github.com/Sophomoresty/mediago) - 核心下载引擎
- [FFmpeg](https://ffmpeg.org/) - 视频处理工具

---

## 💡 常见问题

### Q: 会影响原有 CLI 功能吗？
**A**: 不会。所有原有功能完整保留，WebUI 是可选的附加功能。

### Q: 需要手动配置 mediago 路径吗？
**A**: 不需要。WebUI 直接调用内部函数，无需外部 exe。

### Q: 体积会增加多少？
**A**: 约 10KB（压缩后的 HTML/CSS/JS）。

### Q: 支持哪些平台？
**A**: 与 MediaGo 相同，支持所有 92 个中文平台。

### Q: 如何更新？
**A**: 更新 MediaGo 源码后重新应用集成补丁即可。

---

**立即体验！** 

查看 [集成指南](docs/INTEGRATION_GUIDE.md) 开始使用。
