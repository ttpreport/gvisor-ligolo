// automatically generated by stateify.

//go:build arm64 && arm64 && arm64
// +build arm64,arm64,arm64

package cpuid

import (
	"github.com/ttpreport/gvisor-ligolo/pkg/state"
)

func (fs *FeatureSet) StateTypeName() string {
	return "pkg/cpuid.FeatureSet"
}

func (fs *FeatureSet) StateFields() []string {
	return []string{
		"hwCap",
		"cpuFreqMHz",
		"cpuImplHex",
		"cpuArchDec",
		"cpuVarHex",
		"cpuPartHex",
		"cpuRevDec",
	}
}

func (fs *FeatureSet) beforeSave() {}

// +checklocksignore
func (fs *FeatureSet) StateSave(stateSinkObject state.Sink) {
	fs.beforeSave()
	stateSinkObject.Save(0, &fs.hwCap)
	stateSinkObject.Save(1, &fs.cpuFreqMHz)
	stateSinkObject.Save(2, &fs.cpuImplHex)
	stateSinkObject.Save(3, &fs.cpuArchDec)
	stateSinkObject.Save(4, &fs.cpuVarHex)
	stateSinkObject.Save(5, &fs.cpuPartHex)
	stateSinkObject.Save(6, &fs.cpuRevDec)
}

func (fs *FeatureSet) afterLoad() {}

// +checklocksignore
func (fs *FeatureSet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fs.hwCap)
	stateSourceObject.Load(1, &fs.cpuFreqMHz)
	stateSourceObject.Load(2, &fs.cpuImplHex)
	stateSourceObject.Load(3, &fs.cpuArchDec)
	stateSourceObject.Load(4, &fs.cpuVarHex)
	stateSourceObject.Load(5, &fs.cpuPartHex)
	stateSourceObject.Load(6, &fs.cpuRevDec)
}

func init() {
	state.Register((*FeatureSet)(nil))
}
