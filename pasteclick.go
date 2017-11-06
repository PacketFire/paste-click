package main

import (
	"errors"
	"fmt"
	"github.com/rakyll/magicmime"
	"io/ioutil"
	"log"
	"log/syslog"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var Config = map[string]string{
	"savePath": "/www/paste.click/",
}

type MimeMap struct {
	sync.RWMutex
	m map[string]string
}

func (this *MimeMap) New() {
	this.Lock()
	defer this.Unlock()

	this.m = map[string]string{
		"audio/midi":                                                                ".mid",
		"audio/mp4":                                                                 ".aac",
		"audio/mpeg":                                                                ".mp3",
		"audio/ogg":                                                                 ".oga",
		"audio/x-realaudio":                                                         ".ra",
		"audio/x-wav":                                                               ".wav",
		"image/bmp":                                                                 ".bmp",
		"image/gif":                                                                 ".gif",
		"image/jpeg":                                                                ".jpeg",
		"image/png":                                                                 ".png",
		"image/svg+xml":                                                             ".svg",
		"image/tiff":                                                                ".tif",
		"image/vnd.wap.wbmp":                                                        ".wbmp",
		"image/webp":                                                                ".webp",
		"image/x-icon":                                                              ".ico",
		"image/x-jng":                                                               ".jng",
		"application/javascript":                                                    ".js",
		"application/json":                                                          ".json",
		"application/x-web-app-manifest+json":                                       ".webapp",
		"text/cache-manifest":                                                       ".manifest",
		"application/msword":                                                        ".doc",
		"application/vnd.ms-excel":                                                  ".xls",
		"application/vnd.ms-powerpoint":                                             ".ppt",
		"application/vnd.openxmlformats-officedocument.wordprocessingml.document":   ".docx",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         ".xlsx",
		"application/vnd.openxmlformats-officedocument.presentationml.presentation": ".pptx",
		"video/3gpp":                           ".3gpp",
		"video/mp4":                            ".mp4",
		"video/mpeg":                           ".mpeg",
		"video/ogg":                            ".ogv",
		"video/quicktime":                      ".mov",
		"video/webm":                           ".webm",
		"video/x-flv":                          ".flv",
		"video/x-mng":                          ".mng",
		"video/x-ms-asf":                       ".asx",
		"video/x-ms-wmv":                       ".wmv",
		"video/x-msvideo":                      ".avi",
		"application/xml":                      ".atom",
		"application/font-woff":                ".woff",
		"application/font-woff2":               ".woff2",
		"application/vnd.ms-fontobject":        ".eot",
		"application/x-font-ttf":               ".ttc",
		"font/opentype":                        ".otf",
		"application/java-archive":             ".jar",
		"application/mac-binhex40":             ".hqx",
		"application/pdf":                      ".pdf",
		"application/postscript":               ".ps",
		"application/rtf":                      ".rtf",
		"application/vnd.wap.wmlc":             ".wmlc",
		"application/xhtml+xml":                ".xhtml",
		"application/vnd.google-earth.kml+xml": ".kml",
		"application/vnd.google-earth.kmz":     ".kmz",
		"application/x-7z-compressed":          ".7z",
		"application/x-chrome-extension":       ".crx",
		"application/x-opera-extension":        ".oex",
		"application/x-xpinstall":              ".xpi",
		"application/x-cocoa":                  ".cco",
		"application/x-java-archive-diff":      ".jardiff",
		"application/x-java-jnlp-file":         ".jnlp",
		"application/x-makeself":               ".run",
		"application/x-perl":                   ".pl",
		"application/x-pilot":                  ".prc",
		"application/x-rar-compressed":         ".rar",
		"application/x-redhat-package-manager": ".rpm",
		"application/x-sea":                    ".sea",
		"application/x-shockwave-flash":        ".swf",
		"application/x-stuffit":                ".sit",
		"application/x-tcl":                    ".tcl",
		"application/x-x509-ca-cert":           ".der",
		"application/x-bittorrent":             ".torrent",
		"application/zip":                      ".zip",
		"application/octet-stream":             "",
		"text/css":                             "",
		"text/html":                            "",
		"text/mathml":                          "",
		"text/plain":                           ".txt",
		"text/vnd.sun.j2me.app-descriptor":     ".jad",
		"text/vnd.wap.wml":                     ".wml",
		"text/vtt":                             ".vtt",
		"text/x-component":                     ".htc",
		"text/x-vcard":                         ".vcf",
	}
}

func (this MimeMap) Extension(t string) (string, error) {
	this.RLock()
	defer this.RUnlock()

	if s, p := this.m[t]; p == true {
		return s, nil
	}

	return "", errors.New("Unable to find extension.")
}

func (this *MimeMap) NewExtension(t string, e string) error {
	this.Lock()
	defer this.Unlock()

	if t == "" {
		return errors.New("Mime Type cannot be empty.")
	}

	if _, p := this.m[t]; p == true {
		return errors.New("Mime Type already exists.")
	}

	this.m[t] = e

	return nil
}
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		write := savePost(w, r)
		w.Write(write)
		return
	}
	http.Error(w, "Not a POST request", 405)
	return
}
func getMimeString(data []byte) string {
	if err := magicmime.Open(magicmime.MAGIC_MIME_TYPE | magicmime.MAGIC_ERROR); err != nil {
		log.Fatal(err)
	}
	defer magicmime.Close()

	mimetype, err := magicmime.TypeByBuffer(data)
	if err != nil {
		log.Fatalf("error occured during type lookup: %v", err)
	}
	return mimetype
}

func savePost(w http.ResponseWriter, post *http.Request) []byte {
	//hash it and write to disk;
	code := string(randSeq(6, Config["savePath"]))
	rawVal, err := ioutil.ReadAll(post.Body)
	if err != nil {
		log.Print(err)
	}
	if len(rawVal) < 1 {
		http.Error(w, "Content body cannot be empty", 400)
		log.Fatalf("[ %v ] Filesize is 0, not writing empty file", post.Header.Get("X-Real-IP"))
	}
	mMap := new(MimeMap)
	mMap.New()
	mimeType := getMimeString(rawVal)
	ext, err := mMap.Extension(mimeType)
	filePath := strings.Join([]string{Config["savePath"], code, ext}, "")
	err = ioutil.WriteFile(filePath, rawVal, 0644)
	if err != nil {
		log.Print(err)
	}

	err = saveMeta(code, filePath, mimeType, ext)
	if err != nil {
		log.Print(err)
	}

	log.Printf("[ %v ] New File: %v", post.Header.Get("X-Real-IP"), filePath)
	fileUrl := strings.Join([]string{post.Header.Get("X-Scheme"), "://", post.Host, post.RequestURI, code, "\n"}, "")
	return []byte(fileUrl)
}

func randSeq(n int, path string) []rune {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	if _, err := os.Stat(strings.Join([]string{path, "_", string(b)}, "")); err == nil {
		b = randSeq(n, path)
	}
	return b
}

// saveMeta attempts to write the FileMetadata to disk.
func saveMeta(shortCode, fp, mimeType, ext string) error {
	var err error
	mp := strings.Join([]string{Config["savePath"], "_", shortCode}, "")

	stat, err := os.Stat(fp)
	if err != nil {
		return err
	}

	fm := NewFileMetadata(stat.Size(), mimeType, "", fp)

	json, err := fm.JSON()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(mp, json, 0644)

	return err
}

// bindAddr returns the bind address for the server. The bind address can be
// set by specifying the PASTE_CLICK_BINDADDR env variable.
func bindAddr() string {
	if addr := os.Getenv("PASTE_CLICK_BINDADDR"); addr != "" {
		return addr
	}

	return "127.0.0.1"
}

// bindPort returns the port to bind to for the server. The bind port can be
// set by specifying the PASTE_CLICK_BINDPORT env variable.
func bindPort() string {
	if port := os.Getenv("PASTE_CLICK_PORT"); port != "" {
		return port
	}

	return "8001"
}

func main() {
	logger, err := syslog.New(syslog.LOG_DAEMON, "pasteclickd: ")
	if err != nil {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(logger)
	}

	http.HandleFunc("/", handler)
	err = http.ListenAndServe(fmt.Sprintf("%s:%s", bindAddr(), bindPort()), nil)
	if err != nil {
		log.Print(err)
	}
}
