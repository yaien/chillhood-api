package models

type Product struct {
	Name        string
	Price       int
	Active      bool
	Tags        []string
	Pictures    []string
	Description string
	CreatedAt   int
	Sizes       []Size
}

type Size struct {
	Label     string
	Existence int
}
