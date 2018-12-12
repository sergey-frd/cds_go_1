package lib

import (   
    "fmt"
    "github.com/valyala/fastjson"
    "os" 
    "github.com/boltdb/bolt"
    "errors"

//    "strconv"
//  // "io/ioutil"
//  //"bytes"
//  //"runtime"
    "encoding/json"
//  //"encoding/gob"
//  
//  "time"
//  "math/rand"
//
    S "cds_go_1/config"

)

//----------------------------------------------
func Gen_UmNbDtTi_Bucket(byteValues  []byte, 
    data        map[string]map[string]string,
    Ow_Um_Map   map[string]float64,
    Ow_UmNbDsTi_Map map[string]float64,
    Ow_UmNbDs_Map   map[string]float64,
    Payd_Slots_Map  map[string]float64,
    Free_Slots_Map   map[string]float64,
    ) error {


    //fmt.Println("Gen_UmNbDtTi_Bucket Ow_Um_Map =", Ow_Um_Map)

    var bucket_Name string     = "User_Media" 
    var err         error
    var um          S.User_Media_STC




    //............................................................
    dbFileName := fastjson.GetString(byteValues, "Base", "dbFileName")
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
			
                nnn := 0
				for k, v := c.First(); k != nil; k, v = c.Next() {
                    nnn  += 1

                    //fmt.Println(nnn, string(k[:]), "=>",string(v[:]))

                    //.................................................
                    byt_k := []byte(k)
                    //err := json.Unmarshal(byt_ds, &k_Ds)
                    err = json.Unmarshal(byt_k, &um.UsMd)
                    if err != nil {
                        fmt.Println("There was an error:", err)
                    }

                    byt_v := []byte(v)
                    err = json.Unmarshal(byt_v, &um.UsMdVl)
                    if err != nil {
                        fmt.Println("There was an error:", err)
                    }
                    //fmt.Println("um =", um)

                    //.................................................
                    //keys_City := Get_keys_City(byteValues, 
                    Get_keys_City(byteValues, 
                        data,
                        um,
                        Ow_Um_Map,
                        Ow_UmNbDsTi_Map   ,
                        Ow_UmNbDs_Map   ,
                        Payd_Slots_Map,
                        Free_Slots_Map  ,
                        )

                    //fmt.Println("keys_City =", keys_City)

                    //!!!!!!!!!!
                    //break


				}
				return nil
			})


            return nil
    })



    return err

}


