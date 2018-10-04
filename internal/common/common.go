package common

import "fmt"

func Name(path, ver string) string {
   return fmt.Sprintf("%s version %v", path, ver)
}
