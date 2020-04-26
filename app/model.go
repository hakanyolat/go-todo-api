package app

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Seeder interface {
	WillSeed() []ModelInterface
}

type ModelInterface interface {
	Migrate(db *gorm.DB)
}

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type migration struct {
	Model
	Table  string
	Seeded bool
}

type ModelManager struct {
	db *gorm.DB
}

func NewModelManager(db *gorm.DB) *ModelManager {
	db.AutoMigrate(&migration{})
	return &ModelManager{db: db}
}

func (m *ModelManager) Migrate(model ModelInterface) {
	if !m.db.HasTable(model) {
		row := &migration{
			Table:  m.db.NewScope(model).TableName(),
			Seeded: false,
		}
		m.db.Save(row)
		model.Migrate(m.db)
	}
}

func (m *ModelManager) Seed(value interface{}) int {
	if s, ok := value.(Seeder); ok {
		seeds := s.WillSeed()

		if len(seeds) > 0 {
			var mig migration

			m.db.FirstOrInit(&mig, migration{
				Table:     m.db.NewScope(value).TableName(),
				Seeded:    false,
			})

			if !mig.Seeded{
				for _, data := range seeds {
					m.db.Save(data)
				}
				m.db.Model(mig).Update("seeded", true)
				return len(seeds)
			}
		}
	}

	return 0
}
