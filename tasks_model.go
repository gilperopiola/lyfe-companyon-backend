package main

import (
	"fmt"

	"github.com/gilperopiola/lyfe-companyon-backend/utils"
)

func (task *Task) Create() (*Task, error) {
	result, err := db.DB.Exec(`INSERT INTO tasks (name, description, importance, duration, dueDate) VALUES (?, ?, ?, ?, ?)`, task.Name, task.Description, task.Importance, task.Duration, task.DueDate)
	if err != nil {
		return &Task{}, err
	}

	task.ID = utils.GetEntryID(result)

	task.Tags, err = task.createTags()
	if err != nil {
		return &Task{}, err
	}

	return task.Get()
}

func (task *Task) Get() (*Task, error) {
	err := db.DB.QueryRow(`SELECT name, description, importance, status, duration, dueDate, dateCreated FROM tasks WHERE id = ?`, task.ID).Scan(&task.Name, &task.Description, &task.Importance, &task.Status, &task.Duration, &task.DueDate, &task.DateCreated)
	if err != nil {
		return &Task{}, err
	}

	task.Tags, err = task.getTags()
	if err != nil {
		return &Task{}, err
	}

	return task, nil
}

func (task *Task) Update() (*Task, error) {
	_, err := db.DB.Exec(`UPDATE tasks SET name = ?, description = ?, importance = ?, status = ?, duration = ?, dueDate = ? WHERE id = ?`, task.Name, task.Description, task.Importance, task.Status, task.Duration, task.DueDate, task.ID)
	if err != nil {
		return &Task{}, err
	}

	task.Tags, err = task.updateTags()
	if err != nil {
		return &Task{}, err
	}

	return task.Get()
}

func (task *Task) UpdateStatus() (*Task, error) {
	_, err := db.DB.Exec(`UPDATE tasks SET status = ? WHERE id = ?`, task.Status, task.ID)
	if err != nil {
		return &Task{}, err
	}

	return task.Get()
}

func (task *Task) Search(params *SearchParameters) ([]*Task, error) {
	query := fmt.Sprintf(`SELECT id FROM tasks WHERE (id LIKE ? OR name LIKE ?) AND status != %d AND status != %d AND importance >= ? ORDER BY %s LIMIT ? OFFSET ?`, Done, Archived, getSearchOrderBy(params))

	params.Filter = "%" + params.Filter + "%"
	rows, err := db.DB.Query(query, params.Filter, params.Filter, params.FilterImportance, params.Limit, params.Offset)
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

		hasRequiredTag := func(tags []*Tag, filterTagID int) bool {
			hasRequiredTag := true
			if filterTagID != 0 {
				hasRequiredTag = false

				for _, tag := range tags {
					if tag.ID == filterTagID {
						hasRequiredTag = true
						break
					}
				}
			}

			return hasRequiredTag
		}

		hasRequiredPrivacy := func(tags []*Tag, showPrivate bool) bool {
			if !showPrivate {
				for _, tag := range tags {
					if !tag.Public {
						return false
					}
				}
			}
			return true
		}

		if hasRequiredTag(tempTask.Tags, params.FilterTagID) && hasRequiredPrivacy(tempTask.Tags, params.ShowPrivate) {
			tasks = append(tasks, tempTask)
		}
	}

	return tasks, nil
}

//tasks_tags

func (task *Task) createTags() ([]*Tag, error) {
	tags := []*Tag{}

	for _, tag := range task.Tags {
		_, err := db.DB.Exec(`INSERT INTO tasks_tags (idTask, idTag) VALUES (?, ?)`, task.ID, tag.ID)
		if err != nil {
			return []*Tag{}, err
		}

		tempTag, _ := tag.Get()

		tags = append(tags, tempTag)
	}

	return tags, nil
}

func (task *Task) getTags() ([]*Tag, error) {
	rows, err := db.DB.Query(`SELECT idTag FROM tasks_tags WHERE idTask = ?`, task.ID)
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

func (task *Task) updateTags() ([]*Tag, error) {
	db.DB.Exec(`DELETE FROM tasks_tags WHERE idTask = ?`, task.ID)
	return task.createTags()
}
