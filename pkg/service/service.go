package service

import (
	"cmd/pkg/repository"
	"cmd/pkg/repository/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	CheckUsername(username string) error
}

type Chat interface {
	Create(chat models.Chat) (int, error)
	Delete(chatId int) error
	Get(chatId int) (models.Chat, error)
	AddUser(users models.ChatUsers) (int, error)
	DeleteUser(userId, chatId int) error
	GetUsers(chatId int) ([]models.User, error)
	GetUserChats(userId int) ([]models.Chat, error)
}

type Status interface {
	AddStatus(status models.Status) (int, error)
	GetStatus(senderId, recipientId int) (models.Status, error)
	DeleteStatus(status models.Status) error
	UpdateStatus(status models.Status) error
	GetFriends(userId int) ([]int, error)
	GetBlackList(userId int) ([]int, error)
	GetBlackListToUser(userId int) ([]int, error)
	GetSentInvites(userId int) ([]int, error)
	GetInvites(userId int) ([]int, error)
}

type Message interface {
	Create(msg models.Message) (int, error)
	Get(msgId int) (models.Message, error)
	GetAll(chatId int) ([]models.Message, error)
}

type Service struct {
	Authorization
	Chat
	Status
	Message
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Chat:          NewChatService(repos.Chat),
		Status:        NewStatusService(repos.Status),
		Message:       NewMessageService(repos.Message),
	}
}
