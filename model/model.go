package model

import "go-playground/apps/api/chatgpt"

type ListUsers map[string][]chatgpt.ChatCompletionMessage

var (
	Users = make(ListUsers)
)
