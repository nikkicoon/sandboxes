// dup1 prints the text of each line that appears more than
// once in the standard input, preced by its count.
package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  // create a new empty map (using "make").
  // a map is a hash table.
  // it looks like:
  // map[KeyType]ValueType
  // So in our case we create a map where the
  // key type is string and the value type is int.
  counts := make(map[string]int)
  input := bufio.NewScanner(os.Stdin)
  // as long as the bufio scanner reads a line, it returns true,
  // if there is no more line it returns false.
  for input.Scan() {
    // the below is equivalent to:
    // line := input.Text()
    // counts[line] = counts[line] + 1
    // Each time dup reads a line of input, the line is used as
    // a key into the map and the corresponding value is
    // incremented.
    counts[input.Text()]++
  }
  // NOTE: Ignoring potential error from input.Err()
  for line, n := range counts {
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, line)
    }
  }
}
