# MediaGo - Web Interface Guide

MediaGo now includes a built-in web interface for easier video downloading!

## 🎯 Quick Start

### Command Line Mode (Original)
```bash
mediago https://www.bilibili.com/video/BV1GJ411x7h7
```

### Web Interface Mode (New!)
```bash
mediago --webui
```

The web interface will automatically:
1. Start a local server on http://localhost:8080
2. Open your default browser
3. Display a beautiful GUI for video downloading

## 🖥️ Web Interface Features

- **Visual Interface**: No need to remember command-line flags
- **Form-based Input**: Easy URL, format, cookies, and proxy configuration
- **Real-time Feedback**: See download progress and results in the browser
- **Supports All Platforms**: Same 92 Chinese sites as CLI mode

## 📖 Web Interface Options

```bash
mediago --webui                    # Start on default port 8080
mediago --webui --webui-port 9000  # Start on custom port
```

Once the browser opens, you can:
- Paste video URL
- Select quality (auto/1080p/720p/480p)
- Provide cookies file path (for paid content)
- Configure proxy settings
- Click "Download" and see results

## 💡 Use Cases

**Use CLI when:**
- Batch downloading multiple URLs
- Scripting and automation
- SSH/remote server usage

**Use WebUI when:**
- Occasional downloads
- Prefer visual interface
- New to command-line tools
- Want to see supported sites easily

## 🔧 Technical Details

The web interface is embedded in the binary using Go's `embed` package - no external files needed. Everything runs locally on your machine, no data is sent to external servers.

## 📦 Requirements

Same as CLI mode:
- **ffmpeg** - required for most video platforms (HLS/DASH streams)

Place `ffmpeg` executable in the same directory as `mediago` or add it to your system PATH.
