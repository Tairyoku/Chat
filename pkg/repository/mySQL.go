package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	UsersTable        = "users"
	StatusesTable     = "users_relationship"
	ChatMessagesTable = "chat_messages"
	ChatUsersList     = "chat_users"
	ChatsTable        = "chats"
	MessagesTable     = "messages"
	StatusFriends     = "friends"
	StatusBL          = "black_list"
	StatusInvitation  = "invitation"
)

type Config struct {
	Username string
	Password string
	Host     string
	Url      string
	DBName   string
}

func NewRepositoryDB(cnf Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s@%s%s(%s)/%s", cnf.Username, cnf.Password, cnf.Host, cnf.Url, cnf.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
