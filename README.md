# Epub parser

This parser only supports epub version 3.0 so far.
This parser also only reads MetaData and Manifest.

### How to use?
```go
zipReader, err := zip.OpenReader("./test_epub_v3_0.epub")
if err != nil {
    t.Log(err.Error())
    t.Fail()
}
defer zipReader.Close()

book, err := OpenBook(zipReader)
if err != nil {
    t.Log(err.Error())
    t.Fail()
}
```
After that, all metadata and manifest data can be found in the book Object