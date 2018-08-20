package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var ResultMedia = MediaType("application/vnd.com.jossemargt.sao.result+json", func() {
	Description("The representation of the result of an entry compile, execute or evaluation process")
	Attributes(func() {
		Attribute("id", String, "Unique result ID", func() {
			Example("re-1236-5689")
			Example("ut-1236-5689")
		})
		Attribute("href", String, "API href for making requests on the result", func() {
			Example("/result/re-1236-5689")
			Example("/result/ut-1236-5689")
		})
		Attribute("compilation", CompilationResult, "Entry compilation result")
		Attribute("execution", ArrayOf(ExecutionResult), "Entry execution result")
		Attribute("evaluation", EvaluationResult, "Entry evaluation result")
		Attribute("score", ScoreMedia, "The entry grading score if has any")

		Required("id", "href", "evaluation")
	})

	Links(func() {
		Link("score", "link")
	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("evaluation")
	})

	View("full", func() {
		Attribute("id")
		Attribute("href")
		Attribute("compilation")
		Attribute("execution")
		Attribute("evaluation")
		Attribute("links")
	})
})

var ScoreMedia = MediaType("application/vnd.com.jossemargt.sao.score+json", func() {
	Description("The representation of the entry's scoring after being evaluated")
	Attributes(func() {
		Attribute("id", String, "Unique score ID", func() {
			Example("1236-5689")
		})
		Attribute("href", String, "API href for making requests on the score", func() {
			Example("/score/1236-5689")
		})
		Attribute("untokenedValue", Number, "An un-official graded score", func() {
			Example(20.00)
		})
		Attribute("value", Number, "An official graded score with a token", func() {
			Example(10.50)
		})

	})

	View("link", func() {
		Attribute("id")
		Attribute("href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("value")
	})

	View("full", func() {
		Attribute("id")
		Attribute("href")
		Attribute("untokenedValue")
		Attribute("value")
	})
})

var EntryMedia = MediaType("application/vnd.com.jossemargt.sao.entry+json", func() {
	Description("Any source code or input to be compiled, executed and evaluated")
	Reference(AbstractEntry)

	Attributes(func() {
		Attribute("id", String, "Unique entry ID", func() {
			Example("ut-1236")
			Example("re-1236")
		})
		Attribute("href", String, "API href for making requests on the entry", func() {
			Example("/entry/re-1236")
		})
		Attribute("contestSlug")
		Attribute("taskSlug")
		Attribute("ranked")
		Attribute("result", ResultMedia, "The entry processing result")
		Attribute("score", ScoreMedia, "The entry grading score if has any")

		Required("id", "href")
	})

	Links(func() {
		Link("result", "link")
		Link("score", "link")
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
		Attribute("ranked")
	})

	View("full", func() {
		Attribute("id")
		Attribute("href")
		Attribute("contestSlug")
		Attribute("taskSlug")
		Attribute("ranked")
		Attribute("links")
	})
})

// Embedded types -----------------------------------------------------------------------------------------------------

var ExecutionResult = Type("ExecutionResult", func() {
	Description("Embedded reprensentation of an entry execution result")
	Attribute("status", String, "Execution result status", func() {
		//TODO: Update with the enum tokens used on CMS platform
		// Enum()
		Example("ok")
	})
	Attribute("time", Number, "The spent execution CPU time", func() {
		Example(0.035)
	})
	Attribute("wallClockTime", Number, "The spent execution human perceived time", func() {
		Example(0.568)
	})
	Attribute("memory", Integer, "Memory consumed", func() {
		Example(64)
	})
})

var CompilationResult = Type("CompilationResult", func() {
	Description("Embedded reprensentation of an entry compilation result")
	Attribute("status", String, "Execution result status", func() {
		//TODO: Update with the enum tokens used on CMS platform
		// Enum()
		Example("ok")
	})
	Attribute("tries", Integer, "Compilation retries", func() {
		Minimum(0)
	})
	Attribute("stdout", String, "Compilation process' standard output")
	Attribute("stderr", String, "Compilation process' standard error")
	Attribute("time", Number, "The spent execution CPU time", func() {
		Example(0.035)
	})
	Attribute("wallClockTime", Number, "The spent execution human perceived time", func() {
		Example(0.568)
	})
	Attribute("memory", Integer, "Memory consumed", func() {
		Example(64)
	})
})

var EvaluationResult = Type("EvaluationResult", func() {
	Description("Embedded reprensentation of an entry evaluation result")
	Attribute("status", String, "Execution result status", func() {
		//TODO: Update with the enum tokens used on CMS platform
		// Enum()
		Example("ok")
	})
	Attribute("tries", Integer, "Evaluation retries", func() {
		Minimum(0)
	})

})