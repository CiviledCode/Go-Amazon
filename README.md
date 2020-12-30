# Go-Amazon

Go Amazon is a Query based Amazon API designed to easily execute queries in a readable, scalable, and speedy manner.

The following example prints all prices, titles, and links for all of the products found with the name silverware using the "top" command

```golang
func main() {
 client, err := amazon.NewFromEnvionment()

	if err != nil {
		log.Fatal(err)
	}

	actions.HandleAction("top silverware", os.Stderr, client)
}
```
