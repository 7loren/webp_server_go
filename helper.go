package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ChanErr(ccc chan int) {
	if ccc != nil {
		ccc <- 1
	}
}

func GetFileContentType(buffer []byte) string {
	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)
	return contentType
}

func FileCount(dir string) int {
	count := 0
	_ = filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				count += 1
			}
			return nil
		})
	return count
}

func ImageExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	log.Debugf("file %s exists!", filename)
	return !info.IsDir()
}

func GenWebpAbs(RawImagePath string, ExhaustPath string, ImgFilename string, reqURI string) (string, string) {
	// get file mod time
	STAT, err := os.Stat(RawImagePath)
	if err != nil {
		log.Error(err.Error())
	}
	ModifiedTime := STAT.ModTime().Unix()
	// webpFilename: abc.jpg.png -> abc.jpg.png1582558990.webp
	var WebpFilename = fmt.Sprintf("%s.%d.webp", ImgFilename, ModifiedTime)
	cwd, _ := os.Getwd()

	// /home/webp_server/exhaust/path/to/tsuki.jpg.1582558990.webp
	// Custom Exhaust: /path/to/exhaust/web_path/web_to/tsuki.jpg.1582558990.webp
	WebpAbsolutePath := path.Clean(path.Join(ExhaustPath, path.Dir(reqURI), WebpFilename))
	return cwd, WebpAbsolutePath
}

func goOrigin(UA string) bool {
	// for more information, please check test case
	if strings.Contains(UA, "Firefox") || strings.Contains(UA, "Chrome") {
		// Chrome or firefox on macOS Windows
	} else if strings.Contains(UA, "Android") || strings.Contains(UA, "Windows") || strings.Contains(UA, "Linux") {
		// on Android, Windows and Linux
	} else if strings.Contains(UA, "FxiOS") || strings.Contains(UA, "CriOS") {
		//firefox and Chrome on iOS
	} else {
		return true
	}
	if strings.Contains(UA, "rv:11.0") || strings.Contains(UA, "MSIE") {
		// MSIE
		return true
	}
	return false
}
