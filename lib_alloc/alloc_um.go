package lib_alloc

import (   
    "fmt"
	"log"
	"sort"
//    "github.com/valyala/fastjson"
//    "os" 
//    "github.com/boltdb/bolt"
//    "errors"

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
//    L "cds_go_1/lib"
    S "cds_go_1/config"
)
//---------------------------------------------------------------
func __err_panic(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
//----------------------------------------------
func Alloc_Um(byteValues  []byte, 
    data            map[string]map[string]string,
    ) error {


    var err         error
    //var total       float64

    var um          S.User_Media_STC

    var keys []string
    for k, _ := range data["User_Media"] {
        //fmt.Println(k, "1 =>", v)
        keys = append(keys, k)
    } // for k, v

    //fmt.Println("keys =",keys)
    sort.Strings(keys)  // or sort.Ints(keys), sort.Sort(...), etc., per <K>
    //fmt.Println("sort keys =",keys)

    //total = 0
    for _, k := range keys  {
        v := data["User_Media"][k]
        

        //fmt.Println(k, "=>", v)

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
        fmt.Println("um =", um)



        ow_total, err := Alloc_Ow(byteValues, 
                         data ,
                         um ,
            ); __err_panic(err) 


        //fmt.Println("***************",ow_total)
        fmt.Println("########## um TOTAL",ow_total)
        //total += ow_total

        //!!!!!!!!!!!!!!!!!!!!
        //break

    }


    return  err

} // func alloc_ow
