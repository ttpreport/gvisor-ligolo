// automatically generated by stateify.

//go:build (amd64 || 386) && go1.1
// +build amd64 386
// +build go1.1

package arch

import (
	"github.com/ttpreport/gvisor-ligolo/pkg/state"
)

func (s *State) StateTypeName() string {
	return "pkg/sentry/arch.State"
}

func (s *State) StateFields() []string {
	return []string{
		"Regs",
		"fpState",
	}
}

func (s *State) beforeSave() {}

// +checklocksignore
func (s *State) StateSave(stateSinkObject state.Sink) {
	s.beforeSave()
	stateSinkObject.Save(0, &s.Regs)
	stateSinkObject.Save(1, &s.fpState)
}

// +checklocksignore
func (s *State) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &s.Regs)
	stateSourceObject.LoadWait(1, &s.fpState)
	stateSourceObject.AfterLoad(s.afterLoad)
}

func init() {
	state.Register((*State)(nil))
}
