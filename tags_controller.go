package main

import (
	"net/http"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
	"github.com/gin-gonic/gin"
)

func CreateTag(c *gin.Context) {
	var tag *Tag
	c.BindJSON(&tag)

	if tag.Name == "" || tag.PrimaryColor == "" || tag.SecondaryColor == "" {
		c.JSON(http.StatusBadRequest, "all fields required")
		return
	}

	tag, err := tag.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, tag)
}

func GetTag(c *gin.Context) {
	tag := &Tag{ID: utils.ToInt(c.Param("id_tag"))}

	tag, err := tag.Get()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, tag)
}

func UpdateTag(c *gin.Context) {
	var tag *Tag
	c.BindJSON(&tag)
	tag.ID = utils.ToInt(c.Param("id_tag"))

	if tag.Name == "" || tag.PrimaryColor == "" || tag.SecondaryColor == "" {
		c.JSON(http.StatusBadRequest, "all fields required")
		return
	}

	tag, err := tag.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, tag)
}

func SearchTags(c *gin.Context) {
	tag := &Tag{}

	params := &SearchParameters{
		Filter:        c.Query("filter"),
		SortField:     c.Query("sortField"),
		SortDirection: c.Query("sortDirection"),
		Limit:         utils.ToInt(c.Query("limit")),
		Offset:        utils.ToInt(c.Query("offset")),
	}

	tags, err := tag.Search(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, db.BeautifyError(err))
		return
	}

	c.JSON(http.StatusOK, tags)
}
