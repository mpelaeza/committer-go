package main

type commitType struct{
	name string
	emoji string
	description string
	release bool
}

var COMMIT_TYPES = []commitType{
  commitType{
		name: "feat",
    emoji: "🆕",
    description: "Add new feature",
    release: true},
  commitType{
		name: "fix",
    emoji: "🐛",
    description: "Submit a fix to a bug",
    release: true},
  commitType{
		name: "perf",
    emoji: "⚡",
    description: "Improve performance",
    release: true},
  commitType{
		name: "refactor",
    emoji: "🛠️ ",
    description: "Refactor code",
    release: false},
  commitType{
		name: "docs",
    emoji: "📚",
    description: "Add or update documentation",
    release: false},
  commitType{
		name: "test",
    emoji: "🧪",
    description: "Add or update tests",
    release: false},
   commitType{
		name: "build",
    emoji: "🏗️ ",
    description: "Change build files",
    release: false},
}