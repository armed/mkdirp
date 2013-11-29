mkdirp
======

Make directory hierarchy with one function call.

####Example

    import (
      "github.com/armed/mkdirp"
      "log"
    )

    func main() {
      err := mkdirp.Mk("/Users/armed/{somefolder1,somefolder2}/test/{data1,data2}")
      if err != nil {
        log.Fatal(err)
      }
    }

Following directory tree will be created:

    Users
    |__armed
       |__somefolder1
       |  |__test
       |     |__data1
       |     |__data2
       |__somefolder2
          |__test
             |__data1
             |__data2

`mkdirp.Mk` function internally uses `os.MkdirAll`, and returns errors from it if any.
