package rule

import "github.com/gonebot-dev/gonebot/message"

// Describes the way the rules are combined.
type RuleOp bool

const (
	Or  RuleOp = false
	And RuleOp = true
)

// Describes a filter rule, you can create your own rule by implementing this.
type Rule struct {
	// Return true if the message should be handled by the handler, false otherwise.
	Filter func(msg message.Message) bool
	Next   *Rule
	// Only useful when Next is not nil.
	Operator RuleOp
}

// Describes a list of rules, all the Filter results are "OR"ed together, all the Rules results are "AND"ed together.
type Rules struct {
	FilterHead *Rule
	FilterTail *Rule
	Next       *Rules
	Last       *Rules
	// Only useful when Next is not nil.
	Operator RuleOp
}

func (ru *Rule) SubFilter(msg message.Message) (result bool) {
	if ru == nil {
		return true
	}
	result = ru.Filter(msg)
	if ru.Operator == Or && ru.Next != nil {
		result = result || ru.Next.SubFilter(msg)
	} else if ru.Operator == And && ru.Next != nil {
		result = result && ru.Next.SubFilter(msg)
	}
	return
}

func (ru *Rules) Filter(msg message.Message) (result bool) {
	if ru == nil {
		return true
	}
	result = ru.FilterHead.SubFilter(msg)
	if ru.Operator == Or && ru.Next != nil {
		result = result || ru.Next.Filter(msg)
	} else if ru.Operator == And && ru.Next != nil {
		result = result && ru.Next.Filter(msg)
	}
	return
}

func NewRules(rule *Rule) *Rules {
	rules := &Rules{
		FilterHead: rule,
		FilterTail: rule,
		Next:       nil,
	}
	rules.Last = rules
	return rules
}

func (ru *Rules) Or(rule *Rule) *Rules {
	ru.FilterTail.Operator = Or
	ru.FilterTail.Next = rule
	ru.FilterTail = ru.FilterTail.Next
	return ru
}

func (ru *Rules) And(rule *Rule) *Rules {
	ru.FilterTail.Operator = And
	ru.FilterTail.Next = rule
	ru.FilterTail = ru.FilterTail.Next
	return ru
}

func (ru *Rules) OrRules(rules *Rules) *Rules {
	ru.Last.Operator = Or
	ru.Last.Next = rules
	ru.Last = rules
	return ru
}

func (ru *Rules) AndRules(rules *Rules) *Rules {
	ru.Last.Operator = And
	ru.Last.Next = rules
	ru.Last = rules
	return ru
}
