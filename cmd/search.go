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
        rank, _:= cmd.Flags().GetBool("rank")
        MakeRequest(args[0], rank);
    },
}

func MakeRequest(city string, rank bool) {
    var data map[string]interface{}

    if (rank == true) {
        data = map[string]interface{}{
            "city": city,
            "rank": true,
        }
    } else {
        data = map[string]interface{}{
            "city": city,
        }
    }

    bytesRepresentation, err := json.Marshal(data)
    if err != nil {
        log.Fatalln(err)
    }

    // use with local development
    // url := "http://localhost:5000/city-search"
    // use for production
    url := "https://voyager-index.herokuapp.com/city-search"

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
    var rank bool
    rootCmd.AddCommand(searchCmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // searchCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

    rootCmd.PersistentFlags().BoolVarP(&rank, "rank", "r", false, "return rank information.")
}
