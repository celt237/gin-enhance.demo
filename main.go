package main

import (
	gin_enhance "github.com/celt237/gin-enhance"
	"github.com/celt237/gin-enhance.demo/internal"
	"github.com/celt237/gin-enhance.demo/internal/handler"
	"github.com/celt237/gin-enhance.demo/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func main() {
	app := gin.Default()

	app.GET("/", func(ctx *gin.Context) {
		htmlContent := "<h1>Welcome to Gin!</h1>"
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
	})
	//app.GET("/doc/index", cc)
	//RegisterSwaggerDoc(app, "docs/swagger.json", "/doc")
	//route := "/doc"
	//app.GET(route+"/*any", swagDocHandler(route))
	gin_enhance.RegisterSwaggerDoc(app, "docs/swagger.json", "/doc")
	println("swagger doc: http://localhost:8081/doc/index")
	demoService := service.NewDemoService()
	apiHandler := &internal.ApiHandlerImpl{}
	handler.RegisterDemoHTTPHandler(app.Group("/"), apiHandler, demoService)
	err := app.Run(":8081")
	if err != nil {
		panic(err)
	}
}

func cc(ctx *gin.Context) {
	relativePath := "/doc"
	//docJsonPath := relativePath + "/docJson"
	//servicesPath := relativePath + "/static/service"
	docPath := relativePath + "/index"
	//appjsPath := relativePath + "/static/webjars/js/app.42aa019b.js"
	if ctx.Request.Method != http.MethodGet {
		ctx.AbortWithStatus(http.StatusMethodNotAllowed)
		return
	}
	switch ctx.Request.RequestURI {
	//case appjsPath:
	//	ctx.String(http.StatusOK, appjs)
	//case servicesPath:
	//	ctx.JSON(http.StatusOK, []service{s})
	case docPath:
		doc := "<h1>doc3</h1>"
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(doc))
	//case docJsonPath:
	//	ctx.Header("Content-Type", "application/json")
	//	ctx.String(http.StatusOK, string(config.DocJson))
	default:
		//filePath := strings.TrimPrefix(ctx.Request.RequestURI, config.RelativePath)
		//filePath = strings.TrimPrefix(filePath, "/")
		//file, err := front.Open(filePath)
		//if err != nil {
		//	ctx.String(http.StatusInternalServerError, "Error while opening file: %v", err)
		//	return
		//}
		//defer file.Close()
		//fileInfo, err := file.Stat()
		//if err != nil {
		//	ctx.String(http.StatusInternalServerError, "Error while getting file info: %v", err)
		//	return
		//}
		////ctx.ServeContent(reader, fileInfo.Name(), fileInfo.ModTime())
		//ctx.DataFromReader(http.StatusOK, fileInfo.Size(), mime.TypeByExtension(filepath.Ext(fileInfo.Name())),
		//	file,
		//	map[string]string{
		//		"Content-Disposition": fmt.Sprintf("attachment; filename=%s", fileInfo.Name()),
		//		"Last-Modified":       fileInfo.ModTime().UTC().Format(http.TimeFormat),
		//	})
	}
	ctx.Next()
}

func RegisterSwaggerDoc(app *gin.Engine, jsonPath string, route string) {
	route = strings.TrimPrefix(route, "/")
	route = strings.TrimSuffix(route, "/")
	route = strings.TrimSpace(route)
	if route == "" {
		log.Println("route is empty")
		return
	}
	route = "/" + route

	if jsonPath == "" {
		log.Println(jsonPath + "jsonPath is empty")
		return
	}
	//docJson, err := os.ReadFile(jsonPath)
	//if err != nil {
	//	log.Println("no swagger.json found in " + jsonPath)
	//	return
	//}
	//strDocJson := string(docJson)
	app.GET(route+"/*any", swagDocHandler(route))
}

func swagDocHandler(path string) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		relativePath := "/doc"
		//docJsonPath := relativePath + "/docJson"
		//servicesPath := relativePath + "/static/service"
		docPath := relativePath + "/index"
		//appjsPath := relativePath + "/static/webjars/js/app.42aa019b.js"
		if ctx.Request.Method != http.MethodGet {
			ctx.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}
		switch ctx.Request.RequestURI {
		//case appjsPath:
		//	ctx.String(http.StatusOK, appjs)
		//case servicesPath:
		//	ctx.JSON(http.StatusOK, []service{s})
		case docPath:
			doc := "<h1>doc3</h1>"
			ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(doc))
		//case docJsonPath:
		//	ctx.Header("Content-Type", "application/json")
		//	ctx.String(http.StatusOK, string(config.DocJson))
		default:
			//filePath := strings.TrimPrefix(ctx.Request.RequestURI, config.RelativePath)
			//filePath = strings.TrimPrefix(filePath, "/")
			//file, err := front.Open(filePath)
			//if err != nil {
			//	ctx.String(http.StatusInternalServerError, "Error while opening file: %v", err)
			//	return
			//}
			//defer file.Close()
			//fileInfo, err := file.Stat()
			//if err != nil {
			//	ctx.String(http.StatusInternalServerError, "Error while getting file info: %v", err)
			//	return
			//}
			////ctx.ServeContent(reader, fileInfo.Name(), fileInfo.ModTime())
			//ctx.DataFromReader(http.StatusOK, fileInfo.Size(), mime.TypeByExtension(filepath.Ext(fileInfo.Name())),
			//	file,
			//	map[string]string{
			//		"Content-Disposition": fmt.Sprintf("attachment; filename=%s", fileInfo.Name()),
			//		"Last-Modified":       fileInfo.ModTime().UTC().Format(http.TimeFormat),
			//	})
		}
		ctx.Next()
	}

}
