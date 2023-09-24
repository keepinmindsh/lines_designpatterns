package domain

type Inode interface {
	print(string)
	clone() Inode
}
