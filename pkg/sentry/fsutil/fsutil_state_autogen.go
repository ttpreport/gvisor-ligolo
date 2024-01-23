// automatically generated by stateify.

package fsutil

import (
	"github.com/ttpreport/gvisor-ligolo/pkg/state"
)

func (d *DirtyInfo) StateTypeName() string {
	return "pkg/sentry/fsutil.DirtyInfo"
}

func (d *DirtyInfo) StateFields() []string {
	return []string{
		"Keep",
	}
}

func (d *DirtyInfo) beforeSave() {}

// +checklocksignore
func (d *DirtyInfo) StateSave(stateSinkObject state.Sink) {
	d.beforeSave()
	stateSinkObject.Save(0, &d.Keep)
}

func (d *DirtyInfo) afterLoad() {}

// +checklocksignore
func (d *DirtyInfo) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &d.Keep)
}

func (f *HostFileMapper) StateTypeName() string {
	return "pkg/sentry/fsutil.HostFileMapper"
}

func (f *HostFileMapper) StateFields() []string {
	return []string{
		"refs",
	}
}

func (f *HostFileMapper) beforeSave() {}

// +checklocksignore
func (f *HostFileMapper) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.refs)
}

// +checklocksignore
func (f *HostFileMapper) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.refs)
	stateSourceObject.AfterLoad(f.afterLoad)
}

func init() {
	state.Register((*DirtyInfo)(nil))
	state.Register((*HostFileMapper)(nil))
}
