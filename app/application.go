package app

import "github.com/gin-gonic/gin"

//روتر فقط و فقط در اپلیکیشن میباشد و پرایویت میباشد
var (
	router = gin.Default()
)

// StartApplication برای هر ریکوئستی که میاد باید اینجا سرو بشه و منتقل بشه به کنترلر خودش
func StartApplication() {

	MapUrls()
	router.Run(":8080")
}
