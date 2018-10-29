package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ListBuckets(t *testing.T) {
	list, err := ListBuckets()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(list))
	assert.Equal(t, "bucket0", list[0].Name())
	assert.Equal(t, "bucket1", list[1].Name())
}

func Test_ListObjects_empty(t *testing.T) {
	list, err := ListObjects("bucket0", "")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(list))
}

func Test_ListObjects_fileexists(t *testing.T) {
	list, err := ListObjects("bucket1", "")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(list))
	assert.Equal(t, "object1.txt", list[0].Name())
}
