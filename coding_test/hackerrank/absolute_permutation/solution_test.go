package absolute_permutation

import (
	"fmt"
	"testing"
)

type testCase struct {
	n      int32
	k      int32
	result []int32
}

func Test_absolutePermutation(t *testing.T) {
	testCaseArr := []testCase{
		{
			n:      12,
			k:      3,
			result: []int32{2, 1},
		},
		{
			n:      2,
			k:      1,
			result: []int32{2, 1},
		},
		{
			n:      3,
			k:      0,
			result: []int32{1, 2, 3},
		},
		{
			n:      3,
			k:      2,
			result: []int32{-1},
		},
		{
			n:      10,
			k:      6,
			result: []int32{-1},
		},
	}

	for i, tc := range testCaseArr {
		res := fmt.Sprintf("%+v", absolutePermutation(tc.n, tc.k))
		if res == fmt.Sprintf("%+v", tc.result) {
			t.Logf("PASS - %d (%s)", i, res)
		} else {
			t.Errorf("FAIL - %d (%s)", i, res)
		}
	}
}
//         1X      2X         3X           4X           5
// -----------------------------------------------------------
// 1   |   2  2    3  3       4  X         5  X         6              
// 2   |   13      4  4       5            6            7                                          
// 3   |   24 4    15         6            7            8                         
// 4   |   35      26         17           8            9                            
// 5   |   46 6    37 X       28           19           X                            
// 6   |   57      48 8       39           2                                       
// 7   |   68 8    59         4  X         3                                    
// 8   |   79      6          5            4                                    
// 9   |   8  X    7  7       6            5  X                                   
// 
// 
// 
// 
// 
//  n % (2*k)
//  12 % (2*k)
//  1 2 3 6
//        1(O)          2()           3(O)           4()         5()                                                     
// -------------------------------------------------------------------------------
// 1   |  2    2        3     3       4    4                                                                                                                
// 2   |  13   1        4     4       5    5                                                                                                                
// 3   |  24   4        15    1       6    6                                                                                                                
// 4   |  35   3        26    2       17   1                                                                                                                 
// 5   |  46   6        37    7       28   2                                                                                                                 
// 6   |  57   5        48    8       39   3                                                                                                                 
// 7   |  68   8        59    5       410  10                                                                                                                  
// 8   |  79   7        610   6       511  11                                                                                                                  
// 9   |  810  10       711   11      612  12                                                                                                                  
// 10  |  911  9        812   12      7    7                                                                                                                
// 11  |  1012 12       9     9       8    8                                                                                                                 
// 12  |  11   11       10    10      9    9                                                                                                                 