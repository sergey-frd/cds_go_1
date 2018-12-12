package lib_alloc

import (   
    "fmt"
	"sort"
//	"log"
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

//----------------------------------------------
func Alloc_OwUmNbDsTiSl(byteValues  []byte, 
    oundt_Media_Cost float64,
    data             map[string]map[string]string,
    um               S.User_Media_STC,
    ou               S.Ow_Um_STC     ,
    ound             S.Ow_UmNbDs_STC ,
    oundt            S.Ow_UmNbDsTi_STC ,
    ) (float64, int, error) {


    var dsfs        S.NbDsTiSl_STC     // Free Ds Slots
    var err         error
    var tbl_name    string

    var total       float64
    var count       int

    //var dsfs  S.NbDsTiSl_STC     // Free Ds Slots
    //var Ds          S.Digital_Signage_STC ,

    //fmt.Println("Alloc_OwUmNbDsTiSl oundt =",oundt)
    //fmt.Println("Alloc_OwUmNbDs ou.OwUm_Key.ID_Owner =",ou.OwUm_Key.ID_Owner)
    //fmt.Println("Alloc_Ow um.UsMd =",um.UsMd)


    tbl_name   = "Free_Slots"
    var keys []string
    for k, _ := range data[tbl_name] {
        keys = append(keys, k)
    } // for k, v

    sort.Strings(keys)  // or sort.Ints(keys), sort.Sort(...), etc., per <K>



    total = 0
    count = 0
    for _, k := range keys  {
        v := data[tbl_name][k]


        //fmt.Println(k, "=>", v)

        //.................................................
        byt_k := []byte(k)
        //err := json.Unmarshal(byt_ds, &k_Ds)
        err = json.Unmarshal(byt_k, &dsfs.NbDsTiSl_Key)
        if err != nil {
            fmt.Println("There was an error:", err)
        }
        
        byt_v := []byte(v)
        err = json.Unmarshal(byt_v, &dsfs.Slot_Price)
        if err != nil {
            fmt.Println("There was an error:", err)
        }

        //if dsfs.NbDsTiSl_Key.UsMd != oundt.OwUmNbDsTi_Key.UsMd  {
        //    continue
        //}

        //fmt.Println("ound.OwUmNbDs_Key.CnCtNbDs.ID_Owner =", ound.OwUmNbDs_Key.CnCtNbDs.ID_Owner)
        if dsfs.NbDsTiSl_Key.NbDsTi_key.CnCtNbDs != oundt.OwUmNbDsTi_Key.NbDsTi_key.CnCtNbDs {
            continue
        }


       if oundt.OwUmNbDsTi_Key.NbDsTi_key.ID_Time_Interval != dsfs.NbDsTiSl_Key.NbDsTi_key.ID_Time_Interval {
            continue
       } // if oundt.OwUmNbDsTi_Key.NbDsTi_key.ID_Time_Interval 

        
        if total < oundt_Media_Cost {
            if total + dsfs.Slot_Price < oundt_Media_Cost {  
                      
                count++
                total += dsfs.Slot_Price
                //fmt.Println(dsfs.NbDsTiSl_Key.ID_Slot,"=>", dsfs.Slot_Price,total)
            } //if total + dsfs.Slot_Price > oundt_Media_Cost {

        } //if total < oundt_Media_Cost {

       //!!!!!!!!!!!!!!!!!!!!
       //break

    } // for _, k := range keys  {


    //delta := oundt_Media_Cost - total
    //fmt.Println("#",count,oundt_Media_Cost ,"- total", total," = ",delta)

    return total,count, err

}
