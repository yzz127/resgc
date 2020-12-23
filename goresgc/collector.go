package goresgc

// Collector interface defined to collect holders whose resources need to get released
// after utilization
type Collector interface {
	Register(...interface{})
	Unregister(interface{})
	IsRegistered(interface{}) bool
	DestroyRes(interface{})
	Close(int64, int64) bool
}
