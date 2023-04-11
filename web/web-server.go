package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"moppo.com/gsbt/log"
)

// Echo web server
func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/submit", submit)
	e.GET("/list", list)

	log.GetLogger().Info("server start!")

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

var resultMap map[string]int /*创建集合 */

func init() {
	resultMap = make(map[string]int)
}

// Handler
func submit(c echo.Context) error {
	name := c.QueryParam("name")
	if resultMap[name] != 0 {
		return c.String(http.StatusOK, "不要重复提交!")
	}

	time.Sleep(6 * time.Second)
	randomInt := randInt(1, 100)
	fmt.Printf("提交人: %s:%d\n", name, randomInt)
	result := fmt.Sprintf("<font size=\"20\">你掷出了:%d</font>", randomInt)
	resultMap[name] = randomInt
	return c.HTML(http.StatusOK, result)
}

func list(c echo.Context) error {
	resultStr := ""
	r := sortMapByValue(resultMap)
	fmt.Printf("%v\n", r)
	for i := 0; i < len(r); i++ {
		user := r[i]
		fmt.Printf("%s:%d\n", user.Name, user.Value)
		resultStr = resultStr + fmt.Sprintf("<font size=\"20\" color=\"blue\">%s: %d</font><br>", user.Name, user.Value)
	}
	//for name := range resultMap {
	//	resultStr = resultStr + "\n" + fmt.Sprintf("<font size=\"20\" color=\"blue\">%s: %d</font><br>", name, resultMap[name])
	//}
	return c.HTML(http.StatusOK, resultStr)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

type Pair struct {
	Name  string
	Value int
}

type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}
