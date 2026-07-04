package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/spf13/cobra"

	"github.com/Sophomoresty/mediago/internal/cookie"
	"github.com/Sophomoresty/mediago/internal/download"
	"github.com/Sophomoresty/mediago/internal/extractor"
	"github.com/Sophomoresty/mediago/internal/util"

	_ "github.com/Sophomoresty/mediago/internal/extractor/ahu"
	_ "github.com/Sophomoresty/mediago/internal/extractor/aishangke"
	_ "github.com/Sophomoresty/mediago/internal/extractor/baijiayunxiao"
	_ "github.com/Sophomoresty/mediago/internal/extractor/bilibili"
	_ "github.com/Sophomoresty/mediago/internal/extractor/caixuetang"
	_ "github.com/Sophomoresty/mediago/internal/extractor/cctalk"
	_ "github.com/Sophomoresty/mediago/internal/extractor/cctv"
	_ "github.com/Sophomoresty/mediago/internal/extractor/chaoge"
	_ "github.com/Sophomoresty/mediago/internal/extractor/chaoxing"
	_ "github.com/Sophomoresty/mediago/internal/extractor/ckjr"
	_ "github.com/Sophomoresty/mediago/internal/extractor/classin"
	_ "github.com/Sophomoresty/mediago/internal/extractor/cnmooc"
	_ "github.com/Sophomoresty/mediago/internal/extractor/cto51"
	_ "github.com/Sophomoresty/mediago/internal/extractor/dingtalk"
	_ "github.com/Sophomoresty/mediago/internal/extractor/dongao"
	_ "github.com/Sophomoresty/mediago/internal/extractor/douyin"
	_ "github.com/Sophomoresty/mediago/internal/extractor/duanshu"
	_ "github.com/Sophomoresty/mediago/internal/extractor/enetedu"
	_ "github.com/Sophomoresty/mediago/internal/extractor/eoffcn"
	_ "github.com/Sophomoresty/mediago/internal/extractor/feishu"
	_ "github.com/Sophomoresty/mediago/internal/extractor/fenbi"
	_ "github.com/Sophomoresty/mediago/internal/extractor/gaodun"
	_ "github.com/Sophomoresty/mediago/internal/extractor/gaotu"
	_ "github.com/Sophomoresty/mediago/internal/extractor/gongxuanwang"
	_ "github.com/Sophomoresty/mediago/internal/extractor/haiyangknow"
	_ "github.com/Sophomoresty/mediago/internal/extractor/haozaixian"
	_ "github.com/Sophomoresty/mediago/internal/extractor/houda"
	_ "github.com/Sophomoresty/mediago/internal/extractor/houdu"
	_ "github.com/Sophomoresty/mediago/internal/extractor/hqwx"
	_ "github.com/Sophomoresty/mediago/internal/extractor/htknow"
	_ "github.com/Sophomoresty/mediago/internal/extractor/huatu"
	_ "github.com/Sophomoresty/mediago/internal/extractor/huke88"
	_ "github.com/Sophomoresty/mediago/internal/extractor/icourse163"
	_ "github.com/Sophomoresty/mediago/internal/extractor/icourses"
	_ "github.com/Sophomoresty/mediago/internal/extractor/icve"
	_ "github.com/Sophomoresty/mediago/internal/extractor/imooc"
	_ "github.com/Sophomoresty/mediago/internal/extractor/itbaizhan"
	_ "github.com/Sophomoresty/mediago/internal/extractor/jianshe99"
	_ "github.com/Sophomoresty/mediago/internal/extractor/jinbangshidai"
	_ "github.com/Sophomoresty/mediago/internal/extractor/jingtongxue"
	_ "github.com/Sophomoresty/mediago/internal/extractor/kaimingzhixue"
	_ "github.com/Sophomoresty/mediago/internal/extractor/kaoyanvip"
	_ "github.com/Sophomoresty/mediago/internal/extractor/keqq"
	_ "github.com/Sophomoresty/mediago/internal/extractor/koolearn"
	_ "github.com/Sophomoresty/mediago/internal/extractor/kuke"
	_ "github.com/Sophomoresty/mediago/internal/extractor/ledu"
	_ "github.com/Sophomoresty/mediago/internal/extractor/lexueyun"
	_ "github.com/Sophomoresty/mediago/internal/extractor/lizhiweike"
	_ "github.com/Sophomoresty/mediago/internal/extractor/luffycity"
	_ "github.com/Sophomoresty/mediago/internal/extractor/magedu"
	_ "github.com/Sophomoresty/mediago/internal/extractor/mashibing"
	_ "github.com/Sophomoresty/mediago/internal/extractor/mddclass"
	_ "github.com/Sophomoresty/mediago/internal/extractor/med66"
	_ "github.com/Sophomoresty/mediago/internal/extractor/meeting"
	_ "github.com/Sophomoresty/mediago/internal/extractor/minshi"
	_ "github.com/Sophomoresty/mediago/internal/extractor/nmkjxy"
	_ "github.com/Sophomoresty/mediago/internal/extractor/open163"
	_ "github.com/Sophomoresty/mediago/internal/extractor/orangevip"
	_ "github.com/Sophomoresty/mediago/internal/extractor/plaso"
	_ "github.com/Sophomoresty/mediago/internal/extractor/qihang"
	_ "github.com/Sophomoresty/mediago/internal/extractor/qlchat"
	_ "github.com/Sophomoresty/mediago/internal/extractor/renrenjiang"
	_ "github.com/Sophomoresty/mediago/internal/extractor/sanjieke"
	_ "github.com/Sophomoresty/mediago/internal/extractor/shanxiang"
	_ "github.com/Sophomoresty/mediago/internal/extractor/sier"
	_ "github.com/Sophomoresty/mediago/internal/extractor/sites"
	_ "github.com/Sophomoresty/mediago/internal/extractor/smartedu"
	_ "github.com/Sophomoresty/mediago/internal/extractor/speiyou"
	_ "github.com/Sophomoresty/mediago/internal/extractor/tmooc"
	_ "github.com/Sophomoresty/mediago/internal/extractor/unipus"
	_ "github.com/Sophomoresty/mediago/internal/extractor/wallstreets"
	_ "github.com/Sophomoresty/mediago/internal/extractor/wangxiao"
	_ "github.com/Sophomoresty/mediago/internal/extractor/wangxiao233"
	_ "github.com/Sophomoresty/mediago/internal/extractor/wendao"
	_ "github.com/Sophomoresty/mediago/internal/extractor/wowtiku"
	_ "github.com/Sophomoresty/mediago/internal/extractor/xiaoeapp"
	_ "github.com/Sophomoresty/mediago/internal/extractor/xiaoetech"
	_ "github.com/Sophomoresty/mediago/internal/extractor/xiwang"
	_ "github.com/Sophomoresty/mediago/internal/extractor/xsteach"
	_ "github.com/Sophomoresty/mediago/internal/extractor/xueersi"
	_ "github.com/Sophomoresty/mediago/internal/extractor/xuelang"
	_ "github.com/Sophomoresty/mediago/internal/extractor/xuetang"
	_ "github.com/Sophomoresty/mediago/internal/extractor/yangcong"
	_ "github.com/Sophomoresty/mediago/internal/extractor/yikaobang"
	_ "github.com/Sophomoresty/mediago/internal/extractor/yixiaoerguo"
	_ "github.com/Sophomoresty/mediago/internal/extractor/yizhiknow"
	_ "github.com/Sophomoresty/mediago/internal/extractor/youdao"
	_ "github.com/Sophomoresty/mediago/internal/extractor/youyuan"
	_ "github.com/Sophomoresty/mediago/internal/extractor/youzan"
	_ "github.com/Sophomoresty/mediago/internal/extractor/zhaozhao"
	_ "github.com/Sophomoresty/mediago/internal/extractor/zhengbao"
	_ "github.com/Sophomoresty/mediago/internal/extractor/zhihuishu"
	_ "github.com/Sophomoresty/mediago/internal/extractor/zlketang"
)

var version = "0.2.0"

var (
	formatSpec     string
	outputTemplate string
	cookieFile     string
	cookieBrowser  string
	listFormats    bool
	dumpJSON       bool
	simulate       bool
	writeInfoJSON  bool
	writeSubs      bool
	noOverwrites   bool
	concurrency    int
	listExtractors bool
	downloadAll    bool
	mergeOutputFmt string
	noProgress     bool
	proxy          string
)

func main() {
	// If no arguments, launch GUI
	if len(os.Args) == 1 {
		runGUI()
		return
	}

	// Otherwise, use CLI mode
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	rootCmd := &cobra.Command{
		Use:   "mediago [flags] URL [URL...]",
		Short: "Download media from 92 Chinese platforms",
		Long: `MediaGo - download videos from Chinese educational and media platforms.
Similar to yt-dlp but focused on Chinese internet platforms.`,
		RunE:              runMain,
		Args:              cobra.ArbitraryArgs,
		DisableAutoGenTag: true,
		SilenceUsage:      true,
		SilenceErrors:     true,
	}
	rootCmd.SetContext(ctx)
	rootCmd.Version = version
	rootCmd.SetVersionTemplate("mediago {{.Version}}\n")

	// Format selection (yt-dlp: -f, --format)
	rootCmd.Flags().StringVarP(&formatSpec, "format", "f", "best", "format selection (best/worst/1080p/720p/480p)")

	// Output (yt-dlp: -o, --output)
	rootCmd.Flags().StringVarP(&outputTemplate, "output", "o", "%(title)s.%(ext)s", "output filename template")

	// Cookie options (same as yt-dlp)
	rootCmd.Flags().StringVar(&cookieFile, "cookies", "", "Netscape cookie file path")
	rootCmd.Flags().StringVar(&cookieBrowser, "cookies-from-browser", "", "read cookies from browser (chrome/edge/firefox)")

	// Info/listing (yt-dlp: -F, -j, --write-info-json)
	rootCmd.Flags().BoolVarP(&listFormats, "list-formats", "F", false, "list available formats and exit")
	rootCmd.Flags().BoolVarP(&dumpJSON, "dump-json", "j", false, "dump info JSON to stdout and exit")
	rootCmd.Flags().BoolVar(&simulate, "simulate", false, "show extracted info without downloading")
	rootCmd.Flags().BoolVar(&writeInfoJSON, "write-info-json", false, "write .info.json file alongside download")
	rootCmd.Flags().BoolVar(&writeSubs, "write-subs", false, "write subtitle files alongside download")

	// Download options
	rootCmd.Flags().BoolVar(&noOverwrites, "no-overwrites", false, "do not overwrite existing files")
	rootCmd.Flags().IntVarP(&concurrency, "concurrent-fragments", "N", 10, "number of concurrent fragment downloads")
	rootCmd.Flags().BoolVar(&downloadAll, "yes-playlist", false, "download all items in a playlist/course")
	rootCmd.Flags().StringVar(&mergeOutputFmt, "merge-output-format", "mp4", "merge output container (mp4/mkv/webm)")
	rootCmd.Flags().BoolVar(&noProgress, "no-progress", false, "suppress progress bar")
	rootCmd.Flags().StringVar(&proxy, "proxy", "", "HTTP/SOCKS proxy URL")

	// Extractor listing (yt-dlp: --list-extractors)
	rootCmd.Flags().BoolVar(&listExtractors, "list-extractors", false, "list all supported sites and exit")

	// Version
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("mediago %s\n", version)
		},
	})

	if err := rootCmd.Execute(); err != nil {
		if errors.Is(err, context.Canceled) {
			interruptedf()
			os.Exit(130)
		}
		errorf("%v", err)
		os.Exit(1)
	}
}

// Global GUI progress state
var (
	guiProgressCallback func(string)
	currentVideoTitle   string
	currentVideoTotal   int64
	progressMutex       sync.Mutex
)

func runGUI() {
	myApp := app.New()
	myWindow := myApp.NewWindow("MediaGo 下载器")

	// Default download directory
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	defaultDownloadDir := filepath.Join(exeDir, "downloads")

	// URL input
	urlEntry := widget.NewEntry()
	urlEntry.SetPlaceHolder("粘贴视频链接...")

	// Format selection
	formatSelect := widget.NewSelect([]string{"自动 (best)", "1080p", "720p", "480p"}, nil)
	formatSelect.SetSelected("自动 (best)")

	// Download directory
	downloadDirEntry := widget.NewEntry()
	downloadDirEntry.SetText(defaultDownloadDir)
	downloadDirEntry.SetPlaceHolder("下载保存路径")

	// Browse button for directory
	browseBtn := widget.NewButton("浏览...", func() {
		// TODO: Add directory picker when available in Fyne
		// For now, users can manually edit the path
	})

	// Cookies input
	cookiesEntry := widget.NewEntry()
	cookiesEntry.SetPlaceHolder("Cookies 文件路径（可选）")

	// Proxy input
	proxyEntry := widget.NewEntry()
	proxyEntry.SetPlaceHolder("代理地址（可选）")

	// Output display
	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("下载日志将显示在这里...")
	output.Wrapping = fyne.TextWrapWord

	// Scroll container for output
	outputScroll := container.NewScroll(output)
	outputScroll.SetMinSize(fyne.NewSize(560, 300))

	// Download log accumulator
	var downloadLog strings.Builder
	var logMutex sync.Mutex

	// Helper function to append log (thread-safe)
	appendLog := func(msg string) {
		logMutex.Lock()
		defer logMutex.Unlock()
		downloadLog.WriteString(msg)
		output.SetText(downloadLog.String())
		outputScroll.ScrollToBottom()
	}

	// Set global callback
	guiProgressCallback = appendLog

	// Download button (declare variable first)
	var downloadBtn *widget.Button
	downloadBtn = widget.NewButton("开始下载", func() {
		url := urlEntry.Text
		if url == "" {
			appendLog("[错误] 请输入视频链接\n\n")
			return
		}

		downloadDir := downloadDirEntry.Text
		if downloadDir == "" {
			downloadDir = defaultDownloadDir
		}

		// Create download directory if not exists
		if err := os.MkdirAll(downloadDir, 0755); err != nil {
			appendLog(fmt.Sprintf("[错误] 无法创建下载目录: %v\n\n", err))
			return
		}

		// Get format
		format := "best"
		switch formatSelect.Selected {
		case "1080p":
			format = "1080p"
		case "720p":
			format = "720p"
		case "480p":
			format = "480p"
		}

		// Set global options
		formatSpec = format
		cookieFile = cookiesEntry.Text
		proxy = proxyEntry.Text
		outputTemplate = filepath.Join(downloadDir, "%(title)s.%(ext)s")
		noProgress = true // Disable console progress bar

		// Add separator for new download
		appendLog("════════════════════════════════════════\n")
		appendLog(fmt.Sprintf("[链接] %s\n", url))
		downloadBtn.Disable()

		// Download in background
		go func() {
			defer func() {
				downloadBtn.Enable()
				guiProgressCallback = nil
			}()

			startTime := time.Now()
			appendLog(fmt.Sprintf("[开始] %s\n", startTime.Format("15:04:05")))

			err := processURL(context.Background(), url)

			endTime := time.Now()
			duration := endTime.Sub(startTime)

			appendLog(fmt.Sprintf("[完成] %s (耗时: %s)\n", endTime.Format("15:04:05"), duration.Round(time.Second)))

			if err != nil {
				appendLog(fmt.Sprintf("[失败] %v\n", err))
				appendLog("[提示] 请检查:\n")
				appendLog("  1. 网络连接是否正常\n")
				appendLog("  2. 视频链接是否有效\n")
				appendLog("  3. 是否需要提供 Cookies\n\n")
			} else {
				appendLog("[成功] 下载完成\n")
				appendLog(fmt.Sprintf("[位置] %s\n\n", downloadDir))
			}
		}()
	})

	// Layout
	content := container.NewVBox(
		widget.NewLabel("MediaGo 视频下载器"),
		widget.NewSeparator(),
		widget.NewLabel("视频链接:"),
		urlEntry,
		widget.NewLabel("画质选择:"),
		formatSelect,
		widget.NewLabel("保存路径:"),
		container.NewBorder(nil, nil, nil, browseBtn, downloadDirEntry),
		widget.NewLabel("Cookies 文件:"),
		cookiesEntry,
		widget.NewLabel("代理:"),
		proxyEntry,
		downloadBtn,
		widget.NewSeparator(),
		widget.NewLabel("下载日志:"),
		outputScroll,
	)

	myWindow.SetContent(container.NewPadded(content))
	myWindow.Resize(fyne.NewSize(640, 650))
	myWindow.ShowAndRun()
}

func runMain(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	if listExtractors {
		return printExtractors()
	}

	if len(args) == 0 {
		return cmd.Help()
	}

	if proxy != "" {
		if err := util.SetDefaultProxy(proxy); err != nil {
			return fmt.Errorf("invalid --proxy value: %w", err)
		}
	}

	failures := 0
	for _, url := range args {
		if err := processURL(ctx, url); err != nil {
			if errors.Is(err, context.Canceled) {
				return err
			}
			errorf("%v", err)
			failures++
		}
	}
	if failures > 0 {
		return fmt.Errorf("%d of %d URLs failed", failures, len(args))
	}
	return nil
}

func processURL(ctx context.Context, url string) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	ext, site, err := extractor.MatchWithSite(url)
	if err != nil {
		return fmt.Errorf("unsupported URL: %s\nUse --list-extractors to see supported sites.", url)
	}
	infof("Extracting: %s %s", site.Name, url)

	store := cookie.NewStore()
	if cookieFile != "" {
		if err := store.LoadFromFile(cookieFile); err != nil {
			return fmt.Errorf("failed to load cookies: %w", err)
		}
	}
	if cookieBrowser != "" {
		if err := store.LoadFromBrowser(cookieBrowser); err != nil {
			return fmt.Errorf("failed to read browser cookies: %w", err)
		}
	}

	opts := &extractor.ExtractOpts{
		Cookies:  store.Jar(),
		Quality:  formatSpec,
		ListOnly: listFormats,
	}

	info, err := ext.Extract(url, opts)
	if err != nil {
		return fmt.Errorf("[%s] %w", url, err)
	}

	if err := ctx.Err(); err != nil {
		return err
	}

	if dumpJSON {
		return printJSON(info)
	}

	if info.IsPlaylist() {
		infof("Playlist: %s (%d items)", info.Title, len(info.Entries))
		if !downloadAll {
			warnf("Downloading only the first item. Use --yes-playlist to download all.")
			if len(info.Entries) > 0 && info.Entries[0] != nil {
				infof("%s", info.Entries[0].Title)
				return downloadEntry(ctx, 0, 1, info.Entries[0])
			}
			return fmt.Errorf("playlist is empty")
		}
		if listFormats {
			warnf("use a single-item URL with -F to inspect formats")
			return nil
		}
		if simulate {
			for i, entry := range info.Entries {
				if entry == nil {
					continue
				}
				if err := printSimulation(entry, i+1, len(info.Entries)); err != nil {
					return err
				}
			}
			return nil
		}
		entryFailures := 0
		for i, entry := range info.Entries {
			if entry == nil {
				continue
			}
			if err := downloadEntry(ctx, i, len(info.Entries), entry); err != nil {
				if errors.Is(err, context.Canceled) {
					return err
				}
				errorf("[%d/%d %s]: %v", i+1, len(info.Entries), firstNonEmpty(entry.Title, fmt.Sprintf("item-%d", i+1)), err)
				entryFailures++
			}
		}
		if entryFailures > 0 {
			return fmt.Errorf("%d of %d playlist items failed", entryFailures, len(info.Entries))
		}
		return nil
	}

	infof("%s", info.Title)
	if simulate {
		return printSimulation(info, 0, 0)
	}
	return downloadOne(ctx, info)
}

func downloadEntry(ctx context.Context, itemIndex, totalItems int, info *extractor.MediaInfo) error {
	downloadf("%s", downloadItemMessage(itemIndex+1, totalItems, firstNonEmpty(info.Title, fmt.Sprintf("item-%d", itemIndex+1))))
	return downloadOneFn(ctx, info)
}

var downloadOneFn = downloadOne

func downloadOne(ctx context.Context, info *extractor.MediaInfo) error {
	if listFormats {
		return printFormats(info)
	}

	_, stream := download.SelectBestStream(info.Streams, formatSpec)
	if len(stream.URLs) == 0 && stream.Format == "" {
		return fmt.Errorf("no formats available: %s", info.Title)
	}

	outFilename := applyTemplate(outputTemplate, info, stream)

	// Set current video info for GUI
	progressMutex.Lock()
	currentVideoTitle = info.Title
	currentVideoTotal = 0
	progressMutex.Unlock()

	// Create progress callback for GUI
	var progressCallback func(current, total int64, speed float64)
	if guiProgressCallback != nil {
		progressCallback = func(current, total int64, speed float64) {
			percent := float64(current) / float64(total) * 100
			currentMB := float64(current) / 1024 / 1024
			totalMB := float64(total) / 1024 / 1024

			// Calculate remaining time
			remaining := ""
			if speed > 0 {
				remainingBytes := total - current
				remainingSec := float64(remainingBytes) / (speed * 1024 * 1024)
				remaining = fmt.Sprintf(" | 剩余: %ds", int(remainingSec))
			}

			// Build progress bar
			barWidth := 30
			filled := int(percent / 100 * float64(barWidth))
			bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)

			msg := fmt.Sprintf("[进度] %.1f%% (%.1f MB / %.1f MB)\n", percent, currentMB, totalMB)
			msg += fmt.Sprintf("[速度] %.2f MB/s%s\n", speed, remaining)
			msg += fmt.Sprintf("[%s]\n", bar)

			guiProgressCallback(msg)
		}
	}

	engine := download.New(download.Opts{
		Concurrency:       concurrency,
		OutputDir:         outputDirFromTemplate(outFilename),
		Overwrite:         !noOverwrites,
		Retries:           3,
		NoProgress:        noProgress,
		Proxy:             proxy,
		Context:           ctx,
		MergeOutputFormat: mergeOutputFmt,
		ProgressCallback:  progressCallback,
	})

	info.Title = baseFromTemplate(outFilename)

	if strings.EqualFold(stream.Format, "dash") && engine.HasFFmpeg() {
		mergerf("Merging formats into %s", outFilename)
	}
	outPath, err := engine.Download(info, stream)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}

	downloadf("100%% of %s", sizeStringForPath(outPath, stream.Size))
	if writeInfoJSON {
		writeInfoJSONFile(outPath, info)
	}
	if writeSubs {
		if subs, err := engine.DownloadSubtitles(info, outPath); err != nil {
			return fmt.Errorf("download subtitles: %w", err)
		} else {
			for _, sub := range subs {
				subtitlef("%s", sub)
			}
		}
	}
	return nil
}

func printJSON(info *extractor.MediaInfo) error {
	data, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	return nil
}

func printExtractors() error {
	sites := extractor.ListSites()
	for _, s := range sites {
		auth := ""
		if s.NeedAuth {
			auth = " (auth)"
		}
		fmt.Printf("%s: %s%s\n", s.Name, s.URL, auth)
	}
	fmt.Printf("\n%d extractors\n", len(sites))
	return nil
}

func applyTemplate(tmpl string, info *extractor.MediaInfo, stream extractor.Stream) string {
	ext := stream.Format
	if ext == "m3u8" || ext == "dash" {
		ext = mergeOutputFmt
	}
	if ext == "" {
		ext = "mp4"
	}

	r := strings.NewReplacer(
		"%(title)s", info.Title,
		"%(ext)s", ext,
		"%(site)s", info.Site,
		"%(artist)s", info.Artist,
		"%(quality)s", stream.Quality,
	)
	return r.Replace(tmpl)
}

func outputDirFromTemplate(filename string) string {
	dir := filepath.Dir(filename)
	if dir == "" || dir == "." {
		return "."
	}
	return dir
}

func baseFromTemplate(filename string) string {
	filename = filepath.Base(filename)
	if idx := strings.LastIndex(filename, "."); idx > 0 {
		filename = filename[:idx]
	}
	return filename
}

func writeInfoJSONFile(videoPath string, info *extractor.MediaInfo) {
	jsonPath := videoPath + ".info.json"
	data, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return
	}
	os.WriteFile(jsonPath, data, 0o644)
}
