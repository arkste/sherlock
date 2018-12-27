package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/arkste/sherlock/data"
	strip "github.com/grokify/html-strip-tags-go"
)

const (
	userAgent        = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:64.0) Gecko/20100101 Firefox/64.0"
	errorMessage     = "message"
	errorStatusCode  = "status_code"
	errorResponseURL = "response_url"
)

type socialNetwork struct {
	URL       string `json:"url,omitempty"`
	ErrorType string `json:"errorType,omitempty"`
	ErrorMsg  string `json:"errorMsg,omitempty"`
	ErrorURL  string `json:"errorUrl,omitempty"`
	NoPeriod  string `json:"noPeriod,omitempty"`
}
type socialNetworks map[string]*socialNetwork

var client = &http.Client{
	Timeout: time.Second * 30,
}

func successLine(name, message string) {
	if runtime.GOOS == "windows" {
		fmt.Printf("[+] %s: %s\r\n", name, message)
	} else {
		fmt.Printf("\033[37;1m[\033[92;1m+\033[37;1m]\033[92;1m %s:\033[0m %s\n", name, message)
	}
}

func errorLine(name, message string) {
	if runtime.GOOS == "windows" {
		fmt.Printf("[-] %s: %s\r\n", name, message)
	} else {
		fmt.Printf("\033[37;1m[\033[91;1m-\033[37;1m]\033[92;1m %s:\033[93;1m %s\033[0m\n", name, message)
	}
}

func isAvailable(s *socialNetwork, res *http.Response) bool {
	if s.ErrorType == errorMessage {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return true
		}
		if strings.Contains(strip.StripTags(string(bodyBytes)), s.ErrorMsg) {
			return true
		}
	} else if s.ErrorType == errorStatusCode {
		if res.StatusCode != 200 {
			return true
		}
	} else if s.ErrorType == errorResponseURL {
		if strings.Contains(res.Request.URL.String(), s.ErrorURL) {
			return true
		}
	}

	return false
}

func makeRequest(wg *sync.WaitGroup, username, name string, s *socialNetwork) {
	defer wg.Done()

	if s.NoPeriod == "True" && strings.Contains(username, ".") {
		errorLine(name, "User Name Not Allowed!")
		return
	}

	s.URL = strings.Replace(string(s.URL), "{}", username, 1)

	req, err := http.NewRequest("GET", s.URL, nil)
	if err != nil {
		errorLine(name, fmt.Sprintf("can't create request: %v", err))
		return
	}

	req.Header.Set("User-Agent", userAgent)

	res, err := client.Do(req)
	if err != nil {
		errorLine(name, fmt.Sprintf("request failed: %v", err))
		return
	}
	defer res.Body.Close()

	if isAvailable(s, res) {
		errorLine(name, fmt.Sprintf("Not Found (%s)", s.URL))
	} else {
		successLine(name, fmt.Sprintf("Found (%s)", res.Request.URL.String()))
	}
}

func sherlock(username string) {
	// get all social networks
	socialNetworks := socialNetworks{}
	err := json.Unmarshal([]byte(data.Get()), &socialNetworks)
	if err != nil {
		errorLine("JSON", "Failed to parse JSON-Data")
		os.Exit(1)
	}

	// start checking ...
	var wg sync.WaitGroup
	wg.Add(len(socialNetworks))
	for name, socialNetwork := range socialNetworks {
		go makeRequest(&wg, username, name, socialNetwork)
	}
	wg.Wait()
}

func main() {
	// Print Header
	if runtime.GOOS == "windows" {
		fmt.Println("                                              .\"\"\"-.")
		fmt.Println("                                             /      \\")
		fmt.Println(" ____  _               _            _        |  _..--'-.")
		fmt.Println("/ ___|| |__   ___ _ __| | ___   ___| |__    >.`__.-\"\"\\;\"`")
		fmt.Println("\\___ \\| '_ \\ / _ \\ '__| |/ _ \\ / __| |/ /   / /(     ^\\")
		fmt.Println(" ___) | | | |  __/ |  | | (_) | (__|   <    '-`)     =|-.")
		fmt.Println("|____/|_| |_|\\___|_|  |_|\\___/ \\___|_|\\_\\    /`--.'--'   \\ .-.")
		fmt.Println("                                           .'`-._ `.\\    | J /")
		fmt.Println("                                          /      `--.|   \\__/")
	} else {
		fmt.Println("\033[37;1m                                              .\"\"\"-.")
		fmt.Println("\033[37;1m                                             /      \\")
		fmt.Println("\033[37;1m ____  _               _            _        |  _..--'-.")
		fmt.Println("\033[37;1m/ ___|| |__   ___ _ __| | ___   ___| |__    >.`__.-\"\"\\;\"`")
		fmt.Println("\033[37;1m\\___ \\| '_ \\ / _ \\ '__| |/ _ \\ / __| |/ /   / /(     ^\\")
		fmt.Println("\033[37;1m ___) | | | |  __/ |  | | (_) | (__|   <    '-`)     =|-.")
		fmt.Println("\033[37;1m|____/|_| |_|\\___|_|  |_|\\___/ \\___|_|\\_\\    /`--.'--'   \\ .-.")
		fmt.Println("\033[37;1m                                           .'`-._ `.\\    | J /")
		fmt.Println("\033[37;1m                                          /      `--.|   \\__/\033[0m")
	}
	fmt.Println("")

	// Parse FLags
	username := flag.String("username", "", "check services with given username")
	flag.Parse()

	if *username == "" {
		// Read Username, if flags is empty
		reader := bufio.NewReader(os.Stdin)
		if runtime.GOOS == "windows" {
			fmt.Print("Username: ")
			*username, _ = reader.ReadString('\r')
		} else {
			fmt.Print("\033[37;1mUsername:\033[0m ")
			*username, _ = reader.ReadString('\n')
		}
	}

	*username = strings.ToLower(strings.Replace(strings.Trim(*username, " \r\n"), " ", "", -1))

	sherlock(*username)
}
