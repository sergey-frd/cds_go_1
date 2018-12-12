package lib

import (   
    "fmt"
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
    S "cds_go_1/config"

)

//----------------------------------------------
 func Get_keys_Ds(byteValues []byte, 
    um    S.User_Media_STC   ,
    ct    S.City_STC         ,
    Nb    S.Neighborhoods_STC,
    data  map[string]map[string]string,
    Ow_Um_Map   map[string]float64,
    Ow_UmNbDsTi_Map map[string]float64,
    Ow_UmNbDs_Map   map[string]float64,
    Payd_Slots_Map  map[string]float64,
    Free_Slots_Map   map[string]float64,
    ) []string {

    keys_Ds  := []string{}
    //keys_Ti  := []string{}

    var err error
    var Ds       S.Digital_Signage_STC
    var ou       S.Ow_Um_STC
    var ound     S.Ow_UmNbDs_STC



    //............................................................
    //fmt.Println("Get_keys_Ds Nb =",Nb)

    // byt_nb := []byte(nb_key)
    // err = json.Unmarshal(byt_nb, &k_Nb)
    // if err != nil {
    //     fmt.Println("There was an error:", err)
    // }
    // 
    // fmt.Println("k_Nb =", k_Nb)

    nnn := 0
    //for ds_key, ds_value := range data["Neighborhoods"] {
    //for ds_key, _ := range data["Digital_Signage"] {
    for ds_key, ds_value:= range data["Digital_Signage"] {
    
        //fmt.Println(nnn, "nb_key =", nb_key,ds_value)

        byt_ds := []byte(ds_key)
        //err := json.Unmarshal(byt_ds, &k_Ds)
        err = json.Unmarshal(byt_ds, &Ds.CnCtNbDs)
        if err != nil {
            fmt.Println("There was an error:", err)
        }
        //fmt.Println(nnn, "k_Nb =", k_Nb,ds_value)
    
        if Ds.CnCtNbDs.CnCtNb != Nb.CnCtNb {
            continue
        }

        byt_dsv := []byte(ds_value)
        //err := json.Unmarshal(byt_dsv, &k_Ds)
        err = json.Unmarshal(byt_dsv, &Ds.DsVal)
        if err != nil {
            fmt.Println("There was an error:", err)
        }

        nnn  += 1
        //fmt.Println(nnn, "Neighborhoods Key =", nb_key, "Value =", value)
        //fmt.Println(nnn, "k_Nb =", k_Nb,value)
        //fmt.Println(nnn, "Ds =", Ds)

        // type Ow_Um_KEY   struct {
        // UsMd                User_Media_KEY
        // ID_Owner            string
        // }
        // 
        // type Ow_Um_STC   struct {
        // OwUm_Key            Ow_Um_KEY
        // Total_Cost          float64
        // }
        // 
        
        ou.OwUm_Key.UsMd       = um.UsMd
        ou.OwUm_Key.ID_Owner   = Ds.CnCtNbDs.ID_Owner
        ou.Total_Cost          = 0

        enc_OwUm_Key, err := json.Marshal(ou.OwUm_Key)
        __err_panic(err)
        //fmt.Println(nnn, "ou =", ou)
        Ow_Um_Map[string(enc_OwUm_Key)] = ou.Total_Cost
        //fmt.Println("Get_keys_Ds Ow_Um_Map =",Ow_Um_Map)

        ound.OwUmNbDs_Key.UsMd       = um.UsMd
        ound.OwUmNbDs_Key.CnCtNbDs   = Ds.CnCtNbDs
        ound.Total_Cost              = 0

        enc_OwUmNbDs_Key, err := json.Marshal(ound.OwUmNbDs_Key)
        __err_panic(err)
        Ow_UmNbDs_Map[string(enc_OwUmNbDs_Key)] = ound.Total_Cost


        //keys_Ti  = Get_keys_Ti(byteValues,
        _  = Get_keys_Ti(byteValues,
            um       ,
            ct       ,
            Nb       ,
            Ds       ,
            data     , 
            Ow_Um_Map,
            Ow_UmNbDsTi_Map ,
            Ow_UmNbDs_Map   ,
            Payd_Slots_Map,
            Free_Slots_Map  ,
            )

        __err_panic(err)
        //fmt.Println("keys_Ti =",keys_Ti)

        //!!!!!!!!!!!!!!
        //break

    } // for nb_key, _ := range data["Neighborhoods"] 

    return keys_Ds

} 
 
