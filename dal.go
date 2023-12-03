package main

import (
	"fmt"
	"os"
)

type pgnum int

type dal struct {
	file     *os.File
	pageSize int
}

type page struct {
	num  pgnum
	data []byte
}

func newDal(path string, pageSize int) (*dal, error) {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Failed to open file")
		return nil, err
	}
	dal := &dal{
		file,
		pageSize,
	}
	return dal, nil
}

func (d *dal) close() error {
	if d.file != nil {
		err := d.file.Close()
		if err != nil {
			return fmt.Errorf("could not close file: %s", err)
		}
		d.file = nil
	}
	return nil
}

func (d *dal) allocateEmptyPage() *page {
	return &page{
		data: make([]byte, 0),
	}
}

func (d *dal) readPage() (*page, error) {

	p := d.allocateEmptyPage()

	//calculate the offset position pagenum * size eg 3*4096

	offset := d.pageSize * int(p.num)
	_, err := d.file.ReadAt(p.data, int64(offset))
	if err != nil {
		return nil, err
	}
	return p, nil

}

func (d *dal) writePage(p *page) error {
	offset := int64(p.num) * int64(d.pageSize)
	_, err := d.file.WriteAt(p.data, offset)
	return err
}
