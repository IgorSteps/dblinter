package koanfconfig_test

import (
	"fmt"
	"testing"

	"github.com/IgorSteps/dblinter/internal/adapters/koanfconfig"
	"github.com/stretchr/testify/assert"
)

func TestMaxOpenConnsConfigToDomain_HappyPath(t *testing.T) {
	// Assemble
	expectedEnabled := true
	expectedRequired := 10
	maxOpenConnsConfig := koanfconfig.MaxOpenConnsConfig{
		Enabled:  &expectedEnabled,
		Required: &expectedRequired,
	}

	// Act
	domainCfg, err := maxOpenConnsConfig.ToDomain()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedEnabled, domainCfg.Enabled)
	assert.Equal(t, expectedRequired, domainCfg.Required)
}

func TestMaxOpenConnsConfigToDomain_NilOption(t *testing.T) {
	// Assemble
	expectedEnabled := true
	expectedRequired := 10
	type testCase struct {
		name          string
		enabled       *bool
		required      *int
		expectedError error
	}

	testCases := []testCase{
		{
			name:          "enabled is nil",
			enabled:       nil,
			required:      &expectedRequired,
			expectedError: fmt.Errorf("max_open_conns: enabled cannot be nil"),
		},
		{
			name:          "required is nil",
			enabled:       &expectedEnabled,
			required:      nil,
			expectedError: fmt.Errorf("max_open_conns: required cannot be nil"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			maxOpenConnsConfig := koanfconfig.MaxOpenConnsConfig{
				Enabled:  testCase.enabled,
				Required: testCase.required,
			}

			// Act
			domainCfg, err := maxOpenConnsConfig.ToDomain()

			// Assert
			assert.EqualError(t, err, testCase.expectedError.Error())
			assert.Empty(t, domainCfg)
		})
	}
}
