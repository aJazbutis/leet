func notEnoughLetters(board [][]byte, word string) bool {
	freq := make(map[byte]int)
	for i := range board {
		for j := range board[i] {
			freq[board[i][j]]++
		}
	}
	for i := range word {
		v, ok := freq[word[i]]
		if !ok {
			return true
		}
		if v == 0 {
			return true
		} else {
			freq[word[i]]--
		}
	}
	return false
}

func isPresent(board [][]byte, word string, i, j, idx int) bool {
	if idx == len(word) {
		return true
	}
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) ||
		board[i][j] != word[idx] {
		return false
	}
	board[i][j] = '#'
	if isPresent(board, word, i-1, j, idx+1) || isPresent(board, word, i, j-1, idx+1) ||
		isPresent(board, word, i+1, j, idx+1) || isPresent(board, word, i, j+1, idx+1) {
		return true
	}
	board[i][j] = word[idx]
	return false
}

func exist(board [][]byte, word string) bool {
	if notEnoughLetters(board, word) {
		return false
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == word[0] {
				if isPresent(board, word, i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}
