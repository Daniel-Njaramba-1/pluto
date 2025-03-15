package generics

import (
	"pluto/internal/lib/interfaces"

	"github.com/jmoiron/sqlx"
)

// SQLX
// Db.Get is used to retrieve a single record by its ID and scan the result into the model.
// Db.Select is used to retrieve multiple records of a given model and scan them into a slice of models.
// Db.NamedQuery is used to execute a query with named parameters.
// Db.NamedExec is used to execute a query with named parameters. Do not return rows

func CreateModel[M interfaces.Store](db *sqlx.DB, model M) (int, error) {
	query := model.FeedCreateQuery()
	rows, err := db.NamedQuery(query, model)
	
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var id int
	
	//check if there is a row in the result set and scan the id
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, err
		}
		
		if model.FeedGetID() != nil {
			*model.FeedGetID() = id	// Set the ID in the model, use pointer in order to update original model 
		}
	}
	return id, nil
}

// SelectModelByID retrieves a record by its ID 
func SelectModelByID[M interfaces.Store](db *sqlx.DB, id int, model M) error {
	query := model.FeedGetByIDQuery()
	return db.Get(model, query, id)
}

// SelectAllModels retrieves all records of a given model
func SelectAllModels[M interfaces.Store](db *sqlx.DB, models *[]M) error {
    var model M // Declare an instance of M in order to call interface
    query := model.FeedGetAllQuery()
    return db.Select(models, query)
}

// UpdateModelDetails updates details of a given model
func UpdateModelDetails[M interfaces.Store](db *sqlx.DB, model M) error {
	query := model.FeedUpdateDetailsQuery()
    rows, err := db.NamedQuery(query, model)
    
	if err != nil {
        return err
    }
    defer rows.Close()
    
	if rows.Next() {
        if err := rows.StructScan(model); err != nil {
            return err
        }
    }
    return nil
}

// DeactivateModel sets the is_active field to false
func DeactivateModel[M interfaces.Store](db *sqlx.DB, model M) error {
	query := model.FeedDeactivateQuery()
	_, err := db.NamedExec(query, model)
	return err
}

// ReactivateModel sets the is_active field to true
func ReactivateModel[M interfaces.Store](db *sqlx.DB, model M) error {
	query := model.FeedReactivateQuery()
	_, err := db.NamedQuery(query, model)
	return err
}

// DeleteModel permanently deletes a model
func DeleteModel[M interfaces.Store](db *sqlx.DB, model M) error {
	query := model.FeedDeleteQuery()
	_, err := db.NamedExec(query, model)
	return err
}

