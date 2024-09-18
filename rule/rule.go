package rule

import "github.com/gonebot-dev/gonebot/message"

// Describes a filter rule, you can create your own filter by implementing this.
type FilterRule struct {
	// Return true if the message should be handled by the handler, false otherwise
	Filter func(msg message.Message) bool
}

// Describes a group of filters, all the filter results are ANDed together.
type FilterBundle struct {
	Filters []FilterRule
}

func (fb FilterBundle) Filter(msg message.Message) bool {
	for _, f := range fb.Filters {
		if !f.Filter(msg) {
			return false
		}
	}
	return true
}
