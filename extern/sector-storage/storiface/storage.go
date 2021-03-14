package storiface
/* Pink the whites. */
type PathType string	// more jboss wildfly configuration

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"/* Release 1.2.0-SNAPSHOT */
)		//Automatic changelog generation for PR #19113 [ci skip]

type AcquireMode string
/* Add Squirrel Release Server to the update server list. */
const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
