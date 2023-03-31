package repository

import (
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *repo) FindAllWithWhere(i interface{}, where map[string]interface{}) error {
	result := r.apps.Find(i, where)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) FindAllWithWhereV2(i interface{}, cond ...interface{}) error {

	result := r.apps.Find(i, cond)
	// .Preload(clause.Associations)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Response is error and count result
func (r *repo) FindOne(i interface{}, where map[string]interface{}) (error, int64) {
	result := r.apps.Preload(clause.Associations).Find(i, where)

	if result.Error == gorm.ErrRecordNotFound {
		return nil, 0
	}
	if result.Error != nil {
		return result.Error, 0
	}
	return nil, result.RowsAffected
}

func (r *repo) InsertData(i interface{}) error {
	result := r.apps.Create(i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) FindOneWithTableName(i interface{}, where map[string]interface{}, tablename string) error {
	result := r.apps.Table(tablename).First(i, where)
	if result.Error == gorm.ErrRecordNotFound {
		return nil
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repo) DinamicFindQueryRaw(i interface{}, query string) (*sql.Rows, error) {
	rows, err := r.apps.Raw(query).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		r.apps.ScanRows(rows, i)
	}

	return rows, nil
}

func (r *repo) DeleteData(i interface{}, where map[string]interface{}) error {
	err := r.apps.Where(where).Delete(i)

	if err.Error != nil {
		return err.Error
	}
	return nil
}
func (r *repo) UpdateData(i interface{}, where map[string]interface{}, data map[string]interface{}) error {
	err := r.apps.Model(i).Where(where).Updates(data)

	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (r *repo) FindAll(i interface{}) error {
	result := r.apps.Find(i).Preload(clause.Associations)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *repo) FindAllWithWhereV3(i interface{}, where string) error {
	//new_version
	result := r.apps.Preload(clause.Associations).Where(where).Find(i)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
