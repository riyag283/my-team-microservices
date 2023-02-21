package services

import (
	"database/sql"
	"fmt"
	"teams/graph/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetTeamMembersService(t *testing.T) {
    // Create a new mock database and defer its closure
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    // Define the columns that will be returned by the mock rows
    columns := []string{"ID", "Name", "Role", "City"}

    // Define the rows that will be returned by the mock database
    rows := sqlmock.NewRows(columns).
        AddRow("1", "Riya Gupta", "Developer", "Ranchi").
        AddRow("2", "John Doe", "Manager", "New York")

    // Expect a query to be made to the database and return the mock rows
    mock.ExpectQuery("SELECT (.+) FROM team_members").WillReturnRows(rows)

    // Create an *sqlx.DB instance using the mock database
    sqlxDB := sqlx.NewDb(db, "sqlmock")

    // Call the function being tested
    teamMembers, err := GetTeamMembersService(sqlxDB)
    assert.NoError(t, err)

	teamMember1 := model.TeamMember{
		ID:   "1",
		Name: "Riya Gupta",
		Role: "Developer",
		City: "Ranchi",
	}
	teamMember2 := model.TeamMember{
		ID:   "2",
		Name: "John Doe",
		Role: "Manager",
		City: "New York",
	}

    // Verify that the correct data was returned
    expectedTeamMembers := []*model.TeamMember{
		&teamMember1,
		&teamMember2,
	}
	
    assert.Equal(t, len(expectedTeamMembers), len(teamMembers))
    for i, tm := range teamMembers {
        assert.Equal(t, expectedTeamMembers[i], tm)
    }

    // Verify that all expected queries were made
    assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetTeamMemberServiceFound(t *testing.T) {
    // Create a new mock database and defer its closure
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    // Define the columns that will be returned by the mock rows
    columns := []string{"id", "name", "role", "city"}

    // Define the row that will be returned by the mock database
    rows := sqlmock.NewRows(columns).
        AddRow("1", "Riya Gupta", "Developer", "Ranchi")

    // Expect a query to be made to the database and return the mock rows
	mock.ExpectQuery("SELECT\\s+id,\\s+name,\\s+role,\\s+city\\s+FROM\\s+team_members\\s+WHERE\\s+id\\s+=\\s+\\$1").
		WithArgs("1").
		WillReturnRows(rows)


    // Create an *sqlx.DB instance using the mock database
    sqlxDB := sqlx.NewDb(db, "sqlmock")

    // Call the function being tested
    teamMember, err := GetTeamMemberService(sqlxDB, "1")
    assert.NoError(t, err)

    // Verify that the correct data was returned
    expectedTeamMember := &model.TeamMember{
        ID:   "1",
        Name: "Riya Gupta",
        Role: "Developer",
        City: "Ranchi",
    }
    assert.Equal(t, expectedTeamMember, teamMember)

    // Verify that all expected queries were made
    assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateTeamMemberService(t *testing.T) {
    // Create a new mock database and defer its closure
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    // Create a new team member with some test data
    newTeamMember := model.NewTeamMember{
        Name: "John Doe",
        Role: "Manager",
        City: "New York",
    }

    // Define the columns that will be returned by the mock rows
    columns := []string{"ID", "Name", "Role", "City"}

    // Define the rows that will be returned by the mock database
    rows := sqlmock.NewRows(columns).
        AddRow("1", newTeamMember.Name, newTeamMember.Role, newTeamMember.City)

    // Expect a query to be made to the database and return the mock rows
    mock.ExpectQuery("INSERT INTO team_members (.+)").WillReturnRows(rows)

    // Create an *sqlx.DB instance using the mock database
    sqlxDB := sqlx.NewDb(db, "sqlmock")

    // Call the function being tested
    teamMember, err := CreateTeamMemberService(sqlxDB, newTeamMember)
    assert.NoError(t, err)

    // Verify that the correct data was returned
    assert.Equal(t, "1", teamMember.ID)
    assert.Equal(t, newTeamMember.Name, teamMember.Name)
    assert.Equal(t, newTeamMember.Role, teamMember.Role)
    assert.Equal(t, newTeamMember.City, teamMember.City)

    // Verify that all expected queries were made
    assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateTeamMemberService_Error(t *testing.T) {
    // Create a new mock database and defer its closure
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    // Expect an error to be returned when a NewTeamMember object is empty
    emptyRequest := model.NewTeamMember{}

	sqlxDB := sqlx.NewDb(db, "sqlmock")

    _, err = CreateTeamMemberService(sqlxDB, emptyRequest)
    assert.Error(t, err)

    // Expect an error to be returned when the city of the team member does not exist
    requestWithNonExistingCity := model.NewTeamMember{
        Name: "John Doe",
        Role: "Developer",
    }
    mock.ExpectQuery("INSERT INTO team_members (.+)").
        WithArgs(requestWithNonExistingCity.Name, requestWithNonExistingCity.Role, requestWithNonExistingCity.City).
        WillReturnError(fmt.Errorf("city does not exist"))
    _, err = CreateTeamMemberService(sqlxDB, requestWithNonExistingCity)
    assert.Error(t, err)

    // Verify that all expected queries were made
    assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRemoveTeamMemberService_Success(t *testing.T) {
    // Create a new mock database and defer its closure
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    // Define the team member that will be deleted
    teamMember := model.TeamMember{
        ID: "1",
        Name: "Riya Gupta",
        Role: "Developer",
        City: "Ranchi",
    }

    // Expect a query to be made to delete the team member
    mock.ExpectQuery("DELETE FROM team_members WHERE id=\\$1 RETURNING id, name, role, city").
        WithArgs("1").
        WillReturnRows(sqlmock.NewRows([]string{"id", "name", "role", "city"}).
            AddRow(teamMember.ID, teamMember.Name, teamMember.Role, teamMember.City))

    // Create an *sqlx.DB instance using the mock database
    sqlxDB := sqlx.NewDb(db, "sqlmock")

    // Call the function being tested
    deletedTeamMember, err := RemoveTeamMemberService(sqlxDB, "1")
    assert.NoError(t, err)

    // Verify that the correct data was returned
    assert.Equal(t, teamMember.ID, deletedTeamMember.ID)
    assert.Equal(t, teamMember.Name, deletedTeamMember.Name)
    assert.Equal(t, teamMember.Role, deletedTeamMember.Role)
    assert.Equal(t, teamMember.City, deletedTeamMember.City)

    // Verify that all expected queries were made
    assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRemoveTeamMemberService_TeamMemberNotFound(t *testing.T) {
    // Create a new mock database and defer its closure
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    // Expect a query to be made to delete the team member
    mock.ExpectQuery("DELETE FROM team_members WHERE id=\\$1 RETURNING id, name, role, city").
        WithArgs("1").
        WillReturnError(sql.ErrNoRows)

    // Create an *sqlx.DB instance using the mock database
    sqlxDB := sqlx.NewDb(db, "sqlmock")

    // Call the function being tested
    _, err = RemoveTeamMemberService(sqlxDB, "1")
    assert.EqualError(t, err, "sql: no rows in result set")

    // Verify that all expected queries were made
    assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRemoveTeamMemberService_DatabaseError(t *testing.T) {
    // Create a new mock database and defer its closure
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    // Expect a query to be made to delete the team member
    mock.ExpectQuery("DELETE FROM team_members WHERE id=\\$1 RETURNING id, name, role, city").
        WithArgs("1").
        WillReturnError(fmt.Errorf("database error"))

    // Create an *sqlx.DB instance using the mock database
    sqlxDB := sqlx.NewDb(db, "sqlmock")

    // Call the function being tested
    _, err = RemoveTeamMemberService(sqlxDB, "1")
    assert.EqualError(t, err, "database error")

    // Verify that all expected queries were made
    assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateTeamMemberService(t *testing.T) {
    // Create a new mock database and defer its closure
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)
    defer db.Close()

    // Define the columns that will be returned by the mock rows
    columns := []string{"id", "name", "role", "city"}

    // Define the rows that will be returned by the mock database
    rows := sqlmock.NewRows(columns).
        AddRow("1", "Updated Name", "Updated Role", "Updated City")

    // Expect an update query to be made to the database and return the mock row
    mock.ExpectQuery("UPDATE team_members SET name=\\$2, role=\\$3, city=\\$4 WHERE id=\\$1 RETURNING id, name, role, city").
        WithArgs("1", "Updated Name", "Updated Role", "Updated City").
        WillReturnRows(rows)

    // Create an *sqlx.DB instance using the mock database
    sqlxDB := sqlx.NewDb(db, "sqlmock")

    // Call the function being tested
    requestBody := model.UpdateTeamMember{
        ID:   "1",
        Name: pointerToString("Updated Name"),
        Role: pointerToString("Updated Role"),
        City: pointerToString("Updated City"),
    }
    teamMember, err := UpdateTeamMemberService(sqlxDB, requestBody)
    assert.NoError(t, err)

    expectedTeamMember := &model.TeamMember{
        ID:   "1",
        Name: "Updated Name",
        Role: "Updated Role",
        City: "Updated City",
    }

    // Verify that the correct data was returned
    assert.Equal(t, expectedTeamMember, &teamMember)

    // Verify that all expected queries were made
    assert.NoError(t, mock.ExpectationsWereMet())
}

func pointerToString(s string) *string {
    return &s
}
