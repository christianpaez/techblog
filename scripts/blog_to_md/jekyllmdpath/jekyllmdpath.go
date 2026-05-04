package jekyllmdpathhelper

// data definitions
// string need config md path where files are located
//
// string|template need header string to insert when content does not have header image
//
// string need a blog with md content
//
// string[] for each blog need list of image paths for lookup and replace

type JekyllMdPathHelper struct {
	contentUrls []string
}

func NewJekyllMdPathHelper(contentUrls []string) *JekyllMdPathHelper {
	return &JekyllMdPathHelper{
		contentUrls: contentUrls,
	}
}

func (jekyllMdPathHelper *JekyllMdPathHelper) Normalize() {
	// stub this reads from md folder and calls helper
}

func insertHeader() {} // stub takes image path and inserts header withmain image url

func replaceImagePath() {} // stub replace path with new path

func replacePaths() {} // stub helper that loops over content and calls replace on main image and content images
