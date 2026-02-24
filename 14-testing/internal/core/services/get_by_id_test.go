package services_test

import (
	"testing"

	"github.com/Nextjingjing/go-god/14-testing/internal/adapters/repository"
	"github.com/Nextjingjing/go-god/14-testing/internal/core/domain"
	"github.com/Nextjingjing/go-god/14-testing/internal/core/services"
	"github.com/stretchr/testify/assert"
)

func TestFindById(t *testing.T) {
	mockRepo := new(repository.TestifyMockRepository)
	service := services.NewProductService(mockRepo)

	expectedProduct := &domain.Product{Id: 1, Name: "Gadget"}
	mockRepo.On("FindById", uint(1)).Return(expectedProduct, nil)

	result, err := service.GetById(1)
	assert.Equal(t, expectedProduct, result)
	assert.NotNil(t, result)
	assert.NoError(t, err)
}
