package repository

import (
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/traPtitech/traQ/utils"
	"github.com/traPtitech/traQ/utils/storage"
	"strings"
	"testing"
)

func TestRepositoryImpl_GetFS(t *testing.T) {
	t.Parallel()
	fs := storage.NewInMemoryFileStorage()
	repo := &GormRepository{fileImpl: fileImpl{FS: fs}}
	assert.Equal(t, fs, repo.GetFS())
}

func TestGormRepository_Group(t *testing.T) {
	t.Parallel()
	repo, _, _, user := setupWithUser(t, common)

	g1 := mustMakeUserGroup(t, repo, random, user.GetID())

	t.Run("Found", func(t *testing.T) {
		t.Parallel()
		id, ok := repo.Group(g1.Name)
		assert.True(t, ok)
		assert.EqualValues(t, g1.ID, id)
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()
		_, ok := repo.Group(utils.RandAlphabetAndNumberString(32))
		assert.False(t, ok)
	})
}

func TestGormRepository_Channel(t *testing.T) {
	t.Parallel()
	repo, _, _ := setup(t, common)

	c1 := mustMakeChannel(t, repo, random)
	c2 := mustMakeChannelDetail(t, repo, uuid.Nil, random, c1.ID)
	c3 := mustMakeChannelDetail(t, repo, uuid.Nil, random, c2.ID)
	c4 := mustMakeChannelDetail(t, repo, uuid.Nil, random, c3.ID)
	c5 := mustMakeChannelDetail(t, repo, uuid.Nil, random, c4.ID)

	t.Run("Found1", func(t *testing.T) {
		t.Parallel()
		path := c1.Name + "/" + c2.Name + "/" + c3.Name + "/" + c4.Name + "/" + c5.Name
		id, ok := repo.Channel(path)
		assert.True(t, ok)
		assert.EqualValues(t, c5.ID, id)
	})

	t.Run("Found2", func(t *testing.T) {
		t.Parallel()
		path := strings.ToUpper(c1.Name + "/" + c2.Name + "/" + c3.Name + "/" + c4.Name + "/" + c5.Name)
		id, ok := repo.Channel(path)
		assert.True(t, ok)
		assert.EqualValues(t, c5.ID, id)
	})

	t.Run("NotFound1", func(t *testing.T) {
		t.Parallel()
		path := strings.ToUpper(c1.Name + "/" + c2.Name + "/" + "a" + "/" + c4.Name + "/" + c5.Name)
		_, ok := repo.Channel(path)
		assert.False(t, ok)
	})

	t.Run("NotFound2", func(t *testing.T) {
		t.Parallel()
		path := ""
		_, ok := repo.Channel(path)
		assert.False(t, ok)
	})

	t.Run("NotFound3", func(t *testing.T) {
		t.Parallel()
		path := "/aaa/a"
		_, ok := repo.Channel(path)
		assert.False(t, ok)
	})

	t.Run("NotFound4", func(t *testing.T) {
		t.Parallel()
		path := c1.Name + "//" + c2.Name
		_, ok := repo.Channel(path)
		assert.False(t, ok)
	})
}

func TestGormRepository_User(t *testing.T) {
	t.Parallel()
	repo, _, _ := setup(t, common)

	u1 := mustMakeUser(t, repo, random)

	t.Run("Found1", func(t *testing.T) {
		t.Parallel()
		id, ok := repo.User(u1.GetName())
		assert.True(t, ok)
		assert.EqualValues(t, u1.GetID(), id)
	})

	t.Run("Found2", func(t *testing.T) {
		t.Parallel()
		id, ok := repo.User(strings.ToUpper(u1.GetName()))
		assert.True(t, ok)
		assert.EqualValues(t, u1.GetID(), id)
	})

	t.Run("NotFound", func(t *testing.T) {
		t.Parallel()
		_, ok := repo.User(utils.RandAlphabetAndNumberString(20))
		assert.False(t, ok)
	})
}
