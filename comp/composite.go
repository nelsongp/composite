package comp

//Item
type Item interface {
	Add(Item)
	String() string
	Price() int
}
