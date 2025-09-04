package tag

import (
	"net/http"

	"blog/api/internal/logic/tag"
	"blog/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListTagsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := tag.NewListTagsLogic(r.Context(), svcCtx)
		resp, err := l.ListTags()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
