package cmd

import (
    "fmt"
    "net/http"
    "log"
    "github.com/spf13/cobra"
    //"net/url"
    "io/ioutil"
    "encoding/json"
    "bytes"
    "github.com/tidwall/pretty"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
    Use:   "search",
    Short: "Search the database for cities matching a query.",
    Long: `Try searching for cities called Portland:
    voy search "Portland"
    `,
    Run: func(cmd *cobra.Command, args []string) {
        MakeRequest(args[0]);
    },
}

func MakeRequest(city string) {
    data := map[string]interface{}{
        "city": city,
    }

    bytesRepresentation, err := json.Marshal(data)
    if err != nil {
        log.Fatalln(err)
    }

    // use with local development
    // url := "http://localhost:5000/city-search"
    // use for production
    url := "http://voyager-index.herokuapp.com/city-search"

    bytes := bytes.NewBuffer(bytesRepresentation)
    resp, err := http.Post(url, "application/json", bytes)
    if err != nil {
        log.Fatalln(err)
    }

    src, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    pre := string(pretty.Color(pretty.Pretty(src), nil))
    fmt.Println(pre)
}

func init() {
    rootCmd.AddCommand(searchCmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // searchCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
