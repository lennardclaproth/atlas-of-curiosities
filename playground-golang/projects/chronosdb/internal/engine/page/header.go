/*
	File: 		header.go
	Author: 	Lennard Claproth
	Date: 		20230514

	The header serves as metadata for the pages, it contains data
	so that pages can function as doubly linked lists and you can
	traverse them that way easily. Besides that it also serves as
	an overview of how much space there is still left in the page.
	This is so that we can easily check if there is still enough
	space in the page.
*/

package page

import (
	"encoding/binary"
	"errors"
)

const PageHeaderSize = 40

// PageHeader represents the fixed-size metadata stored at the beginning of each page.
// It contains structural and statistical information needed to interpret the page.
type Header struct {
	RecordCount     uint16
	PageID          uint32
	PrevPageID      uint32
	NextPageID      uint32
	FreeSpaceOffset uint16
	UsedSpaceEnd    uint16
	TimestampMin    int64
	TimestampMax    int64
}

var ErrBufferTooSmall = errors.New("buffer too small")

// Writes a byte array to a PageHeader
func (h *Header) WriteTo(buf []byte) error {
	if len(buf) < PageHeaderSize {
		return ErrBufferTooSmall
	}

	binary.LittleEndian.PutUint16(buf[0:], h.RecordCount)
	binary.LittleEndian.PutUint32(buf[2:], h.PageID)
	binary.LittleEndian.PutUint32(buf[6:], h.PrevPageID)
	binary.LittleEndian.PutUint32(buf[10:], h.NextPageID)
	binary.LittleEndian.PutUint16(buf[14:], h.FreeSpaceOffset)
	binary.LittleEndian.PutUint16(buf[16:], h.UsedSpaceEnd)
	binary.LittleEndian.PutUint64(buf[18:], uint64(h.TimestampMin))
	binary.LittleEndian.PutUint64(buf[26:], uint64(h.TimestampMax))

	return nil
}

func (h *Header) ReadFrom(buf []byte) error {
	if len(buf) < PageHeaderSize {
		return ErrBufferTooSmall
	}

	h.RecordCount = binary.LittleEndian.Uint16(buf[0:])
	h.PageID = binary.LittleEndian.Uint32(buf[2:])
	h.PrevPageID = binary.LittleEndian.Uint32(buf[6:])
	h.NextPageID = binary.LittleEndian.Uint32(buf[10:])
	h.FreeSpaceOffset = binary.LittleEndian.Uint16(buf[14:])
	h.UsedSpaceEnd = binary.LittleEndian.Uint16(buf[16:])
	h.TimestampMin = int64(binary.LittleEndian.Uint64(buf[18:]))
	h.TimestampMax = int64(binary.LittleEndian.Uint64(buf[26:]))

	return nil
}
