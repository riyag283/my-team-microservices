package services

import (
	"database/sql"
	"fmt"
	"teams/db"
	"teams/graph/model"
)

func GetTeamMemberService(id string) (*model.TeamMember, error) {
	var teamMember model.TeamMember

	err := db.DBClientInstance.QueryRow("SELECT id, name, role, city FROM team_members WHERE id = $1", id).Scan(&teamMember.ID, &teamMember.Name, &teamMember.Role, &teamMember.City)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("team member not found")
		}
		return nil, err
	}

	return &teamMember, nil
}
