package syntatical

// Parser parses the program
type Parser struct {
	actionTable [][]string

	stateStack []int

	out []int
}

// NewParser returns a parser from action table
func NewParser(actionTableFile string) (*Parser, error) {
	actionTable, err := buildActionTableFromFile(actionTableFile)
	if err != nil {
		return nil, err
	}

	return &Parser{
		actionTable: actionTable,
		stateStack:  []int{0},
		out:         []int{},
	}, nil
}

func buildActionTableFromFile(file string) ([][]string, error) {
	// 	f, err := os.Open(file)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	//
	// 	reader := csv.NewReader(f)
	// 	reader.Comma = '\t'
	//
	// 	for {
	// 		line, err := reader.Read()
	// 	}
	return nil, nil
}
