package main

import "os"

func main() {
	//methods
	/*
		file, err := os.Open("C:/Users/siddharth.more/OneDrive - HCL TECHNOLOGIES LIMITED/Roche_Onedrive/PlayGroundProjects/GoLang/GoLang-Tutorial/23_files/example.txt")
		if err != nil {
			panic(err)
		}

		fileInfo, err := file.Stat()
		if err != nil {
			panic(err)
		}

		fmt.Println("File name:", fileInfo.Name())
		fmt.Println("File permission:", fileInfo.Mode())
		fmt.Println("File size:", fileInfo.Size())
		fmt.Println("is Directory", fileInfo.IsDir())
		fmt.Println("File last modified:", fileInfo.ModTime())
	*/

	//read data using buffer
	/*	data, err := os.Open("C:/Users/siddharth.more/OneDrive - HCL TECHNOLOGIES LIMITED/Roche_Onedrive/PlayGroundProjects/GoLang/GoLang-Tutorial/23_files/example.txt")
		if err != nil {
			panic(err)
		}

		defer data.Close()
		d := make([]byte, 10)
		_, err = data.Read(d)
		if err != nil {
			return
		}

		for i := 0; i < len(d); i++ {
			fmt.Println(d[i], string(d[i]))
		}
	*/

	//for small files, because ReadFile() will load full file data into the memory
	/*	data, err := os.ReadFile("C:/Users/siddharth.more/OneDrive - HCL TECHNOLOGIES LIMITED/Roche_Onedrive/PlayGroundProjects/GoLang/GoLang-Tutorial/23_files/example.txt")
		if err != nil {
			panic(err)
		}

		fmt.Println("data:", string(data))
	*/

	// read folders
	/*data, err := os.ReadDir("C:/Users/siddharth.more/OneDrive - HCL TECHNOLOGIES LIMITED/Roche_Onedrive/PlayGroundProjects/GoLang/GoLang-Tutorial")

	if err != nil {
		panic(err)
	}

	for _, i := range data {
		fmt.Println(i.Name(), i.IsDir())
	}
	*/

	// dir, err := os.Open(".")
	// error check
	// close dir
	// fileInfo, err := dir.ReadDir(i)  //i here is how many names you need, if i<=0 then all files list is returned
	// loop through the fileInfo

	// create and write into file using string, and bytes

	/*	f, err := os.Create("23_files/example2.txt")
		if err != nil {
			panic(err)
		}

		_, err = f.WriteString("Siddharth")
		if err != nil {
			panic(err)
		}

		_, err = f.Write([]byte("Waah"))
	*/
	/*
		// write data of one file to other (streaming way)
		sourceFile, err := os.Open("C:/Users/siddharth.more/OneDrive - HCL TECHNOLOGIES LIMITED/Roche_Onedrive/PlayGroundProjects/GoLang/GoLang-Tutorial/23_files/example.txt")
		if err != nil {
			panic(err)
		}

		defer sourceFile.Close()

		destFile, err := os.Create("C:/Users/siddharth.more/OneDrive - HCL TECHNOLOGIES LIMITED/Roche_Onedrive/PlayGroundProjects/GoLang/GoLang-Tutorial/23_files/example2.txt")

		if err != nil {
			panic(err)
		}

		defer destFile.Close()

		sourceReader := bufio.NewReader(sourceFile)
		destWriter := bufio.NewWriter(destFile)

		for {
			b, err := sourceReader.ReadByte()
			fmt.Println("Read", string(b))
			if err != nil {
				if err.Error() != "EOF" {
					panic(err)
				}
				fmt.Println("EOF reached")
				break // break when end of file is reached
			}

			err = destWriter.WriteByte(b)
			fmt.Println("Written")
			if err != nil {
				panic(err)
			}
		}

		destWriter.Flush()
	*/

	// delete a file
	err := os.Remove("23_files/example2.txt")
	if err != nil {
		panic(err)
	}
}
