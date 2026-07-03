# MediaGo WebUI API 文档

MediaGo WebUI 提供简单的 HTTP API 用于下载视频和查询支持的平台。

## 📍 Base URL

```
http://localhost:8080
```

默认端口为 8080，可通过 `--webui-port` 参数修改。

---

## 🌐 端点

### 1. 下载视频

#### `POST /api/download`

下载指定 URL 的视频。

**请求参数**

| 参数 | 类型 | 必需 | 说明 |
|------|------|------|------|
| url | string | ✅ | 视频链接 |
| format | string | ❌ | 画质选择（best/1080p/720p/480p） |
| cookies | string | ❌ | Cookies 文件路径 |
| proxy | string | ❌ | 代理地址 |

**请求示例**

```bash
curl -X POST http://localhost:8080/api/download \
  -d "url=https://www.bilibili.com/video/BV1xxx" \
  -d "format=1080p" \
  -d "cookies=cookies.txt" \
  -d "proxy=socks5://127.0.0.1:1080"
```

**响应**

成功（200）：
```
Download completed successfully!
```

失败（500）：
```
Error: [错误信息]
```

---

### 2. 获取支持的平台列表

#### `GET /api/sites`

获取所有支持的视频平台列表。

**请求示例**

```bash
curl http://localhost:8080/api/sites
```

**响应示例**

```
Bilibili: https://www.bilibili.com
Douyin: https://www.douyin.com
iCourse163: https://www.icourse163.org (auth)
...

92 extractors
```

---

## 📝 完整示例

### JavaScript (Fetch API)

```javascript
// 下载视频
async function downloadVideo() {
  const formData = new FormData();
  formData.append('url', 'https://www.bilibili.com/video/BV1xxx');
  formData.append('format', '1080p');
  formData.append('cookies', 'cookies.txt');
  formData.append('proxy', 'socks5://127.0.0.1:1080');

  try {
    const response = await fetch('/api/download', {
      method: 'POST',
      body: formData
    });

    const text = await response.text();
    
    if (response.ok) {
      console.log('Success:', text);
    } else {
      console.error('Error:', text);
    }
  } catch (error) {
    console.error('Request failed:', error);
  }
}

// 获取平台列表
async function getSites() {
  try {
    const response = await fetch('/api/sites');
    const text = await response.text();
    console.log(text);
  } catch (error) {
    console.error('Request failed:', error);
  }
}
```

### Python (requests)

```python
import requests

# 下载视频
def download_video():
    url = "http://localhost:8080/api/download"
    data = {
        'url': 'https://www.bilibili.com/video/BV1xxx',
        'format': '1080p',
        'cookies': 'cookies.txt',
        'proxy': 'socks5://127.0.0.1:1080'
    }
    
    response = requests.post(url, data=data)
    
    if response.status_code == 200:
        print('Success:', response.text)
    else:
        print('Error:', response.text)

# 获取平台列表
def get_sites():
    url = "http://localhost:8080/api/sites"
    response = requests.get(url)
    print(response.text)
```

### cURL

```bash
# 下载视频
curl -X POST http://localhost:8080/api/download \
  -d "url=https://www.bilibili.com/video/BV1xxx" \
  -d "format=1080p"

# 获取平台列表
curl http://localhost:8080/api/sites
```

---

## 🎯 参数详解

### format (画质选择)

| 值 | 说明 |
|------|------|
| `best` | 最佳画质（默认） |
| `worst` | 最低画质 |
| `1080p` | 1080p 画质 |
| `720p` | 720p 画质 |
| `480p` | 480p 画质 |

### cookies (Cookies 文件)

用于下载需要登录或付费的内容。

**格式**: Netscape cookies.txt 格式

**生成方式**:
1. 浏览器插件：[Get cookies.txt](https://chrome.google.com/webstore/detail/get-cookiestxt/bgaddhkoddajcdgocldbbfleckgcbcid)
2. 命令行参数：`--cookies-from-browser chrome`

**示例**:
```
# Netscape HTTP Cookie File
.bilibili.com	TRUE	/	FALSE	1234567890	SESSDATA	abcd1234...
```

### proxy (代理)

支持的代理类型：
- HTTP: `http://127.0.0.1:8080`
- HTTPS: `https://127.0.0.1:8080`
- SOCKS5: `socks5://127.0.0.1:1080`

---

## ⚠️ 错误处理

### 常见错误

#### 1. URL 缺失
```
Status: 400
Body: URL is required
```

**解决**: 提供 `url` 参数

#### 2. 不支持的平台
```
Status: 500
Body: Error: unsupported URL: [URL]
Use --list-extractors to see supported sites.
```

**解决**: 检查 URL 是否在支持列表中

#### 3. ffmpeg 缺失
```
Status: 500
Body: Error: download failed: ffmpeg required for DASH streams
```

**解决**: 安装 ffmpeg 并添加到 PATH

#### 4. Cookies 无效
```
Status: 500
Body: Error: failed to load cookies: [错误]
```

**解决**: 检查 cookies 文件格式和路径

---

## 🔒 安全性

### 本地运行

WebUI 仅监听 `127.0.0.1`，不接受外部连接：

```go
addr := "127.0.0.1:8080"  // 仅本地访问
```

### 无认证

WebUI 不提供身份验证，**请勿暴露到公网**。

如需远程访问，建议：
1. 使用 SSH 隧道
2. 配置反向代理（Nginx + 认证）
3. 使用 VPN

---

## 📊 性能考虑

### 并发请求

WebUI 使用 Go 的 `net/http` 包，支持高并发。

但下载本身受限于：
- 网络带宽
- 目标网站限速
- ffmpeg 处理速度

### 超时设置

默认无超时限制，长视频可能需要较长时间。

可在代码中添加超时：

```go
server := &http.Server{
	Addr:         addr,
	Handler:      mux,
	ReadTimeout:  10 * time.Minute,
	WriteTimeout: 10 * time.Minute,
}
```

---

## 🛠️ 扩展 API

### 添加新端点

在 `main.go` 中添加：

```go
func runWebUI(ctx context.Context) error {
	mux := http.NewServeMux()
	
	// 现有端点
	mux.HandleFunc("/api/download", handleDownload)
	mux.HandleFunc("/api/sites", handleGetSites)
	
	// 新端点
	mux.HandleFunc("/api/status", handleStatus)
	
	// ...
}

func handleStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"ok","version":"0.2.0"}`)
}
```

### 添加请求验证

```go
func handleDownload(w http.ResponseWriter, r *http.Request) {
	// 验证方法
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// 验证参数
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}
	
	// ... 处理逻辑
}
```

---

## 📚 相关文档

- [集成指南](INTEGRATION_GUIDE.md)
- [MediaGo 主项目](https://github.com/Sophomoresty/mediago)

---

**Happy Coding!** 🚀
