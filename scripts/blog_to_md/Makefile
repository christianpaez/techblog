SITE_ROOT := ../..

POSTS_DIR := $(SITE_ROOT)/_posts
IMG_DIR   := $(SITE_ROOT)/assets/img

MD_FILES  := $(wildcard tmp/*/mdFiles/*.md)
IMG_FILES := $(wildcard tmp/images/*)

.PHONY: run publish

clean_dev_to_html:
	go run htmlelementremover/main.go goquery_test.html

run:
	go run .

publish:
	find tmp -path '*/mdFiles/*.md' -exec cp -v {} $(POSTS_DIR)/ \;
	find tmp/images -type f -exec cp -v {} $(IMG_DIR)/ \;

clean:
	rm -rf tmp
