package main
//
//import (
//<<<<<<< HEAD
//	"fmt" // пакет для форматированного ввода вывода
//	"net/http" // пакет для поддержки HTTP протокола
//	"strings" // пакет для работы с  UTF-8 строками
//	"log" // пакет для логирования
//
//)
//
//func HomeRouterHandler(w http.ResponseWriter, r *http.Request) {
//	r.ParseForm() //анализ аргументов,
//	fmt.Println(r.Form)  // ввод информации о форме на стороне сервера
//	fmt.Println("path", r.URL.Path)
//	fmt.Println("scheme", r.URL.Scheme)
//	fmt.Println(r.Form["url_long"])
//	for k, v := range r.Form {
//		fmt.Println("key:", k)
//		fmt.Println("val:", strings.Join(v, ""))
//	}
//	fmt.Fprintf(w, "Hello Maksim!") // отправляем данные на клиентскую сторону
//=======
//	"fmt"
//	"math/rand"
//
//	"runtime"
//
//	//"strings"
//)
//type Tree struct {
//	Left  *Tree
//	Value int
//	Right *Tree
//}
//
//// Walk traverses a tree depth-first,
//// sending each Value on a channel.
//func Walk(t *Tree, ch chan int) {
//	if t == nil {
//		return
//	}
//	Walk(t.Left, ch)
//	ch <- t.Value
//	Walk(t.Right, ch)
//}
//
//// Walker launches Walk in a new goroutine,
//// and returns a read-only channel of values.
//func Walker(t *Tree) <-chan int {
//	ch := make(chan int)
//	go func() {
//		Walk(t, ch)
//		close(ch)
//	}()
//	return ch
//>>>>>>> 403d2e35c4007d05cd16e1ecc71b26748698fb6e
//}
//
//// Compare reads values from two Walkers
//// that run simultaneously, and returns true
//// if t1 and t2 have the same contents.
//func Compare(t1, t2 *Tree) bool {
//	c1, c2 := Walker(t1), Walker(t2)
//	for {
//		v1, ok1 := <-c1
//		v2, ok2 := <-c2
//		if !ok1 || !ok2 {
//			return ok1 == ok2
//		}
//		if v1 != v2 {
//			break
//		}
//	}
//	return false
//}
//
//// New returns a new, random binary tree
//// holding the values 1k, 2k, ..., nk.
//func New(n, k int) *Tree {
//	var t *Tree
//	for _, v := range rand.Perm(n) {
//		t = insert(t, (1+v)*k)
//	}
//	return t
//}
//
//<<<<<<< HEAD
//func main() {
//	http.HandleFunc("/", HomeRouterHandler) // установим роутер
//	err := http.ListenAndServe(":9000", nil) // задаем слушать порт
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//
//}
//=======
//func insert(t *Tree, v int) *Tree {
//	if t == nil {
//		return &Tree{nil, v, nil}
//	}
//	if v < t.Value {
//		t.Left = insert(t.Left, v)
//		return t
//	}
//	t.Right = insert(t.Right, v)
//	return t
//}
//
//func main() {
//	t1 := New(100, 1)
//	fmt.Println(Compare(t1, New(100, 1)), "Same Contents")
//	fmt.Println(Compare(t1, New(99, 1)), "Differing Sizes")
//	fmt.Println(Compare(t1, New(100, 2)), "Differing Values")
//	fmt.Println(Compare(t1, New(101, 2)), "Dissimilar")
//	fmt.Println(runtime.NumCPU())
//}
//>>>>>>> 403d2e35c4007d05cd16e1ecc71b26748698fb6e
