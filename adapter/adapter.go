package adapter

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Adapter struct {
	Postgres *gorm.DB
}

type Option interface {
	Start(a *Adapter) error
	Stop() error
}

var Adapters = &Adapter{}
var Options []Option

func (a *Adapter) Sync(opts ...Option) error {
	var syncErrors []error

	for _, opt := range opts {
		if err := opt.Start(a); err != nil {
			syncErrors = append(syncErrors, err)
		}

		Options = append(Options, opt)
	}

	if len(syncErrors) > 0 {
		return errors.New(fmt.Sprintf("sync adapter errors: %v", syncErrors))
	}

	return nil
}

func (a *Adapter) Unsync() error {
	var unsyncErrors []error

	for _, opt := range Options {
		if err := opt.Stop(); err != nil {
			unsyncErrors = append(unsyncErrors, err)
		}
	}

	if len(unsyncErrors) > 0 {
		return errors.New(fmt.Sprintf("unsync adapter errors: %v", unsyncErrors))
	}

	return nil

}
