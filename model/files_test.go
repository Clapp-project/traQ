package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile_TableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "files", (&File{}).TableName())
}

func TestFileACLEntry_TableName(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "files_acl", (&FileACLEntry{}).TableName())
}
