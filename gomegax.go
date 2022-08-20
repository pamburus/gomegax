package gomegax

import (
	"testing"

	"github.com/onsi/gomega"
)

// New returns a new G.
func New(t *testing.T) G {
	return G{
		gomega.NewWithT(t),
		t,
	}
}

// G combines gomega.WithT and testing.T.
type G struct {
	*gomega.WithT
	*testing.T
}

// Run runs a child test with the give name.
func (g G) Run(name string, f func(g G)) {
	g.T.Run(name, func(t *testing.T) {
		f(New(t))
	})
}

// Fail marks the function as having failed but continues execution.
func (g G) Fail() {
	g.T.Fail()
}
