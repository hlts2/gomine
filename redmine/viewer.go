package redmine

import (
	"fmt"
)

func showIssues(issues Issues) {
	for i, v := range issues {
		fmt.Printf(`
*************************** no.%d ***************************
	#         : %d
	Project   : %s
	Tracker   : %s
	Status    : %s
	Priority  : %s
	Subject   : %s
	Assigned  : %s
	CreatedOn : %s
	DueDate   : %s
	DoneRatio : %d
			`,
			i+1,
			v.ID,
			v.Project.Name,
			v.Tracker.Name,
			v.Status.Name,
			v.Subject,
			v.AssignedTo.Name,
			v.CreatedOn,
			v.DueDate,
			v.DoneRatio)
	}
}

func showDetIssue(issue Issue) {
	fmt.Printf(`
*************************** no.%d ***************************
	#           : %d
	Project     : %s
	Tracker     : %s
	Status      : %s
	Priority    : %s
	Subject     : %s
	Assigned    : %s
	CreatedOn   : %s
	DueDate     : %s
	DoneRatio   : %d
	Description :

%s
		`,
		1,
		issue.ID,
		issue.Project.Name,
		issue.Tracker.Name,
		issue.Status.Name,
		issue.Priority.Name,
		issue.Subject,
		issue.AssignedTo.Name,
		issue.CreatedOn,
		issue.DueDate,
		issue.DoneRatio,
		issue.Description)
}
