package elsearch

import (
	basConf "qiyaowu-go-zero/basic/config"
	//"gopkg.in/olivere/elastic.v7"
	"github.com/olivere/elastic/v7"
	"log"
)

var (
	e *elastic.Client
)

func Init() {
	esConfig := basConf.GetEsConfig()

	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(esConfig.GetHost()))
	if err != nil {
		log.Fatalf("es 链接失败")
	}

	log.Println("es-client", client)
	log.Println("es-err", err)
	e = client
	//检测索引是否已存在  exists= true 存在 false 不存在
	//exists, err := client.IndexExists("my_php").Do(context.Background())
	//if exists == true {
	//	fmt.Println("索引存在")
	//}

}
func GetCli() *elastic.Client {
	return e
}
