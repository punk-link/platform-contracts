package platformcontracts

type Platformer interface {
	GetPlatformName() string
	GetReleaseUrlsByUpc(upcContainers []UpcContainer) []UrlResultContainer
}
