// Code generated by candi v1.6.4.

package usecase

import (
	"sync"

	// @candi:usecaseImport
	contactusecase "be-shark/internal/modules/contact/usecase"
	accountusecase "be-shark/internal/modules/account/usecase"
	"be-shark/pkg/shared/usecase/common"

	"pkg.agungdp.dev/candi/codebase/factory/dependency"
)

type (
	// Usecase unit of work for all usecase in modules
	Usecase interface {
		// @candi:usecaseMethod
		Contact() contactusecase.ContactUsecase
		Account() accountusecase.AccountUsecase
	}

	usecaseUow struct {
		// @candi:usecaseField
		contactusecase.ContactUsecase
		accountusecase.AccountUsecase
	}
)

var usecaseInst *usecaseUow
var once sync.Once

// SetSharedUsecase set singleton usecase unit of work instance
func SetSharedUsecase(deps dependency.Dependency) {
	once.Do(func() {
		usecaseInst = new(usecaseUow)
		var setSharedUsecaseFuncs []func(common.Usecase)
		var setSharedUsecaseFunc func(common.Usecase)

		// @candi:usecaseCommon
		usecaseInst.ContactUsecase, setSharedUsecaseFunc = contactusecase.NewContactUsecase(deps)
		setSharedUsecaseFuncs = append(setSharedUsecaseFuncs, setSharedUsecaseFunc)
		usecaseInst.AccountUsecase, setSharedUsecaseFunc = accountusecase.NewAccountUsecase(deps)
		setSharedUsecaseFuncs = append(setSharedUsecaseFuncs, setSharedUsecaseFunc)

		sharedUsecase := common.SetCommonUsecase(usecaseInst)
		for _, setFunc := range setSharedUsecaseFuncs {
			setFunc(sharedUsecase)
		}
	})
}

// GetSharedUsecase get usecase unit of work instance
func GetSharedUsecase() Usecase {
	return usecaseInst
}

// @candi:usecaseImplementation
func (uc *usecaseUow) Contact() contactusecase.ContactUsecase {
	return uc.ContactUsecase
}

func (uc *usecaseUow) Account() accountusecase.AccountUsecase {
	return uc.AccountUsecase
}
