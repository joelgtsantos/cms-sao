package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var AbstractEntry = Type("AbstractEntry", func() {
	Description("Abstracts the common attributes from EntryPayload and EntryFormPayload")
	Attribute("contestSlug", String, "Contest unique and human readable string identifier", func() {
		Pattern("[_a-zA-Z0-9\\-]+")
		Example("con_test")
		Default("")
	})
	Attribute("taskSlug", String, "Task unique and human readable string identifier", func() {
		Pattern("[_a-zA-Z0-9\\-]+")
		Example("simpleBatch-95")
		Default("")
	})
	Attribute("ranked", Boolean,
		`Identifies when an Entry has been processed using a CMS Entry Token. The default value is true, in other words 
		any submitted Entry will use a CMS Token`, func() {
			Default(true)
		})
})

var SourceType = Type("EntrySource", func() {
	Description("Entry's embed type which represents a source file")
	Attribute("name", String,
		`Source file name including its extension. This field's value should comply with the name format constraint 
		declared by the task resource. Taking the "batch.%l" format as example, the valid source code file names could 
		be "batch.py", "batch.cpp" or "batch.js"`,
		func() {
			Example("my_solution.py")
		})
	Attribute("content", String, "Source content")
	Attribute("language", String,
		`Identifies the programming language used in the entry's content. The special keyword "none" should be used 
		instead when submitting plain text, which are used for user test inputs and  diff based grading`, func() {
			Example("none")
			Example("Python 3")
		})
	Attribute("encoding", String, "Source content's encoding", func() {
		Default("utf8")
	})
})

var EntryPayload = Type("EntryPayload", func() {
	Description("Any source code or input that should be compiled, executed or evaluated")
	Reference(AbstractEntry)
	Attribute("contestSlug")
	Attribute("taskSlug")
	Attribute("ranked")

	Attribute("sources", ArrayOf(SourceType), func() {
		Description("Source files representation. Within this list the source code files and input files can be sent alike.")
		MinLength(1)
	})

	Required("contestSlug", "taskSlug", "ranked", "sources")
})

var EntryFormPayload = Type("EntryFormPayload", func() {
	Description("Any source code or input that should be compiled, executed or evaluated")
	Reference(AbstractEntry)
	Attribute("contestSlug")
	Attribute("taskSlug")
	Attribute("ranked")

	Attribute("sources", ArrayOf(File), func() {
		Description(`Source files representation. Within this list the source code files and input files can be 
						  sent alike.`)
		MinLength(1)
	})

	Required("contestSlug", "taskSlug", "ranked", "sources")
})
