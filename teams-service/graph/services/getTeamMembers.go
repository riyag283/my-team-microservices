package services

import (
	"teams/db"
	"teams/graph/model"
)

// GetTeamMembersService retrieves all team members from the database
func GetTeamMembersService() ([]*model.TeamMember, error) {
	teamMembers := make([]*model.TeamMember, 0)
	rows, err := db.DBClient.Query("SELECT id, name, role, city FROM team_members")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var teamMember model.TeamMember
		if err := rows.Scan(&teamMember.ID, &teamMember.Name, &teamMember.Role, &teamMember.City); err != nil {
			return nil, err
		}
		teamMembers = append(teamMembers, &teamMember)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return teamMembers, nil
}
