package main

import (
	"kafka"
)

type Post struct {
	UID     string `json:"uid" gorm:"primary"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug" gorm:"uniqueIndex`
}

type NewPostMessage struct {
	UID     string `json:"uid"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PublishedPostMessage struct {
	Post
}

type Publisher struct {
	newPostReader       *kafka.Reader
	publishedPostWriter *kafka.Writer
	db                  *gorm.DB
}
