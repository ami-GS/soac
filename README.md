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
* Black
* Red
* Green
* Yellow
* Blue
* Magenda
* Cyan
* White

#### TODO
* Set by arguments

#### Licence
The MIT License (MIT) Copyright (c) 2015 ami-GS
