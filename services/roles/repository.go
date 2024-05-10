package roles

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	db.AutoMigrate(&Role{})
	repo := Repository{db: db}
	repo.seed()
	return &repo
}

func (repo *Repository) seed() {
	roles := []Role{
		{ID: 1, Name: "admin"},
		{ID: 2, Name: "user"},
	}
	for _, role := range roles {
		_, err := repo.FindByID(role.ID)
		if err == gorm.ErrRecordNotFound {
			repo.Create(&role)
		}
	}
}

func (repo *Repository) FindByID(id uint) (*Role, error) {
	var role Role
	if err := repo.db.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (repo *Repository) FindAll() ([]*Role, error) {
	var roles []*Role
	if err := repo.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}
func (repo *Repository) Create(role *Role) error {
	return repo.db.Create(role).Error
}
