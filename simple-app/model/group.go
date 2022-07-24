package model

type Group struct {
	Title     string
	ShortDesc string
	FullDesc  string

	Id int
}

func Groups() ([]Group, error) {
	rows, err := db.Query("SELECT id, title, short_desc, full_desc FROM groups")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []Group
	for rows.Next() {
		var group Group
		err := rows.Scan(&group.Id, &group.Title, &group.ShortDesc, &group.FullDesc)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return groups, nil
}
func RemoveGroupById(id int) error {
	sqlQueue := "DELETE FROM groups WHERE id=?"

	_, err := db.Exec(sqlQueue, id)
	return err
}
func UpdateGroup(g Group) error {
	sqlQueue := `
		UPDATE groups
		SET title=?, short_desc=?, full_desc=?
		WHERE id=?`

	_, err := db.Exec(sqlQueue, g.Title, g.ShortDesc, g.FullDesc, g.Id)
	return err
}
func AddGroup(g Group) error {
	sqlQueue := `
		INSERT INTO groups (title, short_desc, full_desc)
		VALUES (?, ?, ?)`

	_, err := db.Exec(sqlQueue, g.Title, g.ShortDesc, g.FullDesc)
	return err
}
