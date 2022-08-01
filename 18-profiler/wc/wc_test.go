package wc_test

import "testing"
import "wordcount/wc"

func BenchmarkWordCount(b *testing.B) {
	str := `Director Karan Malhotra delivers a masala entertainer that charts the story of a rebel who is shackled more by his lower caste identity than the colonial power
	After dragging its feet for quite a while, this week mainstream Hindi film industry joins its counterparts from the South in celebrating the good old  masala entertainer with a socialist heart. Dressed like  KGF, Karan Malhotra’s  Shamshera carries the politics of  Asuran and  Karnan on its broad shoulders .
	With no history books to bow to, the writers don’t have to rein in their imagination. However, co-writer Neelesh Misra has shown in the Tiger series how facts can enrich a fictional story that reads like a graphic novel. Here, together with dialogue writer Piyush Mishra, who is known to provide social`
	for i := 0; i < 10000000; i++ {
		wc.WordCount(str)
	}
}
