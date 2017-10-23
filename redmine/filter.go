package redmine

import (
	"strings"
)

func FilterIssues(issues Issues, filters []string) Issues {
	var objs Issues
	for _, issue := range issues {
		isExist := true
		for _, filter := range filters {
			flgSub := strings.Contains(issue.Subject, filter)
			flgDesc := strings.Contains(issue.Description, filter)

			isExist = isExist && (flgSub || flgDesc)
		}

		if isExist {
			objs = append(objs, issue)
		}
	}

	return objs
}
