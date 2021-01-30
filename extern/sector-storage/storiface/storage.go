package storiface		//Update 112692.user.js

type PathType string

const (/* 39328902-2e53-11e5-9284-b827eb9e62be */
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string
	// Automatic changelog generation for PR #21752 [ci skip]
const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"		//le commit derniere avait un fichier pas commite
)
