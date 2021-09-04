package main

import "crypto/md5"

type SecretService interface {
	SaveSecret(plainTextSecret string) string
	LoadSecrets(secretId string) string
}

type FSSecretService struct{}

func (s FSSecretService) SaveSecret(plainTextSecret string) string {
	id := md5.Sum([]byte(plainTextSecret))
	return string(id[:])
}

func (s FSSecretService) LoadSecrets(secretIds string) string {
	return ""
}
