package htmx

import "fmt"

type pageInfo struct {
	PageSize int
	PageNum  int
	TotalCount   int64
	CurrentPage string
	HasPrevious bool
	HasNext   bool
	PrevPage int
	NextPage int
}

func calculatePageInfo(pageSize, pageNum int, count int64) pageInfo {
	var pi pageInfo
	pi.PageSize = pageSize
	pi.PageNum = pageNum
	pi.TotalCount = count
	currentStart := (pageNum - 1) * pageSize + 1
	currentEnd := pageNum * pageSize
	pi.CurrentPage = fmt.Sprintf("%d - %d", currentStart, currentEnd)
	if currentStart > 1 {
		pi.HasPrevious = true
		pi.PrevPage = pageNum - 1
	}
	if currentEnd < int(count) {
		pi.HasNext = true
		pi.NextPage = pageNum + 1
	}
	return pi
}