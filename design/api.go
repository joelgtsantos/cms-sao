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
	DefaultMedia(EntryMedia)
	Response(Unauthorized, ErrorMedia)
	Response(BadRequest, ErrorMedia)

	Action("show", func() {
		Description("List all the entries without their sources.")
		Routing(GET("/"))
		Response(OK)
	})

	Action("create", func() {
		Description("Create a new entry")
		Routing(POST("/"))
		Payload(EntryPayload)
		Response(Created)
	})

	Action("submit", func() {
		Description("Submit a new entry as multipart form")
		Routing(POST("/"))
		MultipartForm()
		Payload(EntryFormPayload)
		Response(Created)
	})
})

var _ = Resource("result", func() {
	Description("Represents an entry evaluation and grading process status")
	BasePath("/results")
	DefaultMedia(ResultMedia)
	Response(BadRequest, ErrorMedia)

	Action("show", func() {
		Description("List all the results delimited by the query params")
		Routing(GET("/"))
		Params(func() {
			Param("task", Integer, "Task ID")
			Param("contest", Integer, "Contest ID")
			Param("user", Integer, "User ID")
			Param("entry", Integer, "Entry ID")
			Param("sort", func() {
				Enum("asc", "desc")
				Default("desc")
			})
			Param("ranked", func() {
				Enum("true", "false")
				Default("true")
			})
		})
		Response(OK, CollectionOf(ResultMedia))
	})

	Action("get", func() {
		Description("Returns an specific result with the given entry and testcase ID")
		Routing(GET("/:entryID-:testcaseID"))
		Params(func() {
			Param("resultID", Integer, "Result ID")
			Param("testcaseID", Integer, "Testcase ID")
			Param("ranked", func() {
				Enum("true", "false")
				Default("true")
			})
		})
		Response(OK, ResultMedia)
		Response(NotFound)
	})
})

var _ = Resource("scores", func() {
	Description("Represents an entry grading")
	BasePath("/scores")
	DefaultMedia(ScoreMedia)
	Response(BadRequest, ErrorMedia)

	Action("show", func() {
		Description("List all the scores delimited by the query params")
		Routing(GET("/"))
		Params(func() {
			Param("task", Integer, "Task ID")
			Param("contest", Integer, "Contest ID")
			Param("user", Integer, "User ID")
			Param("entry", Integer, "Entry ID")
			Param("sort", func() {
				Enum("asc", "desc")
				Default("desc")
			})
		})
		Response(OK, CollectionOf(ScoreMedia))
	})

	Action("get", func() {
		Description("Returns an specific score with the given entry and testcase ID")
		Routing(GET("/:entryID-:testcaseID"))
		Params(func() {
			Param("resultID", Integer, "Result ID")
			Param("testcaseID", Integer, "Testcase ID")
		})
		Response(OK, ScoreMedia)
		Response(NotFound)
	})
})

