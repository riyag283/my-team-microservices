package services

import (
	"teams/db"
	"teams/graph/model"
)

func CreateTeamMemberService (requestBody model.NewTeamMember) (model.TeamMember, error) {
	var teamMember model.TeamMember

	err := db.DBClientInstance.QueryRow("INSERT INTO team_members (name, role, city) VALUES ($1, $2, $3) RETURNING id, name, role, city",
		requestBody.Name,
		requestBody.Role,
		requestBody.City,
	).Scan(&teamMember.ID, &teamMember.Name, &teamMember.Role, &teamMember.City)

	if err != nil {
		return teamMember, err
	}

	return teamMember, nil
}