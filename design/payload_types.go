package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var AbstractEntry = Type("AbstractEntry", func() {
	Description("Abstracts the common attributes from EntryPayload and EntryFormPayload")
	Attribute("contestSlug", String, "Contest unique and human readble string identifier", func() {
		Pattern("[a-zA-Z]+")
	})
	Attribute("taskSlug", String, "Task unique and human readble string identifier", func() {
		Pattern("[a-zA-Z]+")
	})
	Attribute("ranked", Boolean, "Indenties if the entry should be ranked or taken as an user test", func() {
		Default(true)
	})
})

var SourceType = Type("EntrySource", func() {
	Description("Entry's embed type which represents a source file")
	Attribute("name", String, "Source file name")
	Attribute("content", String, "Source content")
	Attribute("language", String, "Source programming languague or none when using plain text", func() {
		Example("Python 3")
		Example("none")
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

	Attribute("sources", ArrayOf(SourceType), "Source files representation", func() {
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

	Attribute("sources", ArrayOf(File), "Source files", func() {
		MinLength(1)
	})

	Required("contestSlug", "taskSlug", "ranked", "sources")
})
