package services

import (
	"log"
	"shtem-web/sources/internal/core/domain"
	"shtem-web/sources/internal/repositories"
)

type categoriesService struct {
	categoriesRepository repositories.CategoriesRepository
	shtemsRepository     repositories.ShtemsRepository
}

func (q *categoriesService) GetCategories() ([]*domain.Category, domain.Error) {
	return q.categoriesRepository.GetCategories()
}

func (q *categoriesService) GetCategoryByID(id int64) (*domain.Category, domain.Error) {
	return q.categoriesRepository.GetCategoryByID(id)
}

func (q *categoriesService) GetCategoryByLinkName(c_link_name string) (*domain.Category, domain.Error) {
	return q.categoriesRepository.GetCategoryByLinkName(c_link_name)
}

func (q *categoriesService) GetCategoryByShtemLinkName(s_linkName string) (*domain.Category, domain.Error) {
	shtem, err := q.shtemsRepository.GetShtemByLinkName(s_linkName)
	if err != nil {
		return nil, err
	}

	return q.categoriesRepository.GetCategoryByID(shtem.Category)
}

func (s *categoriesService) GetCategoriesWithShtems() (*domain.Categories, domain.Error) {
	categories, err := s.categoriesRepository.GetCategories()
	if err != nil {
		return nil, err
	}

	final := make(domain.Categories, 0)

	for _, c := range categories {

		var curr domain.SortedCategory

		shtems, err := s.shtemsRepository.GetShtemsByCategoryId(c.C_id)
		if err != nil {
			return nil, err
		}

		curr.Category = c
		curr.Shtemarans = shtems

		log.Println(c.LinkName, c.C_id, shtems)

		final = append(final, curr)
	}

	return &final, nil
}

func (s *categoriesService) GetShtemsByCategoryLinkName(c_linkName string) ([]*domain.Shtemaran, domain.Error) {
	category, err := s.categoriesRepository.GetCategoryByLinkName(c_linkName)
	if err != nil {
		return nil, err
	}

	// Now, retrieve all shtems associated with the category ID
	return s.shtemsRepository.GetShtemsByCategoryId(category.C_id)
}

func NewCategoriesService(categoriesRepository repositories.CategoriesRepository, shtemsRepository repositories.ShtemsRepository) *categoriesService {
	return &categoriesService{categoriesRepository, shtemsRepository}
}
