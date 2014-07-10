package utils

func PageInfo(page, pre, count int) (nopre, nonext bool, prepage, nextpage, start int) {
	nopre = false
	if page == 0 || page == 1 {
		nopre = true
		page = 1
	} else {
		page = page
	}

	prepage = page - 1
	nextpage = page + 1
	start = (page - 1) * pre
	nonext = false
	if page*pre >= count {
		nonext = true
	}
	return
}
