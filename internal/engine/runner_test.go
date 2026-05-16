package engine_test

import (
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
