package writer

import "github.com/lennardclaproth/chronosdb/internal/engine/page"

type Writer interface {
	WritePage(page *page.Page) error
	Close() error
}
