package main

const initialPage pgnum = 0

type freelist struct {
	maxPage       pgnum   // Holds the maximum page allocated. maxPage*PageSize = fileSize
	releasedPages []pgnum // Pages that were previouslly allocated but are now free
}

func newFreeList() *freelist {
	return &freelist{
		maxPage:       initialPage,
		releasedPages: []pgnum{},
	}
}

func (f *freelist) getNextPage() pgnum {
	if len(f.releasedPages) != 0 {
		freePageNum := f.releasedPages[0];
		f.releasedPages = f.releasedPages[1:]
		return freePageNum;
	}
	f.maxPage += 1 // if suppose no free pages are left
	return f.maxPage
}

func (fr *freelist) releasePage(page pgnum) {
	fr.releasedPages = append(fr.releasedPages, page)
}
