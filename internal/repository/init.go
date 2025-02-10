package repository

import (
	"log"

	"github.com/mnsv1511/kasikorn-line-assignment/ent"
)

type Repository interface {
}

type RepositoryImpl struct {
	clientDB *ent.Client
}

func NewRepository() (Repository, error) {
	clientDB, err := InitDatabase()
	if err != nil {
		return nil, err
	}

	dataExists, err := CheckDataExists(clientDB)
	if err != nil {
		log.Fatalf("failed checking data exists: %v", err)
	}

	if dataExists == false {
		log.Println("Mocking up data...")
		MockUp()
		log.Println("Mock up data successfully.")
	}

	return &RepositoryImpl{
		clientDB: clientDB,
	}, nil
}
