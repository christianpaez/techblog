SITE_ROOT := ../..

POSTS_DIR := $(SITE_ROOT)/_posts
IMG_DIR   := $(SITE_ROOT)/assets/img

MD_FILES  := $(wildcard tmp/*/mdFiles/*.md)
IMG_FILES := $(wildcard tmp/images/*)

.PHONY: run publish clean

run:
	go run .

publish:
	cp -v $(MD_FILES) $(POSTS_DIR)/
	cp -v $(IMG_FILES) $(IMG_DIR)/

clean:
	rm -rf tmp
