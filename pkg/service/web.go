package service

import (
	"ai-kcal-agent/pkg/appContext"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func WebHandler(rootDir, mount string) http.Handler {
	ctx := appContext.Get()

	// нормализуем mount: хотим чтобы был с завершающим слешем
	if !strings.HasSuffix(mount, "/") {
		mount += "/"
	}
	absRoot, _ := filepath.Abs(rootDir)
	indexPath := filepath.Join(absRoot, "index.html")

	// file server только для реальных файлов
	fileServer := http.StripPrefix(mount, http.FileServer(http.Dir(absRoot)))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// /web -> /web/
		if r.URL.Path == strings.TrimSuffix(mount, "/") {
			http.Redirect(w, r, mount, http.StatusMovedPermanently)
			return
		}
		if !strings.HasPrefix(r.URL.Path, mount) {
			http.NotFound(w, r)
			return
		}

		rel := strings.TrimPrefix(r.URL.Path, mount) // путь внутри dist
		full := filepath.Join(absRoot, filepath.FromSlash(rel))
		ctx.Logger.Printf("[static] %s -> %s", r.URL.Path, full)

		// если запрошен реальный файл — отдаём его через FileServer
		if fi, err := os.Stat(full); err == nil && !fi.IsDir() {
			fileServer.ServeHTTP(w, r)
			return
		}

		// Fallback для SPA: отдаём index.html НАПРЯМУЮ (без FileServer),
		// чтобы избежать лишних редиректов
		http.ServeFile(w, r, indexPath)
	})
}
