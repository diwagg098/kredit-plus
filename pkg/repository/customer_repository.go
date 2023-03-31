package repository

func (r *repo) CustomerList(i interface{}, where string) error {
	result := r.apps.Preload("Customers").Find(i)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
