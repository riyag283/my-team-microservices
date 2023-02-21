package services

import (
	"database/sql"
	"fmt"
	"teams/graph/model"

	"github.com/jmoiron/sqlx"
)

func GetTeamMemberService(db *sqlx.DB, id string) (*model.TeamMember, error) {
	var teamMember model.TeamMember

	err := db.QueryRow("SELECT id, name, role, city FROM team_members WHERE id = $1", id).Scan(&teamMember.ID, &teamMember.Name, &teamMember.Role, &teamMember.City)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("team member not found")
		}
		return nil, err
	}

	return &teamMember, nil
}

// GetTeamMembersService retrieves all team members from the database
func GetTeamMembersService(db *sqlx.DB) ([]*model.TeamMember, error) {
	rows, err := db.Query("SELECT ID, Name, Role, City FROM team_members")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teamMembers := make([]*model.TeamMember, 0)
	for rows.Next() {
		var tm model.TeamMember
		err := rows.Scan(&tm.ID, &tm.Name, &tm.Role, &tm.City)
		if err != nil {
			return nil, err
		}
		teamMembers = append(teamMembers, &tm)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return teamMembers, nil
}

func CreateTeamMemberService(db *sqlx.DB, requestBody model.NewTeamMember) (model.TeamMember, error) {
	var teamMember model.TeamMember

	err := db.QueryRow("INSERT INTO team_members (name, role, city) VALUES ($1, $2, $3) RETURNING id, name, role, city",
		requestBody.Name,
		requestBody.Role,
		requestBody.City,
	).Scan(&teamMember.ID, &teamMember.Name, &teamMember.Role, &teamMember.City)

	if err != nil {
		return teamMember, err
	}

	return teamMember, nil
}

func RemoveTeamMemberService(db *sqlx.DB, teamMemberID string) (model.TeamMember, error) {
	var teamMember model.TeamMember
	err := db.QueryRow("DELETE FROM team_members WHERE id=$1 RETURNING id, name, role, city", teamMemberID).Scan(&teamMember.ID, &teamMember.Name, &teamMember.Role, &teamMember.City)
	if err != nil {
		return teamMember, err
	}

	return teamMember, nil
}

func UpdateTeamMemberService(db *sqlx.DB, requestBody model.UpdateTeamMember) (model.TeamMember, error) {
    var teamMember model.TeamMember
    
    err := db.QueryRowx("UPDATE team_members SET name=$2, role=$3, city=$4 WHERE id=$1 RETURNING id, name, role, city",
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