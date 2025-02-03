package main

func main() {
	f := &File{
		Name: "teste.txt",
	}

	err := f.CreateFile(f.Name)

	if err != nil {
		panic(err)
	}

}
