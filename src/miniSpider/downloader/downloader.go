package downloader

import (
	"miniSpider/page"
	"miniSpider/request"
)

type Downloader interface {
	Download(req *request.Request) (*page.Page, error)
}
