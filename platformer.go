package platformcontracts

type Platformer interface {
	GetBatchSize() int
	GetPlatformName() string
	GetReleaseUrlsByUpc(upcContainers []UpcContainer) []UrlResultContainer
}
