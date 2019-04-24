package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("SAO", func() {
	Title("Sao v1")
	Description("Exposes CMS platform entry and score resources")
	Version("1.1")
	Host("localhost:8000")
	Scheme("http")
	BasePath("/sao/v1")
})

var _ = Resource("entry", func() {
	Description("A contestant document that has been compiled, evaluated and graded")
	BasePath("/entries")
	Response(Unauthorized, ErrorMedia)
	Response(BadRequest, ErrorMedia)

	Action("show", func() {
		Description("List the ranked entries without their sources.")
		Routing(GET("/"))
		Params(func() {
			Param("contest", Integer, "Contest ID", func() {
				Default(0)
			})
			Param("contest_slug", String, "Contest Slug", func() {
				Default("")
			})
			Param("task", Integer, "Task ID", func() {
				Default(0)
			})
			Param("task_slug", String, "Task Slug", func() {
				Default("")
			})
			Param("user", Integer, "User ID", func() {
				Default(0)
			})
			Param("page", Integer, "Page number", func() {
				Default(1)
				Minimum(1)
			})
			Param("page_size", Integer, "Item amount per page", func() {
				Default(10)
				Minimum(5)
			})
			Param("sort", String, "Sorting order", func() {
				Enum("asc", "desc")
				Default("desc")
			})
		})
		Response(OK, CollectionOf(EntryMedia))
	})

	Action("get", func() {
		Description("Get the complete entry metadata (excluding the associated sources) for the given ID")
		Routing(GET("/:entryID"))
		Params(func() {
			Param("entryID", Integer, "Entry ID", func() {
				Example(123588)
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
	Response(BadRequest, ErrorMedia)

	Action("show", func() {
		Description("List the Results delimited and grouped by contest, task, entry or user identifier")
		Routing(GET("/"))
		Params(func() {
			Param("contest", Integer, "Contest ID", func() {
				Default(0)
			})
			Param("contest_slug", String, "Contest Slug", func() {
				Default("")
			})
			Param("task", Integer, "Task ID", func() {
				Default(0)
			})
			Param("task_slug", String, "Task Slug", func() {
				Default("")
			})
			Param("user", Integer, "User ID", func() {
				Default(0)
			})
			Param("entry", Integer, "Entry ID", func() {
				Default(0)
			})
			Param("max", Boolean, "Filter the results with only their maximum score", func() {
				Default(false)
			})
			Param("view", String, "Filter result sub-schemas", func() {
				Enum("default", "score")
				Default("default")
			})
			Param("page", Integer, "Page number", func() {
				Default(1)
				Minimum(1)
			})
			Param("page_size", Integer, "Item amount per page", func() {
				Default(10)
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
		Description(`Get complete Entry Result data for the given entry and testcase ID.`)
		Routing(GET("/:resultID"))
		Params(func() {
			Param("resultID", String, "Result ID", func() {
				Pattern("\\d+-\\d+")
				Example("1235-6988")
			})
		})
		Response(OK, func() {
			Media(ResultMedia, "full")
		})
		Response(NotFound)
	})
})

var _ = Resource("draft", func() {
	Description("A contestant document draft that has been compiled and evaluated")
	BasePath("/drafts")
	Response(Unauthorized, ErrorMedia)
	Response(BadRequest, ErrorMedia)

	Action("show", func() {
		Description("List the entry drafts without their sources.")
		Routing(GET("/"))
		Params(func() {
			Param("contest", Integer, "Contest ID", func() {
				Default(0)
			})
			Param("contest_slug", String, "Contest Slug", func() {
				Default("")
			})
			Param("task", Integer, "Task ID", func() {
				Default(0)
			})
			Param("task_slug", String, "Task Slug", func() {
				Default("")
			})
			Param("user", Integer, "User ID", func() {
				Default(0)
			})
			Param("page", Integer, "Page number", func() {
				Default(1)
				Minimum(1)
			})
			Param("page_size", Integer, "Item amount per page", func() {
				Default(10)
				Minimum(5)
			})
			Param("sort", String, "Sorting order", func() {
				Enum("asc", "desc")
				Default("desc")
			})
		})
		Response(OK, CollectionOf(DraftMedia))
	})

	Action("get", func() {
		Description("Get the complete Entry Draft metadata (excluding the associated sources) for the given ID")
		Routing(GET("/:draftID"))
		Params(func() {
			Param("draftID", Integer, "Entry draft ID", func() {
				Example(123588)
			})
		})
		Response(OK, func() {
			Media(DraftMedia, "full")
		})
		Response(NotFound)
	})
})

var _ = Resource("draftresult", func() {
	Description("Represents an entry evaluation and grading process status")
	BasePath("/draft-results")
	Response(BadRequest, ErrorMedia)

	Action("show", func() {
		Description("List the Results delimited and grouped by contest, task, entry or user identifier")
		Routing(GET("/"))
		Params(func() {
			Param("contest", Integer, "Contest ID", func() {
				Default(0)
				Minimum(0)
			})
			Param("contest_slug", String, "Contest Slug", func() {
				Default("")
			})
			Param("task", Integer, "Task ID", func() {
				Default(0)
				Minimum(0)
			})
			Param("task_slug", String, "Task Slug", func() {
				Default("")
			})
			Param("user", Integer, "User ID", func() {
				Default(0)
				Minimum(0)
			})
			Param("entry", Integer, "Entry ID", func() {
				Default(0)
				Minimum(0)
			})
			Param("page", Integer, "Page number", func() {
				Default(1)
				Minimum(1)
			})
			Param("page_size", Integer, "Item amount per page", func() {
				Default(10)
				Minimum(5)
			})
			Param("sort", String, "Sorting order", func() {
				Enum("asc", "desc")
				Default("desc")
			})
		})
		Response(OK, CollectionOf(DraftResultMedia))
	})

	Action("get", func() {
		Description(`Get complete Entry Draft Result data for the given entry and testcase ID.`)
		Routing(GET("/:resultID"))
		Params(func() {
			Param("resultID", String, "Result ID", func() {
				Example("4590-1325")
				Pattern("\\d+-\\d+")
			})
		})
		Response(OK, func() {
			Media(DraftResultMedia, "full")
		})
		Response(NotFound)
	})
})

var _ = Resource("actions", func() {
	Description("All the non-REST actions supported by this API")
	BasePath("")
	Response(NotImplemented)
	Response(BadRequest, ErrorMedia)

	Action("submitEntry", func() {
		Description("Orchestrates the resource creation related to a entry submit process (Entry, Token, Result and Score).")
		Routing(POST("/submit-entry"))
		Payload(EntryPayload)

		Response(Created, func() {
			Media(EntryMedia, "full")
			Headers(func() {
				Header("Location", String, "href to created entry", func() {
					Pattern("/entries/\\d+")
					Example("/entries/124588")
				})
			})
		})
	})

	Action("submitEntryDraft", func() {
		Description("Orchestrates the resource creation related to a entry draft submit process (Draft and Result).")
		Routing(POST("/submit-draft"))
		Payload(EntryPayload)

		Response(Created, func() {
			Media(EntryMedia, "full")
			Headers(func() {
				Header("Location", String, "href to created entry", func() {
					Pattern("/drafts/\\d+")
					Example("/drafts/124588")
				})
			})
		})
	})

	Action("summarizeScore", func() {
		//TODO: Temporal "score board" resources should replace this action endpoint
		Description("List scores and its total grouped and filter by contest, task or user")
		Routing(GET("/summarize-score"))
		Params(func() {
			Param("contest", Integer, "Contest ID", func() {
				Minimum(0)
			})
			Param("task", Integer, "Task ID", func() {
				Minimum(0)
			})
			Param("user", Integer, "User ID", func() {
				Minimum(0)
			})
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
