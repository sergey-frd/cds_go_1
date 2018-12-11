package lib

import (   
    "fmt"
    "os" 
    "github.com/boltdb/bolt"
    // "errors"
//    "strconv"
    // "io/ioutil"
    "github.com/valyala/fastjson"

    //"bytes"

    //"runtime"
//    "encoding/json"
    //"encoding/gob"
    
//    "time"
//    "math/rand"

//    S "data_xcls/config"

)

//----------------------------------------------

func Save_Map(byteValues  []byte, 
    bucket_Name string, 
    Ow_Um_Map   map[string]float64,
    ) error {


    // fmt.Println("Save_Map bucket_Name =", bucket_Name)

    dbFileName := fastjson.GetString(byteValues, "Base", "dbFileName")
    if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
        __err_panic(err)

    }
    //fmt.Println(" ----------  L.Buckets Stat IsNotExist ---------- ")
    
    db, err := bolt.Open(dbFileName, 0600, nil)
    if err != nil {
        __err_panic(err)
    }
    defer db.Close()
    


    err = db.Update(func(tx *bolt.Tx) error {

        b, err := tx.CreateBucketIfNotExists([]byte(bucket_Name))
        __err_panic(err)

        //nn = 0

        //i_ID_Digital_Signage = 0

        for key, value := range Ow_Um_Map {
  
            //fmt.Println("Key =", key, "Value =", value)
            //fmt.Println("[]byte(key) =", []byte(key), "Value =", value)


            err = b.Put([]byte(key), []byte(fmt.Sprintf("%f", value)))
            __err_panic(err)

        } // for key, value := range Ow_Um_Map {


            return err

    })  // err = db.Update(func(tx *bolt.Tx) error 



    return err

}


