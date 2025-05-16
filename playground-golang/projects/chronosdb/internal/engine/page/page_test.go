package page

import (
	"testing"
)

func TestCreateFirstPageAndInsertrecord(t *testing.T) {
	// ARRANGE
	p := CreateHead()
	record := []byte("hello chronosdb")

	// ACT
	err := p.InsertRecord(record)

	// ASSERT
	if err != nil {
		t.Fatalf("Insert record failed: %v", err)
	}

	if p.Header.RecordCount != 1 {
		t.Errorf("expected 1 record, got %d", p.Header.RecordCount)
	}

	lastSlot := p.SlotDir[len(p.SlotDir)-1]
	if lastSlot.Length != uint16(len(record)) {
		t.Errorf("expected slot length %d, got %d", len(record), lastSlot.Length)
	}
}
