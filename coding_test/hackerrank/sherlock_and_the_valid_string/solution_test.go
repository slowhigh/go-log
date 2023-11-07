package sherlock_and_the_valid_string

import "testing"

type testCase struct {
	s      string
	result string
}

func Test_isValid(t *testing.T) {
	testCaseArr := []testCase{
		{
			s:      "ibfdgaeadiaefgbhbdghhhbgdfgeiccbiehhfcggchgghadhdhagfbahhddgghbdehidbibaeaagaeeigffcebfbaieggabcfbiiedcabfihchdfabifahcbhagccbdfifhghcadfiadeeaheeddddiecaicbgigccageicehfdhdgafaddhffadigfhhcaedcedecafeacbdacgfgfeeibgaiffdehigebhhehiaahfidibccdcdagifgaihacihadecgifihbebffebdfbchbgigeccahgihbcbcaggebaaafgfedbfgagfediddghdgbgehhhifhgcedechahidcbchebheihaadbbbiaiccededchdagfhccfdefigfibifabeiaccghcegfbcghaefifbachebaacbhbfgfddeceababbacgffbagidebeadfihaefefegbghgddbbgddeehgfbhafbccidebgehifafgbghafacgfdccgifdcbbbidfifhdaibgigebigaedeaaiadegfefbhacgddhchgcbgcaeaieiegiffchbgbebgbehbbfcebciiagacaiechdigbgbghefcahgbhfibhedaeeiffebdiabcifgccdefabccdghehfibfiifdaicfedagahhdcbhbicdgibgcedieihcichadgchgbdcdagaihebbabhibcihicadgadfcihdheefbhffiageddhgahaidfdhhdbgciiaciegchiiebfbcbhaeagccfhbfhaddagnfieihghfbaggiffbbfbecgaiiidccdceadbbdfgigibgcgchafccdchgifdeieicbaididhfcfdedbhaadedfageigfdehgcdaecaebebebfcieaecfagfdieaefdiedbcadchabhebgehiidfcgahcdhcdhgchhiiheffiifeegcfdgbdeffhgeghdfhbfbifgidcafbfcd",
			result: "YES",
		},
		{
			s:      "aaaa",
			result: "YES",
		},
		{
			s:      "aabbbc",
			result: "NO",
		},
		{
			s:      "aaac",
			result: "YES",
		},
		{
			s:      "abcdd",
			result: "YES",
		},
		{
			s:      "abc",
			result: "YES",
		},
		{
			s:      "aabbc",
			result: "YES",
		},
		{
			s:      "aabbcc",
			result: "YES",
		},
		{
			s:      "aaabbbcccddee",
			result: "NO",
		},
		{
			s:      "aaabbc",
			result: "NO",
		},
		{
			s:      "aabbcccd",
			result: "NO",
		},
		{
			s:      "aabbcd",
			result: "NO",
		},
		{
			s:      "aabbccddeefghi",
			result: "NO",
		},
		{
			s:      "abcdefghhgfedecba",
			result: "YES",
		},
	}

	for i, tc := range testCaseArr {
		if isValid(tc.s) == tc.result {
			t.Logf("PASS - %d", i)
		} else {
			t.Errorf("FAIL - %d", i)
		}
	}
}
