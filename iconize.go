package main

import (
  "strings"
  "fmt"
  "os"
  "os/user"
  "flag"
)

func checkArgs() (string, string, string) {
  namePtr := flag.String("n", "", "Name")
  pathPtr := flag.String("p", "", "Path")
  iconPtr := flag.String("i", "", "Icon")
  flag.Parse()
  if *namePtr == "" {
    fmt.Println("Name required! (run iconize -h for help)")
    os.Exit(1)
  }
  if *pathPtr == "" {
    fmt.Println("Path required! (run iconize -h for help)")
    os.Exit(1)
  }
	return *namePtr, *pathPtr, *iconPtr
}

func createDesktopEntry(name string) *os.File {
  name = strings.ToLower(name)
  name = strings.Replace(name, " ", "-", -1) // replace all spaces
  user, err := user.Current()
  home := user.HomeDir
  path := home + "/.local/share/applications/" + name + ".desktop"
  entry, err := os.Create(path)
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  return entry
}

func write(where *os.File, what string) {
  _, err := where.WriteString(what + "\n")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func fill(entry *os.File, name, path, icon string) {
  write(entry, "[Desktop Entry]")
  write(entry, "Type=Application")
  write(entry, "Name=" + name)
  write(entry, "Exec=" + path + " %U")
  write(entry, "Icon=" + icon)
}

func close(file *os.File) {
  file.Sync()
  file.Close()
}

func main() {
  name, path, icon := checkArgs()
  entry := createDesktopEntry(name)
  fill(entry, name, path, icon)
  close(entry)
}
