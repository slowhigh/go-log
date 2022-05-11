package composite

import "testing"

func Test_Composite(t *testing.T) {
	fileA := &file{name: "a"}
	fileB := &file{name: "b"}
	fileC := &file{name: "c"}
	
	folderA := &folder{name: "A"}
	folderA.add(fileA)
	folderA.add(fileB)

	folderB := &folder{name: "B"}
	folderB.add(folderA)
	folderB.add(fileC)

	if !folderB.search("b") {
		t.Error("Invalid search function.")
	}

	if folderB.search("d") {
		t.Error("Invalid search function.")
	}
}