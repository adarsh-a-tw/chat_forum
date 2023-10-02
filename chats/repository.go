package chats

import "github.com/jmoiron/sqlx"

type MesageRepository struct {
	db *sqlx.DB
}

func NewMessageRepository(db *sqlx.DB) *MesageRepository {
	return &MesageRepository{db}
}

func (mr *MesageRepository) GetAllById(roomId string) ([]Message, error) {
	query := `
        SELECT id, room_id, content, created_at
        FROM messages
        WHERE room_id = $1
        ORDER BY created_at DESC
    `
	var messages []Message
	err := mr.db.Select(&messages, query, roomId)
	if err != nil {
		return nil, err
	}
	return messages, nil
}

func (mr *MesageRepository) Create(roomId string, content string) error {
	query := `
        INSERT INTO messages (room_id, content)
        VALUES ($1, $2)
        RETURNING id
    `
	_, err := mr.db.Exec(query, roomId, content)
	return err
}
