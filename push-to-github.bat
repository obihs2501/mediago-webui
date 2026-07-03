@echo off
echo ================================
echo MediaGo WebUI - Push to GitHub
echo ================================
echo.

echo [1/4] Committing code...
git commit -m "Initial commit: MediaGo WebUI with GitHub Actions auto-build"
if errorlevel 1 (
    echo Warning: Commit may have failed, continuing...
)
echo.

echo [2/4] Setting main branch...
git branch -M main
echo.

echo [3/4] Adding remote repository...
git remote add origin https://github.com/obihs2501/mediago-webui.git
if errorlevel 1 (
    echo Note: Remote may already exist, trying to update...
    git remote set-url origin https://github.com/obihs2501/mediago-webui.git
)
echo.

echo [4/4] Pushing to GitHub...
echo You may need to enter GitHub username and Personal Access Token
echo.
git push -u origin main
if errorlevel 1 (
    echo.
    echo Push failed!
    echo.
    echo Possible reasons:
    echo 1. Need GitHub Personal Access Token as password
    echo 2. Network connection issue
    echo 3. Repository permission issue
    echo.
    echo Solution:
    echo - Create Personal Access Token at:
    echo   GitHub ^> Settings ^> Developer settings ^> Personal access tokens
    echo - Use Token as password (not GitHub password)
    echo.
    pause
    exit /b 1
)

echo.
echo ================================
echo Successfully pushed to GitHub!
echo ================================
echo.
echo Next steps:
echo 1. Visit https://github.com/obihs2501/mediago-webui/actions
echo 2. Watch auto-build progress (about 3-5 minutes)
echo 3. Download compiled files from Actions page
echo.
echo Or create a Release:
echo    git tag -a v1.0.0 -m "Release v1.0.0"
echo    git push origin v1.0.0
echo.
echo Then download .exe directly from Releases page!
echo.
pause
