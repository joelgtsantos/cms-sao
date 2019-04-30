package model

type Result struct {
	DatasetID int
	EntryID   int
	Compilation
	Evaluation
	Scoring
}

type DraftResult struct {
	DatasetID int
	EntryID   int
	Compilation
	Evaluation
	Execution
}

type Compilation struct {
	Status        string
	Tries         int
	Stdout        string
	Stderr        string
	Time          float64
	WallClockTime float64
	Memory        int
}

type Evaluation struct {
	Done  bool
	Tries int
}

type Scoring struct {
	TaskScore    float64
	ContestScore float64
}

type Execution struct {
	Time          float64
	WallClockTime float64
	Memory        int
	Output        []byte
}
