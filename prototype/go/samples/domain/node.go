package domain

type Inode interface {
	Print(string)
	Clone() Inode
}
