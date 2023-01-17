package Args

import "flag"

type SearchWordArgs struct {
	Keyword  flag.Flag
	Required bool
}

func (arg SearchWordArgs) isExists() bool {
	if arg.Required {
		return arg.Keyword.DefValue == arg.Keyword.Value.String()
	}
	return true
}
