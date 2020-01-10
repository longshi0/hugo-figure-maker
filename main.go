package main

/*

## example datum that we want to be automatically generated:

{{< gallery >}}

{{< figure link="/img/podejrzane/Kazimierz Dolny(Obrazy podejrzane).jpg"
caption="Kazimierz Dolny." alt="Układanka światła i cienia. Interpretacja dowolna. Elementy „porozrzucane” w kadrze mogą stać się pretekstem do opowiedzenia niejednej historii. Być może z wątkiem kryminalnym?" >}}

what i have is a list of files with jpg files with descriptive names to
be inserted as *caption* key value and matching txt files to be,
optionally, inserted as a *alt* key value.

{{< gallery >}}

*/

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// directory path to the folder with files i want parsed. some
	// photos matching text files, with descriptions as its contents.
	path := "/home/mist/Desktop/kajeczka/MICHAŚ OK/bio"

	// suffix to the filename containing the result.
	resultFilename := "result.txt"

	txtFiles, jpgFiles := getFiles(path)

	// create a file with figure tags, a result
	f, err := os.Create(path + "-" + resultFilename)
	if err != nil {
		log.Fatalf("Cannot create file: %v", err)
	}
	defer f.Close()

	// write a starting tag to our result file `f`
	f.WriteString("{{< gallery >}}\n")

	for _, jpg := range jpgFiles {

		// start a figure node. write path to the jpg file being
		// parsed as a value of *link* key.
		f.WriteString("{{< figure link=\"/img/" + filepath.Base(path) + "/")
		f.WriteString(jpg + "\" ")

		// get picture's name, without extension, and store it
		// inside `jpgName`, but without its extension. Then,
		// write `jpgName` as a value to *caption* key.
		jpgName := jpg[:len(jpg)-4]
		f.WriteString("caption=" + "\"" + jpgName + "\" ")

		// check if there exists a txt file with the same name as
		// in jpg file -- a match. if so, read its contents and
		// write it at as a value (read from `word` var,
		// sequentially, from the `scanner`) to the alt key.
		// else, end the figure by writing a ">}}", and newline.
		if i, ok := existsMatchingTxt(txtFiles, jpgName); ok {

			// get the txt file content
			txt := txtFiles[i]
			file, err := os.Open(path + "/" + txt)
			if err != nil {
				log.Fatalf("Cannot open file %s due to error %v", txt, err)
			}

			defer file.Close()

			// preprare a file with photograph description
			// for scanning
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanWords)

			// write the file to the file at *alt* key, to be
			// the photo's description
			f.WriteString("alt= \"")
			for scanner.Scan() {
				word := scanner.Bytes()
				f.Write(word)
				f.WriteString(" ")
			}
			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}

			// write to end alt tag and markup
			f.WriteString("\" >}}\n")

		} else {
			// write to end figure markup
			f.WriteString(" >}}\n")
		}

	}

	// end gallery tag. list is complete.
	f.WriteString("{{</ gallery >}}\n")

	// update the ready version of a file to persistent storage
	f.Sync()

}

// walk the given path and return a slice of the os.FileInfo type to be
// filtered.
func getFiles(p string) ([]string, []string) {
	files, err := ioutil.ReadDir(p)
	if err != nil {
		panic(err)
	}
	return filterFiles(files)
}

// walk the slice populated with os.FileInfo structs and check for names.
// add ones with jpg and txt suffixes in its names to their respective
// slices.
func filterFiles(f []os.FileInfo) (txt []string, jpg []string) {

	for _, v := range f {
		// Ignore all directories
		if !v.IsDir() {
			// Append files with .txt extension
			if filepath.Ext(v.Name()) == ".txt" {
				txt = append(txt, v.Name())
			}

			// Append files with .jpg extension
			if filepath.Ext(v.Name()) == ".jpg" || filepath.Ext(v.Name()) == ".JPG" {
				jpg = append(jpg, v.Name())
			}
		}
	}

	return
}

// helper function. checks if there is a matching txt file inside slice
// populated with txt files names -- to decide whether the program should
// attempt writing files contents as an `alt` pseudo-html key accepted as
// hugo/golang templating markup.
func existsMatchingTxt(txts []string, jpgName string) (int, bool) {

	for i, txt := range txts {
		txtName := strings.Trim(txt, ".txt")
		if txtName == jpgName {
			return i, true
		}
	}

	return -1, false

}
