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
			v.Priority.Name,
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

func showProjects(projects Projects) {
	for i, project := range projects {
		fmt.Printf(`
*************************** no.%d ***************************
	#           : %d
	Name        : %s
			`,
			i+1,
			project.ID,
			project.Name)
	}
}

func showDetProject(project Project) {
	fmt.Printf(`
*************************** no.%d ***************************
	#           : %s
	Name        : %s
	Identifier  : %d
	Status      : %d
	CreatedOn   : %s
	UpdatedOn   : %s
	Description :

%s
		`,
		1,
		project.Name,
		project.Identifier,
		project.Status,
		project.CreatedOn,
		project.UpdatedOn,
		project.Description)
}
