package txt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWords(t *testing.T) {
	t.Run("I'm a lazy brown fox!", func(t *testing.T) {
		result := Words("I'm a lazy BRoWN fox!")
		assert.Equal(t, []string{"lazy", "BRoWN", "fox"}, result)
	})
	t.Run("no result", func(t *testing.T) {
		result := Words("x")
		assert.Equal(t, []string(nil), result)
	})
}

func TestKeywords(t *testing.T) {
	t.Run("I'm a lazy brown fox!", func(t *testing.T) {
		result := Keywords("I'm a lazy BRoWN img!")
		assert.Equal(t, []string{"lazy", "brown"}, result)
	})
	t.Run("no result", func(t *testing.T) {
		result := Keywords("was")
		assert.Equal(t, []string(nil), result)
	})
}

func TestUniqueWords(t *testing.T) {
	t.Run("many", func(t *testing.T) {
		result := UniqueWords([]string{"lazy", "brown", "apple", "brown"})
		assert.Equal(t, []string{"apple", "brown", "lazy"}, result)
	})
	t.Run("one", func(t *testing.T) {
		result := UniqueWords([]string{"lazy"})
		assert.Equal(t, []string{"lazy"}, result)
	})
}

func TestUniqueKeywords(t *testing.T) {
	t.Run("many", func(t *testing.T) {
		result := UniqueKeywords("lazy, brown, apple, brown, ...")
		assert.Equal(t, []string{"apple", "brown", "lazy"}, result)
	})
	t.Run("one", func(t *testing.T) {
		result := UniqueKeywords("")
		assert.Equal(t, []string(nil), result)
	})
}
