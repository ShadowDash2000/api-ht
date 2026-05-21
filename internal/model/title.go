package model

import (
	"errors"
	"unicode/utf8"
)

type Title string

type TitleConfig struct {
	MinLength uint
	MaxLength uint
	Error     error
}

func NewTitle(value string, cfg TitleConfig) (Title, error) {
	if cfg.Error == nil {
		return "", errors.New("title error must be set")
	}

	valueLen := uint(utf8.RuneCountInString(value))
	if cfg.MinLength > 0 && valueLen < cfg.MinLength {
		return "", cfg.Error
	}

	if cfg.MaxLength > 0 && valueLen > cfg.MaxLength {
		return "", cfg.Error
	}

	return Title(value), nil
}
