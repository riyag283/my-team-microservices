package services

import (
	"teams/db"
	"teams/graph/model"
)

func RemoveTeamMemberService(teamMemberID string) (model.TeamMember, error) {
	var teamMember model.TeamMember
	err := db.DBClientInstance.QueryRow("DELETE FROM team_members WHERE id=$1 RETURNING id, name, role, city", teamMemberID).Scan(&teamMember.ID, &teamMember.Name, &teamMember.Role, &teamMember.City)
	if err != nil {
		return teamMember, err
	}

	return teamMember, nil
}
