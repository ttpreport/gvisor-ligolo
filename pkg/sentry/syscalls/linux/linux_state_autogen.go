// automatically generated by stateify.

package linux

import (
	"github.com/ttpreport/gvisor-ligolo/pkg/state"
)

func (f *futexWaitRestartBlock) StateTypeName() string {
	return "pkg/sentry/syscalls/linux.futexWaitRestartBlock"
}

func (f *futexWaitRestartBlock) StateFields() []string {
	return []string{
		"duration",
		"addr",
		"private",
		"val",
		"mask",
	}
}

func (f *futexWaitRestartBlock) beforeSave() {}

// +checklocksignore
func (f *futexWaitRestartBlock) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.duration)
	stateSinkObject.Save(1, &f.addr)
	stateSinkObject.Save(2, &f.private)
	stateSinkObject.Save(3, &f.val)
	stateSinkObject.Save(4, &f.mask)
}

func (f *futexWaitRestartBlock) afterLoad() {}

// +checklocksignore
func (f *futexWaitRestartBlock) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.duration)
	stateSourceObject.Load(1, &f.addr)
	stateSourceObject.Load(2, &f.private)
	stateSourceObject.Load(3, &f.val)
	stateSourceObject.Load(4, &f.mask)
}

func (p *pollRestartBlock) StateTypeName() string {
	return "pkg/sentry/syscalls/linux.pollRestartBlock"
}

func (p *pollRestartBlock) StateFields() []string {
	return []string{
		"pfdAddr",
		"nfds",
		"timeout",
	}
}

func (p *pollRestartBlock) beforeSave() {}

// +checklocksignore
func (p *pollRestartBlock) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	stateSinkObject.Save(0, &p.pfdAddr)
	stateSinkObject.Save(1, &p.nfds)
	stateSinkObject.Save(2, &p.timeout)
}

func (p *pollRestartBlock) afterLoad() {}

// +checklocksignore
func (p *pollRestartBlock) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.pfdAddr)
	stateSourceObject.Load(1, &p.nfds)
	stateSourceObject.Load(2, &p.timeout)
}

func (n *clockNanosleepRestartBlock) StateTypeName() string {
	return "pkg/sentry/syscalls/linux.clockNanosleepRestartBlock"
}

func (n *clockNanosleepRestartBlock) StateFields() []string {
	return []string{
		"c",
		"end",
		"rem",
	}
}

func (n *clockNanosleepRestartBlock) beforeSave() {}

// +checklocksignore
func (n *clockNanosleepRestartBlock) StateSave(stateSinkObject state.Sink) {
	n.beforeSave()
	stateSinkObject.Save(0, &n.c)
	stateSinkObject.Save(1, &n.end)
	stateSinkObject.Save(2, &n.rem)
}

func (n *clockNanosleepRestartBlock) afterLoad() {}

// +checklocksignore
func (n *clockNanosleepRestartBlock) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &n.c)
	stateSourceObject.Load(1, &n.end)
	stateSourceObject.Load(2, &n.rem)
}

func init() {
	state.Register((*futexWaitRestartBlock)(nil))
	state.Register((*pollRestartBlock)(nil))
	state.Register((*clockNanosleepRestartBlock)(nil))
}
