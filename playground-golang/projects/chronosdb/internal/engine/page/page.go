/*

	File:	page.go
	Author:	Lennard Claproth
	Date: 	20250514

	The page is responsible for storing and managing the
	data that is contained in the record. It is the so called
	Domain model of the page. It functions as an aggregate root
	where it also manages and creates the headers, slots and
	record data.

*/

package page

import (
	"errors"
)

type Page struct {
	Header     Header
	SlotDir    []Slot
	DataRegion []byte
}

const PageSize = 8 * 1024

// Checks if the Page has space for a record.
//
// Checks the length of the slot directory and calculates how
// much space the slot directory takes up.
func (p *Page) hasSpace(record []byte) bool {
	// Slot.Offset and Slot.Length are uint16 which means they're 2 bytes
	slotSpace := len(p.SlotDir) * SlotSize
	dataRegionSize := len(p.DataRegion)
	totalSpace := dataRegionSize + slotSpace + PageHeaderSize

	requiredSpace := len(record) + SlotSize // The + 4 is for the new slot

	return totalSpace+requiredSpace <= PageSize
}

// Inserts a new record into the page
//
// It does this by following these steps:
//  1. Check if there is enough space in the page
//  2. Create a slot based on the offset (FreeSpaceOffset - len(record)) and the length of the record
//  3. Add the slot, metadata and increase the record count.
//  4. Update header freeSpaceOffset and UsedSpaceEnd
//
// It returns an error if there is not enough space in the page.
func (p *Page) InsertRecord(record []byte) error {
	if !p.hasSpace(record) {
		return errors.New("error inserting record, not enough space in Page")
	}

	slot := &Slot{
		Offset: p.Header.FreeSpaceOffset - uint16(len(record)),
		Length: uint16(len(record)),
	}

	p.Header.FreeSpaceOffset += SlotSize
	p.Header.UsedSpaceEnd -= uint16(len(record))
	p.SlotDir = append(p.SlotDir, *slot)
	p.DataRegion = append(p.DataRegion, record...)
	p.Header.RecordCount++

	return nil
}

func (p *Page) Encode() ([]byte, error) {

	// initialize byte buffer
	buf := make([]byte, PageSize)

	// Write the header to the buffer
	if err := p.Header.WriteTo(buf[0:PageHeaderSize]); err != nil {
		return nil, err
	}

	// Write the slot directory to the buffer
	slotPos := PageHeaderSize
	for _, s := range p.SlotDir {
		err := s.WriteTo(buf[slotPos : slotPos+SlotSize])
		if err != nil {
			return nil, err
		}
		slotPos += 4
	}

	// Copy the data region to the buffer note that the starting spot is the end of the used space (which is the beginning of []byte data)
	copy(buf[p.Header.UsedSpaceEnd:], p.DataRegion)
	return buf, nil
}

// Creates a new page in a linked page fashion
//
// Initializes several fields in the header:
//   - PageID: Based on the pageId + 1
//   - PrevPageID: the pageId of the previous page
//   - UsedSpaceEnd: Initializes with the size of a page because that's the maximum available space after which it will shrink with records towards the beginning.
//   - FreeSpaceOffset: initializes with PageHeaderSize because from that place the first data starts after which it will grow with slots towards the end based on the slots.
//
// After a new page is initialized it gets set as the next page of the page passed into the function
func Create(prevPage Page) *Page {
	p := Page{
		Header: Header{
			PageID:          prevPage.Header.PrevPageID + 1,
			PrevPageID:      prevPage.Header.PageID,
			UsedSpaceEnd:    PageSize,
			FreeSpaceOffset: PageHeaderSize,
		},
	}
	prevPage.Header.NextPageID = p.Header.PageID
	return &p
}

// Creates a new page which will be the first page of a doubly linked list also called the head.
func CreateHead() *Page {
	return &Page{
		Header: Header{
			PageID:          1,
			UsedSpaceEnd:    PageSize,
			FreeSpaceOffset: PageHeaderSize,
		},
	}
}
