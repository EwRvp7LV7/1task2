package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func restClient() {

	stdreader := bufio.NewScanner(os.Stdin)
	fmt.Print("c2s > ")
	for stdreader.Scan() {

		// // localAddr, err := net.ResolveIPAddr("ip", "localhost")
		// localAddr, err := net.ResolveIPAddr("ip", "127.0.0.1")
		// if err != nil {
		// 	fmt.Println("Failed to read responce", err)
		// 	return
		// }

		// netClient := &http.Client{
		// 	Transport: &http.Transport{
		// 		Proxy: http.ProxyFromEnvironment,
		// 		// DialContext: (&net.Dialer{
		// 		// 		LocalAddr: &net.TCPAddr{
		// 		// 			// IP: net.ParseIP("127.0.0.1"), //localhost
		// 		// 			// IP: net.ParseIP("localhost"),
		// 		// 			IP:localAddr.IP,
		// 		// 			Port: 8090,
		// 		// 		},
		// 		// 	Timeout:   30 * time.Second,
		// 		// 	KeepAlive: 30 * time.Second,
		// 		// 	DualStack: true,
		// 		// }).DialContext,
		// 		MaxIdleConns:          100,
		// 		IdleConnTimeout:       90 * time.Second,
		// 		TLSHandshakeTimeout:   10 * time.Second,
		// 		ExpectContinueTimeout: 1 * time.Second,
		// 	},
		// 	Timeout: time.Second * 30,
		// }

		// netClient.Transport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		// 	// if addr == "google.com:443" {
		// 	//     addr = "216.58.198.206:443"
		// 	// network
		// 	// fmt.Println("Client:network", network)
		// 	addr = "localhost:8090"
		// 	// }
		// 	dialer := &net.Dialer{
		// 		Timeout:   30 * time.Second,
		// 		KeepAlive: 30 * time.Second,
		// 		DualStack: true,
		// 	}
		// 	return dialer.DialContext(ctx, network, addr)
		// }

		// req, err := http.NewRequest("GET", "http://ione-cloud.net/hhhhh/hhhh", nil)
		// if err != nil {
		// 	fmt.Println("http.NewRequest", err)
		// }

		// fmt.Println("http.NewRequest", req.URL)
		// // fmt.Print("pppp2s > ")
		// response, err := netClient.Do(req)
		// // response, err := netClient.Get("http://ione-cloud.net/")
		// // response, err := netClient.Get("http://zero.client.net/sometestpass/" + stdreader.Text())
		// // response, err := http.Get("http://localhost:8090/sometestpass/" + stdreader.Text())
		response, err := http.Get("http://localhost:8090/sometestpass/" + stdreader.Text())
		if err != nil {
			fmt.Println("Client:Failed to get http", err)
			return
		}

		fmt.Println("response.Status", response.Status)
		fmt.Println("response.Request.URL", response.Request.URL)

		body1, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println("Client:Failed to read response.Body)")
			return
		}
		defer response.Body.Close()
		fmt.Printf("response: %v\n", string(body1))

		// sb := stripRegex(string(body1))
		// // fmt.Println(sb, in.Message)
		// index1 := strings.Index(sb, stdreader.Text())
		// if 0 < index1 {
		// 	sb = sb[index1 : index1+20]
		// } else {
		// 	sb = "Text not found!"
		// }

		// fmt.Printf("response: %v\n", string(sb))

		fmt.Print("c2s > ")

	}
}

func restClient1() {

	fmt.Println("restClient1 start!")

	response, err := http.Get("http://localhost:8090/sometestpass1/")
	if err != nil {
		fmt.Println("Client:Failed to get http", err)
		return
	}

	fmt.Println("response.Status", response.Status)
	fmt.Println("response.Request.URL", response.Request.URL)

	body1, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Client:Failed to read response.Body)")
		return
	}
	defer response.Body.Close()
	fmt.Printf("response: %v\n", string(body1))

}

func restClient2() {

	fmt.Println("restClient2 start!")

	response, err := http.Get("http://localhost:8090/sometestpass2/")
	if err != nil {
		fmt.Println("Client:Failed to get http", err)
		return
	}

	fmt.Println("response.Status", response.Status)
	fmt.Println("response.Request.URL", response.Request.URL)

	body1, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Client:Failed to read response.Body)")
		return
	}
	defer response.Body.Close()
	fmt.Printf("response: %v\n", string(body1))

}

func main() {
	// grpcClient()

	// restClient()

	go restClient1()
	<-time.After(1 * time.Second)
	restClient2()

	// httpJSONClient()
}
