package services

import (
	"teams/db"
	"teams/graph/model"
)

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
