ZetCode
All Go Python C# Java JavaScript Subscribe
Go chromedp
last modified February 8, 2022

Go chromedp tutorial shows how to automate browsers in Golang with chromedp.

The chromedp is a Go library which provides a high-level API to control Chromium over the DevTools Protocol. It allows to use a browser in a headless mode (the default mode), which works without the UI. This is great for scripting.

Get outer HTML
The chromedp.OuterHTML retrieves the outer HTML of the first element node matching the selector.

html.go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/chromedp/chromedp"
)

func main() {

    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    url := "http://webcode.me"

    var data string

    if err := chromedp.Run(ctx,

        chromedp.Navigate(url),
        chromedp.OuterHTML("html", &data, chromedp.ByQuery),
    ); err != nil {

        log.Fatal(err)
    }

    fmt.Println(data)
}
The example retrieves the home page of webcode.me.

ctx, cancel := chromedp.NewContext(context.Background())
defer cancel()
The chromedp.NewContext creates a chromedp context from the parent context. The returned cancellation function must be called to terminate the chromedp context; the function waits for the resources to be cleaned up, and returns any error encountered during that process.

if err := chromedp.Run(ctx,

    chromedp.Navigate(url),
    chromedp.OuterHTML("html", &data, chromedp.ByQuery),
); err != nil {
    
    log.Fatal(err)
}
The Run function runs actions against the context. We navigate to the URL and retrieve the HTML data of the html tag.

Get title
With chromedp.Title, we get the document title.

title.go
package main

import (
    "context"
    "fmt"
    "io"
    "log"
    "net/http"
    "net/http/httptest"
    "strings"

    "github.com/chromedp/chromedp"
)

func writeHTML(content string) http.Handler {

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        w.Header().Set("Content-Type", "text/html")
        io.WriteString(w, strings.TrimSpace(content))
    })
}

func main() {

    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    ts := httptest.NewServer(writeHTML(`
<head>
    <title>Home page</title>
</head>
<body>
    <p>Hello there!</a>
</body>
    `))

    defer ts.Close()

    var title string

    if err := chromedp.Run(ctx,

        chromedp.Navigate(ts.URL),
        chromedp.Title(&title),
    ); err != nil {

        log.Fatal(err)
    }

    fmt.Println(title)
}
In the example, we create our built-in web server that sends a tiny web page. We retrieve its title.

return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "text/html")
    io.WriteString(w, strings.TrimSpace(content))
})
The handler function sends the text/html content back to the client.

ts := httptest.NewServer(writeHTML(`
<head>
    <title>Home page</title>
</head>
<body>
    <p>Hello there!</a>
</body>
    `))
A testing server is created with httptest.NewServer.

defer ts.Close()
At the end of the program, the test server is closed.

if err := chromedp.Run(ctx,

    chromedp.Navigate(ts.URL),
    chromedp.Title(&title),
); err != nil {

    log.Fatal(err)
}
We navigate to the test server's URL and get the document title with chromedp.Title.

$ go run title.go 
Home page
Setting timeouts
We can experience deadlocks with our tasks. To prevent them, we can set up timeouts.

timeout.go
package main

import (
    "context"
    "fmt"
    "log"
    "strings"
    "time"

    "github.com/chromedp/chromedp"
)

func main() {

    ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
    defer cancel()

    url := "http://webcode.me"

    var res string

    err := chromedp.Run(ctx,

        chromedp.Navigate(url),
        chromedp.Text("body", &res, chromedp.NodeVisible),
    )

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(strings.TrimSpace(res))
}
In the example, we get the visible text of the body tag. A timeout is set.

ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
defer cancel()
We set a timeout of 5 seconds.

err := chromedp.Run(ctx,

    chromedp.Navigate(url),
    chromedp.Text("body", &res, chromedp.NodeVisible),
)
We run a task list. We get the text of body with chromedp.Text.

Click action
A click query action is performed with chromedp.Click.

click.go
package main

import (
    "context"
    "log"
    "time"

    "github.com/chromedp/chromedp"
    "github.com/chromedp/chromedp/device"
)

func main() {
    
    ctx, cancel := chromedp.NewContext(
        context.Background(),
    )
    defer cancel()

    ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
    defer cancel()

    url := "http://webcode.me/click.html"

    var ua string

    err := chromedp.Run(ctx,

        chromedp.Emulate(device.IPhone11),
        chromedp.Navigate(url),
        chromedp.Click("button", chromedp.NodeVisible),
        chromedp.Text("#output", &ua),
    )

    if err != nil {
        log.Fatal(err)
    }

    log.Printf("User agent: %s\n", ua)
}
In the example, we click on a button of a web page. The web page shows the client's user agent in the output div.

err := chromedp.Run(ctx,

    chromedp.Emulate(device.IPhone11),
    chromedp.Navigate(url),
    chromedp.Click("button", chromedp.NodeVisible),
    chromedp.Text("#output", &ua),
)
In the task list, we navigate to the URL, click on the button, and retrieve the text output. We get our user agent. We emulate an IPhone11 device with chromedp.Emulate.

Create screenshot
We can create a screenshot of an element with chromedp.Screenshot. The chromedp.FullScreenshot takes a screenshot of the entire browser viewport.

screenshot.go
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "log"

    "github.com/chromedp/chromedp"
)

func main() {
    
    ctx, cancel := chromedp.NewContext(
        context.Background(),
    )
    
    defer cancel()

    url := "http://webcode.me"

    var buf []byte
    if err := chromedp.Run(ctx, ElementScreenshot(url, "body", &buf)); err != nil {
        log.Fatal(err)
    }

    if err := ioutil.WriteFile("body.png", buf, 0o644); err != nil {
        log.Fatal(err)
    }

    if err := chromedp.Run(ctx, FullScreenshot(url, 90, &buf)); err != nil {
        log.Fatal(err)
    }

    if err := ioutil.WriteFile("full.png", buf, 0o644); err != nil {
        log.Fatal(err)
    }

    fmt.Println("screenshots created")
}

func ElementScreenshot(url, sel string, res *[]byte) chromedp.Tasks {

    return chromedp.Tasks{

        chromedp.Navigate(url),
        chromedp.Screenshot(sel, res, chromedp.NodeVisible),
    }
}

func FullScreenshot(url string, quality int, res *[]byte) chromedp.Tasks {

    return chromedp.Tasks{

        chromedp.Navigate(url),
        chromedp.FullScreenshot(res, quality),
    }
}
We create two screenshots. The images are written to disk with ioutil.WriteFile.

Submit form
The chromedp.SendKeys is used to fill in the input fields. A form can be submitted either with chromedp.Click or chromedp.Submit.

submit.go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "github.com/chromedp/chromedp"
)

func main() {

    ctx, cancel := chromedp.NewContext(
        context.Background(),
    )
    defer cancel()

    ctx, cancel = context.WithTimeout(ctx, 6*time.Second)
    defer cancel()

    url := "http://webcode.me/submit/"
    var res string

    err := chromedp.Run(ctx,

        chromedp.Navigate(url),
        chromedp.SendKeys("input[name=name]", "Lucia"),
        chromedp.SendKeys("input[name=message]", "Hello!"),
        // chromedp.Click("button", chromedp.NodeVisible),
        chromedp.Submit("input[name=name]"),
        chromedp.Text("*", &res),
    )

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(res)
}
The example fills in a form and receives a message.

chromedp.SendKeys("input[name=name]", "Lucia"),
chromedp.SendKeys("input[name=message]", "Hello!"),
We set two strings to the specified input tags.

// chromedp.Click("button", chromedp.NodeVisible),
chromedp.Submit("input[name=name]"),
We can submit the form either with chromedp.Click or chromedp.Submit. In the latter case, the function submits the parent form of the first element node matching the selector.

In this tutorial, we have automated browsers in Go with chromedp.

List all Go tutorials.

Home Facebook Twitter Github Subscribe Privacy
© 2007 - 2022 Jan Bodnar admin(at)zetcode.com