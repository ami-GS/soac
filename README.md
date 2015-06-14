# soac
Standard output appearance changer


### Usage
This changer can be used like "method chain"

```go
c := soac.NewChanger()
fmt.Printf("%s\n", c.Underline().White().Bblack().Apply("Hello!!"))
c.Reset() //back to default
```

### Supported colors
##### Basic 8 colors
* Black
* Red
* Green
* Yellow
* Blue
* Magenda
* Cyan
* White

##### 256 colors
The list are [here](http://www.calmar.ws/vim/256-xterm-24bit-rgb-color-chart.html)

#### TODO
* Improve codes

#### Licence
The MIT License (MIT) Copyright (c) 2015 ami-GS
