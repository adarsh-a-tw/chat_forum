package chats

import "time"

type Message struct {
	Id        int       `db:"id"`
	RoomId    string    `db:"room_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}
