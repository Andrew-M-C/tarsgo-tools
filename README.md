# tarsgo-tools
This is a repo of tools for [Tencent TarsGo](https://github.com/TarsCloud/TarsGo/). 

[中文文档](https://cloud.tencent.com/developer/article/1394093)

# ./config

The config package provide tools to access self-defined template configurations. For example, a user-defined configuration are like:

```xml
<tars>
    <application>
        <server>
			myStr=This is a string
            myInt=54321
            myLong=12345
            myErrorInt=abcde
        </server>
    </application>
</tars>
```

The official TarsGo tools does not provide interfaces to read informatons above, you may access them with this tool. Here is the example:

```go
import (
	"github.com/Andrew-M-C/tarsgo-tools/config"
)

func main() {
    tarsconf, err := config.NewConfig()
    if err != nil {
        fmt.Println("Failed to get config: " + err.Error())
    } else {
        myStr, exist := tarsconf.GetString("/tars/application/server", "myStr", "WHAT?")
        fmt.Printf("%t, myStr: %s\n", exist, myStr)

        myInt, exist := tarsconf.GetInt("/tars/application/server", "myInt", -1)
        fmt.Printf("%t, myInt: %d\n", exist, myInt)
        
        myInt2, exist := tarsconf.GetInt("/tars/application/server", "myInt2", -2)
        fmt.Printf("%t, myInt2: %d\n", exist, myInt2)

        myLong, exist := tarsconf.GetLong("/tars/application/server", "myLong", -3)
        fmt.Printf("%t, myLong: %d\n", exist, myLong)
        
        myErrorInt, exist := tarsconf.GetInt("/tars/application/server", "myInt", -4)
        fmt.Printf("%t, myErrorInt: %d\n", exist, myErrorInt)
    }
    return
}
```

Output:

```shell
true, myStr: This is a string
true, myInt: 54321
false, myInt2: -2
true, myLong: 12345
false, myErrorInt: -4
```

Please pay a attention that the value "`myErrorInt`" actually exists, but `Config` still returns `exist = false`. The reason is that the value could not be parsed as an integer.