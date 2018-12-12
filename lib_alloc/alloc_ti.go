package lib_alloc

import (   
    "fmt"
	"sort"
//	"log"
    "github.com/valyala/fastjson"
    "strconv"
//    "os" 
//    "github.com/boltdb/bolt"
//    "errors"

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
func Alloc_OwUmNbDsTi(byteValues  []byte, 
    ound_Media_Cost float64,
    data            map[string]map[string]string,
    um              S.User_Media_STC,
    ou              S.Ow_Um_STC     ,
    ound            S.Ow_UmNbDs_STC ,
    ) (float64,  error) {


    var err         error

    var oundt        S.Ow_UmNbDsTi_STC
    //var Ds          S.Digital_Signage_STC ,

    var total             float64
    var oundt_total       float64
    //var oundt_count       int

    //fmt.Println("Alloc_OwUmNbDsTi ound =",ound)
    //fmt.Println("Alloc_OwUmNbDs ou.OwUm_Key.ID_Owner =",ou.OwUm_Key.ID_Owner)
    //fmt.Println("Alloc_Ow um.UsMd =",um.UsMd)
    Time_Interval_Counter := fastjson.GetInt(byteValues, "Base", "Time_Interval_Counter")


    var keys []string
    for k, _ := range data["Ow_UmNbDsTi"] {
        keys = append(keys, k)
    } // for k, v

    sort.Strings(keys)  // or sort.Ints(keys), sort.Sort(...), etc., per <K>

    total = 0
    for _, k := range keys  {
        v := data["Ow_UmNbDsTi"][k]
        //fmt.Println(k, "=>", v)

        //.................................................
        byt_k := []byte(k)
        //err := json.Unmarshal(byt_ds, &k_Ds)
        err = json.Unmarshal(byt_k, &oundt.OwUmNbDsTi_Key)
        if err != nil {
            fmt.Println("There was an error:", err)
        }
        
        byt_v := []byte(v)
        err = json.Unmarshal(byt_v, &oundt.Total_Cost)
        if err != nil {
            fmt.Println("There was an error:", err)
        }

        if oundt.OwUmNbDsTi_Key.UsMd != ound.OwUmNbDs_Key.UsMd {
            continue
        }

        //fmt.Println("ound.OwUmNbDs_Key.CnCtNbDs.ID_Owner =", ound.OwUmNbDs_Key.CnCtNbDs.ID_Owner)
        if oundt.OwUmNbDsTi_Key.NbDsTi_key.CnCtNbDs != ound.OwUmNbDs_Key.CnCtNbDs {
            continue
        }

        //fmt.Println("Alloc_OwUmNbDsTi oundt =", oundt)
        //fmt.Println("Alloc_OwUmNbDsTi Time_Interval_Counter =",Time_Interval_Counter)

        for i := 1; i <= Time_Interval_Counter+1; i++ {


            //fmt.Println("i =",i)

            if oundt.OwUmNbDsTi_Key.NbDsTi_key.ID_Time_Interval == strconv.Itoa(i) {

                //fmt.Println("ound.OwUmNbDs_Key.CnCtNbDs.ID_Owner =", ound.OwUmNbDs_Key.CnCtNbDs.ID_Owner)
                //fmt.Println("oundt.OwUmNbDsTi_Key.NbDsTi_key =", oundt.OwUmNbDsTi_Key.NbDsTi_key)
                //fmt.Println("oundt.OwUmNbDsTi_Key.NbDsTi_key.ID_Time_Interval =", oundt.OwUmNbDsTi_Key.NbDsTi_key.ID_Time_Interval)
                //fmt.Println("Alloc_OwUmNbDsTi oundt =",i, oundt)
                //fmt.Println("ou.OwUm_Key.UsMd =", ou.OwUm_Key.UsMd)

                //fmt.Println(k, "=>", v)
                oundt_Media_Cost := (ound.Total_Cost/ound_Media_Cost)*oundt.Total_Cost
                //fmt.Println("oundt_Media_Cost =", oundt_Media_Cost)

                oundt_total, _, err = Alloc_OwUmNbDsTiSl(byteValues, 
                                 oundt_Media_Cost ,
                                 data ,
                                 um ,
                                 ou ,
                                 ound ,
                                 oundt ,
                    ); __err_panic(err) 

                    //fmt.Println("###",oundt_total,oundt_count)
                    total += oundt_total

               //!!!!!!!!!!!!!!!!!!!!
               //break

           } // if oundt.OwUmNbDsTi_Key.NbDsTi_key.ID_Time_Interval
           //!!!!!!!!!!!!!!!!!!!!
           //break
        } // for i := 1; i <= Time_Interval_Counter+1; i++ {

    } // for _, k := range keys  {

    return total, err

}
