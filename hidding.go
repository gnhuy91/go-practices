package main

import "fmt"

type Config struct {
	Env       string
	AccessKey string
	SecretKey string
}

// from https://www.youtube.com/watch?v=DMiMp2yIJhk
// func (c *Config) String() string {
// 	type c2 Config
// 	cs := c2(*c)
// 	cs.AccessKey = "(HIDDEN)"
// 	cs.SecretKey = "(HIDDEN)"
// 	return fmt.Sprintf("%+v", cs)
// }

func (c Config) String() string {
	c.AccessKey = "(HIDDEN)"
	c.SecretKey = "(HIDDEN)"
	return fmt.Sprintf("%+v %+v", c.AccessKey, c.SecretKey)
}

func main() {
	c := Config{"my ENV", "--AccessKey--", "--SecretKey--"}
	fmt.Printf("%+v\n", c)
	fmt.Println(c.AccessKey, c.SecretKey)
}
