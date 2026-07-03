@echo off
chcp 65001 >nul
echo ================================
echo MediaGo WebUI - 推送到 GitHub
echo ================================
echo.

echo [1/4] 提交代码到本地仓库...
git commit -m "Initial commit: MediaGo WebUI with GitHub Actions auto-build"
if errorlevel 1 (
    echo 警告: 提交可能失败，但继续执行...
)
echo.

echo [2/4] 设置主分支为 main...
git branch -M main
echo.

echo [3/4] 添加远程仓库...
git remote add origin https://github.com/obihs2501/mediago-webui.git
if errorlevel 1 (
    echo 注意: 远程仓库可能已存在，尝试更新...
    git remote set-url origin https://github.com/obihs2501/mediago-webui.git
)
echo.

echo [4/4] 推送到 GitHub...
echo 提示: 可能需要输入 GitHub 用户名和密码/Token
echo.
git push -u origin main
if errorlevel 1 (
    echo.
    echo ❌ 推送失败！
    echo.
    echo 可能的原因:
    echo 1. 需要 GitHub Personal Access Token 作为密码
    echo 2. 网络连接问题
    echo 3. 仓库权限问题
    echo.
    echo 解决方法:
    echo - 创建 Personal Access Token:
    echo   GitHub ^> Settings ^> Developer settings ^> Personal access tokens
    echo - 使用 Token 作为密码（不是 GitHub 密码）
    echo.
    pause
    exit /b 1
)

echo.
echo ================================
echo ✅ 成功推送到 GitHub！
echo ================================
echo.
echo 接下来:
echo 1. 访问 https://github.com/obihs2501/mediago-webui/actions
echo 2. 查看自动编译进度（约 3-5 分钟）
echo 3. 编译完成后，在 Actions 页面下载编译产物
echo.
echo 或者创建正式 Release:
echo    git tag -a v1.0.0 -m "Release v1.0.0"
echo    git push origin v1.0.0
echo.
echo 然后在 Releases 页面直接下载 .exe 文件！
echo.
pause
