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
	fmt.Printf(`
*************************** no.1 ***************************`)

	for _, project := range projects {
		fmt.Printf(`
	#           : %d
	Name        : %s
			`,
			project.ID,
			project.Name)
	}
}

func showDetProject(project Project) {
	fmt.Printf(`
*************************** no.%d ***************************
	#           : %d
	Name        : %s
	Identifier  : %s
	Status      : %d
	CreatedOn   : %s
	UpdatedOn   : %s
	Description :

%s
		`,
		1,
		project.ID,
		project.Name,
		project.Identifier,
		project.Status,
		project.CreatedOn,
		project.UpdatedOn,
		project.Description)
}

func showMemberships(memberships Memberships) {
	for i, membership := range memberships {
		var role string
		if len(membership.Roles) >= 1 {
			role = membership.Roles[0].Name
		}

		fmt.Printf(`
*************************** no.%d ***************************
	#           : %d
	Name        : %s
	Role        : %s
		    `,
			i+1,
			membership.ID,
			membership.User.Name,
			role)
	}
}
