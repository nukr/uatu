# Watcher

```go
import github.com/nukr/uatu/watcher

w, err := watcher.New(".")
w.Start()
if err != nil {
  log.Fatal("error %v", err)
}

events := w.EventStream

for event := range events {
  fmt.Println(event.Path)
  fmt.Println(event.Type)
}
```
