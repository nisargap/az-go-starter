package models

type Config struct {
	DatabaseName string `json:"database_name" binding:"required"`
	DatabaseUrl  string `json:"database_url" binding:"required"`
	Hostname     string `json:"hostname" binding:"required"`
	Domain       string `json:"domain" binding:"required"`
	Port         string `json:"port" binding:"required"`
}
