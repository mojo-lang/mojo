package openapi

import (
	"sort"
	"strings"

	"github.com/daveshanley/vacuum/model"
	"github.com/daveshanley/vacuum/motor"
	"github.com/daveshanley/vacuum/rulesets"
)

type LintOperations []*LintOperation

func (c LintOperations) Len() int {
	return len(c)
}

func (c LintOperations) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c LintOperations) Less(i, j int) bool {
	return strings.Compare(c[i].Path, c[j].Path) < 0
}

var (
	// ref: https://github.com/daveshanley/vacuum/blob/main/rulesets/rulesets.go
	lintRules = &rulesets.RuleSet{
		Rules: map[string]*model.Rule{
			"operation-parameters": rulesets.GetOperationParametersRule(),
		},
	}
)

func EasyLint(spec []byte) *LintResult {
	result := motor.ApplyRulesToRuleSet(&motor.RuleSetExecution{
		RuleSet:         lintRules,
		Spec:            spec,
		CustomFunctions: map[string]model.RuleFunction{},
	})

	if result.Index == nil {
		return &LintResult{Valid: false}
	}

	operations := result.Index.GetAllPaths()

	lintResult := &LintResult{
		Operations: make([]*LintOperation, 0, len(operations)),
		Valid:      true,
	}

	for path, operation := range operations {
		for method := range operation {
			operationResult := &LintOperation{
				Path:   path,
				Method: method,
				Valid:  true,
			}

			operationPath := strings.Join([]string{"$.paths", path, method, "parameters"}, ".")
			for _, r := range result.Results {
				if r.Path != operationPath {
					continue
				}

				operationResult.Valid = false
				operationResult.Description = r.Rule.Description
				operationResult.HowToFix = r.Rule.HowToFix
				if r.StartNode != nil {
					operationResult.Line = int32(r.StartNode.Line)
				}

				lintResult.Valid = false
				break
			}

			lintResult.Operations = append(lintResult.Operations, operationResult)
		}
	}

	sort.Sort(LintOperations(lintResult.Operations))
	return lintResult
}
