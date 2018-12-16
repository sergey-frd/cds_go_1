package lib_init

import (   
    "fmt"
	"log"
	"sort"
    "encoding/json"
//    "github.com/valyala/fastjson"
//    "os" 
//    "github.com/boltdb/bolt"
//    "errors"

//    "strconv"
//  // "io/ioutil"
//  //"bytes"
//  //"runtime"
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
func Init_Um(byteValues  []byte, 
    data            map[string]map[string]string,
    TotalDict       map[string]map[string]float64,
    )  (err error) {

    var       tbl_name   string
    var       ps         S.Um_NbDsTiSl_STC
    var       od         S.Ow_Day_STC
    var       ods        S.Ow_Day_Ds_STC

    //var total       float64

    //var um          S.User_Media_STC

    // var dsfs        S.NbDsTiSl_STC     // Free Ds Slots
    // var err         error
    // 
    // var total       float64
    // var count       int

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



    //total = 0
    //count = 0
    nnn := 0
    for _, k := range keys  {
        nnn  += 1
        v := data[tbl_name][k]


        //fmt.Println(k, "=>", v)

        //.................................................
        byt_k := []byte(k)
        //err := json.Unmarshal(byt_ds, &k_Ds)
        err = json.Unmarshal(byt_k, &ps.UmNbDsTiSl_Key)
        if err != nil {
            fmt.Println("There was an error:", err)
        }
        
        byt_v := []byte(v)
        err = json.Unmarshal(byt_v, &ps.Slot_Price)
        if err != nil {
            fmt.Println("There was an error:", err)
        }

        // fmt.Println("ps = ", ps)
        // fmt.Println("ps.Slot_Price = ",ps.Slot_Price)
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key)
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.Year  = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.Year )
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.Month = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.Month)
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.Day   = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.Day  )
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.Hour  = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.Hour )
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.Index = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.Index)
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key)
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.ID_Time_Interval = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.ID_Time_Interval)
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs)
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.ID_Digital_Signage = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.ID_Digital_Signage)
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.ID_Owner = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.ID_Owner)
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb)
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.CnCt = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.CnCt            )
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.CnCt.ID_Country  = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.CnCt.ID_Country            )
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.CnCt.ID_City     = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.CnCt.ID_City               )
        // fmt.Println("ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.ID_Neighborhoods = ",ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.CnCtNb.ID_Neighborhoods)

        od.Ow_day_key.Ymd_key.Year  = ps.UmNbDsTiSl_Key.NbDsTiSl_key.Year
        od.Ow_day_key.Ymd_key.Month = ps.UmNbDsTiSl_Key.NbDsTiSl_key.Month    
        od.Ow_day_key.Ymd_key.Day   = ps.UmNbDsTiSl_Key.NbDsTiSl_key.Day   
        od.Ow_day_key.ID_Owner      = ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs.ID_Owner

        ods.Ow_day_ds_key.Ymd_key   = od.Ow_day_key.Ymd_key
        ods.Ow_day_ds_key.CnCtNbDs  = ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.CnCtNbDs
    

        ow_day_key_enc,    err := json.Marshal(od.Ow_day_key);     __err_panic(err)
        ow_day_ds_key_enc, err := json.Marshal(ods.Ow_day_ds_key); __err_panic(err)
        
        //fmt.Println("od = ", od)
        //fmt.Println("od.Ow_day_key = ", od.Ow_day_key )
        //fmt.Println("ow_day_key_enc = ", ow_day_key_enc)
        //fmt.Println("string(ow_day_key_enc) = ", string(ow_day_key_enc))

        _, ok := TotalDict["Ow_Day"][string(ow_day_key_enc)]
        if !ok {
            TotalDict["Ow_Day"][string(ow_day_key_enc)]  = ps.Slot_Price
        } else {
            TotalDict["Ow_Day"][string(ow_day_key_enc)] += ps.Slot_Price
        }

        _, ok = TotalDict["Ow_Day_Ds"][string(ow_day_ds_key_enc)]
        if !ok {
            TotalDict["Ow_Day_Ds"][string(ow_day_ds_key_enc)]  = ps.Slot_Price
        } else {
            TotalDict["Ow_Day_Ds"][string(ow_day_ds_key_enc)] += ps.Slot_Price
        }


        //price, _ := strconv.ParseFloat(data["Price"][ps.Slot_Pric], 64)
        // fmt.Println(nnn,"Ow_Day   = ", TotalDict["Ow_Day"])
        // fmt.Println(nnn,"Ow_Day_Ds= ", TotalDict["Ow_Day_Ds"])
                
		//if data["Ow_Day"][string(ow_day_key_enc)] == nil {
        //    data["Ow_Day"][string(ow_day_key_enc)] = make(map[string]string)
		//} 
		//if m[string(key)] == nil {
		//	m[string(key)] = make(map[string]interface{})
		//}
        
        //!!!!!!!!!!!!!!
        // if nnn >= 2 {
        //     //!!!!!!!!!!!!!!!!!!!!
        // 
        //     break
        // }


    } // for _, k := range keys  {

    fmt.Println(nnn,"Ow_Day   = ", TotalDict["Ow_Day"])
    fmt.Println(nnn,"Ow_Day_Ds= ", TotalDict["Ow_Day_Ds"])


    return  err

} // func alloc_ow
