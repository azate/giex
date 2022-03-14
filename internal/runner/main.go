package runner

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/alitto/pond"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type runner struct {
	Config     *config
	Pool       *pond.WorkerPool
	HttpClient *fasthttp.Client
}

func New(cfg *config) *runner {
	client := &fasthttp.Client{
		Name:               "GIEX/0.0",
		MaxConnWaitTimeout: 6 * time.Second,
		MaxConnDuration:    14 * time.Second,
		ReadTimeout:        6 * time.Second,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	if len(cfg.Proxy) > 0 {
		client.Dial = fasthttpproxy.FasthttpHTTPDialerTimeout(cfg.Proxy, 6*time.Second)
	}

	return &runner{
		Config: cfg,
		Pool: pond.New(
			int(cfg.MaxWorkers),
			int(cfg.MaxTasks),
			pond.Strategy(pond.Lazy()),
		),
		HttpClient: client,
	}
}

func (r *runner) do(domain string) {
	url := fmt.Sprintf("http://%s/.git/config", domain)

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(url)
	req.Header.SetMethod("GET")
	defer fasthttp.ReleaseRequest(req)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := r.HttpClient.DoRedirects(req, resp, 5); err != nil {
		log.Printf("%s: %s\n", domain, err)
		return
	} else if resp.StatusCode() != fasthttp.StatusOK {
		log.Printf("%s: HTTP status %d\n", domain, resp.StatusCode())
		return
	}

	body := resp.Body()

	if bytes.Index(body, []byte("<head>")) != -1 || bytes.Index(body, []byte("[core]")) == -1 {
		log.Printf("%s: not found git config\n", domain)
		return
	}

	if err := ioutil.WriteFile(fmt.Sprintf("%s/%s", r.Config.Output, domain), body, 0600); err != nil {
		log.Printf("%s: %s\n", domain, err)
		return
	}

	log.Printf("%s: successfully\n", domain)
}

func (r *runner) Go() error {
	defer r.Pool.StopAndWait()

	file, err := os.Open(r.Config.Input)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domain := scanner.Text()
		r.Pool.Submit(func() {
			r.do(domain)
		})
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
