package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	//Arrange
	password := "Password!123"

	t.Run("Should assert equal", func(t *testing.T) {
		//Arrange
		//Act
		hash, err := HashPassword(password, "")
		//Assert
		if err != nil {
			t.Fatal(err)
		}
		if !ComparePassword(password, hash, "") {
			t.FailNow()
		}
	})

	t.Run("Should fail on different pepper", func(t *testing.T) {
		//Arrange
		//Act
		hash, err := HashPassword(password, "Pepper123")
		//Assert
		if err != nil {
			t.FailNow()
		}
		if ComparePassword(password, hash, "") {
			t.FailNow()
		}
	})
	t.Run("Should succeed on same pepper", func(t *testing.T) {
		//Arrange
		//Act
		hash, err := HashPassword(password, "Pepper123")
		//Assert
		if err != nil {
			t.FailNow()
		}
		if !ComparePassword(password, hash, "Pepper123") {
			t.FailNow()
		}
	})
}

func TestComparePassword(t *testing.T) {
	//Arrange
	password := "Password!123"
	hash, err := HashPassword(password, "")
	hashPepper, err2 := HashPassword(password, "Pepper123")

	if err != nil || err2 != nil {
		t.FailNow()
	}

	require.NoError(t, err)
	require.NoError(t, err2)

	t.Run("Should succeed", func(t *testing.T) {
		//Act
		res1 := ComparePassword(password, hash, "")
		res2 := ComparePassword(password, hashPepper, "Pepper123")
		//Assert
		assert.True(t, res1)
		assert.True(t, res2)
	})

	t.Run("Should fail on different password", func(t *testing.T) {
		//Act
		res1 := ComparePassword("Password", hash, "")
		res2 := ComparePassword("Password", hashPepper, "Pepper123")
		//Assert
		assert.False(t, res1)
		assert.False(t, res2)

	})

	t.Run("Should fail on different pepper", func(t *testing.T) {
		//Act
		res1 := ComparePassword(password, hash, "Pepper123")
		res2 := ComparePassword(password, hashPepper, "")
		//Assert
		assert.False(t, res1)
		assert.False(t, res2)
	})

}

func TestAddPepper(t *testing.T) {
	//Arrange
	password := "Password!123"
	pepper := "ABC"
	expected := []byte("Password!123ABC")

	//Act
	res := addPepper(password, pepper)
	assert.Equal(t, expected, res)
}
