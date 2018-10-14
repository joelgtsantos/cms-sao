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
		Description("List the ranked entries without their sources.")
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
		Description("Get the complete entry metadata (excluding the associated sources) for the given ID")
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
})

var _ = Resource("result", func() {
	Description("Represents an entry evaluation and grading process status")
	BasePath("/results")
	DefaultMedia(ResultMedia)
	Response(BadRequest, ErrorMedia)

	Action("show", func() {
		Description("List the results delimited and grouped by contest, task, entry or user identifier")
		Routing(GET("/"))
		Params(func() {
			Param("contest", Integer, "Contest ID")
			Param("task", Integer, "Task ID")
			Param("user", Integer, "User ID")
			Param("entry", Integer, "Entry ID")
			Param("ranked", Boolean, "List the ranked entries or the user tests", func() {
				Default(true)
			})
			Param("page", Integer, "Page number", func() {
				Default(1)
				Minimum(1)
			})
			Param("page_size", Integer, "Item amount per page", func() {
				Default(20)
				Minimum(5)
			})
			Param("sort", String, "Sorting order", func() {
				Enum("asc", "desc")
				Default("desc")
			})
		})
		Response(OK, CollectionOf(ResultMedia))
	})

	Action("get", func() {
		Description(`Get complete result data for the given entry and testcase ID.
						The "re" and "ut" prefix delimits the entry type as "ranked entry" or "user test" respectively.`)
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
			Param("contest", Integer, "Contest ID")
			Param("task", Integer, "Task ID")
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
			Param("sort", String, "Sorting order", func() {
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

var _ = Resource("actions", func() {
	Description("All the non RESTful http actions supported by this API")
	BasePath("")
	Response(NotImplemented)
	Response(BadRequest, ErrorMedia)

	Action("submitEntry", func() {
		//TODO: This action endpoint logic will be completely handled with middlewares, so is needed to determine the goa requirements for it
		Description("Orchestrates the resource creation related to a entry submition (Entry, Token, Result and Score).")
		Routing(POST("/submit-entry"))
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

	Action("summarizeScore", func() {
		//TODO: Temporal "score board" resources should replace this action endpoint
		Description("List scores and its total grouped and filter by contest, task or user")
		Routing(GET("/summarize-score"))
		Params(func() {
			Param("contest", Integer, "Contest ID")
			Param("task", Integer, "Task ID")
			Param("user", Integer, "User ID")
			Param("groupBy", String, "", func() {
				Enum("contest", "task", "user", "none")
				Default("none")
			})
			Param("sort", String, "Sorting order based on score value", func() {
				Enum("asc", "desc")
				Default("desc")
			})
			Param("page", Integer, "Page number", func() {
				Default(1)
				Minimum(1)
			})
			Param("page_size", Integer, "Item amount per page", func() {
				Default(20)
				Minimum(5)
			})
		})
		Response(OK, CollectionOf(ScoreSumMedia))
		Response(NotFound)
	})
})
