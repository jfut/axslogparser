package axslogparser

import (
	"fmt"
	"strings"
	"time"

	"github.com/Songmu/go-ltsv"
	"github.com/pkg/errors"
)

// LTSV access log parser
type LTSV struct {
}

// Parse for Parser interface
func (lv *LTSV) Parse(line string) (*Log, error) {
	for _, k := range []string{"apptime", "reqtime", "taken_sec"} {
		line = strings.Replace(line, fmt.Sprintf("\t%s:-", k), "", 1)
	}
	l := &Log{}
	err := ltsv.Unmarshal([]byte(line), l)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse ltsvlog")
	}
	l.Time, _ = time.Parse(clfTimeLayout, l.TimeStr)
	l.breakdownRequest()
	return l, nil
}
