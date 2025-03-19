import (
	"gator/internal/config/config.go"
	"fmt"
)

func main() {
	cfg := config.Read()
	cfg.SetUser("Carly")
	result := Read()
	fmt.Println(result)
}