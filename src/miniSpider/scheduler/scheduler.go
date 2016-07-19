package scheduler

import (
	"miniSpider/request"
)

type Scheduler interface {
	Push(req *request.Request)
	Pop() (*request.Request, error)
	Count() int
	Empty() bool
}
