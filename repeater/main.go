package main

import (
	"fmt"
	"html"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(1000)))

		dice := rand.Float32()

		if dice < 0.15 {
			http.Error(w, http.StatusText(500), 500)
		} else if dice > 0.70 {
			http.Error(w, http.StatusText(400), 400)
		} else {
			fmt.Fprintf(w, `Hello, %q. How are you? I need to make this response larger than a single packet so I can test reconstructing the TCP streStill not larger enough. Need to keep writing. And writing. And keeping it going is tough. You'll see the scissors of life cut down the young ones  Lorem ipsum dolor sit amet, consectetur adipiscing elit. Ut in mi erat. Nam arcu turpis, ornare sed lorem quis, posuere consectetur enim. Sed vulputate quam non purus sodales ultrices. In blandit iaculis rutrum. Nulla quam dolor, convallis id turpis sit amet, rhoncus egestas lorem. Sed suscipit dolor sit amet elit congue viverra. Aliquam suscipit rutrum elit sit amet ullamcorper.

Praesent erat lacus, faucibus eget posuere non, vulputate ac diam. Fusce vel ipsum at quam sagittis elementum. Nulla dapibus at quam et iaculis. Sed in lobortis nibh. Aenean velit ligula, sodales vel molestie sed, venenatis eu lectus. Sed dapibus tellus vel elementum porttitor. Duis vel magna urna. Phasellus porttitor elementum odio, quis euismod nunc faucibus sed. Mauris eu tortor nibh. Mauris id mollis ligula. Sed fermentum, nibh sit amet pharetra elementum, augue purus faucibus dui, a facilisis purus risus at est. Aliquam vestibulum accumsan erat ac ornare.

Vivamus congue, justo eu facilisis viverra, augue odio semper leo, vel dignissim ipsum nibh id dolor. Nulla vehicula orci nulla, et aliquet lectus consequat nec. Aliquam dignissim mauris sed LARGER PACKET eros varius, eu mollis nunc ultrices. Integer in elementum turpis. Etiam interdum dapibus aliquet. Integer euismod auctor leo, sit amet mattis nisl vestibulum eget. Nam vel urna tempus, ullamcorper neque ac, lobortis velit. Aliquam blandit, dolor quis euismod porttitor, neque nulla sollicitudin lectus, vitae ullamcorper augue arcu at ante. Duis sit amet justo enim. Donec imperdiet adipiscing mattis. Nullam tincidunt mauris est, ut condimentum sem porttitor non.`, html.EscapeString(r.URL.Path))
		}
	})

	for i := 0; i < 5; i++ {
		go func() {
			for {
				log.Println("Ping")
				resp, _ := http.Get("http://localhost:8080/bar")
				resp.Body.Close()
			}
		}()
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
