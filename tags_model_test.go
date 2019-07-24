package main

import (
	"testing"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateTag(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()

	tag := &Tag{
		Name:           "name",
		PrimaryColor:   "primaryColor",
		SecondaryColor: "secondaryColor",
	}

	tag, err := tag.Create()

	assert.NoError(t, err)
	assert.NotZero(t, tag.ID)
	assert.Equal(t, "name", tag.Name)
	assert.Equal(t, "primaryColor", tag.PrimaryColor)
	assert.Equal(t, "secondaryColor", tag.SecondaryColor)
	assert.Equal(t, true, tag.Enabled)
}

func TestGetTag(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()

	tag := &Tag{
		Name:           "name",
		PrimaryColor:   "primaryColor",
		SecondaryColor: "secondaryColor",
	}

	tag, _ = tag.Create()
	tag, err := tag.Get()

	assert.NoError(t, err)
	assert.NotZero(t, tag.ID)
	assert.Equal(t, "name", tag.Name)
	assert.Equal(t, "primaryColor", tag.PrimaryColor)
	assert.Equal(t, "secondaryColor", tag.SecondaryColor)
	assert.Equal(t, true, tag.Enabled)
}

func TestUpdateTag(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()

	tag := &Tag{
		Name:           "name",
		PrimaryColor:   "primaryColor",
		SecondaryColor: "secondaryColor",
	}
	tag, _ = tag.Create()

	tag.Name = "name2"
	tag.PrimaryColor = "primaryColor2"
	tag.SecondaryColor = "secondaryColor2"
	tag.Enabled = !tag.Enabled

	tag, err := tag.Update()
	assert.NoError(t, err)
	assert.NotZero(t, tag.ID)
	assert.Equal(t, "name2", tag.Name)
	assert.Equal(t, "primaryColor2", tag.PrimaryColor)
	assert.Equal(t, "secondaryColor2", tag.SecondaryColor)
	assert.False(t, tag.Enabled)
}

func TestSearchTags(t *testing.T) {
	cfg.Setup("testing")
	db.Setup(cfg)
	defer db.Close()

	tag := &Tag{}
	for i := 15; i <= 25; i++ {
		tag = &Tag{
			Name:           "name" + utils.ToString(i),
			PrimaryColor:   "primaryColor" + utils.ToString(i),
			SecondaryColor: "secondaryColor" + utils.ToString(i),
		}
		tag, _ = tag.Create()
	}

	params := &SearchParameters{
		Filter:        "name1",
		SortField:     "id",
		SortDirection: "DESC",
		Limit:         3,
		Offset:        1,
	}

	tags, err := tag.Search(params)
	assert.NoError(t, err)

	assert.Equal(t, params.Limit, len(tags))
	assert.NotZero(t, tags[0].ID)
	assert.Equal(t, "name18", tags[0].Name)
	assert.Equal(t, "primaryColor18", tags[0].PrimaryColor)
	assert.Equal(t, "secondaryColor18", tags[0].SecondaryColor)
	assert.True(t, tags[0].Enabled)
}

//used for testing entities that use tags
func createTestingTags(n int) []*Tag {
	tags := []*Tag{}
	for i := 1; i <= n; i++ {
		tag := &Tag{
			Name:           "name" + utils.ToString(i),
			PrimaryColor:   "primaryColor" + utils.ToString(i),
			SecondaryColor: "secondaryColor" + utils.ToString(i),
		}
		tag, _ = tag.Create()
		tags = append(tags, tag)
	}

	return tags
}
