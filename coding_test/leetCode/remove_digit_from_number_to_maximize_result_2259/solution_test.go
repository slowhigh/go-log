package remove_digit_from_number_to_maximize_result_2259

import "testing"

type testCase struct {
	number string
	digit  byte
	result string
}

func Test_removeDigit(t *testing.T) {
	testCaseArr := []testCase{
		{
			number: "133235",
			digit: '3',
			result: "13325",
		},
		{
			number: "123",
			digit:  '3',
			result: "12",
		},
		{
			number: "1231",
			digit:  '1',
			result: "231",
		},
		{
			number: "1231",
			digit:  '1',
			result: "231",
		},
		{
			number: "2998589353917872714814599237991174513476623756395992135212546127959342974628712329595771672911914471",
			digit:  '3',
			result: "299858953917872714814599237991174513476623756395992135212546127959342974628712329595771672911914471",
		},
	}

	for i, tc := range testCaseArr {
		if removeDigit(tc.number, tc.digit) == tc.result {
			t.Logf("PASS - %d", i)
		} else {
			t.Errorf("FAIL - %d", i)
		}
	}
}
