# Epub parser

This parser only supports epub version 3.0 so far.
This parser also only reads MetaData and Manifest.

### How to use?
```go
binaryFile, err := os.ReadFile("./test_epub_v3_0.epub")
if err != nil {
    println(err.Error())
	return;
}
zipReader, err := zip.NewReader(bytes.NewReader(binaryFile), int64(len(binaryFile)))
if err != nil {
    println(err.Error())
	return;
}
book, err := OpenBook(zipReader)
if err != nil {
    println(err.Error())
	return;
}
```
After that, all metadata and manifest data can be found in the book Object