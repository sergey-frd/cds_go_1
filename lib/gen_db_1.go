package lib

import (   
    "fmt"
    "os" 
    "github.com/boltdb/bolt"
    "errors"
//    "strconv"
    // "io/ioutil"
    "github.com/valyala/fastjson"

    //"bytes"

    //"runtime"
    "encoding/json"
    //"encoding/gob"
    
    "time"
    "math/rand"

    S "cds_go_1/config"
    //L "cds_go_1/lib"

)
//---------------------------------------------------------------
func random(min, max int) int {

    //s1 := rand.NewSource(time.Now().UnixNano())
    //r1 := rand.New(s1)

    // rand.Seed(time.Now().Unix())
    // return rand.Intn(max - min) + min

    //r1.Seed(time.Now().Unix())
    //return r1.Intn(max - min) + min

    rand.Seed(int64(time.Now().Nanosecond()))
    return rand.Intn(max - min) + min

}

//----------------------------------------------
func GenAllFiles(byteValues []byte) {


    project_name := fastjson.GetString(byteValues, "Base", "project_name")
    fmt.Printf("*** GenAllFiles project_name = %s\n", project_name)

    // var  myrand int
    // myrand = random (1, 6)
    // fmt.Println (myrand)

    dbFileName, err := GetDbName(byteValues);  __err_panic(err) 
    if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
        fmt.Println(err)
        return
    }
    //fmt.Println(" ----------  L.Buckets Stat IsNotExist ---------- ")
    
    db, err := bolt.Open(dbFileName, 0600, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()
    

    //var buf bytes.Buffer        
    //dec := gob.NewDecoder(&buf) // Will read .

    //fmt.Println(" ----------  Open ---------- ")
    err = db.View(func(tx *bolt.Tx) error {
        return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
    
        sheet_Name := string(name)
        //fmt.Println(string(name))
        //fmt.Printf(" *********  %s  ************* \n", sheet_Name)
             
        //data[sheet_Name] = make(map[string]string)

        switch sheet_Name {
        //................................................
        //case "Neighborhoods":
        case "Time_Interval":
            fmt.Printf(" ********* GenAllFiles %s  ************* \n", sheet_Name)

            err = db.View(func(tx *bolt.Tx) error {
                b := tx.Bucket([]byte(sheet_Name))
                c := b.Cursor()
                if b == nil {
                    return errors.New("Bucket not found")
                }
                //log.Printf("Value of %s is %s", "answer", b.Get([]byte("answer")))
            
                for k, v := c.First(); k != nil; k, v = c.Next() {
                    //log.Printf("Value of %s is %s", k, v)
                    //log.Printf(" %s => %s", k, string(v))
                    //fmt.Println( k, "=>",v)

                    //var ti S.Time_Interval_STC
                    var ti S.Time_Interval_VAL

                    err = json.Unmarshal(v, &ti)
                    if err != nil {
                        fmt.Println("There was an error:", err)
                    }
                    fmt.Println(  string(k), "=>",ti,string(v))
                   //fmt.Println("1 ti =",ti)


                    //data[sheet_Name][string(k[:])] = string(v[:])
                }
                return nil
            })
        } // switch sheet_Name

        return nil
        }) //db.View(func(tx *bolt.Tx)

    }) // db.View(func(tx *bolt.Tx) error {
    
    if err != nil {
        fmt.Println(err)
        return
    }

}

//------------------------------------------------------------
func DemoGenDB(byteValues []byte) {
    fmt.Println("DemoGenDB HI")
    //fmt.Println("byteValues =",byteValues)

    project_name := fastjson.GetString(byteValues, "Base", "project_name")
    fmt.Printf("DemoGenDB project_name = %s\n", project_name)

}
