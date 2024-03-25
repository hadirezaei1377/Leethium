package max


func main() {
    var words string

    fmt.Println("Enter words separated by commas:")

    fmt.Scanln(&words)
    wordList := strings.Split(words, ", ")

    maxWord := findMaxWord(wordList)

    fmt.Printf("%s has max characters\n", maxWord)
}

func findMaxWord(words []string) string {
    if len(words) == 0 {
        return ""
    }
    maxWord := words[0]
    maxLen := len(maxWord)
    for _, word := range words {
        if len(word) > maxLen {
            maxWord = word
            maxLen = len(word)
        }
    }
    return maxWord
}