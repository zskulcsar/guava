package cmd

import(
  "regexp"
  //yml "gopkg.in/yaml.v2"
)

type GuavaConfig struct {
  ServerPort  string   `yaml:server_port`
}

func (gc *GuavaConfig) validatePort() (bool) {
  matched, err := regexp.MatchString("\\d*", gc.ServerPort)
  if err != nil {
    return false
  } else {
    return matched
  }
}

func (gc *GuavaConfig) getServerPort() (string) {
  return ":" + gc.ServerPort
}