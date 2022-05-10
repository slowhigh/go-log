package clone

type inode interface {
	print(string)
	clone() inode // prototype pattern
	equalsClone(inode) bool
}
