package main

import (
    "bufio"
    "bytes"
    "flag"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "regexp"
	"sync"
	"context"
	"strings"
	//"time"

	"github.com/chromedp/chromedp"
)

// global variable find
var findVariables []string

func addToFindVariables(variable string) {
    for _, v := range findVariables {
        if v == variable {
            return
        }
    }
    findVariables = append(findVariables, variable)
}

func printVariable(variables []string) {
	for _, v :=range variables {
		fmt.Println(v)
	}
}

func main() {
    // Create a regular expression to match variable declarations
    varDeclaration := regexp.MustCompile(`(?:var|let|const)\s+([a-zA-Z_$][a-zA-Z0-9_$]*)\s*=?`)

    // Parse command-line flags
	threads := flag.Int("t",1,"Number of goroutines to use")
    filePath := flag.String("f", "", "path to a file containing a list of URLs")
    help := flag.Bool("h",false,"display usage information")
	silent := flag.Bool("s", false, "hide banner")
	attack := flag.Bool("a",false, "attack mode")
	flag.Parse()

	

	if !*silent {
		fmt.Println(`

$$\    $$\                    $$$$$$\                      $$\                           
$$ |   $$ |                  $$  __$$\                     $$ |                          
$$ |   $$ |$$$$$$\   $$$$$$\ $$ /  \__| $$$$$$\   $$$$$$\  $$ |  $$\  $$$$$$\   $$$$$$\  
\$$\  $$  |\____$$\ $$  __$$\\$$$$$$\  $$  __$$\ $$  __$$\ $$ | $$  |$$  __$$\ $$  __$$\ 
 \$$\$$  / $$$$$$$ |$$ |  \__|\____$$\ $$$$$$$$ |$$$$$$$$ |$$$$$$  / $$$$$$$$ |$$ |  \__|
  \$$$  / $$  __$$ |$$ |     $$\   $$ |$$   ____|$$   ____|$$  _$$<  $$   ____|$$ |      
   \$  /  \$$$$$$$ |$$ |     \$$$$$$  |\$$$$$$$\ \$$$$$$$\ $$ | \$$\ \$$$$$$$\ $$ |      
	\_/    \_______|\__|      \______/  \_______| \_______|\__|  \__| \_______|\__|      																																 
		`)
		fmt.Println("\nRemember that bug bounty and security tools should only be used ethically and responsibly.")
		fmt.Println("Misuse of these tools can lead to harm and legal consequences.")
		fmt.Println("Use these tools with caution and obtain permission before performing any testing or analysis.\n")
	}

	if *help {
		fmt.Fprintf(os.Stderr,"Usage: VarSeeker [OPTIONS]\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		return
	}


	type AttackPayload struct{
		url string
		payload string
	}

	attackJobs := make(chan AttackPayload)
	jobs := make(chan string)
	var wg sync.WaitGroup
	//var attackWg sync.WaitGroup
    // Create a slice to hold the URLs
    var urls []string

    // Read URLs from standard input if no file was specified with -f
    if *filePath == "" {
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            urls = append(urls, scanner.Text())
        }
        if err := scanner.Err(); err != nil {
            fmt.Fprintf(os.Stderr, "Error reading URLs from standard input: %v\n", err)
            os.Exit(1)
        }
    } else {
        // Read URLs from the specified file
        data, err := ioutil.ReadFile(*filePath)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", *filePath, err)
            os.Exit(1)
        }
        scanner := bufio.NewScanner(bytes.NewReader(data))
        for scanner.Scan() {
            urls = append(urls, scanner.Text())
        }
        if err := scanner.Err(); err != nil {
            fmt.Fprintf(os.Stderr, "Error reading URLs from file %s: %v\n", *filePath, err)
            os.Exit(1)
        }
    }

	//fmt.Println(urls)
    // Iterate over the URLs
	for i := 0; i < *threads; i++ {
		wg.Add(1)
		go func() {
			for url := range jobs {
				// Send a GET request to the URL
				resp, err := http.Get(url)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error fetching %s: %v\n", url, err)
					continue
				}
				defer resp.Body.Close()
	
				// Read the response body
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error reading response body for %s: %v\n", url, err)
					continue
				}
	
				// Find all variable declarations in the response body
				matches := varDeclaration.FindAllStringSubmatch(string(body), -1)
				for _, match := range matches {
					addToFindVariables(match[1])
					if (*attack){
						ap := AttackPayload{url:url,payload:match[1]}
						attackJobs <- ap
					}
				}
				
			}
			wg.Done()
		}()

	}


	

	
	
	go func() {
		for value := range attackJobs {
			attack_url := value.url+"?"+value.payload+"=spyx"
			ctx, cancel := chromedp.NewContext(context.Background())
			defer cancel()
			var responseBody string
			err := chromedp.Run(ctx, chromedp.Navigate(attack_url), chromedp.OuterHTML("html", &responseBody))
			if err != nil {
				panic(err)
			}
			if strings.Contains(responseBody, "spyx") {
				fmt.Println(attack_url)
			}
		}
	}()
	


	for _, url := range urls {
		jobs <- url
		
	}
	close(jobs)
	wg.Wait()
	
	if !(*attack){
		printVariable(findVariables)
	}
}

 