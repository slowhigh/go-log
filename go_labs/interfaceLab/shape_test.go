package interface_lib

import "testing"

type testCase struct {
	width int
	height int
	rectangleArea int
	rightTriangleArea int
}

func Test_Figure(t *testing.T) {
	testCaseArr := []testCase {
		{
			width: 4,
			height: 5,
			rectangleArea: 20,
			rightTriangleArea: 10, 
		},
		{
			width: 6,
			height: 10,
			rectangleArea: 60,
			rightTriangleArea: 30,
		},
	}

	for _, tc := range testCaseArr {
		rect := Rectangle{width: tc.width, height: tc.height}
		if rect.GetAreaSize() != tc.rectangleArea {
			t.Error("Rectangle's size is invalid")
		}

		tri := RightTriangle{width: tc.width, height: tc.height}
		if tri.GetAreaSize() != tc.rightTriangleArea {
			t.Error("a RightTriangle's size is invalid")
		}
	}
}
