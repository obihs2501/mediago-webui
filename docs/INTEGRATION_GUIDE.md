# MediaGo WebUI 集成指南

本指南详细说明如何将 WebUI 功能集成到 MediaGo 项目中。

## 📋 前置要求

- Go 1.21+
- Git
- MediaGo 源码

---

## 🚀 快速开始

### 1. 克隆 MediaGo 源码

```bash
git clone https://github.com/Sophomoresty/mediago.git
cd mediago
```

### 2. 下载集成文件

从本仓库下载以下文件：
- `integration/main.go`
- `integration/web/index.html`
- `integration/WEBUI.md`

### 3. 应用集成文件

```bash
# 备份原文件
cp cmd/mediago/main.go cmd/mediago/main.go.bak

# 复制新文件
cp path/to/integration/main.go cmd/mediago/main.go
cp -r path/to/integration/web cmd/mediago/
cp path/to/integration/WEBUI.md .
```

### 4. 编译测试

```bash
# 编译
go build -o mediago ./cmd/mediago

# 测试 CLI 模式
./mediago --list-extractors

# 测试 WebUI 模式
./mediago --webui
```

---

## 🔍 修改内容详解

### 1. main.go 修改

#### 添加导入包

```go
import (
	"embed"      // 新增：嵌入文件
	"net/http"   // 新增：HTTP 服务器
	"time"       // 新增：延迟等待
)
```

#### 添加 embed 声明

```go
//go:embed web/*
var webFiles embed.FS
```

#### 添加 WebUI 标志

```go
var (
	// ... 原有变量 ...
	webui     bool   // 新增
	webPort   string // 新增
)
```

#### 添加 WebUI 参数

在 `rootCmd.Flags()` 中添加：

```go
rootCmd.Flags().BoolVar(&webui, "webui", false, "start web interface")
rootCmd.Flags().StringVar(&webPort, "webui-port", "8080", "web interface port")
```

#### 修改 runMain 函数

```go
func runMain(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	// 新增：WebUI 模式检测
	if webui {
		return runWebUI(ctx)
	}

	// ... 原有逻辑 ...
}
```

#### 添加新函数

1. **runWebUI()** - WebUI 服务器启动
2. **handleDownload()** - 下载 API 端点
3. **handleGetSites()** - 站点列表 API
4. **openBrowser()** - 自动打开浏览器

---

## 📂 文件结构

集成后的项目结构：

```
mediago/
├── cmd/
│   └── mediago/
│       ├── main.go          # 修改（添加 WebUI）
│       ├── web/             # 新增
│       │   └── index.html   # 新增
│       ├── extractors.go    # 不变
│       ├── output.go        # 不变
│       └── ...
├── WEBUI.md                 # 新增
├── README.md                # 不变
└── ...
```

---

## 🧪 测试验证

### 功能测试清单

#### CLI 模式（原功能）
```bash
# 基本下载
./mediago https://www.bilibili.com/video/BV1xxx

# 列出格式
./mediago -F https://www.bilibili.com/video/BV1xxx

# 使用 Cookies
./mediago --cookies cookies.txt URL

# 列出提取器
./mediago --list-extractors

# 版本信息
./mediago --version
```

✅ 所有原有功能应正常工作

#### WebUI 模式（新功能）
```bash
# 启动 WebUI
./mediago --webui
```

验证项：
- ✅ 服务器成功启动在 localhost:8080
- ✅ 浏览器自动打开
- ✅ 界面正常显示
- ✅ 表单可以提交
- ✅ 下载功能正常
- ✅ 错误提示正常显示

#### 自定义端口
```bash
./mediago --webui --webui-port 9000
```

✅ 应在 localhost:9000 启动

---

## 🔧 常见问题

### Q: 编译失败："pattern web/*: no matching files found"

**原因**: `web` 目录不存在或为空

**解决**:
```bash
# 确认文件存在
ls -la cmd/mediago/web/

# 应该看到 index.html
```

### Q: 浏览器打开失败

**原因**: `openBrowser()` 函数依赖系统命令

**解决**: 手动打开 http://localhost:8080/web/

### Q: 下载失败

**原因**: 
1. ffmpeg 未安装
2. Cookies 未提供（付费内容）
3. 网络问题

**解决**: 检查日志输出，按提示操作

---

## 📦 发布编译

### 使用 GoReleaser

原项目的 `.goreleaser.yaml` 无需修改，会自动打包 `embed` 文件。

```bash
# 创建标签
git tag -a v0.3.0 -m "Add WebUI support"

# 推送标签（触发 GitHub Actions）
git push origin v0.3.0
```

GitHub Actions 会自动：
1. 编译所有平台版本
2. 打包嵌入的 web 文件
3. 创建 Release
4. 上传二进制文件

### 手动编译

```bash
# Windows
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o mediago.exe ./cmd/mediago

# Linux
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o mediago ./cmd/mediago

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o mediago ./cmd/mediago

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o mediago ./cmd/mediago
```

---

## 🎯 最佳实践

### 1. 保持代码同步

当 MediaGo 更新时：

```bash
# 更新源码
git pull origin main

# 重新应用集成文件
cp path/to/integration/main.go cmd/mediago/main.go

# 测试编译
go build ./cmd/mediago
```

### 2. 自定义界面

修改 `cmd/mediago/web/index.html`:

```html
<!-- 修改标题 -->
<title>我的 MediaGo</title>

<!-- 修改主题色 -->
<style>
  :root {
    --terracotta: #YOUR_COLOR;
  }
</style>
```

### 3. 添加新功能

在 `handleDownload()` 中添加新参数：

```go
func handleDownload(w http.ResponseWriter, r *http.Request) {
	// 获取新参数
	newParam := r.FormValue("new_param")
	
	// 使用新参数
	// ...
}
```

---

## 🔄 回滚

如果需要移除 WebUI 功能：

```bash
# 恢复原文件
mv cmd/mediago/main.go.bak cmd/mediago/main.go

# 删除 web 目录
rm -rf cmd/mediago/web

# 删除文档
rm WEBUI.md

# 重新编译
go build ./cmd/mediago
```

---

## 📞 支持

如有问题：
1. 查看本文档的「常见问题」
2. 提交 Issue 到本仓库
3. 查看 MediaGo 原项目文档

---

## 📜 变更日志

### v1.0.0 (2026-07-03)
- ✅ 初始版本
- ✅ 添加 `--webui` 参数
- ✅ 实现 Web 界面
- ✅ 实现下载 API
- ✅ 自动打开浏览器

---

**祝您集成顺利！** 🎉
