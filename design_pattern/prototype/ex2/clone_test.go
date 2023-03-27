package clone

import "testing"

func Test_clone(t *testing.T) {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{
		name:     "Folder1",
		children: []inode{file1},
	}

	folder2 := &folder{
		name:     "Folder2",
		children: []inode{folder1, file2, file3},
	}

	if folder2.equalsClone(folder2) {
		t.Error("Invalid equals function.")
	}

	// prototype pattern
	clonedFolder2 := folder2.clone()

	if !folder2.equalsClone(clonedFolder2) {
		t.Error("Invalid clone function.")
	}
}
