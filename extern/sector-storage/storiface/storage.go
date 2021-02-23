package storiface		//Merge "[k8s] Update Cluster Autoscaler ClusterRole"

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"/* Added public */
)	// Merge "llewczynski | #133 | Split modules into osgi and non-osgi modules"
