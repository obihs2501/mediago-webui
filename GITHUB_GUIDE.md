# 🚀 MediaGo WebUI - GitHub 发布指南

## 📋 准备工作清单

✅ 项目已初始化 Git 仓库  
✅ 所有文件已暂存  
✅ GitHub Actions 工作流已配置  
✅ 多平台编译支持（Windows/Linux/macOS）  

---

## 🎯 发布到 GitHub 的步骤

### 第 1 步：在 GitHub 创建仓库

1. 访问 https://github.com/new
2. 填写信息：
   - **Repository name**: `mediago-webui`
   - **Description**: `Modern Web UI for MediaGo downloader - 支持 92 个中文平台的可视化下载器`
   - **Public** 或 **Private**（推荐 Public）
   - ❌ 不要勾选 "Add README" / "Add .gitignore" / "Add license"（我们已经创建了）

3. 点击 **Create repository**

### 第 2 步：推送代码到 GitHub

复制 GitHub 显示的命令，或使用以下命令：

```bash
# 确保在 mediago-webui 目录中
cd mediago-webui

# 首次提交
git commit -m "Initial commit: MediaGo WebUI with GitHub Actions"

# 设置主分支为 main（GitHub Actions 配置使用 main）
git branch -M main

# 添加远程仓库（替换 YOUR_USERNAME 为你的 GitHub 用户名）
git remote add origin https://github.com/YOUR_USERNAME/mediago-webui.git

# 推送到 GitHub
git push -u origin main
```

### 第 3 步：触发自动编译

推送完成后：

1. 访问你的仓库页面
2. 点击 **Actions** 标签
3. 你会看到 "Build and Release" 工作流正在运行
4. 等待 3-5 分钟，三个平台将自动编译完成

编译产物会出现在 **Actions** → 选择具体运行 → **Artifacts** 下方

### 第 4 步：创建正式版本（可选）

如果要创建正式 Release：

```bash
# 打标签
git tag -a v1.0.0 -m "Release v1.0.0: Initial release"

# 推送标签
git push origin v1.0.0
```

推送标签后，GitHub Actions 会：
- 自动编译所有平台
- 自动创建 Release
- 自动上传编译好的二进制文件

然后你和其他人就可以从 **Releases** 页面直接下载 `.exe` 文件！

---

## 📦 编译产物说明

GitHub Actions 会自动生成：

| 文件名 | 平台 | 大小（约） |
|--------|------|------------|
| `mediago-webui-windows-amd64.exe` | Windows 64位 | ~15 MB |
| `mediago-webui-linux-amd64` | Linux 64位 | ~15 MB |
| `mediago-webui-darwin-amd64` | macOS 64位 | ~15 MB |

---

## 🔧 后续维护

### 更新代码
```bash
# 修改代码后
git add .
git commit -m "描述你的修改"
git push origin main
```

每次 push 都会自动触发编译。

### 发布新版本
```bash
git tag -a v1.0.1 -m "Release v1.0.1: 修复 bug"
git push origin v1.0.1
```

---

## ⚙️ 自定义配置

### 修改 README 中的用户名

编辑 `README.md`，替换所有 `YOUR_USERNAME` 为你的 GitHub 用户名：

```bash
# 可以用 VSCode 或任何文本编辑器全局替换
YOUR_USERNAME -> 你的GitHub用户名
```

### 调整编译选项

编辑 `.github/workflows/build.yml` 可以：
- 添加更多平台（如 ARM）
- 调整编译参数
- 添加测试步骤

---

## 🎉 完成！

完成上述步骤后：

1. ✅ 代码托管在 GitHub
2. ✅ 自动编译 Windows/Linux/macOS 版本
3. ✅ 其他人可以直接下载 `.exe` 运行
4. ✅ 无需任何 Go 环境

---

## 📌 快速命令汇总

```bash
# 在 mediago-webui 目录执行

# 1. 首次提交
git commit -m "Initial commit: MediaGo WebUI"
git branch -M main

# 2. 添加远程仓库（替换用户名）
git remote add origin https://github.com/YOUR_USERNAME/mediago-webui.git

# 3. 推送
git push -u origin main

# 4. （可选）发布版本
git tag -a v1.0.0 -m "Release v1.0.0"
git push origin v1.0.0
```

---

## ❓ 常见问题

**Q: 我没有 GitHub 账号？**  
A: 免费注册：https://github.com/signup

**Q: 推送失败，提示认证错误？**  
A: 需要设置 Personal Access Token：
1. GitHub → Settings → Developer settings → Personal access tokens
2. Generate new token (classic)
3. 勾选 `repo` 权限
4. 使用 token 作为密码

**Q: 编译失败了怎么办？**  
A: 查看 Actions 页面的错误日志，或联系我帮你排查。

**Q: 能不能不公开仓库？**  
A: 可以设为 Private，但 GitHub Actions 免费额度有限制。

---

## 🌟 推荐下一步

1. 在 README.md 中添加项目截图
2. 创建 Issues 模板
3. 添加 Contributing 指南
4. 设置 GitHub Pages 展示项目

有问题随时问我！🚀
