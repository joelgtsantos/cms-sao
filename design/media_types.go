package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var EntryMedia = MediaType("application/vnd.com.jossemargt.sao.entry+json", func() {
	Description("Any source code or input to be compiled, executed and evaluated")
	Reference(AbstractEntry)

	Attributes(func() {
		Attribute("id", Integer, "Unique entry ID")
		Attribute("href", String, "API href for making requests on the entry")
		Attribute("contestSlug")
		Attribute("taskSlug")
		Attribute("ranked")

		Required("id", "href")
	})

	View("default", func() {
		Attribute("id")
		Attribute("contestSlug")
		Attribute("taskSlug")
		Attribute("ranked")
		Attribute("href")
	})
})

var ResultMedia = MediaType("application/vnd.com.jossemargt.sao.result+json", func() {
	Description("The representation of the result of an entry compile, execute or evaluation process")
	Attributes(func() {
		Attribute("id", String, "Unique result ID")
		Attribute("href", String, "API href for making requests on the result")
		Attribute("compilation", CompilationResult, "Entry compilation result")
		Attribute("execution", ArrayOf(ExecutionResult), "Entry execution result")
		Attribute("evaluation", EvaluationResult, "Entry evaluation result")

		Required("id", "href", "evaluation")
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
	})
})

var ScoreMedia = MediaType("application/vnd.com.jossemargt.sao.score+json", func() {
	Description("The representation of the entry's scoring after being evaluated")
	Attributes(func() {
		Attribute("id", String, "Unique bottle ID")
		Attribute("href", String, "API href for making requests on the score")
		Attribute("untokenedValue", Number, "An un-official graded score", func() {
			Example(20.00)
		})
		Attribute("value", Number, "An official graded score with a token", func() {
			Example(10.50)
		})

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

// Embedded types -----------------------------------------------------------------------------------------------------

var ExecutionResult = Type("ExecutionResult", func() {
	Description("Embedded reprensentation of an entry execution result")
	Attribute("status", String, "Execution result status") //TODO: Update with the enum tokens used on CMS platform
	Attribute("time", Number, "The spent execution CPU time", func() {
		Example(0.035)
	})
	Attribute("wallClockTime", Number, "The spent execution human perceived time", func() {
		Example(0.568)
	})
	Attribute("memory", Integer, "Memory consumed")
})

var CompilationResult = Type("CompilationResult", func() {
	Description("Embedded reprensentation of an entry compilation result")
	Attribute("status", String, "Execution result status") //TODO: Update with the enum tokens used on CMS platform
	Attribute("tries", Integer, "Compilation retries")
	Attribute("stdout", String, "Compilation process' standard output")
	Attribute("stderr", String, "Compilation process' standard error")
	Attribute("time", Number, "The spent execution CPU time", func() {
		Example(0.035)
	})
	Attribute("wallClockTime", Number, "The spent execution human perceived time", func() {
		Example(0.568)
	})
	Attribute("memory", Integer, "Memory consumed")
})

var EvaluationResult = Type("EvaluationResult", func() {
	Description("Embedded reprensentation of an entry evaluation result")
	Attribute("status", String, "Execution result status") //TODO: Update with the enum tokens used on CMS platform
	Attribute("tries", Integer, "Evaluation retries")

})