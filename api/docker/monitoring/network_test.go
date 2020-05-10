package monitoring

import "testing"

type conversionTest struct {
	bitValue      float64
	expectedValue float64
	expectedUnit  string
}

func TestConvertNetworkStat(t *testing.T) {

	convertToBit := conversionTest{
		650.0,
		650.0,
		"b",
	}
	convertToKo := conversionTest{
		6000.0,
		6.0,
		"Kb",
	}

	convertToMo := conversionTest{
		10000000.0,
		10.0,
		"Mb",
	}

	convertToGo := conversionTest{
		2000000000.0,
		2.0,
		"Gb",
	}

	conversionTests := []conversionTest{convertToBit, convertToKo, convertToMo, convertToGo}

	for _, test := range conversionTests {
		converted, unit := convertNetworkStat(test.bitValue)

		if converted != test.expectedValue || unit != test.expectedUnit {
			t.Fatalf("Conversion to %s test FAILED", test.expectedUnit)
		} else {
			t.Logf("Conversion to %s test PASSED", test.expectedUnit)
		}
	}
}
