package gormpkg

import (
	"testing"
	"time"

	pagepkg "github.com/ikaiguang/go-srv-kit/kit/page"
	"github.com/stretchr/testify/require"
)

// go test -v ./data/gorm/ -count=1 -run TestPaging
func TestPaging(t *testing.T) {
	var data = []struct {
		name          string
		pageReq       *pagepkg.PageRequest
		paginatorArgs *PaginatorArgs
	}{
		{
			name: "#分页：顺序：第 1 页",
			pageReq: &pagepkg.PageRequest{
				Page:     1,
				PageSize: 5,
			},
			paginatorArgs: &PaginatorArgs{
				PageOrders: []*Order{
					{
						Field: "id",
						Order: "asc",
					},
				},
				PageWheres: []*Where{
					{
						Field:       "id",
						Operator:    ">",
						Placeholder: DefaultPlaceholder,
						Value:       5,
					},
				},
			},
		},
		{
			name: "#分页：顺序：第 2 页",
			pageReq: &pagepkg.PageRequest{
				Page:     2,
				PageSize: 5,
			},
			paginatorArgs: &PaginatorArgs{
				PageOrders: []*Order{
					{
						Field: "id",
						Order: "asc",
					},
				},
				PageWheres: []*Where{
					{
						Field:       "id",
						Operator:    ">",
						Placeholder: DefaultPlaceholder,
						Value:       5,
					},
				},
			},
		},
		{
			name: "#分页：顺序：第 100 页",
			pageReq: &pagepkg.PageRequest{
				Page:     100,
				PageSize: 20,
			},
			paginatorArgs: &PaginatorArgs{
				PageOrders: []*Order{
					{
						Field: "id",
						Order: "asc",
					},
				},
				PageWheres: []*Where{
					{
						Field:       "id",
						Operator:    ">",
						Placeholder: DefaultPlaceholder,
						Value:       5,
					},
				},
			},
		},
		{
			name: "#分页：倒序：第 1 页",
			pageReq: &pagepkg.PageRequest{
				Page:     1,
				PageSize: 5,
			},
			paginatorArgs: &PaginatorArgs{
				PageOrders: []*Order{
					{
						Field: "id",
						Order: "desc",
					},
				},
				PageWheres: []*Where{
					{
						Field:       "id",
						Operator:    ">",
						Placeholder: DefaultPlaceholder,
						Value:       5,
					},
				},
			},
		},
		{
			name: "#分页：倒序：第 2 页",
			pageReq: &pagepkg.PageRequest{
				Page:     2,
				PageSize: 5,
			},
			paginatorArgs: &PaginatorArgs{
				PageOrders: []*Order{
					{
						Field: "id",
						Order: "desc",
					},
				},
				PageWheres: []*Where{
					{
						Field:       "id",
						Operator:    ">",
						Placeholder: DefaultPlaceholder,
						Value:       5,
					},
				},
			},
		},
		{
			name: "#分页：倒序：第 100 页",
			pageReq: &pagepkg.PageRequest{
				Page:     100,
				PageSize: 20,
			},
			paginatorArgs: &PaginatorArgs{
				PageOrders: []*Order{
					{
						Field: "id",
						Order: "desc",
					},
				},
				PageWheres: []*Where{
					{
						Field:       "id",
						Operator:    ">",
						Placeholder: DefaultPlaceholder,
						Value:       5,
					},
				},
			},
		},
	}

	for _, dd := range data {
		t.Run(dd.name, func(t *testing.T) {
			var (
				dataModels []*User
				counter    int64
			)
			// 数据库
			db := dbConn.Table((&User{}).TableName())

			// 分页
			pageReq, pageOption := pagepkg.ParsePageRequest(dd.pageReq)
			t.Logf("pageReq:  第 %d 页", pageReq.Page)

			// 条件
			db = AssembleWheres(db, dd.paginatorArgs.PageWheres)

			// 计算总数
			if db.Count(&counter).Error != nil {
				t.Errorf("分页：计算总数失败！")
				t.FailNow()
			}

			// 排序
			db = AssembleOrders(db, dd.paginatorArgs.PageOrders)
			// 分页
			db = Paginator(db, pageOption)

			// 总数为 0，无需分页！
			if counter == 0 {
				t.Log("分页：总数为 0，无需分页！")
				return
			}

			// 分页查询
			if db.Find(&dataModels).Error != nil {
				t.Errorf("分页：查询失败！")
				t.FailNow()
			}

			pageResp := pagepkg.CalcPageResponse(pageReq, uint32(counter))
			t.Log("==> pageResp.TotalNumber", pageResp.TotalNumber)
			t.Log("==> pageResp.TotalPage", pageResp.TotalPage)
			t.Log("==> pageResp.Page", pageResp.Page)
			t.Log("==> pageResp.PageSize", pageResp.PageSize)

			var idList = make([]uint64, len(dataModels))
			for i := range dataModels {
				idList[i] = dataModels[i].Id
			}
			t.Log("==> idList", idList)
		})
	}
}

// go test -v ./data/gorm/ -count=1 -run TestTest_CreateUser
func TestTest_CreateUser(t *testing.T) {
	userModel := &User{
		Id:   0,
		Name: "😊_" + time.Now().Format(time.RFC3339),
		Age:  30,
	}
	err := dbConn.Create(userModel).Error
	require.Nil(t, err)

	var dataModel = &User{}
	err = dbConn.Where("id = ?", userModel.Id).First(dataModel).Error
	require.Nil(t, err)
	t.Logf("==> dataModel : %#v\n", dataModel)
}
