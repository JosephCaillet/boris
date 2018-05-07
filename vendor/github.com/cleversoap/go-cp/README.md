go-cp
=====

Copy file function in golang. Returns nil on success otherwise an error if there were any issues.

```golang
import "github.com/cleversoap/go-cp"

func main() {
  err := cp.Copy("somefile.txt", "somefile.copy.txt")
  if err != nil {
    panic(err)
  } else {
    // Successfully copied
  }
}
```
