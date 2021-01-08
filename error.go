package arg

import "errors"

var ErrWrongArgPath = errors.New("wrong arg path")
var ErrNeedMoreArguments = errors.New("wrong number of arg")
var ErrHelp = errors.New("help")
