// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "SAO": draftresult TestHelpers
//
// Command:
// $ goagen
// --design=github.com/jossemargt/cms-sao/design
// --notool=true
// --out=$(GOPATH)/src/github.com/jossemargt/cms-sao
// --version=v1.4.1

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/jossemargt/cms-sao/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
)

// GetDraftresultBadRequest runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetDraftresultBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.DraftresultController, resultID string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/sao/v1/draft-results/%v", resultID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["resultID"] = []string{fmt.Sprintf("%v", resultID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "DraftresultTest"), rw, req, prms)
	getCtx, _err := app.NewGetDraftresultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.Get(getCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// GetDraftresultNotFound runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetDraftresultNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.DraftresultController, resultID string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/sao/v1/draft-results/%v", resultID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["resultID"] = []string{fmt.Sprintf("%v", resultID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "DraftresultTest"), rw, req, prms)
	getCtx, _err := app.NewGetDraftresultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.Get(getCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// GetDraftresultOK runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetDraftresultOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.DraftresultController, resultID string) (http.ResponseWriter, *app.ComJossemargtSaoDraftResult) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/sao/v1/draft-results/%v", resultID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["resultID"] = []string{fmt.Sprintf("%v", resultID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "DraftresultTest"), rw, req, prms)
	getCtx, _err := app.NewGetDraftresultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Get(getCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.ComJossemargtSaoDraftResult
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.ComJossemargtSaoDraftResult)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoDraftResult", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// GetDraftresultOKFull runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetDraftresultOKFull(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.DraftresultController, resultID string) (http.ResponseWriter, *app.ComJossemargtSaoDraftResultFull) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/sao/v1/draft-results/%v", resultID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["resultID"] = []string{fmt.Sprintf("%v", resultID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "DraftresultTest"), rw, req, prms)
	getCtx, _err := app.NewGetDraftresultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Get(getCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.ComJossemargtSaoDraftResultFull
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.ComJossemargtSaoDraftResultFull)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoDraftResultFull", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// GetDraftresultOKLink runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetDraftresultOKLink(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.DraftresultController, resultID string) (http.ResponseWriter, *app.ComJossemargtSaoDraftResultLink) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/sao/v1/draft-results/%v", resultID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["resultID"] = []string{fmt.Sprintf("%v", resultID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "DraftresultTest"), rw, req, prms)
	getCtx, _err := app.NewGetDraftresultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Get(getCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.ComJossemargtSaoDraftResultLink
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.ComJossemargtSaoDraftResultLink)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoDraftResultLink", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ShowDraftresultBadRequest runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowDraftresultBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.DraftresultController, contest int, contestSlug string, entry int, page int, pageSize int, sort string, task int, taskSlug string, user int) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		query["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		query["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(entry)}
		query["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		query["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		query["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		query["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		query["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		query["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		query["user"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/sao/v1/draft-results/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		prms["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		prms["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(entry)}
		prms["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		prms["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		prms["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		prms["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		prms["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		prms["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		prms["user"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "DraftresultTest"), rw, req, prms)
	showCtx, _err := app.NewShowDraftresultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// ShowDraftresultOK runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowDraftresultOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.DraftresultController, contest int, contestSlug string, entry int, page int, pageSize int, sort string, task int, taskSlug string, user int) (http.ResponseWriter, app.ComJossemargtSaoDraftResultCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		query["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		query["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(entry)}
		query["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		query["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		query["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		query["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		query["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		query["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		query["user"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/sao/v1/draft-results/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		prms["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		prms["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(entry)}
		prms["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		prms["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		prms["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		prms["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		prms["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		prms["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		prms["user"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "DraftresultTest"), rw, req, prms)
	showCtx, _err := app.NewShowDraftresultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.ComJossemargtSaoDraftResultCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.ComJossemargtSaoDraftResultCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoDraftResultCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ShowDraftresultOKFull runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowDraftresultOKFull(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.DraftresultController, contest int, contestSlug string, entry int, page int, pageSize int, sort string, task int, taskSlug string, user int) (http.ResponseWriter, app.ComJossemargtSaoDraftResultFullCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		query["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		query["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(entry)}
		query["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		query["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		query["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		query["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		query["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		query["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		query["user"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/sao/v1/draft-results/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		prms["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		prms["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(entry)}
		prms["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		prms["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		prms["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		prms["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		prms["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		prms["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		prms["user"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "DraftresultTest"), rw, req, prms)
	showCtx, _err := app.NewShowDraftresultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.ComJossemargtSaoDraftResultFullCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.ComJossemargtSaoDraftResultFullCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoDraftResultFullCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ShowDraftresultOKLink runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowDraftresultOKLink(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.DraftresultController, contest int, contestSlug string, entry int, page int, pageSize int, sort string, task int, taskSlug string, user int) (http.ResponseWriter, app.ComJossemargtSaoDraftResultLinkCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		query["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		query["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(entry)}
		query["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		query["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		query["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		query["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		query["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		query["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		query["user"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/sao/v1/draft-results/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(contest)}
		prms["contest"] = sliceVal
	}
	{
		sliceVal := []string{contestSlug}
		prms["contest_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(entry)}
		prms["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		prms["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		prms["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		prms["sort"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(task)}
		prms["task"] = sliceVal
	}
	{
		sliceVal := []string{taskSlug}
		prms["task_slug"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(user)}
		prms["user"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "DraftresultTest"), rw, req, prms)
	showCtx, _err := app.NewShowDraftresultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.ComJossemargtSaoDraftResultLinkCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.ComJossemargtSaoDraftResultLinkCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoDraftResultLinkCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}
