package redmine

import (
	"strings"
)

func filterIssues(issues Issues, filters []string) Issues {
	var objs Issues
	for _, issue := range issues {
		isExist := true
		for _, filter := range filters {
			isExist = isExist && strings.Contains(filter, issue.Subject) || strings.Contains(filter, issue.Description)
		}

		if isExist {
			objs = append(objs, issue)
		}
	}

	return objs
}
