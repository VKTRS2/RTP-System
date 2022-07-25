package main

var fileToRead = [2]string{"sct1.xml", "sct2.xml"}

func client(address string) {
	for _, path := range fileToRead {
		data := readFile(path)
		makeRequest(address, data)
	}
}
