package builder

import (
	"design_pattern/pattern/builder/igloo"
	"design_pattern/pattern/builder/normal"
	builderDomain "design_pattern/domain/builder"
)


func GetBuilder(builderType string) builderDomain.Builder {
	if builderType == "normal" {
		return &normal.Builder{}
	}

	if builderType == "igloo" {
		return &igloo.Builder{}
	}
	return nil
}
