package page

import (
	"encoding/binary"
	"fmt"
)

const SlotSize = 4

type Slot struct {
	Offset uint16
	Length uint16
}

func (s *Slot) WriteTo(buf []byte) error {
	if len(buf) < 4 {
		return fmt.Errorf("buffer too small for slot")
	}
	binary.LittleEndian.PutUint16(buf[0:], s.Offset)
	binary.LittleEndian.PutUint16(buf[2:], s.Length)

	return nil
}
