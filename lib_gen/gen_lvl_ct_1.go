package lib_gen

import (   
//    "fmt"
//    "log"
	"sort"
    "github.com/valyala/fastjson"
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


//----------------------------------------------
//func Alloc_Um(byteValues  []byte, 
func Gen_Lvl_Ct(byteValues  []byte, 
    data       map[string]map[string]string,
    ps         S.Um_NbDsTiSl_STC,
    um         S.User_Media_STC,
    ) error {


    var err         error
    //var total       float64

    var ct       S.City_STC
    //var um   S.User_Media_STC


    //............................................................
    Clip_4_ALL_Country := fastjson.GetInt(byteValues,    "Base", "Clip_4_ALL_Country")
    Clip_Code_Country  := fastjson.GetString(byteValues, "Base", "Clip_Code_Country")
    Clip_Code_City     := fastjson.GetString(byteValues, "Base", "Clip_Code_City")

    var keys []string
    for k, _ := range data["City"] {
        keys = append(keys, k)
    } // for k, v
    sort.Strings(keys)  // or sort.Ints(keys), sort.Sort(...), etc., per <K>
    n := 0
    for _, k := range keys  {
        n++

        if Clip_4_ALL_Country == 0 {

            byt_k := []byte(k)
            err = json.Unmarshal(byt_k, &ct.CnCt);  __err_panic(err)
            //if err != nil {
            //    fmt.Println("There was an error:", err)
            //}

            if Clip_Code_Country != ct.CnCt.ID_Country {
                continue
            }

            //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
            if Clip_Code_City != ct.CnCt.ID_City {
                continue
            }

            ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.CnCt = ct.CnCt


            //fmt.Println("um =", um)
            //fmt.Println("  ct ps =", ps)
            err = Gen_Lvl_Nb(byteValues,data,ps,um,);  __err_panic(err)

        } //if Clip_4_ALL_Country == 0 {
        //.................................................



        //!!!!!!!!!!!!!!!!!!!!
        //if n >=5 { break }


    }

    return  err

} // func alloc_ow

