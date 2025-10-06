// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"runner/internal/routes"

// 	"github.com/joho/godotenv"
// 	"github.com/valyala/fasthttp"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	// Logger sozlash
// 	fmt.Println("Server is Running.. ", os.Getenv("PORT"))
// 	if err := fasthttp.ListenAndServe(":"+os.Getenv("PORT"), routes.Router); err != nil {
// 		fmt.Printf("Server xatosi: %s\n", err)
// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	// Ping buyrug‘ini ishga tushiramiz
	cmd := exec.Command("ping", "-c", "3", "google.com")

	stdout, _ := cmd.StdoutPipe()
	cmd.Start() // Processni ishga tushurish

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		fmt.Println("Ping natija:", scanner.Text())
	}

	cmd.Wait()
}
