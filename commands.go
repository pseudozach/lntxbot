package main

import (
	"strings"

	"git.alhur.es/fiatjaf/lntxbot/t"
	"github.com/docopt/docopt-go"
	"github.com/kballard/go-shellquote"
)

type def struct {
	aliases        []string
	argstr         string
	inline         bool
	inline_example string
}

func (d def) help(lang string) string {
	key := t.Key(strings.ReplaceAll(d.aliases[0], " ", "_") + "Help")
	if _, ok := t.EN[key]; !ok {
		return ""
	}
	return translateTemplate(key, lang, t.T{"BotName": s.ServiceId})
}

var methods = []def{
	def{
		aliases: []string{"start"},
	},
	def{
		aliases:        []string{"receive", "invoice", "fund"},
		argstr:         "(lnurl <lnurl> | (<satoshis> | any) [<description>...] [--preimage=<preimage>])",
		inline:         true,
		inline_example: "invoice <satoshis>",
	},
	def{
		aliases: []string{"pay", "decode", "paynow", "withdraw"},
		argstr:  "(lnurl [<satoshis>] | [now] [<invoice>])",
	},
	def{
		aliases: []string{"send", "tip", "sendanonymously"},
		argstr:  "[anonymously] <satoshis> [<receiver>...] [--anonymous]",
	},
	def{
		aliases: []string{"balance"},
	},
	def{
		aliases: []string{"transactions"},
	},
	def{
		aliases:        []string{"giveaway"},
		argstr:         "<satoshis>",
		inline:         true,
		inline_example: "giveaway <satoshis>",
	},
	def{
		aliases:        []string{"coinflip", "lottery"},
		argstr:         "<satoshis> [<num_participants>]",
		inline:         true,
		inline_example: "coinflip <satoshis> <num_participants>",
	},
	def{
		aliases:        []string{"giveflip"},
		argstr:         "<satoshis> [<num_participants>]",
		inline:         true,
		inline_example: "giveflip <satoshis> <num_participants>",
	},
	def{
		aliases: []string{"fundraise", "crowdfund"},
		argstr:  "<satoshis> <num_participants> <receiver>...",
	},
	def{
		aliases: []string{"hide"},
		argstr:  "<satoshis> <message>... [--revealers=<num_revealers>] [--crowdfund=<num_participants>] [--public] [--private]",
	},
	def{
		aliases:        []string{"reveal"},
		argstr:         "<hidden_message_id>",
		inline:         true,
		inline_example: "reveal [hidden_message_id]",
	},
	def{
		aliases: []string{"app"},
	},
	def{
		aliases: []string{"microbet", "app microbet"},
		argstr:  "[bets | balance | withdraw | bet]",
	},
	def{
		aliases: []string{"bitflash", "app bitflash"},
		argstr:  "(orders | status | rate | <satoshis> <address>)",
	},
	def{
		aliases: []string{"satellite", "app satellite"},
		argstr:  "(transmissions | queue | bump <satoshis> <transmission_id> | delete <transmission_id> | <satoshis> <message>...)",
	},
	def{
		aliases: []string{"golightning", "app golightning"},
		argstr:  "<satoshis>",
	},
	def{
		aliases: []string{"gifts", "app gifts"},
		argstr:  "(list | [<satoshis>])",
	},
	def{
		aliases: []string{"paywall", "app paywall"},
		argstr:  "[list | <url> <satoshis> <memo>... | balance | withdraw]",
	},
	def{
		aliases:        []string{"poker", "app poker"},
		argstr:         "deposit <satoshis> | balance | withdraw | status | url | play | (available|watch|wait) <minutes>",
		inline:         true,
		inline_example: "poker",
	},
	def{
		aliases: []string{"bluewallet", "lndhub"},
		argstr:  "[refresh]",
	},
	def{
		aliases: []string{"toggle"},
		argstr:  "(ticket [<price>]|spammy)",
	},
	def{
		aliases: []string{"help"},
		argstr:  "[<command>]",
	},
	def{
		aliases: []string{"stop"},
	},
}

var commandList []string
var commandIndex = make(map[string]def)

func setupCommands() {
	s.Usage = docoptFromMethodDefinitions()

	for _, def := range methods {
		for _, alias := range def.aliases {
			commandList = append(commandList, alias)
			commandIndex[alias] = def
		}
	}
}

func docoptFromMethodDefinitions() string {
	var lines []string

	for _, def := range methods {
		for _, alias := range def.aliases {
			lines = append(lines, "  c "+alias+" "+def.argstr)
		}
	}

	return s.ServiceId + "\n\nUsage:\n" + strings.Join(lines, "\n")
}

var parser = docopt.Parser{
	HelpHandler:   func(_ error, _ string) {},
	OptionsFirst:  false,
	SkipHelpFlags: true,
}

func parse(message string) (opts docopt.Opts, isCommand bool, err error) {
	if strings.HasPrefix(message, "/") {
		isCommand = true
		message = message[1:]
	} else {
		return
	}

	// turn /app_microbet_bets, for example, into /app microbet bet
	parts := strings.SplitN(message, " ", 2)
	parts[0] = strings.ReplaceAll(parts[0], "_", " ")
	message = strings.Join(parts, " ")

	// apply this to get shell-like quoting rules
	var argv []string
	argv, err = shellquote.Split(message)
	if err != nil {
		// fallback to just a split when people do `/app satellite 23 can't quote`
		// which would normally break here
		argv = strings.Split(message, " ")
	}

	// parse using docopt
	opts, err = parser.ParseArgs(s.Usage, argv, "")
	return
}

func handleHelp(u User, method string) (handled bool) {
	var def def
	var mainName string
	var aliases []map[string]string
	var helpString string
	var ok bool
	method = strings.ToLower(strings.TrimSpace(method))
	if method == "" {
		helpString = translateTemplate(t.HELPINTRO, u.Locale, t.T{
			"Help": escapeHTML(strings.Replace(s.Usage, "  c ", "  /", -1)),
		})
		goto gothelpstring
	}

	// render specific help instructions for the given method
	def, ok = commandIndex[method]
	if ok {
		mainName = method
		aliases := ""
		for _, alias := range def.aliases {
			if alias != mainName {
				aliases += " " + alias
			}
		}
	} else {
		similar := findSimilar(method, commandList)
		if len(similar) > 0 {
			u.notify(t.HELPSIMILAR, t.T{
				"Method":  method,
				"Similar": similar,
			})
			return true
		} else {
			return false
		}
	}

	// here we have a working method definition
	helpString = translateTemplate(t.HELPMETHOD, u.Locale, t.T{
		"MainName":      mainName,
		"Argstr":        escapeHTML(def.argstr),
		"Help":          def.help(u.Locale),
		"HasInline":     def.inline,
		"InlineExample": escapeHTML(def.inline_example),
		"Aliases":       aliases,
		"ServiceId":     s.ServiceId,
	})

	goto gothelpstring

gothelpstring:
	sendMessage(u.ChatId, helpString)
	return true
}
