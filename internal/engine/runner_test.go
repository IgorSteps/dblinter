package engine_test

import (
	"errors"
	"testing"

	"github.com/IgorSteps/dblinter/internal/diagnostics"
	"github.com/IgorSteps/dblinter/internal/engine"
	"github.com/IgorSteps/dblinter/internal/rules"
	mock_rules "github.com/IgorSteps/dblinter/mocks/rules"
	"github.com/stretchr/testify/assert"
)

func TestRunner_RunHappyPath(t *testing.T) {
	// Assemble
	mockRule := mock_rules.NewMockRule(t)
	testCallSites := []rules.CallSite{}
	testRunner := engine.New([]rules.Rule{mockRule})
	expectedIssues := []diagnostics.Issue{
		{
			RuleID:  "test",
			Message: "test",
			Doc:     "test",
		},
	}
	mockRule.EXPECT().Check(testCallSites).Return(expectedIssues, nil).Once()

	// Act
	actualIssues, errs := testRunner.Run(testCallSites)

	// Assert
	assert.Len(t, errs, 0)
	assert.Equal(t, expectedIssues, actualIssues)
}

func TestRunner_RunWithErrors(t *testing.T) {
	// Assemble
	testCallSite := []rules.CallSite{}
	emptyIssues := []diagnostics.Issue(nil)
	firstMockRule := mock_rules.NewMockRule(t)
	firstError := errors.New("boom 1")
	firstMockRule.EXPECT().Check(testCallSite).Return(emptyIssues, firstError).Once()
	secondError := errors.New("boom 2")
	secondMockRule := mock_rules.NewMockRule(t)
	secondMockRule.EXPECT().Check(testCallSite).Return(emptyIssues, secondError).Once()
	expectedErrors := []error{firstError, secondError}
	testRunner := engine.New([]rules.Rule{firstMockRule, secondMockRule})

	// Act
	actualIssues, actualErrors := testRunner.Run(testCallSite)

	// Assert
	assert.Equal(t, emptyIssues, actualIssues)
	assert.Equal(t, expectedErrors, actualErrors)
}
