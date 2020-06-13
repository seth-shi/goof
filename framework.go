package main

func GetFrameworkFiles() []file {

	return []file{
		newDirectory("commands"),

		newDirectory("configs"),

		newDirectory("database"),
		newDirectory("database/data"),
		newDirectory("database/migrations"),
		newDirectory("database/seeds"),

		newDirectory("http"),
		newDirectory("http/controllers"),
		newDirectory("http/requests"),

		newDirectory("models"),
		newDirectory("resources"),
		newDirectory("routes"),
		newDirectory("services"),
	}
}


type file interface {

	Name() string
	IsDir() bool
	Content() string
}


type directory struct {
	name string
}

func newDirectory(name string) directory {

	return directory{
		name: name,
	}
}

func (d directory) Name() string {

	return d.name
}

func (d directory) IsDir() bool {

	return true
}

func (d directory) Content() string {

	return ""
}
