package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	databases "moppo.com/gsbt/mysql"
	"os"
	"strings"
)

type MessageRecordImage struct {
	Id  int    `json:"id"`
	Url string `json:"url"`
}

func HandleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
func main() {
	// 创建OSSClient实例。
	client, err := oss.New("", "", "")
	if err != nil {
		HandleError(err)
	}
	// 获取存储空间。
	bucketName := "yourbucket"
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		HandleError(err)
	}
	guids := make([]string, 19)
	guids[0] = "b5c4d9e177224b249cb1e06616afedf8"
	guids[1] = "15e23a58629140baad7077c8ab5be0ee"
	guids[2] = "a8e642dd70144880aae741d6951ca1ec"
	guids[3] = "4655b765ff30433e8e48d08b0454a820"
	guids[4] = "bb19d2e4abc34172a0214c5fb92dc36e"
	guids[5] = "d232d34c029a4a8d8f17f6c120661896"
	guids[6] = "0a94017ef5de492d9ad9ae365f03e1f8"
	guids[7] = "2824566f62de4599bb838e6ab51a4832"
	guids[8] = "5ebb1ee9fc444e529fb51d166e3e1ff5"
	guids[9] = "42de5ae40648449a82f106d5d415ddc3"
	guids[10] = "77f10b104f01478b8b9b92d8299ff667"
	guids[11] = "9dde61a79c1d43ffadf3515b83166ed4"
	guids[12] = "e55516dd8fcf4f34bb0bb72a54a73ba6"
	guids[13] = "aeef424cc80542bbadfa7aed853f6896"
	guids[14] = "9731a9e70c7b4d07bd476240b5769ef8"
	guids[15] = "e285c2d6af9740d5bab3e12efcf22b67"
	guids[16] = "7e607c23953442a4bc77c89ba295fb79"
	guids[17] = "53ef46d006eb49859d43bd065b20c4ab"
	guids[18] = "fb6af7ab1c0d4c59833f415419476f50"

	for i := 0; i < len(guids); i++ {
		guid := guids[i]
		// 列举文件。
		marker := guid + "/image"
		baseUrl := "http://yourbucket.oss-ap-southeast-1.aliyuncs.com/"

		for {
			lsRes, err := bucket.ListObjects(oss.Marker(marker))
			if err != nil {
				HandleError(err)
			}
			// 打印列举文件，默认情况下一次返回100条记录。
			for _, object := range lsRes.Objects {
				fullUrl := baseUrl + object.Key
				if !strings.Contains(fullUrl, "image") {
					break
				}
				mri := MessageRecordImage{
					Url: fullUrl,
				}
				fmt.Println(fullUrl)
				databases.DB.Select("Url").Create(&mri)
			}
			if lsRes.IsTruncated {
				marker = lsRes.NextMarker
			} else {
				break
			}
		}
	}

}
