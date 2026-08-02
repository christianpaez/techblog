package jekyllmdpathhelper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const assetBasePath = "{{ site.url }}{{ site.baseurl }}/assets/img"

type JekyllMdPathHelper struct {
	filePath    string
	contentUrls []string
}

func NewJekyllMdPathHelper(filePath string, contentUrls []string) *JekyllMdPathHelper {
	return &JekyllMdPathHelper{
		filePath:    filePath,
		contentUrls: contentUrls,
	}
}

func (h *JekyllMdPathHelper) Normalize() error {
	content, err := os.ReadFile(h.filePath)
	if err != nil {
		return err
	}

	slug := slugFromFileName(h.filePath)

	updated := string(content)
	for i, url := range h.contentUrls {
		if url == "" {
			continue // no main_image -> downloader puts empty string at index 0
		}
		localPath := fmt.Sprintf("%s/%s-0%d.png", assetBasePath, slug, i)
		updated = strings.ReplaceAll(updated, url, localPath)
	}

	if err := os.WriteFile(h.filePath, []byte(updated), 0644); err != nil {
		return err
	}
	return nil
}

// slugFromFileName turns "2026-03-08-the-case-for-boring-tools.md" into "the-case-for-boring-tools"
func slugFromFileName(filePath string) string {
	base := strings.TrimSuffix(filepath.Base(filePath), ".md")
	parts := strings.SplitN(base, "-", 4)
	if len(parts) == 4 {
		return parts[3]
	}
	return base
}
