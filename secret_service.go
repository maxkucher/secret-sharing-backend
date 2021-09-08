package main

import (
	"crypto/md5"
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

type SecretService interface {
	SaveSecret(plainTextSecret string) (string, error)
	LoadSecrets(secretId string) (string, error)
}

type FSSecretService struct {
	Mutex    sync.Mutex
	FilePath string
}

func (s *FSSecretService) SaveSecret(plainTextSecret string) (string, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	secrets, err := s.loadSecretsFromFile()
	if err != nil {
		return "", err
	}
	idBytes := md5.Sum([]byte(plainTextSecret))
	id := string(idBytes[:])
	secrets[id] = plainTextSecret
	err = s.writeSecretsToFile(secrets)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *FSSecretService) LoadSecrets(secretId string) (string, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	secrets, err := s.loadSecretsFromFile()
	if err != nil {
		return "", err
	}
	secret := secrets[secretId]
	if secret != "" {
		delete(secrets, secretId)
	}
	err = s.writeSecretsToFile(secrets)
	if err != nil {
		return "", err
	}
	return secret, nil
}

func (s *FSSecretService) writeSecretsToFile(secrets map[string]string) error {
	secretBytes, err := json.Marshal(secrets)
	file, err := os.Create(s.FilePath)
	_, err = file.Write(secretBytes)
	if err != nil {
		return err
	}
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func (s *FSSecretService) loadSecretsFromFile() (map[string]string, error) {
	file, err := os.Open(s.FilePath)
	if err != nil {
		return nil, err
	}
	fileDataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var secrets map[string]string
	err = json.Unmarshal(fileDataBytes, &secrets)
	if err != nil {
		return nil, err
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	return secrets, nil
}
