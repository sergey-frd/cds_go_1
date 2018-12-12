package lib

import (   
    "fmt"
	"github.com/boltdb/bolt"
    "os" 
	"errors"
//	"log"

    // "github.com/tealeg/xlsx"
	// "io/ioutil"

    //"encoding/gob"

    //L "cds_go_1/lib"

    //"encoding/json"

    // "time"
    // "math/rand"

//    "path/filepath"
//    "github.com/valyala/fastjson"
//    L "cds_go_1/lib"
    //S "cds_go_1/config"

)

//----------------------------------------------
func Print_DB_Bucket(byteValues []byte,  bucket_Name string) error {


    dbFileName, err := GetDbName(byteValues);  __err_panic(err) 

    if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
        fmt.Println(err)
        return err
    }

    db, err := bolt.Open(dbFileName, 0600, nil)
    if err != nil {
        fmt.Println(err)
        return err
    }
    defer db.Close()

    //fmt.Println(" ----------  Open ---------- ")
    err = db.View(func(tx *bolt.Tx) error {

			sheet_Name := string(bucket_Name)
            //fmt.Println(string(name))
			fmt.Printf(" *** Print_DB_Bucket %s *** \n", sheet_Name)
				 
			//data[sheet_Name] = make(map[string]string)
			err = db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(sheet_Name))
				c := b.Cursor()
				if b == nil {
					return errors.New("Bucket not found")
				}
				//log.Printf("Value of %s is %s", "answer", b.Get([]byte("answer")))
			
				for k, v := c.First(); k != nil; k, v = c.Next() {

                    fmt.Println( string(k[:]), "=>",string(v[:]))


				}
				return nil
			})


            return nil
    })

    if err != nil {
        fmt.Println(err)
        return err
    }
    return err

}
