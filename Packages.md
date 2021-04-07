Packages are a group of related source files bundled to be used as a dependency. <br>
Each package can be imported by other packages. <br>
Enables software re-use. <br>

## Usage

 - Package source files have to have the package name specified in the first line.
   ```
   package abcpkg
   ...
   ...
   ```
 - Packages are imported as follows:
   ```
   package ...
   import {
     "abcpkg"
   }
   ```

## Main package
In a go project, there has to be at least one package called `main` containing a `main()` function. <br>
This will be the entry point for execution.
Eg:
```
package main
import "fmt"
func main() {
  fmt.Printf("Hello, world\n")
}
```
