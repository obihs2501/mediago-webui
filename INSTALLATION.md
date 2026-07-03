# MediaGo WebUI - 使用说明

## ⚠️ 重要：需要的组件

MediaGo WebUI 需要以下组件才能正常工作：

### 1. MediaGo 核心程序 ✅
**下载地址**: https://github.com/Sophomoresty/mediago/releases

- 下载 `mediago.exe`
- 放在与 `mediago-webui.exe` **同一目录**

### 2. FFmpeg ⭐ **必需**
**为什么需要**: Bilibili 等平台使用 DASH 流格式，必须用 ffmpeg 合并音视频

#### 方法 A：快速配置（推荐）
1. 下载 FFmpeg：https://www.gyan.dev/ffmpeg/builds/ffmpeg-release-essentials.zip
2. 解压后，将 `bin` 目录中的 `ffmpeg.exe` 复制到与 `mediago-webui.exe` **同一目录**

#### 方法 B：添加到系统 PATH
1. 下载并解压 FFmpeg
2. 将 FFmpeg 的 `bin` 目录路径添加到系统环境变量 PATH
3. 重启应用

#### 方法 C：放在常见位置
程序会自动检测以下路径：
- `C:\ffmpeg\bin\ffmpeg.exe`
- `C:\Program Files\ffmpeg\bin\ffmpeg.exe`
- `C:\softwares\ffmpeg\ffmpeg.exe`
- `C:\softwares\ffmpeg\bin\ffmpeg.exe`

---

## 📁 推荐的文件结构

```
你的文件夹/
├── mediago-webui.exe    # 桌面应用（本项目）
├── mediago.exe          # 核心下载器
├── ffmpeg.exe           # 视频处理工具
└── downloads/           # 下载的视频（自动创建）
```

---

## 🚀 启动应用

1. 确认上述三个 `.exe` 文件在同一目录
2. 双击 `mediago-webui.exe`
3. GUI 窗口自动打开
4. 开始下载！

---

## 🎯 下载视频

1. **粘贴链接** - 支持 Bilibili、知乎、小红书等 92 个平台
2. **选择画质**（可选）- 1080p / 720p / 480p
3. **高级选项**（可选）：
   - Cookies 文件（付费内容需要）
   - 代理设置
4. **点击"开始下载"**
5. 下载的视频保存在 `downloads/` 目录

---

## ❓ 常见问题

### Q: 提示 "ffmpeg not found"？
**A**: 下载 FFmpeg 并放在与 `mediago-webui.exe` 同一目录。  
下载地址: https://www.gyan.dev/ffmpeg/builds/

### Q: 提示 "mediago binary not found"？
**A**: 下载 `mediago.exe` 并放在与 `mediago-webui.exe` 同一目录。  
下载地址: https://github.com/Sophomoresty/mediago/releases

### Q: 下载失败？
**A**: 检查：
1. FFmpeg 是否已安装
2. 链接是否正确
3. 付费内容是否提供了 Cookies
4. 是否需要代理访问

### Q: 哪些平台需要 FFmpeg？
**A**: 大部分现代视频平台都需要，包括：
- Bilibili（B站）
- 知乎视频
- 小红书
- YouTube
- 等等...

---

## 📦 完整下载清单

| 文件 | 下载地址 | 说明 |
|------|----------|------|
| mediago-webui.exe | [本项目 Releases](https://github.com/obihs2501/mediago-webui/releases) | 桌面应用界面 |
| mediago.exe | [MediaGo Releases](https://github.com/Sophomoresty/mediago/releases) | 核心下载引擎 |
| ffmpeg.exe | [FFmpeg 官网](https://www.gyan.dev/ffmpeg/builds/) | 视频处理工具 |

---

## 🎊 现在开始

1. 下载三个 exe 文件
2. 放在同一文件夹
3. 双击 mediago-webui.exe
4. 享受可视化下载体验！

有问题？访问：https://github.com/obihs2501/mediago-webui/issues
