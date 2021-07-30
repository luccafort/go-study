package iterator

// Booker is returned book name interface.
type Booker interface {
	GetName() string
	Book(name string) bool
}

// Book is book object.
// @Java
// public class Book {
// 	private String name;
// 	public Book(String name) {
// 		this.name = name
// 	}
// 	public String getName() {
// 		return name;
// 	}
// }
type Book struct {
	name string
}

// GetName is returned book name.
func (b Book) GetName() string {
	return b.name
}

// Book is set new name.
func (b *Book) Book(name string) bool {
	if b.name == name {
		return true
	}
	b.name = name
	return true
}

// Aggregater is aggregate interface.
type Aggregater interface {
	Iterator() Iteractor
}

// BookShelf is book shelf object.
type BookShelf struct {
	books []*Book
}

// Add is adding book to books.
func (bs *BookShelf) Add(book *Book) {
	bs.books = append(bs.books, book)
}

// Iteractor is incremernt index interface.
// Javaでいうところの↓
// public interface Iterator {
// 	public abstract boolean hasNext();
// 	public abstract Object next();
// }
type Iteractor interface {
	HasNext() bool
	Next() (Book, bool)
}

// BookShelfIterator is scanned book shelf object.
type BookShelfIterator struct {
	bookShelf *BookShelf
	current   int
}

// HasNext returned
func (bsi BookShelfIterator) HasNext() bool {
	if len(bsi.bookShelf.books) > bsi.current {
		return true
	}
	return false
}

// Next is returned next book object and result.
func (bsi *BookShelfIterator) Next() (*Book, bool) {
	book := bsi.bookShelf.books[bsi.current]
	if !bsi.HasNext() {
		return book, false
	}
	bsi.current++
	return book, true
}
