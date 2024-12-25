package main

import (
	"fmt"
	"log"

	"github.com/Ramcache/git-helper/commit"
	"github.com/Ramcache/git-helper/git"
)

func main() {
	repoPath := "." // Путь к вашему репозиторию

	// Сканируем изменения
	changes, err := git.ScanChanges(repoPath)
	if err != nil {
		log.Fatalf("Error scanning changes: %v", err)
	}

	// Генерируем сообщение для коммита
	message, err := commit.GenerateCommitMessage(changes)
	if err != nil {
		log.Fatalf("Error generating commit message: %v", err)
	}

	fmt.Println("Commit Message:", message)
}
