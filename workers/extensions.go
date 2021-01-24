package workers

import (
	"github.com/PuerkitoBio/goquery"
	uuid "github.com/nu7hatch/gouuid"
)

// WrapperSelection wrapps the goquery selection
type WrapperSelection struct {
	selection  *goquery.Selection
	identifier *uuid.UUID
}

// WrapperMessage wrapps the all message
type WrapperMessage struct {
	text           string
	querySelection *WrapperSelection
	configuration  *HandlerConfiguration
}
