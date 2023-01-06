package handlers

import (
	_ "database/sql"
	"fmt"
	"github.com/Click2Cloud/m2k-external-api/database"
	"github.com/Click2Cloud/m2k-external-api/external/types"
)

func insertWorkspace(reqWorkspace types.Workspace) error {
	fmt.Println(reqWorkspace.Id)
	dbconn, _ := database.CreateDatabaseConnection().DB()
	sqlStatement := `INSERT INTO workspace (w_id, w_name, w_description, w_timestamp) VALUES ($1, $2, $3, $4)`
	result, err := dbconn.Exec(sqlStatement, reqWorkspace.Id, reqWorkspace.Name, reqWorkspace.Description, reqWorkspace.Timestamp)
	defer dbconn.Close()
	if err != nil {
		return err
		panic(err)
	} else {
		fmt.Println(result)
	}
	return err
}
