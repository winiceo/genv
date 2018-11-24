package cmd

import (
	"github.com/winiceo/genv/internal/config"
	"github.com/winiceo/genv/internal/db"
	"github.com/winiceo/genv/pkg/container"
	"github.com/google/uuid"
)

type memStore struct {
	env db.Environment
}

func (s *memStore) Create(e db.Environment) error {
	s.env = e

	return nil
}

func (s *memStore) Read() (db.Environment, error) {
	return s.env, nil
}

func (s *memStore) Delete() error {
	s.env = db.Environment{}
	return nil
}

type mockCtl struct {
	current *container.Metadata

	// These allow the specific tests to override the underlying behavior if
	// necessary to test alternative code-paths.
	createFn func(container.Metadata) (container.Metadata, error)
	removeFn func(container.Metadata) error
	attachFn func(container.Metadata) error
	runFn    func(container.Metadata, []string) error
}

func newMockCtl(init *container.Metadata) *mockCtl {
	ctl := &mockCtl{
		current: init,
	}

	ctl.createFn = func(m container.Metadata) (container.Metadata, error) {
		ctl.current = &m

		if ctl.current.ID == "" {
			ctl.current.ID = uuid.New().String()
		}

		if ctl.current.ImageID == "" {
			ctl.current.ImageID = uuid.New().String()
		}

		return *ctl.current, nil
	}

	ctl.removeFn = func(m container.Metadata) error {
		ctl.current = nil

		return nil
	}

	ctl.attachFn = func(m container.Metadata) error {
		return nil
	}

	ctl.runFn = func(m container.Metadata, cmds []string) error {
		return nil
	}

	return ctl
}

func (ctl *mockCtl) Create(m container.Metadata) (container.Metadata, error) {
	return ctl.createFn(m)
}

func (ctl *mockCtl) Remove(m container.Metadata) error {
	return ctl.removeFn(m)
}

func (ctl *mockCtl) Attach(m container.Metadata) error {
	return ctl.attachFn(m)
}

func (ctl *mockCtl) Run(m container.Metadata, cmds []string) error {
	return ctl.runFn(m, cmds)
}

type memConfig struct {
	opts config.Opts
}

func (c memConfig) Load() (config.Opts, error) {
	if c.opts.CacheImage == nil {
		c.opts.CacheImage = config.CacheImage
	}

	return c.opts, nil
}
