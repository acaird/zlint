// lint_ev_valid_time_too_long.go

package lints

import (

	"github.com/zmap/zlint/util"
	"github.com/zmap/zgrab/ztools/x509"
)

type evValidTooLong struct {
	// Internal data here
}

func (l *evValidTooLong) Initialize() error {
	return nil
}

func (l *evValidTooLong) CheckApplies(c *x509.Certificate) bool {
	return util.IsEv(c.PolicyIdentifiers)
}

func (l *evValidTooLong) RunTest(c *x509.Certificate) (ResultStruct, error) {
	if c.NotBefore.AddDate(2, 3, 0).Before(c.NotAfter) {
		return ResultStruct{Result: Error}, nil
	}
	return ResultStruct{Result: Pass}, nil
}

func init() {
	RegisterLint(&Lint{
		Name:          "ev_valid_time_too_long",
		Description:   "EV certificates must be 27 months in validity or less",
		Providence:    "",
		EffectiveDate: util.ZeroDate,
		Test:          &evValidTooLong{}})
}