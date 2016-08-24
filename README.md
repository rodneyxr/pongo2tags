# pongo2tags
A collection of my custom pongo2 tags and filters.

# static
The static tag provides an easy and flexible way to tell pongo2 where to find your static files similar to Django.

All we have to do is do something like this in **func init()**:  
`pongo2tags.StaticURL("/static")`  
or  
`pongo2tags.StaticURL("https://storage.googleapis.com/<your-cloud-bucket>/static")`

If we had the following directory structure:

    .
    ├── main.go
    ├── static
    │   ├── css
    │   │   ├── base.css
    │   └── images
    │       └── favicon.png
    └── templates
        ├── bases
        │   └── base.html
        └── index.html

**/templates/bases/base.html**  

    <link rel="stylesheet" href="{% static "/css/base.css" %}">
    <link rel="icon" type="image/png" href="{% static "/images/favicon.png" %}">

**/main.go**  
    
    package main
    
    import (
    	"net/http"
    
    	"github.com/flosch/pongo2"
    	"github.com/gin-gonic/gin"
    	"github.com/robvdl/pongo2gin"
    
    	"github.com/rodneyxr/pongo2tags"
    )
    
    func init() {
    	// Set static url
    	pongo2tags.StaticURL("/static")
    }
    
    func main() {
    	router := gin.Default()
    	router.HTMLRender = pongo2gin.Default()
    	router.Static("/static", "./static")
    
    	// Index route
    	router.GET("/", func(c *gin.Context) {
    		c.HTML(http.StatusOK, "index.html", pongo2.Context{"title": "Home"})
    	})
    
    	// Listen and serve on 0.0.0.0:8000
    	router.Run()
    }
    
If we are **debugging**, like the example above, and want to serve from the **./static** directory we can simply do:

    func init() {
      pongo2tags.StaticURL("/static")
    }
    
    func main() {
      router := gin.Default()
      router.Static("/static", "./static")
      router.Run()
    }
    
If we want to run in **production mode** and serve our static files from **somewhere else** (ex: Google Cloud Platform), we can do:

    func init() {
      pongo2tags.StaticURL("https://storage.googleapis.com/<your-cloud-bucket>/static")
    }
    
    func main() {
      router := gin.Default()
      router.Run()
    }
