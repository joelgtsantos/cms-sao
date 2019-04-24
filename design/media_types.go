package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var EntryMedia = MediaType("application/vnd.com.jossemargt.sao.entry+json", func() {
	Description("Any source code or input to be compiled, executed and evaluated")
	Reference(AbstractEntry)

	Attributes(func() {
		Attribute("id", Integer, "Unique entry ID", func() {
			Example(1236)
		})
		Attribute("href", String, "API href for making requests on the entry", func() {
			Example("/entries/1236")
		})
		Attribute("contestSlug")
		Attribute("contestID", Integer, "Contest ID where this Entry has been submitted", func() {
			Default(0)
		})
		Attribute("taskSlug")
		Attribute("taskID", Integer, "Task ID where this Entry has been submitted", func() {
			Default(0)
		})
		Attribute("userID", Integer, "User ID of the Entry's owner", func() {
			Default(0)
		})
		Attribute("token")
		Attribute("language")
		Attribute("result", ResultMedia, "The entry processing result")

		Required("id", "href")
	})

	Links(func() {
		Link("result", "link")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("contestSlug")
		Attribute("taskSlug")
		Attribute("token")
	})

	View("full", func() {
		Attribute("id")
		Attribute("href")
		Attribute("contestID")
		Attribute("contestSlug")
		Attribute("taskID")
		Attribute("taskSlug")
		Attribute("userID")
		Attribute("token")
		Attribute("language")
		Attribute("links")
	})
})

var DraftMedia = MediaType("application/vnd.com.jossemargt.sao.draft+json", func() {
	Description("Any source code or input to be compiled and executed against the user test case")
	Reference(AbstractEntry)

	Attributes(func() {
		Attribute("id", Integer, "Unique entry ID", func() {
			Example(1236)
		})
		Attribute("href", String, "API href for making requests on the entry", func() {
			Example("/drafts/1236")
		})
		Attribute("contestSlug")
		Attribute("contestID", Integer, "Contest ID where this Entry has been submitted", func() {
			Default(0)
		})
		Attribute("taskSlug")
		Attribute("taskID", Integer, "Task ID where this Entry has been submitted", func() {
			Default(0)
		})
		Attribute("userID", Integer, "User ID of the Entry's owner", func() {
			Default(0)
		})
		Attribute("language")
		Attribute("result", DraftResultMedia, "The entry processing result")

		Required("id", "href")
	})

	Links(func() {
		Link("result", "link")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("contestSlug")
		Attribute("taskSlug")
	})

	View("full", func() {
		Attribute("id")
		Attribute("href")
		Attribute("contestID")
		Attribute("contestSlug")
		Attribute("taskID")
		Attribute("taskSlug")
		Attribute("userID")
		Attribute("language")
		Attribute("links")
	})
})

var ResultMedia = MediaType("application/vnd.com.jossemargt.sao.result+json", func() {
	Description("The representation of the result of an entry compile, evaluation and grading process")
	Attributes(func() {
		Attribute("id", String, "Compound Result ID", func() {
			Example("1236-5689")
		})
		Attribute("href", String, "API href for making requests on the result", func() {
			Example("/results/1236-5689")
		})
		Attribute("compilation", CompilationResult, "Entry compilation result")
		Attribute("evaluation", EvaluationResult, "Entry evaluation result")
		Attribute("score", ScoreResult, "Entry graded score")

		Required("id", "href", "evaluation")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("evaluation")
		Attribute("score")
	})

	View("full", func() {
		Attribute("id")
		Attribute("href")
		Attribute("compilation")
		Attribute("evaluation")
		Attribute("score")
	})
})

var DraftResultMedia = MediaType("application/vnd.com.jossemargt.sao.draft-result+json", func() {
	Description("The representation of the result of an entry draft compile, execution and evaluation process")
	Attributes(func() {
		Attribute("id", String, "Compound Result ID", func() {
			Example("1236-5689")
		})
		Attribute("href", String, "API href for making requests on the result", func() {
			Example("/draft-results/1236-5689")
		})
		Attribute("compilation", CompilationResult, "Entry compilation result")
		Attribute("execution", ExecutionResult, "Entry execution result")
		Attribute("evaluation", EvaluationResult, "Entry evaluation result")

		Required("id", "href", "execution")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("execution")
	})

	View("full", func() {
		Attribute("id")
		Attribute("href")
		Attribute("compilation")
		Attribute("execution")
		Attribute("evaluation")
	})
})

var ScoreSumMedia = MediaType("application/vnd.com.jossemargt.sao.score-sum+json", func() {
	Description("The representation of a summarized entry's score")
	Attributes(func() {
		Attribute("contestID", Integer, "Contest Identifier associated with this score", func() {
			Example(1)
			Minimum(1)
		})
		Attribute("userID", Integer, "Contest Identifier associated with this score", func() {
			Example(1)
			Minimum(1)
		})
		Attribute("username", String, "Contest Identifier associated with this score", func() {
			Example("foobar")
		})
		Attribute("taskID", Integer, "Contest Identifier associated with this score", func() {
			Example(1)
			Minimum(1)
		})
		Attribute("taskValue", Number, "The graded value relative to the Task total score", func() {
			Example(100.00)
		})
		Attribute("contestValue", Number, "The graded value relative to the contest", func() {
			Example(25.00)
		})

		Required("contestValue", "taskValue")
	})

	View("default", func() {
		Attribute("contestID")
		Attribute("userID")
		Attribute("taskID")
		Attribute("taskValue")
		Attribute("contestValue")
	})
})
