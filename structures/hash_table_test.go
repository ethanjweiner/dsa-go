package structures

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type HashTableTestSuite struct {
	suite.Suite
	table1 *HashTable
}

func (suite *HashTableTestSuite) SetupTest() {
	suite.table1 = InitHashTable()
	suite.table1.Insert("Ethan")
	suite.table1.Insert("Etan")  // Same bucket as "Etthan"
	suite.table1.Insert("Ethah") // Same bucket as "Etthan"
	suite.table1.Insert("Will")
	suite.table1.Insert("Noah")
	suite.table1.Insert("Nohh") // Same bucket as Noah
	suite.table1.Insert("Drew")
}

// Tests
func (suite *HashTableTestSuite) TestSearch() {
	suite.True(suite.table1.Search("Ethan"))
	suite.True(suite.table1.Search("Etan"))
	suite.True(suite.table1.Search("Ethah"))
	suite.True(suite.table1.Search("Will"))
	suite.True(suite.table1.Search("Noah"))
	suite.True(suite.table1.Search("Nohh"))
	suite.True(suite.table1.Search("Drew"))
	suite.False(suite.table1.Search("Grace"))
	suite.False(suite.table1.Search("Mom"))
}

func (suite *HashTableTestSuite) TestDelete() {
	suite.table1.Delete("Etan")
	suite.False(suite.table1.Search("Etan"))

	suite.table1.Delete("Drew")
	suite.False(suite.table1.Search("Drew"))
}

func (suite *HashTableTestSuite) TestDeleteNonExistentKey() {
	suite.table1.Delete("abc")
	suite.False(suite.table1.Search("abc"))
}

func TestHashTableTestSuite(t *testing.T) {
	suite.Run(t, new(HashTableTestSuite))
}
