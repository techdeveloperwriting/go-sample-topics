# Json parsing

Go offers built-in support for JSON encoding and decoding, including to and from built-in and custom data types.

Package json implements encoding and decoding of JSON as defined in RFC 7159. The mapping between JSON and Go values is described in the documentation for the Marshal and Unmarshal functions.

* Unmarshal
Once we’ve used the os.Open function to read our file into memory, we then have to convert it toa byte array using ioutil.ReadAll. Once it’s in a byte array we can pass it to our json.Unmarshal() method.


# Difference bertween YAML and JSON

* JSON syntax is somewhat cumbersome as it uses special syntax like curly braces {} and square brackets  []  to represent objects and arrays.
* YAML's syntax looks a bit friendlier as it uses blank spaces to denote relations between objects and ‘–‘ to represent array elements.
* There is no guarenty that json will take high size compared to the yaml, its based on how many lines of code you are writing in the file (In some cases json we can write in one line also, In that moment json will take less size compared to YAML)
* YAML allows comments by using #, a feature that is often desired when working with JSON files:
* Another feature missing in JSON but present in YAML is multi-line strings:
* JavaScript is the standard for the web, meaning that it's almost impossible to find a language that doesn't fully support JSON.

* On the other hand, YAML is widely supported, but it's not a standard. This means that libraries exist for most popular programming languages, but due to its complexity, they might not fully implement the specification.

# Run the example 

go run ./json/json.go