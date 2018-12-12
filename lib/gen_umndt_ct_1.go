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
func Get_keys_City(byteValues  []byte, 
        data            map[string]map[string]string,
        um              S.User_Media_STC,
        Ow_Um_Map       map[string]float64,
        Ow_UmNbDsTi_Map map[string]float64,
        Ow_UmNbDs_Map   map[string]float64,
        Payd_Slots_Map  map[string]float64,
        Free_Slots_Map  map[string]float64,
     ) []string {


    //fmt.Println("Get_keys_City um =", um)

    //............................................................
    var ct       S.City_STC
    //var k_Nb     S.Neighborhoods_KEY
    
    var err         error

    keys_City  := []string{}
    //keys_Nb    := []string{}

    //............................................................
    Clip_4_ALL_Country := fastjson.GetInt(byteValues,    "Base", "Clip_4_ALL_Country")
    Clip_Code_Country  := fastjson.GetString(byteValues, "Base", "Clip_Code_Country")
    Clip_Code_City     := fastjson.GetString(byteValues, "Base", "Clip_Code_City")

    nn := 0
    //for key, _ := range data["City"] {
    for CityKey, CityValue:= range data["City"] {
        nn  += 1
        //fmt.Println(nn, "CityKey =", CityKey, "Value =", value)
        if Clip_4_ALL_Country == 0 {

            byt := []byte(CityKey)
            err = json.Unmarshal(byt, &ct.CnCt)
            if err != nil {
                fmt.Println("There was an error:", err)
            }
            if Clip_Code_Country != ct.CnCt.ID_Country {
                continue
            }

            //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
            if Clip_Code_City != ct.CnCt.ID_City {
                continue
            }

            ct.City  = CityValue
            //fmt.Println(nn, "Get_keys_City ct =", ct)

            //keys_Nb  = Get_keys_Nb(byteValues,
            _  = Get_keys_Nb(byteValues,
                um,
                ct,
                data,
                Ow_Um_Map,
                Ow_UmNbDsTi_Map ,
                Ow_UmNbDs_Map   ,
                Payd_Slots_Map,
                Free_Slots_Map  ,
                )

            __err_panic(err)

            //fmt.Println("keys_Nb =",keys_Nb)

            keys_City = append(keys_City, CityKey)

            //!!!!!!!!!!!!!!
            break	


        }
    } // for CityKey, _ := range data["City"] {

    return keys_City

}
