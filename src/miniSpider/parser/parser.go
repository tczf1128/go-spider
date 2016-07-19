package parser

import (
	"miniSpider/page"
)

type Parser interface {
	Parse(page *page.Page) error // 解析页面内容
}
