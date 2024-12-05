package rule

import (
	"log"
	"regexp"
	"strings"

	"github.com/gonebot-dev/gonebot/configurations"
	"github.com/gonebot-dev/gonebot/message"
)

// Always always returns true
func Always() *Rule {
	return &Rule{
		Filter: func(msg message.Message) bool {
			return true
		},
	}
}

// Never never returns true
func Never() *Rule {
	return &Rule{
		Filter: func(msg message.Message) bool {
			return false
		},
	}
}

// Command creates a filter rule that matches if the raw message is a command and is in the prefixList.
func Command(prefixList ...string) *Rule {
	return &Rule{
		Filter: func(msg message.Message) bool {
			for _, prefix := range prefixList {
				if strings.HasPrefix(
					msg.GetRawText(),
					configurations.GetConf("COMMAND_START")+prefix,
				) {
					return true
				}
			}
			return false
		},
	}
}

// FullMatch creates a filter rule that matches if the raw message is the same with one of the strings.
func FullMatch(strs ...string) *Rule {
	return &Rule{
		Filter: func(msg message.Message) bool {
			for _, str := range strs {
				if str == msg.GetRawText() {
					return true
				}
			}
			return false
		},
	}
}

// Keyword creates a filter rule that matches if the raw message contains one of the keywords.
//
// If forceStart is true, the keyword must be at the start of the message.
func Keyword(forceStart bool, keywords ...string) *Rule {
	return &Rule{
		Filter: func(msg message.Message) bool {
			for _, keyword := range keywords {
				if forceStart && strings.HasPrefix(msg.GetRawText(), keyword) {
					return true
				}
				if !forceStart && strings.Contains(msg.GetRawText(), keyword) {
					return true
				}
			}
			return false
		},
	}
}

// RegEx creates a filter rule that matches if the raw message does match one of the RegEx expressions.
//
// If you wrote a wrong RegEx expression, an error message with plugin name will be printed.
func RegEx(pluginName string, exprs ...string) *Rule {
	return &Rule{
		Filter: func(msg message.Message) bool {
			for _, expr := range exprs {
				reg, err := regexp.Compile(expr)
				if err != nil {
					log.Printf("[GONEBOT] | %s: RegEx filter rule compilation error!\n", pluginName)
					return false
				}
				if reg.FindStringIndex(msg.GetRawText()) != nil {
					return true
				}
			}
			return false
		},
	}
}

// ToMe filters messages that are directed to the bot.(@bot or private message, should be identified by adapters)
//
// If your adapter don't, what can i say?
func ToMe() *Rule {
	return &Rule{
		Filter: func(msg message.Message) bool {
			return msg.IsToMe
		},
	}
}

// OfType filters messages that has the specified type for specified adapter
//
// # You must be sure the message has at least one segment
func OfType(typeName, adapterName string) *Rule {
	return &Rule{
		Filter: func(msg message.Message) bool {
			return msg.GetSegments()[0].Type == typeName && msg.GetSegments()[0].Data.AdapterName() == adapterName
		},
	}
}

// Notice filters messages with specified notice types.
//
// Combines it with `ToMe` to filter notice messages directed to the bot.
func Notice(typeList ...string) *Rule {
	return &Rule{
		Filter: func(msg message.Message) bool {
			for _, typeName := range typeList {
				if msg.GetSegments()[0].Type == typeName {
					return true
				}
			}
			return false
		},
	}
}
