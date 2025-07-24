package main

import "github.com/akmalrizaev/gists/internal/models"

type templateData struct {
	Gist  models.Gist
	Gists []models.Gist
}
