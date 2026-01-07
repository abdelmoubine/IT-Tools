package main

import ( "encoding/json" "fmt" "io" "io/ioutil" "net/http" "os" "path" "strings" "time" )

type Config struct { ManagementServerURL string json:"management_server_url" SignedResourcePublicKey string json:"signed_resource_public_key" AuthToken string json:"auth_token,omitempty" }

func loadConfig(filename string) (*Config, error) { b, err := ioutil.ReadFile(filename) if err != nil { return nil, err } var c Config if err := json.Unmarshal(b, &c); err != nil { return nil, err } return &c, nil }

func httpGetBytes(url, token string) ([]byte, error) { client := &http.Client{Timeout: 30 * time.Second} req, err := http.NewRequest("GET", url, nil) if err != nil { return nil, err } if token != "" { req.Header.Set("Authorization", "token "+token) } resp, err := client.Do(req) if err != nil { return nil, err } defer resp.Body.Close() if resp.StatusCode != 200 { return nil, fmt.Errorf("HTTP %d fetching %s", resp.StatusCode, url) } return io.ReadAll(resp.Body) }

func saveFileAtomic(dir, filename string, data []byte) error { if err := os.MkdirAll(dir, 0700); err != nil { return err } tmp := path.Join(dir, filename+".tmp") if err := ioutil.WriteFile(tmp, data, 0600); err != nil { return err } final := path.Join(dir, filename) if err := os.Rename(tmp, final); err != nil { return err } return nil }

func main() { cfg, err := loadConfig("config.json") if err != nil { fmt.Println("failed to load config.json:", err) os.Exit(1) } base := cfg.ManagementServerURL if base == "" { fmt.Println("management_server_url empty in config.json; nothing to fetch") return } if !strings.HasSuffix(base, "/") { base += "/" }