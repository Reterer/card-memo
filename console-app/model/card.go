package model

type Card struct {
	Title     string
	ShortDesc string
	FullDesc  string

	Id       int
	LearnVal float64
}

func CardsOfGroup(groupId int) ([]Card, error) {
	sqlQueue := `
		SELECT id, learn_val, title, short_dec, full_desc 
		FROM cards
		WHERE group_id=?`

	rows, err := db.Query(sqlQueue, groupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cards []Card
	for rows.Next() {
		var card Card
		err := rows.Scan(&card.Id, &card.LearnVal, &card.Title, &card.ShortDesc, &card.FullDesc)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cards, nil
}
func RemoveCardById(id int) error {
	sqlQueue := "DELETE FROM cards WHERE id=?"

	_, err := db.Exec(sqlQueue, id)
	return err
}
func UpdateCard(c Card) error {
	sqlQueue := `
		UPDATE cards
		SET learn_val=?, title=?, short_desc=?, full_desc=?
		WHERE id=?`

	_, err := db.Exec(sqlQueue, c.LearnVal, c.Title, c.ShortDesc, c.FullDesc, c.Id)
	return err
}
