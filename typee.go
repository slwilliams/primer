package main

import "github.com/google/uuid"

// ServiceResponse holds Prime
type ServiceResponse struct {
	ID    string      `json:"id"`
	Prime PrimeObject `json:"prime"`
	Host  string      `json:"host"`
	Ts    string      `json:"ts"`
	Envs  []string    `json:"envs"`
}

// PrimeObject holds prime response context
type PrimeObject struct {
	Max   int `json:"max"`
	Value int `json:"val"`
}

func getUUIDv4() string {
	id, err := uuid.NewRandom()
	if err != nil {
		logger.Fatalf("Error while getting id: %v\n", err)
	}
	return id.String()
}
