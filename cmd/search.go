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
    voy search --color --rank --name "portland"
    `,
    Run: func(cmd *cobra.Command, args []string) {
        color, _:= cmd.Flags().GetBool("color")
        id, _:= cmd.Flags().GetInt("id")
        name, _:= cmd.Flags().GetString("name")
        address, _:= cmd.Flags().GetString("address")
        rank, _:= cmd.Flags().GetBool("rank")
        MakeRequest(name, rank, color, id, address);
    },
}

func MakeRequest(name string, rank bool, color bool, id int, address string) {
    var data map[string]interface{}

    if (id > 0) {
        data = map[string]interface{}{
            "id": id,
        }
    } else if (rank == true) {
        data = map[string]interface{}{
            "city": name,
            "rank": true,
        }
    } else {
        data = map[string]interface{}{
            "city": name,
        }
    }

    bytesRepresentation, err := json.Marshal(data)
    if err != nil {
        log.Fatalln(err)
    }

    bytes := bytes.NewBuffer(bytesRepresentation)
    resp, err := http.Post(address, "application/json", bytes)
    if err != nil {
        log.Fatalln(err)
    }

    src, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    pre := ""
    if (color == true) {
        pre = string(pretty.Color(pretty.Pretty(src), nil))
    } else {
        pre = string(pretty.Pretty(src))
    }
    fmt.Println(pre)
}

func init() {
    var color bool
    var id int
    var address string
    var name string
    var rank bool
    rootCmd.AddCommand(searchCmd)

    // Here you will define your flags and configuration settings.

    // Cobra supports Persistent Flags which will work for this command
    // and all subcommands, e.g.:
    // searchCmd.PersistentFlags().String("foo", "", "A help for foo")

    // Cobra supports local flags which will only run when this command
    // is called directly, e.g.:
    // searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

    rootCmd.PersistentFlags().StringVarP(&address, "address", "a", "https://voyager-index.herokuapp.com/city-search", "address to send requests to.")
    rootCmd.PersistentFlags().BoolVarP(&color, "color", "c", false, "enable colored output.")
    rootCmd.PersistentFlags().IntVarP(&id, "id", "i", 0, "search based on city id.")
    rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "search based on city name.")
    rootCmd.PersistentFlags().BoolVarP(&rank, "rank", "r", false, "return rank information.")
}
