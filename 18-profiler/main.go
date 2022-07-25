package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	// 	str := `Director Karan Malhotra delivers a masala entertainer that charts the story of a rebel who is shackled more by his lower caste identity than the colonial power
	// After dragging its feet for quite a while, this week mainstream Hindi film industry joins its counterparts from the South in celebrating the good old  masala entertainer with a socialist heart. Dressed like  KGF, Karan Malhotra’s  Shamshera carries the politics of  Asuran and  Karnan on its broad shoulders .
	// With no history books to bow to, the writers don’t have to rein in their imagination. However, co-writer Neelesh Misra has shown in the Tiger series how facts can enrich a fictional story that reads like a graphic novel. Here, together with dialogue writer Piyush Mishra, who is known to provide social`

	// str := `Apart from counting words and characters, our online editor can help you to improve word choice and writing style, and, optionally, help you to detect grammar mistakes and plagiarism. To check word count, simply place your cursor into the text box above and start typing. You'll see the number of characters and words increase or decrease as you type, delete, and edit them. You can also copy and paste text from another program over into the online editor above. The Auto-Save feature will make sure you won't lose any changes while editing, even if you leave the site and come back later. Tip: Bookmark this page now.

	// Knowing the word count of a text can be important. For example, if an author has to write a minimum or maximum amount of words for an article, essay, report, story, book, paper, you name it. WordCounter will help to make sure its word count reaches a specific requirement or stays within a certain limit.

	// In addition, WordCounter shows you the top 10 keywords and keyword density of the article you're writing. This allows you to know which keywords you use how often and at what percentages. This can prevent you from over-using certain words or word combinations and check for best distribution of keywords in your writing.

	// In the Details overview you can see the average speaking and reading time for your text, while Reading Level is an indicator of the education level a person would need in order to understand the words you’re using.

	// Disclaimer: We strive to make our tools as accurate as possible but we cannot guarantee it will always be so.`

	str := `Apart from counting words and characters,
	our online editor can help you to improve word choice and writing style,
	and, optionally, help you to detect grammar mistakes and plagiarism.`
	//str := ""

	// 	str := `Hello How are you
	// I am Jiten .`

	lines := strings.Split(str, "\n")
	wc := make(chan int)
	words := make(chan string)
	done := make(chan struct{})
	wg := &sync.WaitGroup{}
	for i := 0; i < len(lines); i++ {
		wg.Add(1)
		go WordCount(lines[i], wg, wc, words)
	}
	var count = 0
	go func() {
		for {
			select {
			case c := <-wc:
				count += c
			case word := <-words:
				fmt.Println(word)
			case <-done:
				break
			}
		}
	}()
	wg.Wait()
	done <- struct{}{}
	fmt.Println(count)
}

func WordCount(line string, wg *sync.WaitGroup, wc chan<- int, words chan<- string) {
	//var words int
	hasChar := false
	si := 0
	count := 0
	for i := 0; i < len(line); i++ {
		// if string(line[i]) == "\n" {
		// 	lines++
		// }
		if string(line[i]) != " " && string(line[i]) != "\n" {
			hasChar = true
			if count == 0 {
				count++
			}
		}
		if (string(line[i]) == " " || string(line[i]) == "\n") && hasChar {
			words <- line[si:i]
			si = i
			//wc <- 1
			count++
			hasChar = false
		}
	}
	wc <- count
	wg.Done()
	// if hasChar {
	// 	words++
	// }
}
