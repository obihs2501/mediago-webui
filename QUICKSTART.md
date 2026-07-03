# MediaGo WebUI - 快速启动指南

## 📦 项目已创建完成！

目录结构：
```
mediago-webui/
├── server/main.go              # Go 后端服务
├── web/
│   ├── templates/index.html    # 主页面
│   └── static/
│       ├── css/style.css       # 样式（温暖的陶土色主题）
│       └── js/app.js           # 前端 JavaScript
├── go.mod                      # Go 依赖管理
├── README.md                   # 详细文档
├── run.bat                     # Windows 快速启动
└── run.sh                      # Linux/Mac 快速启动
```

## 🚀 启动步骤

### Windows 用户

**方法 1: 使用启动脚本（推荐）**
```cmd
run.bat
```

**方法 2: 手动启动**
```cmd
# 1. 初始化依赖（首次运行）
go mod download

# 2. 构建并运行
go build -o mediago-webui.exe server/main.go
mediago-webui.exe
```

### Linux/Mac 用户

```bash
# 添加执行权限
chmod +x run.sh

# 运行
./run.sh
```

或手动运行：
```bash
go mod download
go build -o mediago-webui server/main.go
./mediago-webui
```

## 🌐 访问界面

启动后，在浏览器中访问：
```
http://localhost:8080
```

## ✨ 功能特性

✅ **可视化下载** - 通过 Web 表单提交下载任务  
✅ **实时更新** - WebSocket 实时推送任务进度  
✅ **高级配置** - 支持格式、Cookies、代理等选项  
✅ **任务管理** - 查看历史、取消任务  
✅ **平台列表** - 查看 92 个支持的平台  
✅ **精美设计** - 温暖陶土色主题，响应式布局  

## 📝 使用示例

1. **简单下载**
   - 粘贴 Bilibili 视频链接
   - 点击"开始下载"

2. **下载付费内容**
   - 展开"高级选项"
   - 选择"从浏览器读取" → Chrome/Edge/Firefox
   - 提交下载

3. **下载整个课程**
   - 粘贴课程/播放列表链接
   - 勾选"下载整个播放列表"
   - 开始下载

4. **自定义输出**
   - 高级选项 → 输出模板
   - 例如: `videos/%(site)s/%(title)s.%(ext)s`

## 🔧 配置

### 自定义端口
```bash
# Windows
set PORT=3000
mediago-webui.exe

# Linux/Mac
PORT=3000 ./mediago-webui
```

### MediaGo 位置
确保 `mediago.exe` 在以下位置之一：
- 当前目录 (`mediago-webui/`)
- 上级目录 (`mediago_0.2.0_windows_amd64/`)
- 系统 PATH 中

## 🛠️ 技术栈

- **后端**: Go + Gorilla Mux + WebSocket
- **前端**: 原生 HTML/CSS/JavaScript
- **设计**: 温暖陶土色调 + 衬线/无衬线混搭

## 📊 API 端点

- `POST /api/download` - 提交下载任务
- `GET /api/tasks` - 获取任务列表
- `GET /api/tasks/{id}` - 获取单个任务
- `POST /api/tasks/{id}/cancel` - 取消任务
- `GET /api/extractors` - 支持的平台
- `WebSocket /ws` - 实时更新

## ⚠️ 注意事项

1. **首次运行需要 Go 环境** (1.21+)
2. **下载文件默认保存在 `downloads/` 目录**
3. **付费内容需要提供 Cookies**
4. **部分平台可能需要代理访问**

## 🐛 故障排除

### 提示 "mediago binary not found"
- 确认 mediago.exe 在正确位置
- 或将完整路径添加到 PATH

### 无法连接 WebSocket
- 检查防火墙设置
- 确认端口 8080 未被占用

### 下载失败
- 查看任务错误信息
- 尝试使用代理
- 检查是否需要 Cookies

## 📖 更多信息

详细文档请查看 `README.md`

---

**享受使用 MediaGo WebUI！** 🎉
