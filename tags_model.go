package main

import (
	"fmt"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
)

func (tag *Tag) Create() (*Tag, error) {
	result, err := db.DB.Exec(`INSERT INTO tags (name, primaryColor, secondaryColor, public) VALUES (?, ?, ?, ?)`, tag.Name, tag.PrimaryColor, tag.SecondaryColor, tag.Public)
	if err != nil {
		return &Tag{}, err
	}

	tag.ID = utils.GetEntryID(result)

	return tag.Get()
}

func (tag *Tag) Get() (*Tag, error) {
	if err := db.DB.QueryRow(`SELECT name, primaryColor, secondaryColor, public, enabled FROM tags WHERE id = ?`, tag.ID).Scan(
		&tag.Name, &tag.PrimaryColor, &tag.SecondaryColor, &tag.Public, &tag.Enabled); err != nil {
		return &Tag{}, err
	}

	return tag, nil
}

func (tag *Tag) Update() (*Tag, error) {
	_, err := db.DB.Exec(`UPDATE tags SET name = ?, primaryColor = ?, secondaryColor = ?, public = ?, enabled = ? WHERE id = ?`,
		tag.Name, tag.PrimaryColor, tag.SecondaryColor, tag.Public, tag.Enabled, tag.ID)
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

		//shows only public tags if the param is set
		if !params.ShowPrivate && !tempTag.Public {
			continue
		}

		tags = append(tags, tempTag)
	}

	return tags, nil
}

func (tag *Tag) GetTasks() ([]*Task, error) {
	rows, err := db.DB.Query(`SELECT idTask FROM tasks_tags WHERE idTag = ?`, tag.ID)
	defer rows.Close()
	if err != nil {
		return []*Task{}, err
	}

	tasks := []*Task{}
	for rows.Next() {
		tempTask := &Task{}
		if err = rows.Scan(&tempTask.ID); err != nil {
			return []*Task{}, err
		}

		tempTask, err = tempTask.Get()
		if err != nil {
			return []*Task{}, err
		}

		tasks = append(tasks, tempTask)
	}

	return tasks, nil
}
