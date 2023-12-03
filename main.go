package main

import (
	 "fmt"
	"os"
)

func main() {
	fl := newFreeList()
	// pgnum := fl.getNextPage()
	// dal, _ := newDal("db.db", os.Getpagesize())
	// p := dal.allocateEmptyPage()
	// p.num = dal.getNextPage()

	// fmt.Println("pgnum", pgnum)
	// fmt.Println(dal, os.Getpagesize())

	dal, _ := newDal("db.db", os.Getpagesize())
	

	// create a new page
	p := dal.allocateEmptyPage()
	p.num = fl.getNextPage()
	fmt.Println("pgnum", dal)
	fmt.Println("pgnum", p.data)
	//p.data = []byte(string("Acacaca"))
	//fmt.Println("pgnum", p.data)
	copy(p.data, []byte(string("naman")))

	// commit it
	_ = dal.writePage(p)
}
