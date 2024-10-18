package main

import "fmt"

func main() {
    bookmarks := map[string]string{}
    fmt.Println("Приложение для закладок")
    for {
        variant := getMenu()
        switch variant {
        case 1:
            printBookmarks(bookmarks)
        case 2:
            bookmarks = addBookmark(bookmarks)
        case 3:
            bookmarks = deleteBookmark(bookmarks)
        default:
            break
        }
    }
}
func getMenu() int {
    var variant int
    fmt.Println("Выберите ваариант")
    fmt.Println("1. Посмотреть закладки")
    fmt.Println("2. Добавить закладку")
    fmt.Println("3. Удалить закладку")
    fmt.Println("4. Выход")
    fmt.Scan(&variant)
    return variant
}

func printBookmarks(bookmarks map[string]string) {
    if len(bookmarks) == 0 {
        fmt.Println("Пока нет закладок")
    }
    for key, value := range bookmarks{
        fmt.Println(key, ": ", value)
    }
}

func addBookmark(bookmarks map[string]string)  map[string]string {
    var newBookmarkKey string
    var newBookmarkValue string
    fmt.Println("Введите название:")
    fmt.Scan(&newBookmarkKey)
    fmt.Println("Введите ссылку:")
    fmt.Scan(&newBookmarkValue)
    bookmarks[newBookmarkKey] = newBookmarkValue
    return bookmarks
}

func deleteBookmark(bookmarks map[string]string) map[string]string {
    var delBookmarkKey string
    fmt.Println("Введите название:")
    fmt.Scan(&delBookmarkKey)
    delete(bookmarks, delBookmarkKey)
    return bookmarks
}
