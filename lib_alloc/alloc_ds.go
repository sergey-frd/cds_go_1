package lib_alloc

import (   
    "fmt"
	"sort"
//	"log"
//    "github.com/valyala/fastjson"
//    "os" 
//    "github.com/boltdb/bolt"
//    "errors"

    "strconv"
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
func Alloc_OwUmNbDs(byteValues  []byte, 
    data            map[string]map[string]string,
    um              S.User_Media_STC,
    ou              S.Ow_Um_STC     ,
    ) (float64,  error) {


    var err         error
    var tbl_name    string

    var ound        S.Ow_UmNbDs_STC
    //var Ds          S.Digital_Signage_STC ,

    var total       float64

    //fmt.Println("Alloc_OwUmNbDs ou =",ou)
    //fmt.Println("Alloc_OwUmNbDs ou.OwUm_Key.ID_Owner =",ou.OwUm_Key.ID_Owner)
    //fmt.Println("Alloc_Ow um.UsMd =",um.UsMd)


    tbl_name   = "Ow_UmNbDs"
    var keys []string
    for k, _ := range data[tbl_name] {
        keys = append(keys, k)
    } // for k, v

    sort.Strings(keys)  // or sort.Ints(keys), sort.Sort(...), etc., per <K>

    total = 0
    for _, k := range keys  {
        v := data[tbl_name][k]

        //fmt.Println(k, "=>", v)

        //.................................................
        byt_k := []byte(k)
        //err := json.Unmarshal(byt_ds, &k_Ds)
        err = json.Unmarshal(byt_k, &ound.OwUmNbDs_Key)
        if err != nil {
            fmt.Println("There was an error:", err)
        }
        
        byt_v := []byte(v)
        err = json.Unmarshal(byt_v, &ound.Total_Cost)
        if err != nil {
            fmt.Println("There was an error:", err)
        }

        if ound.OwUmNbDs_Key.UsMd != ou.OwUm_Key.UsMd {
            continue
        }

        //fmt.Println("ound.OwUmNbDs_Key.CnCtNbDs.ID_Owner =", ound.OwUmNbDs_Key.CnCtNbDs.ID_Owner)
        if ound.OwUmNbDs_Key.CnCtNbDs.ID_Owner != ou.OwUm_Key.ID_Owner {
            continue
        }

        //fmt.Println("ound.OwUmNbDs_Key.CnCtNbDs.ID_Owner =", ound.OwUmNbDs_Key.CnCtNbDs.ID_Owner)
        //fmt.Println("ound =", ound.OwUmNbDs_Key.CnCtNbDs.ID_Owner,ound)
        //fmt.Println("ou.OwUm_Key.UsMd =", ou.OwUm_Key.UsMd)

        //fmt.Println(k, "=>", v)
        Media_Cost,err := strconv.ParseFloat(um.UsMdVl.Media_Cost, 64)

        //ou_Media_Cost := (Media_Cost/ou.Total_Cost)*ound.Total_Cost
        //ou_a_Media_Cost := (Media_Cost/ou.Total_Cost)
        //ou_Media_Cost := ou_a_Media_Cost*ound.Total_Cost

        //fmt.Println("Media_Cost =", Media_Cost)
        //fmt.Println("ou_a_Media_Cost =", ou_a_Media_Cost)
        //fmt.Println("ound.Total_Cost =", ound.Total_Cost)

        ound_Media_Cost := (ou.Total_Cost/Media_Cost)*ound.Total_Cost
        //fmt.Println("ound_Media_Cost =", ound_Media_Cost)

        ound_total, err := Alloc_OwUmNbDsTi(byteValues, 
                         ound_Media_Cost,
                         data ,
                         um ,
                         ou ,
                         ound ,
            ); __err_panic(err) 

        //fmt.Println("###",ound_total)
        total += ound_total

       //!!!!!!!!!!!!!!!!!!!!
       //break

    }

    return total, err

}
