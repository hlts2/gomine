package redmine

import (
	"fmt"
)

func viewer(issues Issues) {
	for i, v := range issues {
		fmt.Printf("*************************** no.%d ***************************\n", i+1)
		fmt.Printf("  #           : %d\n", v.ID)
		fmt.Printf("  プロジェクト: %s\n", v.Project.Name)
		fmt.Printf("  トラッカー  : %s\n", v.Tracker.Name)
		fmt.Printf("  ステータス  : %s\n", v.Status.Name)
		fmt.Printf("  優先度      : %s\n", v.Priority.Name)
		fmt.Printf("  題名        : %s\n", v.Subject)
		fmt.Printf("  担当者      : %s\n", v.AssignedTo.Name)
		fmt.Printf("  更新日      : %s\n", v.CreatedOn)
		fmt.Printf("  期限        : %s\n", v.DueDate)
		fmt.Printf("  予定工数    : %d\n\n", v.DoneRatio)
	}
}
