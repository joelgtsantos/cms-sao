package model

type Result struct {
	DatasetID int `db:"dataset_id"`
	EntryID   int `db:"entry_id"`
	Compilation
	Evaluation
	Scoring
}

type Compilation struct {
	Status        *string `db:"compilation_status"`
	Tries         int     `db:"compilation_tries"`
	Stdout        string  `db:"compilation_stdout"`
	Stderr        string  `db:"compilation_stderr"`
	Time          float32 `db:"compilation_time"`
	WallClockTime float32 `db:"compilation_wall_clock_time"`
	Memory        int     `db:"compilation_memory"`
}

type Evaluation struct {
	Done  bool `db:"evaluation_done"`
	Tries int  `db:"evaluation_tries"`
}

type Scoring struct {
	TaskScore    float32 `db:"score"`
	ContestScore float32 `db:"public_score"`
}
