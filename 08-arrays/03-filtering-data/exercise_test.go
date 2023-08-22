package filteringdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var keys []string
var indices []int

func TestFilterData(t *testing.T) {
	// even
	

	// odd
	

	// x > 5
	

	keys = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	indices = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	assert.Equal(t, [10]string{"a", "b", "c", "d", "e", "f", "", "", "", ""}    , filterData(keys, indices))

	keys = []string{"a", "b"}
	indices = []int{0, 1}
	assert.Equal(t, [10]string{"a", "b", "", "", "", "", "", "", "", ""}	    , filterData(keys, indices))

	keys = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	indices = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	assert.Equal(t, [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, filterData(keys, indices))

	keys = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	indices = []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	assert.Equal(t, [10]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, filterData(keys, indices))

	keys = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	indices = []int{1}
	assert.Equal(t, [10]string{"", "", "", "", "", "", "", "", "", ""}	    , filterData(keys, indices))

				      	 									  

	// x <= 4
				      	 									  
}
