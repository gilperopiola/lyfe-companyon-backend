package main

import (
	"fmt"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
)

func (tag *Tag) Create() (*Tag, error) {
	result, err := db.DB.Exec(`INSERT INTO tags (name, primaryColor, secondaryColor) VALUES (?, ?, ?)`, tag.Name, tag.PrimaryColor, tag.SecondaryColor)
	if err != nil {
		return &Tag{}, err
	}

	tag.ID = utils.GetEntryID(result)

	return tag.Get()
}

func (tag *Tag) Get() (*Tag, error) {
	if err := db.DB.QueryRow(`SELECT name, primaryColor, secondaryColor, enabled FROM tags WHERE id = ?`, tag.ID).Scan(
		&tag.Name, &tag.PrimaryColor, &tag.SecondaryColor, &tag.Enabled); err != nil {
		return &Tag{}, err
	}

	return tag, nil
}

func (tag *Tag) Update() (*Tag, error) {
	_, err := db.DB.Exec(`UPDATE tags SET name = ?, primaryColor = ?, secondaryColor = ?, enabled = ? WHERE id = ?`,
		tag.Name, tag.PrimaryColor, tag.SecondaryColor, tag.Enabled, tag.ID)
	if err != nil {
		return &Tag{}, err
	}

	return tag.Get()
}

func (tag *Tag) Search(params *SearchParameters) ([]*Tag, error) {
	query := fmt.Sprintf(`SELECT id FROM tags WHERE id LIKE ? OR name LIKE ? ORDER BY %s LIMIT ? OFFSET ?`, getSearchOrderBy(params))

	params.Filter = "%" + params.Filter + "%"
	rows, err := db.DB.Query(query, params.Filter, params.Filter, params.Limit, params.Offset)
	defer rows.Close()
	if err != nil {
		return []*Tag{}, err
	}

	tags := []*Tag{}
	for rows.Next() {
		tempTag := &Tag{}
		if err = rows.Scan(&tempTag.ID); err != nil {
			return []*Tag{}, err
		}

		tempTag, err = tempTag.Get()
		if err != nil {
			return []*Tag{}, err
		}

		tags = append(tags, tempTag)
	}

	return tags, nil
}
