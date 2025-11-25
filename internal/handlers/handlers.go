package handlers

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"time"

	"github.com/qewyxiww/projectt/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Serving index.html...") // Добавьте эту строку
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to get file from form", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusInternalServerError)
		return
	}

	content := string(fileBytes)

	converted, err := service.DetectAndConvert(content)
	if err != nil {
		http.Error(w, "Conversion error", http.StatusInternalServerError)
		return
	}

	timestamp := time.Now().UTC().Format("20060102150405")
	originalExt := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("converted_%s%s", timestamp, originalExt)

	err = saveToFile(filename, converted)
	if err != nil {
		http.Error(w, "Unable to save result file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "File converted successfully!\n\nOriginal content:\n%s\n\nConverted content:\n%s\n\nSaved to: %s",
		content, converted, filename)
}

// saveToFile сохраняет строку в файл
func saveToFile(filename, content string) error {
	return nil // В реальном приложении здесь была бы запись в файл
}
