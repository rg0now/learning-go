package richterscale

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDescribeEarthquake(t *testing.T) {

	assert.Equal(t, describeEarthquake(0.5), "micro")

	assert.Equal(t, describeEarthquake(2.5), "very minor")

	assert.Equal(t, describeEarthquake(3), "minor")

	assert.Equal(t, describeEarthquake(4.5), "light")

	assert.Equal(t, describeEarthquake(5), "moderate")

	assert.Equal(t, describeEarthquake(6), "strong")

	assert.Equal(t, describeEarthquake(7), "major")

	assert.Equal(t, describeEarthquake(8), "great")

	assert.Equal(t, describeEarthquake(11), "massive")

}
