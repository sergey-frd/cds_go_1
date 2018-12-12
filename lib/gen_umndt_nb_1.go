package lib

import (   
    "fmt"
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
    S "cds_go_1/config"

)

//----------------------------------------------
 func Get_keys_Nb(byteValues  []byte, 
     um    S.User_Media_STC,
     ct    S.City_STC,
     data  map[string]map[string]string,
     Ow_Um_Map   map[string]float64,
    Ow_UmNbDsTi_Map map[string]float64,
    Ow_UmNbDs_Map   map[string]float64,
    Payd_Slots_Map  map[string]float64,
        Free_Slots_Map   map[string]float64,
      ) []string {
 

    //fmt.Println("Get_keys_Nb ct =", ct)

    //............................................................
    var Nb       S.Neighborhoods_STC
    //var k_Nb     S.Neighborhoods_KEY
    var iCounter int
    
    //keys_City  := []string{}
    keys_Nb    := []string{}
    //keys_Ds    := []string{}


    Clip_4_ALL_Nb := Random(0, 2)
    //fmt.Println("Clip_4_ALL_Nb =",Clip_4_ALL_Nb)

    Neighborhoods_Counter := fastjson.GetInt(byteValues, "Base", "Neighborhoods_Counter")
    //fmt.Println("Neighborhoods_Counter =",Neighborhoods_Counter)

    Clip_Nb_Count := Random(1, Neighborhoods_Counter)
    //fmt.Println("Clip_Nb_Count =",Clip_Nb_Count)

    //............................................................
    nnn := 0
    for nb_key, nb_value := range data["Neighborhoods"] {
    //for nb_key, _ := range data["Neighborhoods"] {
    
        //fmt.Println(nnn, "nb_key =", nb_key,value)

        byt_nb := []byte(nb_key)
        err := json.Unmarshal(byt_nb, &Nb.CnCtNb)
        if err != nil {
            fmt.Println("There was an error:", err)
        }
        //fmt.Println(nnn, "k_Nb =", k_Nb,value)
    
        if Nb.CnCtNb.CnCt != ct.CnCt {
            continue
        }
        Nb.Neighborhoods = nb_value
    
        nnn  += 1
        //fmt.Println(nnn, "Neighborhoods Key =", nb_key, "Value =", value)
        //fmt.Println(nnn, "Nb =", Nb)

        //keys_Ds  = Get_keys_Ds(byteValues,
        _  = Get_keys_Ds(byteValues,
            um    ,
            ct    ,
            Nb    ,
            data, 
            Ow_Um_Map,
            Ow_UmNbDsTi_Map ,
            Ow_UmNbDs_Map   ,
            Payd_Slots_Map,
            Free_Slots_Map  ,
            )
        __err_panic(err)
        //fmt.Println("keys_Ds =",keys_Ds)


        keys_Nb = append(keys_Nb, nb_key)
    
        if Clip_4_ALL_Nb == 1 {
            iCounter = Neighborhoods_Counter
        } else {
            iCounter = Clip_Nb_Count
        }

        if nnn >= iCounter {
            break	
        } 

        //!!!!!!!!!!!!!!
        //break	


    } // for nb_key, value := range data["Neighborhoods"] {

 
    return keys_Nb
 }

