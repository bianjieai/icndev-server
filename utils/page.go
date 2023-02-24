package utils

const defaultSize = 10

func ParsePage(page int, size int, total bool) (skip, limit int, useTotal bool) {
	if page == 0 && size == 0 {
		size = defaultSize
	}

	skip = (page - 1) * size
	limit = size
	useTotal = total
	return
}
