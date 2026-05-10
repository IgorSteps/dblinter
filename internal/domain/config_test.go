package domain_test

import (
	"testing"

	"github.com/IgorSteps/dblinter/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewMaxOpenConnsConfig_ErrorsOnInvalidParam(t *testing.T) {
	// Assemble
	invalidRequired := -1

	// Act
	cfg, err := domain.NewMaxOpenConnsConfig(true, invalidRequired)

	// Assert
	assert.Empty(t, cfg)
	assert.Error(t, err)
}
