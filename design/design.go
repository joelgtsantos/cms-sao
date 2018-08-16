package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("SAO v1", func() {
	Title("Sao")
	Description("Exposes CMS platform entry and score resources")
	Version("1.0")
	Host("localhost:8080")
	Scheme("http")
	BasePath("/sao/v1")
})

var _ = Resource("entry", func() {
	Description("A document to be evaluated and ranked")
	BasePath("/entries")
	Response(Unauthorized, ErrorMedia)
	Response(BadRequest, ErrorMedia)

	Action("show", func() {
		Description("List all the entries, which is not allowed in the current implementation")
		Routing(GET("/"))
		Response(Forbidden)
	})

	Action("create", func() {
		Description("Create a new entry")
		Routing(POST("/"))
		Payload(EntryPayload)
		Response(Accepted)
	})

	Action("submit", func() {
		Description("Submit a new entry as multipart form")
		MultipartForm()
		Routing(POST("/"))
		Payload(EntryFormPayload)
		Response(Accepted)
	})
})

var _ = Resource("testentry", func() {
	Description("A document to be only evaluated")
	BasePath("/testentries")
	Response(Unauthorized, ErrorMedia)
	Response(BadRequest, ErrorMedia)

	Action("show", func() {
		Description("List all the entries, which is not allowed in the current implementation")
		Routing(GET("/"))
		Response(Forbidden)
	})

	Action("create", func() {
		Description("Create a new entry")
		Routing(POST("/"))
		Payload(TestEntryPayload)
		Response(Accepted)
	})

	Action("submit", func() {
		Description("Submit a new test entry as multipart form")
		MultipartForm()
		Routing(POST("/"))
		Payload(TestEntryFormPayload)
		Response(Accepted)
	})
})

var _ = Resource("result", func() {
	Description("Represents an entry evaluation and grading process status")
	BasePath("/results")
	DefaultMedia(ResultMedia)
	Response(OK)
	Response(NotFound)
	Response(BadRequest, ErrorMedia)

	Action("show", func() {
		Description("List all the results delimited by the query params")
		Routing(GET("/"))
		Params(func() {
			Param("entry", Integer, "Entry ID")
			Param("user", Integer, "User ID")
			Param("task", Integer, "Task ID")
		})
	})

	Action("get", func() {
		Description("Returns an specific result with the given ID")
		Routing(GET("/:resultID"))
		Params(func() {
			Param("resultID", String, "Result ID")
		})
		Response(OK)
		Response(NotFound)
	})
})

var _ = Resource("scores", func() {
	Description("Represents an entry grading")
	BasePath("/scores")
	DefaultMedia(ScoreMedia)
	Response(OK)
	Response(NotFound)
	Response(BadRequest, ErrorMedia)

	Action("show", func() {
		Description("List all the scores delimited by the query params")
		Routing(GET("/"))
		Params(func() {
			Param("entry", Integer, "Entry ID")
			Param("user", Integer, "User ID")
			Param("task", Integer, "Task ID")
			Param("contest", Integer, "Contest ID")
		})
	})

	Action("get", func() {
		Description("Returns an specific score with the given ID")
		Routing(GET("/:scoreID"))
		Params(func() {
			Param("scoreID", String, "Score ID")
		})
	})
})

// ResultMedia defines the media type used to render resutls.
var ResultMedia = MediaType("application/vnd.com.jossemargt.sao.result+json", func() {
	Description("A bottle of wine")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique bottle ID")
		Attribute("href", String, "API href for making requests on the bottle")
		Attribute("name", String, "Name of wine")
		Required("id", "href", "name")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")   // Media types may have multiple views and must
		Attribute("href") // have a "default" view.
		Attribute("name")
	})
})

// ScoreMedia defines the media type used to render scores.
var ScoreMedia = MediaType("application/vnd.com.jossemargt.sao.score+json", func() {
	Description("A bottle of wine")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique bottle ID")
		Attribute("href", String, "API href for making requests on the bottle")
		Attribute("name", String, "Name of wine")
		Required("id", "href", "name")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")   // Media types may have multiple views and must
		Attribute("href") // have a "default" view.
		Attribute("name")
	})
})

// EntryPayload ...
var EntryPayload = Type("EntryPayload", func() {
	Attribute("name", String, "Name")
	Attribute("birthday", DateTime, "Birthday")
	Required("name", "birthday")
})

// TestEntryPayload ...
var TestEntryPayload = Type("TestEntryPayload", func() {
	Attribute("name", String, "Name")
	Attribute("birthday", DateTime, "Birthday")
	Required("name", "birthday")
})

// EntryFormPayload ...
var EntryFormPayload = Type("EntryFormPayload", func() {
	Attribute("name", String, "Name")
	Attribute("birthday", DateTime, "Birthday")
	Attribute("icon", File, "Icon")
	Required("name", "birthday", "icon")
})

// TestEntryFormPayload ...
var TestEntryFormPayload = Type("TestEntryFormPayload", func() {
	Attribute("name", String, "Name")
	Attribute("birthday", DateTime, "Birthday")
	Attribute("icon", File, "Icon")
	Required("name", "birthday", "icon")
})
