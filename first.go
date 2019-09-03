package main


import "fmt"

type Books struct{
	title string
	author string
	book_id int
}
func main(){

	var book1 Books
	var book2 Books

	book1.title = "Go语言教程"
	book1.author = "author_go"
	book1.book_id = 1111

	book2.title = "Python语言教程"

	var booklist []Books
	fmt.Println(booklist == nil)
	//booklist =[] Books {book1, book2}
	booklist = append(booklist,book1)
	booklist = append(booklist,book2)
	fmt.Println("结构体数组输出为：")
	printBook(booklist[0])
	printBook(booklist[1])
	fmt.Println(len(booklist),cap(booklist))

	//printBook(book1)
	//printBook(book2)
}

func printBook(book Books){
	fmt.Println(book.title, book.author,book.book_id)
}