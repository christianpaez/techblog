package jekyllmdpathhelper

import (
	"fmt"
	"os"
	"strings"
)

const assetBasePath = "{{ site.url }}{{ site.baseurl }}/assets/img"

type JekyllMdPathHelper struct {
	filePath          string
	slug              string
	contentUrls       []string
	contentExtensions []string
}

func NewJekyllMdPathHelper(filePath string, slug string, contentUrls []string, contentExtensions []string) *JekyllMdPathHelper {
	return &JekyllMdPathHelper{
		filePath:          filePath,
		slug:              slug,
		contentUrls:       contentUrls,
		contentExtensions: contentExtensions,
	}
}

func (h *JekyllMdPathHelper) Normalize() error {
	content, err := os.ReadFile(h.filePath)
	if err != nil {
		return err
	}

	updated := string(content)
	for i, url := range h.contentUrls {
		if url == "" {
			continue // no main_image -> downloader puts empty string at index 0
		}
		localPath := fmt.Sprintf("%s/%s-0%d.%s", assetBasePath, h.slug, i, h.contentExtensions[i])
		updated = strings.ReplaceAll(updated, url, localPath)
	}

	if err := os.WriteFile(h.filePath, []byte(updated), 0644); err != nil {
		return err
	}
	return nil
}
