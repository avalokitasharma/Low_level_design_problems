package filesystem

func init() {
	fs := NewFileSystem()

	// Setup test directory structure
	fs.CreateDirectory("/home")
	fs.CreateDirectory("/home/coding")

	fs.CreateFile("/home/Hello.txt", []byte("Hi there!"))
	fs.CreateFile("/home/GoodDay.txt", []byte("Have a good day!"))

	// Test listing root directory
	entries, err := fs.ListDirectory("/")

}
