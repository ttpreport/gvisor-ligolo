// automatically generated by stateify.

package fuse

import (
	"github.com/ttpreport/gvisor-ligolo/pkg/state"
)

func (conn *connection) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.connection"
}

func (conn *connection) StateFields() []string {
	return []string{
		"fd",
		"attributeVersion",
		"initialized",
		"initializedChan",
		"connected",
		"connInitError",
		"connInitSuccess",
		"aborted",
		"numWaiting",
		"asyncNum",
		"asyncCongestionThreshold",
		"asyncNumMax",
		"maxRead",
		"maxWrite",
		"maxPages",
		"maxActiveRequests",
		"minor",
		"atomicOTrunc",
		"asyncRead",
		"writebackCache",
		"bigWrites",
		"dontMask",
		"noOpen",
	}
}

func (conn *connection) beforeSave() {}

// +checklocksignore
func (conn *connection) StateSave(stateSinkObject state.Sink) {
	conn.beforeSave()
	var initializedChanValue bool
	initializedChanValue = conn.saveInitializedChan()
	stateSinkObject.SaveValue(3, initializedChanValue)
	stateSinkObject.Save(0, &conn.fd)
	stateSinkObject.Save(1, &conn.attributeVersion)
	stateSinkObject.Save(2, &conn.initialized)
	stateSinkObject.Save(4, &conn.connected)
	stateSinkObject.Save(5, &conn.connInitError)
	stateSinkObject.Save(6, &conn.connInitSuccess)
	stateSinkObject.Save(7, &conn.aborted)
	stateSinkObject.Save(8, &conn.numWaiting)
	stateSinkObject.Save(9, &conn.asyncNum)
	stateSinkObject.Save(10, &conn.asyncCongestionThreshold)
	stateSinkObject.Save(11, &conn.asyncNumMax)
	stateSinkObject.Save(12, &conn.maxRead)
	stateSinkObject.Save(13, &conn.maxWrite)
	stateSinkObject.Save(14, &conn.maxPages)
	stateSinkObject.Save(15, &conn.maxActiveRequests)
	stateSinkObject.Save(16, &conn.minor)
	stateSinkObject.Save(17, &conn.atomicOTrunc)
	stateSinkObject.Save(18, &conn.asyncRead)
	stateSinkObject.Save(19, &conn.writebackCache)
	stateSinkObject.Save(20, &conn.bigWrites)
	stateSinkObject.Save(21, &conn.dontMask)
	stateSinkObject.Save(22, &conn.noOpen)
}

func (conn *connection) afterLoad() {}

// +checklocksignore
func (conn *connection) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &conn.fd)
	stateSourceObject.Load(1, &conn.attributeVersion)
	stateSourceObject.Load(2, &conn.initialized)
	stateSourceObject.Load(4, &conn.connected)
	stateSourceObject.Load(5, &conn.connInitError)
	stateSourceObject.Load(6, &conn.connInitSuccess)
	stateSourceObject.Load(7, &conn.aborted)
	stateSourceObject.Load(8, &conn.numWaiting)
	stateSourceObject.Load(9, &conn.asyncNum)
	stateSourceObject.Load(10, &conn.asyncCongestionThreshold)
	stateSourceObject.Load(11, &conn.asyncNumMax)
	stateSourceObject.Load(12, &conn.maxRead)
	stateSourceObject.Load(13, &conn.maxWrite)
	stateSourceObject.Load(14, &conn.maxPages)
	stateSourceObject.Load(15, &conn.maxActiveRequests)
	stateSourceObject.Load(16, &conn.minor)
	stateSourceObject.Load(17, &conn.atomicOTrunc)
	stateSourceObject.Load(18, &conn.asyncRead)
	stateSourceObject.Load(19, &conn.writebackCache)
	stateSourceObject.Load(20, &conn.bigWrites)
	stateSourceObject.Load(21, &conn.dontMask)
	stateSourceObject.Load(22, &conn.noOpen)
	stateSourceObject.LoadValue(3, new(bool), func(y any) { conn.loadInitializedChan(y.(bool)) })
}

func (f *fuseDevice) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.fuseDevice"
}

func (f *fuseDevice) StateFields() []string {
	return []string{}
}

func (f *fuseDevice) beforeSave() {}

// +checklocksignore
func (f *fuseDevice) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
}

func (f *fuseDevice) afterLoad() {}

// +checklocksignore
func (f *fuseDevice) StateLoad(stateSourceObject state.Source) {
}

func (fd *DeviceFD) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.DeviceFD"
}

func (fd *DeviceFD) StateFields() []string {
	return []string{
		"vfsfd",
		"FileDescriptionDefaultImpl",
		"DentryMetadataFileDescriptionImpl",
		"NoLockFD",
		"waitQueue",
		"fullQueueCh",
		"nextOpID",
		"queue",
		"numActiveRequests",
		"completions",
		"writeBuf",
		"conn",
	}
}

func (fd *DeviceFD) beforeSave() {}

// +checklocksignore
func (fd *DeviceFD) StateSave(stateSinkObject state.Sink) {
	fd.beforeSave()
	var fullQueueChValue int
	fullQueueChValue = fd.saveFullQueueCh()
	stateSinkObject.SaveValue(5, fullQueueChValue)
	stateSinkObject.Save(0, &fd.vfsfd)
	stateSinkObject.Save(1, &fd.FileDescriptionDefaultImpl)
	stateSinkObject.Save(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSinkObject.Save(3, &fd.NoLockFD)
	stateSinkObject.Save(4, &fd.waitQueue)
	stateSinkObject.Save(6, &fd.nextOpID)
	stateSinkObject.Save(7, &fd.queue)
	stateSinkObject.Save(8, &fd.numActiveRequests)
	stateSinkObject.Save(9, &fd.completions)
	stateSinkObject.Save(10, &fd.writeBuf)
	stateSinkObject.Save(11, &fd.conn)
}

func (fd *DeviceFD) afterLoad() {}

// +checklocksignore
func (fd *DeviceFD) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fd.vfsfd)
	stateSourceObject.Load(1, &fd.FileDescriptionDefaultImpl)
	stateSourceObject.Load(2, &fd.DentryMetadataFileDescriptionImpl)
	stateSourceObject.Load(3, &fd.NoLockFD)
	stateSourceObject.Load(4, &fd.waitQueue)
	stateSourceObject.Load(6, &fd.nextOpID)
	stateSourceObject.Load(7, &fd.queue)
	stateSourceObject.Load(8, &fd.numActiveRequests)
	stateSourceObject.Load(9, &fd.completions)
	stateSourceObject.Load(10, &fd.writeBuf)
	stateSourceObject.Load(11, &fd.conn)
	stateSourceObject.LoadValue(5, new(int), func(y any) { fd.loadFullQueueCh(y.(int)) })
}

func (fsType *FilesystemType) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.FilesystemType"
}

func (fsType *FilesystemType) StateFields() []string {
	return []string{}
}

func (fsType *FilesystemType) beforeSave() {}

// +checklocksignore
func (fsType *FilesystemType) StateSave(stateSinkObject state.Sink) {
	fsType.beforeSave()
}

func (fsType *FilesystemType) afterLoad() {}

// +checklocksignore
func (fsType *FilesystemType) StateLoad(stateSourceObject state.Source) {
}

func (f *filesystemOptions) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.filesystemOptions"
}

func (f *filesystemOptions) StateFields() []string {
	return []string{
		"mopts",
		"uid",
		"gid",
		"rootMode",
		"maxActiveRequests",
		"maxRead",
		"defaultPermissions",
		"allowOther",
	}
}

func (f *filesystemOptions) beforeSave() {}

// +checklocksignore
func (f *filesystemOptions) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.mopts)
	stateSinkObject.Save(1, &f.uid)
	stateSinkObject.Save(2, &f.gid)
	stateSinkObject.Save(3, &f.rootMode)
	stateSinkObject.Save(4, &f.maxActiveRequests)
	stateSinkObject.Save(5, &f.maxRead)
	stateSinkObject.Save(6, &f.defaultPermissions)
	stateSinkObject.Save(7, &f.allowOther)
}

func (f *filesystemOptions) afterLoad() {}

// +checklocksignore
func (f *filesystemOptions) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.mopts)
	stateSourceObject.Load(1, &f.uid)
	stateSourceObject.Load(2, &f.gid)
	stateSourceObject.Load(3, &f.rootMode)
	stateSourceObject.Load(4, &f.maxActiveRequests)
	stateSourceObject.Load(5, &f.maxRead)
	stateSourceObject.Load(6, &f.defaultPermissions)
	stateSourceObject.Load(7, &f.allowOther)
}

func (fs *filesystem) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.filesystem"
}

func (fs *filesystem) StateFields() []string {
	return []string{
		"Filesystem",
		"devMinor",
		"conn",
		"opts",
		"clock",
	}
}

func (fs *filesystem) beforeSave() {}

// +checklocksignore
func (fs *filesystem) StateSave(stateSinkObject state.Sink) {
	fs.beforeSave()
	stateSinkObject.Save(0, &fs.Filesystem)
	stateSinkObject.Save(1, &fs.devMinor)
	stateSinkObject.Save(2, &fs.conn)
	stateSinkObject.Save(3, &fs.opts)
	stateSinkObject.Save(4, &fs.clock)
}

func (fs *filesystem) afterLoad() {}

// +checklocksignore
func (fs *filesystem) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &fs.Filesystem)
	stateSourceObject.Load(1, &fs.devMinor)
	stateSourceObject.Load(2, &fs.conn)
	stateSourceObject.Load(3, &fs.opts)
	stateSourceObject.Load(4, &fs.clock)
}

func (f *fileHandle) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.fileHandle"
}

func (f *fileHandle) StateFields() []string {
	return []string{
		"new",
		"handle",
		"flags",
	}
}

func (f *fileHandle) beforeSave() {}

// +checklocksignore
func (f *fileHandle) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.new)
	stateSinkObject.Save(1, &f.handle)
	stateSinkObject.Save(2, &f.flags)
}

func (f *fileHandle) afterLoad() {}

// +checklocksignore
func (f *fileHandle) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.new)
	stateSourceObject.Load(1, &f.handle)
	stateSourceObject.Load(2, &f.flags)
}

func (i *inode) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.inode"
}

func (i *inode) StateFields() []string {
	return []string{
		"inodeRefs",
		"InodeAlwaysValid",
		"InodeNotAnonymous",
		"InodeNotSymlink",
		"InodeWatches",
		"OrderedChildren",
		"CachedMappable",
		"fs",
		"nodeID",
		"attrVersion",
		"attrTime",
		"link",
		"fh",
		"locks",
		"watches",
		"attrMu",
		"ino",
		"uid",
		"gid",
		"mode",
		"atime",
		"mtime",
		"ctime",
		"size",
		"nlink",
		"blockSize",
	}
}

func (i *inode) beforeSave() {}

// +checklocksignore
func (i *inode) StateSave(stateSinkObject state.Sink) {
	i.beforeSave()
	stateSinkObject.Save(0, &i.inodeRefs)
	stateSinkObject.Save(1, &i.InodeAlwaysValid)
	stateSinkObject.Save(2, &i.InodeNotAnonymous)
	stateSinkObject.Save(3, &i.InodeNotSymlink)
	stateSinkObject.Save(4, &i.InodeWatches)
	stateSinkObject.Save(5, &i.OrderedChildren)
	stateSinkObject.Save(6, &i.CachedMappable)
	stateSinkObject.Save(7, &i.fs)
	stateSinkObject.Save(8, &i.nodeID)
	stateSinkObject.Save(9, &i.attrVersion)
	stateSinkObject.Save(10, &i.attrTime)
	stateSinkObject.Save(11, &i.link)
	stateSinkObject.Save(12, &i.fh)
	stateSinkObject.Save(13, &i.locks)
	stateSinkObject.Save(14, &i.watches)
	stateSinkObject.Save(15, &i.attrMu)
	stateSinkObject.Save(16, &i.ino)
	stateSinkObject.Save(17, &i.uid)
	stateSinkObject.Save(18, &i.gid)
	stateSinkObject.Save(19, &i.mode)
	stateSinkObject.Save(20, &i.atime)
	stateSinkObject.Save(21, &i.mtime)
	stateSinkObject.Save(22, &i.ctime)
	stateSinkObject.Save(23, &i.size)
	stateSinkObject.Save(24, &i.nlink)
	stateSinkObject.Save(25, &i.blockSize)
}

func (i *inode) afterLoad() {}

// +checklocksignore
func (i *inode) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &i.inodeRefs)
	stateSourceObject.Load(1, &i.InodeAlwaysValid)
	stateSourceObject.Load(2, &i.InodeNotAnonymous)
	stateSourceObject.Load(3, &i.InodeNotSymlink)
	stateSourceObject.Load(4, &i.InodeWatches)
	stateSourceObject.Load(5, &i.OrderedChildren)
	stateSourceObject.Load(6, &i.CachedMappable)
	stateSourceObject.Load(7, &i.fs)
	stateSourceObject.Load(8, &i.nodeID)
	stateSourceObject.Load(9, &i.attrVersion)
	stateSourceObject.Load(10, &i.attrTime)
	stateSourceObject.Load(11, &i.link)
	stateSourceObject.Load(12, &i.fh)
	stateSourceObject.Load(13, &i.locks)
	stateSourceObject.Load(14, &i.watches)
	stateSourceObject.Load(15, &i.attrMu)
	stateSourceObject.Load(16, &i.ino)
	stateSourceObject.Load(17, &i.uid)
	stateSourceObject.Load(18, &i.gid)
	stateSourceObject.Load(19, &i.mode)
	stateSourceObject.Load(20, &i.atime)
	stateSourceObject.Load(21, &i.mtime)
	stateSourceObject.Load(22, &i.ctime)
	stateSourceObject.Load(23, &i.size)
	stateSourceObject.Load(24, &i.nlink)
	stateSourceObject.Load(25, &i.blockSize)
}

func (r *inodeRefs) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.inodeRefs"
}

func (r *inodeRefs) StateFields() []string {
	return []string{
		"refCount",
	}
}

func (r *inodeRefs) beforeSave() {}

// +checklocksignore
func (r *inodeRefs) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.refCount)
}

// +checklocksignore
func (r *inodeRefs) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.refCount)
	stateSourceObject.AfterLoad(r.afterLoad)
}

func (l *requestList) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.requestList"
}

func (l *requestList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *requestList) beforeSave() {}

// +checklocksignore
func (l *requestList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *requestList) afterLoad() {}

// +checklocksignore
func (l *requestList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *requestEntry) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.requestEntry"
}

func (e *requestEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *requestEntry) beforeSave() {}

// +checklocksignore
func (e *requestEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *requestEntry) afterLoad() {}

// +checklocksignore
func (e *requestEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func (r *Request) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.Request"
}

func (r *Request) StateFields() []string {
	return []string{
		"requestEntry",
		"id",
		"hdr",
		"data",
		"async",
		"noReply",
	}
}

func (r *Request) beforeSave() {}

// +checklocksignore
func (r *Request) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.requestEntry)
	stateSinkObject.Save(1, &r.id)
	stateSinkObject.Save(2, &r.hdr)
	stateSinkObject.Save(3, &r.data)
	stateSinkObject.Save(4, &r.async)
	stateSinkObject.Save(5, &r.noReply)
}

func (r *Request) afterLoad() {}

// +checklocksignore
func (r *Request) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.requestEntry)
	stateSourceObject.Load(1, &r.id)
	stateSourceObject.Load(2, &r.hdr)
	stateSourceObject.Load(3, &r.data)
	stateSourceObject.Load(4, &r.async)
	stateSourceObject.Load(5, &r.noReply)
}

func (f *futureResponse) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.futureResponse"
}

func (f *futureResponse) StateFields() []string {
	return []string{
		"opcode",
		"ch",
		"hdr",
		"data",
		"async",
	}
}

func (f *futureResponse) beforeSave() {}

// +checklocksignore
func (f *futureResponse) StateSave(stateSinkObject state.Sink) {
	f.beforeSave()
	stateSinkObject.Save(0, &f.opcode)
	stateSinkObject.Save(1, &f.ch)
	stateSinkObject.Save(2, &f.hdr)
	stateSinkObject.Save(3, &f.data)
	stateSinkObject.Save(4, &f.async)
}

func (f *futureResponse) afterLoad() {}

// +checklocksignore
func (f *futureResponse) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &f.opcode)
	stateSourceObject.Load(1, &f.ch)
	stateSourceObject.Load(2, &f.hdr)
	stateSourceObject.Load(3, &f.data)
	stateSourceObject.Load(4, &f.async)
}

func (r *Response) StateTypeName() string {
	return "pkg/sentry/fsimpl/fuse.Response"
}

func (r *Response) StateFields() []string {
	return []string{
		"opcode",
		"hdr",
		"data",
	}
}

func (r *Response) beforeSave() {}

// +checklocksignore
func (r *Response) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.opcode)
	stateSinkObject.Save(1, &r.hdr)
	stateSinkObject.Save(2, &r.data)
}

func (r *Response) afterLoad() {}

// +checklocksignore
func (r *Response) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.opcode)
	stateSourceObject.Load(1, &r.hdr)
	stateSourceObject.Load(2, &r.data)
}

func init() {
	state.Register((*connection)(nil))
	state.Register((*fuseDevice)(nil))
	state.Register((*DeviceFD)(nil))
	state.Register((*FilesystemType)(nil))
	state.Register((*filesystemOptions)(nil))
	state.Register((*filesystem)(nil))
	state.Register((*fileHandle)(nil))
	state.Register((*inode)(nil))
	state.Register((*inodeRefs)(nil))
	state.Register((*requestList)(nil))
	state.Register((*requestEntry)(nil))
	state.Register((*Request)(nil))
	state.Register((*futureResponse)(nil))
	state.Register((*Response)(nil))
}
