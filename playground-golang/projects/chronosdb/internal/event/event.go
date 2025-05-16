/*

A potential json would look like this:
{
	NodeName: PopcornMarket.MyService
	Timestamp: 1747390174,
	Type: my.type,
	CorrelationId: "d9ebfa43-fa17-499a-806f-335d7da0340a",
	CausationId: "abf355f3-dda6-44f4-a3f5-1d03b275691e",
	DataType: "json",
	Data: [dc 07 61 2a 84 db 47 f6 3b ad ]
}

*/

package event

type Event struct {
	NodeName      string
	Timestamp     int64
	Type          string
	CorrelationID string
	CausationID   string
	DataType      string // Fully qualified type name
	Data          []byte
}

func (e *Event) Encode() ([]byte, error) {
	
}

// func Decode() (*Event, error)