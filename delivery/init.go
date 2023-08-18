package delivery

import (
	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/upload-api")
	{
		grp1.GET("getUploadLists", GetUploadLists)
		grp1.POST("uploadFile/:docTypeId/:groupId", UploadFile)
		//grp1.GET("product/:id", GetProductByID)
		//grp1.PUT("product/:id", UpdateProduct)
		//grp1.DELETE("product/:id", DeleteProduct)
	}

	grp2 := r.Group("/downloadFile")
	{
		grp2.GET(":docTypeId/:groupId/:fileName", DownloadFile)
		grp2.GET("byId/:fileId", DownloadFileById)
	}

	return r
}
