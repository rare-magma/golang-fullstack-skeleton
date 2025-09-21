package eventbus

import (
	"golang-fullstack-skeleton/internal/contexts/shared/src/domain"
	"testing"
)

func TestObjectMother(t *testing.T) {
	objectMother := domain.NewObjectMother()

	t.Run("works", func(t *testing.T) {
		uuid := objectMother.Uuid()
		if uuid == "" {
			t.Errorf("Expected uuid to be non-empty, got: %v", uuid)
		}
	})
}