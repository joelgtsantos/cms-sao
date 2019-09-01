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
	Attribute("language", String,
		`Identifies the programming language used in the entry's content. The special keyword "none" should be used 
		instead when submitting plain text, which are used for user test inputs and  diff based grading`, func() {
			Default("")
			Example("Python 3")
		})
	Attribute("token", Boolean,
		`Identifies when an Entry has been processed using a CMS Entry Token. The default value is true, in other words 
		any submitted Entry will use a CMS Token`, func() {
			Default(true)
		})
})

var SourceType = Type("EntrySource", func() {
	Description("Entry's embed type which represents a source file")
	Attribute("filename", String,
		`Source file name including its extension. This field's value should comply with the name format (fileid) 
		constraint declared by the Task resource. Taking the "batch.%l" format as example, the valid source code file 
		names could be "batch.py", "batch.cpp" or "batch.js"`,
		func() {
			Example("my_solution.py")
			Default("")
		})
	Attribute("fileid", String,
		`Also known as filepattern, and is expected to be sent along with the filename. This field is defined by the 
		Task resource`,
		func() {
			Example("my_solution.%l")
			Example("batch.cpp")
			Default("")
		})
	Attribute("content", String, "Source content", func() {
		Default("")
	})
	Attribute("language", String,
		`Identifies the programming language used in the entry's content. This attribute can be ommited for "plain text" files`,
		func() {
			Example("Python 3")
			Default("")
		})
})

var EntryPayload = Type("EntryPayload", func() {
	Description("Any source code or input that should be compiled, executed or evaluated")
	Reference(AbstractEntry)
	Attribute("contestSlug")
	Attribute("taskSlug")
	Attribute("token")

	Attribute("sources", ArrayOf(SourceType), func() {
		Description("Source files representation. Within this list the source code files and input files can be sent alike.")
		MinLength(1)
	})

	Required("contestSlug", "taskSlug", "token", "sources")
})
