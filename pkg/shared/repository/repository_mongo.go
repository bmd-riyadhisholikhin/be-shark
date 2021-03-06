// Code generated by candi v1.6.4. DO NOT EDIT.

package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	// @candi:repositoryImport
	contactrepo "be-shark/internal/modules/contact/repository"
	accountrepo "be-shark/internal/modules/account/repository"
)

type (
	// RepoMongo abstraction
	RepoMongo interface {
		// @candi:repositoryMethod
		ContactRepo() contactrepo.ContactRepository
		AccountRepo() accountrepo.AccountRepository
	}

	repoMongoImpl struct {
		readDB, writeDB *mongo.Database

		// register all repository from modules
		// @candi:repositoryField
		contactRepo contactrepo.ContactRepository
		accountRepo accountrepo.AccountRepository
	}
)

var globalRepoMongo RepoMongo

// setSharedRepoMongo set the global singleton "RepoMongo" implementation
func setSharedRepoMongo(readDB, writeDB *mongo.Database) {
	globalRepoMongo = &repoMongoImpl{
		readDB: readDB, writeDB: writeDB,

		// @candi:repositoryConstructor
	}
}

// GetSharedRepoMongo returns the global singleton "RepoMongo" implementation
func GetSharedRepoMongo() RepoMongo {
	return globalRepoMongo
}

// @candi:repositoryImplementation
func (r *repoMongoImpl) ContactRepo() contactrepo.ContactRepository {
	return r.contactRepo
}

func (r *repoMongoImpl) AccountRepo() accountrepo.AccountRepository {
	return r.accountRepo
}

