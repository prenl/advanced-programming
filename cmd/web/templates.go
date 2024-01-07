package main

import "snippetbox.yelnurabdrakhmanov.net/internal/models"

type templateData struct {
	Snippet *models.Snippet
	Snippets []*models.Snippet
}
