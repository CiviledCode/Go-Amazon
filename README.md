# Go-Amazon

Go Amazon is a Query based Amazon API designed to easily execute queries in a readable, scalable, and speedy manner.

The following example prints the top 50 results price, url, and search placement to the console using the account defined inside of your build script

```golang
func main() {
 client, err := amazon.NewFromEnvionment()

	if err != nil {
		log.Fatal(err)
	}

	actions.HandleAction("top silverware 50", os.Stderr, client)
}
```

#Important!

This API requires an AWS access key and ID, which requires an AWS account. **STANDARD AMAZON ACCOUNTS WILL NOT WORK**

#Licenses and Contribution

Go-Amazon is available for anybody to use freely. No ifs, ands, or buts!

Contributing is encouraged. If you run across an issue, make a detailed issue or create a pull request. This project's maintained by myself, so if you have questions, concerns, or recommendations, shoot me a DM on Discord Civiled#1713 or email me civiled@usa.com and I will try to get back to you.
