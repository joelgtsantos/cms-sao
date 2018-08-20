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
		Description("List all the ranked entries without their sources.")
		Routing(GET("/"))
		Params(func() {
			Param("page", Integer, "Page number", func() {
				Default(1)
				Minimum(1)
			})
			Param("page_size", Integer, "Item amount per page", func() {
				Default(20)
				Minimum(5)
			})
		})
		Response(OK, CollectionOf(EntryMedia))
	})

	Action("get", func() {
		Description("Returns all the entry metadata (without the sources) for the given ID")
		Routing(GET("/:entryID"))
		Params(func() {
			Param("entryID", String, "Result ID", func() {
				Example("ut-123588")
				Example("re-124588")
			})
		})
		Response(OK, func() {
			Media(EntryMedia, "full")
		})
		Response(NotFound)
	})

	Action("create", func() {
		Description("Create a new entry")
		Routing(POST("/"))
		Payload(EntryPayload)
		Response(Created, func() {
			Media(EntryMedia, "full")
			Headers(func() {
				Header("Location", String, "href to created entry", func() {
					Pattern("/entries/\\w{2}-\\d+")
					Example("/entries/re-124588")
				})
			})
		})
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
			Param("page", Integer, "Page number", func() {
				Default(1)
				Minimum(1)
			})
			Param("page_size", Integer, "Item amount per page", func() {
				Default(20)
				Minimum(5)
			})
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
		Routing(GET("/:resultID"))
		Params(func() {
			Param("resultID", String, "Result ID", func() {
				Example("re-1235-6988") // For ranked entries
				Example("ut-4590-1325") // For user tests
			})
		})
		Response(OK, func() {
			Media(ResultMedia, "full")
		})
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
			Param("page", Integer, "Page number", func() {
				Default(1)
				Minimum(1)
			})
			Param("page_size", Integer, "Item amount per page", func() {
				Default(20)
				Minimum(5)
			})
			Param("sort", func() {
				Enum("asc", "desc")
				Default("desc")
			})
		})
		Response(OK, CollectionOf(ScoreMedia))
	})

	Action("get", func() {
		Description("Returns an specific score with the given entry and testcase ID")
		Routing(GET("/:scoreID"))
		Params(func() {
			Param("scoreID", String, "Score ID", func() {
				Example("1234-5987")
			})
		})
		Response(OK, func() {
			Media(ScoreMedia, "full")
		})
		Response(NotFound)
	})
})

