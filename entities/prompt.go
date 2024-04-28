package entities

type Prompt struct {
	ID          string   `yaml:"id"`
	Description string   `yaml:"description"`
	BasePath    string   `yaml:"base_path"`
	Files       []string `yaml:"files"`
}

type Prompts struct {
	Prompts []Prompt `yaml:"prompts"`
}
