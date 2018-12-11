package lib_alloc

import (   
    "fmt"
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
    S "data_xcls/config"

)

//----------------------------------------------
func Alloc_Ow(byteValues  []byte, 
    data            map[string]map[string]string,
    um              S.User_Media_STC,
    ) (float64,  error) {


    var err         error
    var total       float64

    var ou       S.Ow_Um_STC

    //fmt.Println("Alloc_Ow um =",um)
    //fmt.Println("Alloc_Ow um.UsMd =",um.UsMd)


    var keys []string
    for k, _ := range data["Ow_Um"] {
        keys = append(keys, k)
    } // for k, v

    sort.Strings(keys)  // or sort.Ints(keys), sort.Sort(...), etc., per <K>

    total = 0
    for _, k := range keys  {
        v := data["Ow_Um"][k]

        //fmt.Println(k, "=>", v)

        //.................................................
        byt_k := []byte(k)
        //err := json.Unmarshal(byt_ds, &k_Ds)
        err = json.Unmarshal(byt_k, &ou.OwUm_Key)
        if err != nil {
            fmt.Println("There was an error:", err)
        }
        
        byt_v := []byte(v)
        err = json.Unmarshal(byt_v, &ou.Total_Cost)
        if err != nil {
            fmt.Println("There was an error:", err)
        }

        if ou.OwUm_Key.UsMd != um.UsMd {
            continue
        }

        //fmt.Println("ou =", ou)
        //fmt.Println("ou.OwUm_Key.UsMd =", ou.OwUm_Key.UsMd)

        owds_total, err := Alloc_OwUmNbDs(byteValues, 
                         data ,
                         um ,
                         ou ,
            ); __err_panic(err) 

        //fmt.Println("===========",owds_total)
        total += owds_total
       //!!!!!!!!!!!!!!!!!!!!
       break
    }

    return total, err

} // func alloc_ow
