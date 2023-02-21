package services

import (
	"database/sql"
	"fmt"
	"teams/db"
	"teams/graph/model"
)

type servicesInterface interface {
	GetTeamMemberService(id string) (*model.TeamMember, error)
	GetTeamMembersService() ([]*model.TeamMember, error)
	CreateTeamMemberService(requestBody model.NewTeamMember) (model.TeamMember, error)
	RemoveTeamMemberService(teamMemberID string) (model.TeamMember, error)
	UpdateTeamMemberService(requestBody model.UpdateTeamMember) (model.TeamMember, error)
}

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

// GetTeamMembersService retrieves all team members from the database
func GetTeamMembersService() ([]*model.TeamMember, error) {
	teamMembers := make([]*model.TeamMember, 0)
	rows, err := db.DBClientInstance.Query("SELECT id, name, role, city FROM team_members")
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

func CreateTeamMemberService(requestBody model.NewTeamMember) (model.TeamMember, error) {
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

func RemoveTeamMemberService(teamMemberID string) (model.TeamMember, error) {
	var teamMember model.TeamMember
	err := db.DBClientInstance.QueryRow("DELETE FROM team_members WHERE id=$1 RETURNING id, name, role, city", teamMemberID).Scan(&teamMember.ID, &teamMember.Name, &teamMember.Role, &teamMember.City)
	if err != nil {
		return teamMember, err
	}

	return teamMember, nil
}

func UpdateTeamMemberService(requestBody model.UpdateTeamMember) (model.TeamMember, error) {
    var teamMember model.TeamMember
    
    err := db.DBClientInstance.QueryRowx("UPDATE team_members SET name=$2, role=$3, city=$4 WHERE id=$1 RETURNING id, name, role, city",
        requestBody.ID,
        requestBody.Name,
        requestBody.Role,
        requestBody.City,
    ).StructScan(&teamMember)
    
    if err != nil {
        return teamMember, err
    }
    
    return teamMember, nil
}