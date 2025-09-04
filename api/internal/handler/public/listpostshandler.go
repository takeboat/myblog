package public

import (
	"net/http"

	"blog/api/internal/logic/public"
	"blog/api/internal/svc"
	"blog/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListPostsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PostListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewListPostsLogic(r.Context(), svcCtx)
		resp, err := l.ListPosts(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
