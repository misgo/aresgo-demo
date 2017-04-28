// Common
package Action

//	"aresgo"
//	"fmt"

type (
	GoodsInfo struct {
		GoodsId   int64
		GoodsName string
		Shop      ShopInfo
		Amount    int32
		Price     float32
	}

	ShopInfo struct {
		ShopId      int32
		ShopName    string
		ShopAddress string
		Phone       string
	}
)

func GetGoodsInfo(goodsId int64) *GoodsInfo {
	var (
		goods *GoodsInfo
	)
	if goodsId == 1 {
		goods = &GoodsInfo{
			GoodsId:   1,
			GoodsName: "测试商品1",
			Amount:    100,
			Price:     99.1,
		}

		goods.Shop = ShopInfo{
			ShopId:      1,
			ShopName:    "店铺1",
			ShopAddress: "北京市朝阳区",
		}
	} else {
		goods = &GoodsInfo{
			GoodsId:   2,
			GoodsName: "测试商品2",
			Amount:    200,
			Price:     101,
		}
		goods.Shop = ShopInfo{
			ShopId:      1,
			ShopName:    "店铺2",
			ShopAddress: "北京市海淀区",
		}
	}
	return goods

}
