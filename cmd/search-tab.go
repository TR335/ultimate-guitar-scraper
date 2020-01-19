package cmd

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "github.com/Pilfer/ultimate-guitar-scraper/pkg/ultimateguitar"
    "github.com/urfave/cli"
)

var SearchTab = cli.Command{
        Name:        "search",
        Usage:       "ug s term",
        Description: "Search tabs from ultimate-guitar.com",
        Aliases:     []string{"s"},
        Flags: []cli.Flag{
                cli.StringFlag{
                        Name:  "term",
                        Value: "cream",
                        Usage: "",
                },
        },
        Action: searcher,
}

func searcher(c *cli.Context) {
    var term string

    if c.IsSet("term") {
        term = c.String("term")
        term = strings.ReplaceAll(term, " ", "%20")
    }

    s:= ultimateguitar.New()
    fmt.Println("Searching for:", term)
    totalPages := 1 // bit of a hack?

    encountered := make(map[string]bool)
    var results []string
    for i := 0; i < totalPages; i++{
        resultPage, pageInfo, err := s.Search(term, i)
        totalPages = pageInfo.TotalPages
        if err != nil {
            break
        }
        for q := 0; q < len(resultPage.Tabs); q++{
            if encountered[resultPage.Tabs[q].SongName] == true {
            } else {
                encountered[resultPage.Tabs[q].SongName] = true
                results = append(results, resultPage.Tabs[q].SongName)
                fmt.Println(resultPage.Tabs[q].SongName)
            }
        }
        if i % 4 == 0 || i % 2 == 0{
            reader := bufio.NewReader(os.Stdin)
            fmt.Println("Continue?")
            text, _ := reader.ReadString('\n')
            if text == "\n" {
        //    } else if text == {
                
            } else {
                break
            }
        }
    }
}
