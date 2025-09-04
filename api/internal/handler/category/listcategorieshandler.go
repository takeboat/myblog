package category

import (
	"net/http"

	"blog/api/internal/logic/category"
	"blog/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListCategoriesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := category.NewListCategoriesLogic(r.Context(), svcCtx)
		resp, err := l.ListCategories()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
