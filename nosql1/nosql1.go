package main

import (
		"log"

			"github.com/dgraph-io/badger"
		)

		func main() {
			  // Open the Badger database located in the /tmp/badger directory.
			    // It will be created if it doesn't exist.
			      opts := badger.DefaultOptions
			      opts.Dir = "c:\\test\\db"
			      opts.ValueDir = "c:\\test\\db"
				    db, err := badger.Open(opts)
				      if err != nil {
					      	  log.Fatal(err)
						    }
						      defer db.Close()
						        // Your code here…
						}
