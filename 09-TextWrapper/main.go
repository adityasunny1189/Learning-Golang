package main

import "fmt"

func main() {
	text := `こんにちは。
日本語でメッセージを書かせていただきますね。
私も正確な英語を学ぶ必要があるので、お互いに言葉を学べたらと思います。
伝わる英語ではなく、正確な英語を学びたいのです。
文法的に正しいだけではなく、自然な文章を書きたいのです。
私は日本語の英作文を添削することができます。
あなたに私の書いた英語を添削してもらいたいです。
ところで、`

	const maxWidth = 10
	var lw int

	for i, r := range text {
		fmt.Printf("%c", r)
		if lw++; lw > maxWidth {
			fmt.Printf("[%d]", i)
			fmt.Println()
			lw = 0
		}
	}
}
