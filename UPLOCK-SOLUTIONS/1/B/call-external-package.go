/* package main
group functions */
package main

/* import
when you need your code to do something that might have been implemented by someone else.*/

/* fmt
contains functions for formatting text : Println */
import "fmt"

/* rsc.io/quote
this package collects pithy sayings: package quote*/
import "rsc.io/quote"

/* func main
create a enter point for go
*/
func main() {
	fmt.Println(quote.Go())
}
