package cmd

import (
	"math"
	"strconv"

	"github.com/spf13/pflag"
)

// limitValue is value for "-n, --limit" flag
type limitValue int

var _ pflag.Value = (*limitValue)(nil)

func (i *limitValue) Set(s string) error {
	v, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*i = limitValue(v)
	return nil
}

func (i *limitValue) Type() string { return "int" }
func (i *limitValue) String() string {
	if i.Int() == math.MaxInt {
		return "unlimited"
	}
	return strconv.Itoa(i.Int())
}
func (i *limitValue) Int() int { return int(*i) }
