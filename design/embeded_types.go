package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

const (
	cmsAsyncOK          = "ok"
	cmsAsyncFail        = "fail"
	cmsAsyncUnprocessed = "unprocessed"
)

var ExecutionResult = Type("ExecutionResult", func() {
	Description("Embedded representation of an entry execution result")
	Attribute("status", String, "Execution result status", func() {
		// cms/cms/db/submission.py:721
		Example("ok")
	})
	Attribute("time", Number, "The spent execution CPU time", func() {
		Example(0.035)
		Default(0)
	})
	Attribute("wallClockTime", Number, "The spent execution human perceived time", func() {
		Example(0.568)
		Default(0)
	})
	Attribute("memory", Integer, "Memory consumed", func() {
		Example(64)
		Default(0)
	})
	Attribute("output", String, "Execution output", func() {
		Default("")
	})
})

var CompilationResult = Type("CompilationResult", func() {
	Description("Embedded representation of an entry compilation result")
	Attribute("status", String, "Execution result status", func() {
		// cms/cms/db/submission.py:300
		Enum(cmsAsyncOK, cmsAsyncFail, cmsAsyncUnprocessed)
		Default(cmsAsyncUnprocessed)
	})
	Attribute("tries", Integer, "Compilation retries", func() {
		Default(0)
	})
	Attribute("stdout", String, "Compilation process' standard output", func() {
		Default("")
	})
	Attribute("stderr", String, "Compilation process' standard error", func() {
		Default("")
	})
	Attribute("time", Number, "The spent execution CPU time", func() {
		Example(0.035)
		Default(0)
	})
	Attribute("wallClockTime", Number, "The spent execution human perceived time", func() {
		Example(0.568)
		Default(0)
	})
	Attribute("memory", Integer, "Memory consumed", func() {
		Example(64)
		Default(0)
	})
})

var EvaluationResult = Type("EvaluationResult", func() {
	Description("Embedded representation of an entry evaluation result")
	Attribute("status", String, "Execution result status", func() {
		// cms/cms/db/submission.py:348
		Enum(cmsAsyncOK, cmsAsyncUnprocessed)
		Default(cmsAsyncUnprocessed)
	})
	Attribute("tries", Integer, "Evaluation retries", func() {
		Minimum(0)
		Default(0)
	})

})

var ScoreResult = Type("ScoreResult", func() {
	Description("Embedded representation of the entry's scoring after being evaluated")
	Attribute("taskValue", Number, "The graded value relative to the Task score", func() {
		Example(20.00)
		Default(0)
	})
	Attribute("contestValue", Number, "The graded value relative to the Contest score", func() {
		Example(10.50)
		Default(0)
	})
})
