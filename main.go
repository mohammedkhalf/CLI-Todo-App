package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.load(&todos)
	CmdFlags := newCmdFlags()
	CmdFlags.Execute(&todos)
	storage.save(todos)
}
