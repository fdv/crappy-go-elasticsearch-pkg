# Crappy Code, don't use it!

So I've finally decided to learn Go seriously. It's the first time I'm diving into a new language since 2002, and I'm not a developer so expect crappy code, usual conventions violations and "there's a much better way to do it".

My main problem is I can't just follow tutorials so I needed a project to code on, hence this repo. So now you know.

## Usage

```
package main

import (
	"fmt"
	"github.com/fdv/crappy-go-elasticsearch-pkg"
)

func main() {
	cluster := elasticsearch.Init("http://localhost", 9000)

	fmt.Println(cluster.Geturi())
	fmt.Println(cluster.Getport())
	fmt.Println(cluster.Getsettings())
	fmt.Println(cluster.Getsetting("cluster.routing.allocation.enable"))
}
````
