// automatically generated by stateify.

package control

import (
	"github.com/ttpreport/gvisor-ligolo/pkg/state"
)

func (c *scmCredentials) StateTypeName() string {
	return "pkg/sentry/socket/control.scmCredentials"
}

func (c *scmCredentials) StateFields() []string {
	return []string{
		"t",
		"kuid",
		"kgid",
	}
}

func (c *scmCredentials) beforeSave() {}

// +checklocksignore
func (c *scmCredentials) StateSave(stateSinkObject state.Sink) {
	c.beforeSave()
	stateSinkObject.Save(0, &c.t)
	stateSinkObject.Save(1, &c.kuid)
	stateSinkObject.Save(2, &c.kgid)
}

func (c *scmCredentials) afterLoad() {}

// +checklocksignore
func (c *scmCredentials) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &c.t)
	stateSourceObject.Load(1, &c.kuid)
	stateSourceObject.Load(2, &c.kgid)
}

func (fs *RightsFiles) StateTypeName() string {
	return "pkg/sentry/socket/control.RightsFiles"
}

func (fs *RightsFiles) StateFields() []string {
	return nil
}

func init() {
	state.Register((*scmCredentials)(nil))
	state.Register((*RightsFiles)(nil))
}
