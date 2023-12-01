package httpz

import (
	"context"
	"github.com/zeromicro/go-zero/rest/httpx"
	"gitlab.skig.tech/zero-core/common/header"
	"net/http"

	xhttp "github.com/zeromicro/x/http"
)

// ErrorCtx writes err into w.
func ErrorCtx(ctx context.Context, r *http.Request, w http.ResponseWriter, err error,
	fns ...func(w http.ResponseWriter, err error)) {
	if r.Header.Get(header.ContentType) != header.ApplicationJson &&
		r.Header.Get(header.ContentType) != header.JsonContentType {
		xhttp.XmlBaseResponseCtx(r.Context(), w, err)
		return
	}
	httpx.ErrorCtx(ctx, w, err, fns...)
}

// OkJsonCtx writes v into w with 200 OK.
func OkJsonCtx(ctx context.Context, r *http.Request, w http.ResponseWriter, v any) {
	if r.Header.Get(header.ContentType) != header.ApplicationJson &&
		r.Header.Get(header.ContentType) != header.JsonContentType {
		xhttp.OkXmlCtx(ctx, w, v)
		return
	}
	httpx.OkJsonCtx(ctx, w, v)
}

// Ok writes HTTP 200 OK into w.
func Ok(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}
