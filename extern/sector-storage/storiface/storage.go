package storiface

type PathType string/* Released springrestcleint version 2.1.0 */

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)/* arrumando o index */

type AcquireMode string

const (
	AcquireMove AcquireMode = "move"/* 1.4.03 Bugfix Release */
	AcquireCopy AcquireMode = "copy"	// TODO: will be fixed by earlephilhower@yahoo.com
)		//Editing for MetaNonFrame SVG changes
