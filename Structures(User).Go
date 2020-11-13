1.//Structures
package main
import "fmt"
type Books struct {
   title string
   author string
   subject string
   book_id int
}
func main() {
   var Book1 Books   
   var Book2 Books   
   Book1.title = "java Programming"
   Book1.author = "Haritha"
   Book1.subject = "python Programming Tutorial"
   Book1.book_id = 6495406   
   Book2.title = "HTML Billing"
   Book2.author = "rani"
   Book2.subject = "rani Billing Tutorial"
   Book2.book_id = 6495700
   printBook(&Book1)
   printBook(&Book2)
}
func printBook( book *Books ) {
   fmt.Printf( "Book title : %s\n", book.title);
   fmt.Printf( "Book author : %s\n", book.author);
   fmt.Printf( "Book subject : %s\n", book.subject);
   fmt.Printf( "Book book_id : %d\n", book.book_id);
}
$go run main.go
Book title : java Programming
Book author : Haritha
Book subject : python Programming Tutorial
Book book_id : 6495406
Book title : HTML Billing
Book author : rani
Book subject : rani Billing Tutorial
Book book_id : 6495700
