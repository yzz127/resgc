package goresgc

// Holder interface defines a holder for aritrary resource and methods to access the
// resource it holds
type Holder interface {
	Get() interface{}
	Set(interface{})
	Clear()
	HasResource() bool
	SetCollector(Collector)
	AutoReclaim() bool
	CancelAutoReclaim()
	Destroy(...interface{})
}
