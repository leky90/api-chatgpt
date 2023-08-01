package model

import "go-playground/chatgpt"

type ListUsers map[string][]chatgpt.ChatCompletionMessage

var (
	Users = make(ListUsers)
)
