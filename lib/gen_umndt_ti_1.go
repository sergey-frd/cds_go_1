package lib

import (   
    "fmt"
    "github.com/valyala/fastjson"
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
    S "data_xcls/config"

)

//----------------------------------------------
 func Get_keys_Ti(byteValues  []byte, 
        um              S.User_Media_STC      ,
        ct              S.City_STC            ,
        Nb              S.Neighborhoods_STC   ,
        Ds              S.Digital_Signage_STC ,
        data            map[string]map[string]string,
        Ow_Um_Map       map[string]float64,
        Ow_UmNbDsTi_Map map[string]float64,
        Ow_UmNbDs_Map   map[string]float64,
        Payd_Slots_Map  map[string]float64,
        Free_Slots_Map  map[string]float64,
      ) []string {

    //fmt.Println("Get_keys_Ti Ow_Um_Map =", Ow_Um_Map)

    // fmt.Println("Get_keys_Ti um =", um)
    // fmt.Println("Get_keys_Ti ct =", ct)
    // fmt.Println("Get_keys_Ti Nb =", Nb)
    // fmt.Println("Get_keys_Ti Ds =", Ds)


    keys_Ti     := []string{}
    //keys_Ti_Sl  := []string{}

    var iCounter    int
    var err         error
    var Ti          S.Time_Interval_STC


    //var oundt    S.Ow_UmNbDsTi_STC

    Time_Interval_Counter := fastjson.GetInt(byteValues, "Base", "Time_Interval_Counter")
    //fmt.Println("Time_Interval_Counter =",Time_Interval_Counter)

    Clip_4_ALL_Ti := random(0, 2)
    //fmt.Println("Clip_4_ALL_Ti =",Clip_4_ALL_Ti)


    if Clip_4_ALL_Ti == 1 {
        iCounter = Time_Interval_Counter 
    } else {
        iCounter = random(1, Time_Interval_Counter)
    }

    //fmt.Println("iCounter =",iCounter)

    // type Time_Interval_KEY   struct {
    //     CnCtNbDs          Digital_Signage_KEY
    //     ID_Time_Interval  string
    // }

    for i := 1; i <= iCounter; i++ {
        //fmt.Println("i =",i)
        
        //Ti.TiVl.CnCtNbDs         = Ds.CnCtNbDs
        Ti.ID_Time_Interval = strconv.Itoa(i)


        //fmt.Println("Get_keys_Ti_Slot ID_Time_Interval =",ID_Time_Interval)
        //fmt.Println("data[Time_Interval] =",data["Time_Interval"])

        //Ti.TiVl := data["Time_Interval"][ID_Time_Interval]
        Ti_Val := data["Time_Interval"][Ti.ID_Time_Interval]

        //fmt.Println("Ti_Val =",Ti_Val)
        ////fmt.Println(  string(Ti_Val))


        err = json.Unmarshal([]byte(Ti_Val) , &Ti.TiVl)
        if err != nil {
            fmt.Println("There was an error:", err)
        }
        //fmt.Println("Ti =",Ti)

        keys_Ti  = Get_keys_Ti_Slot(byteValues,
            um       ,
            ct       ,  
            Nb       ,  
            Ds       ,  
            Ti       ,  
            data     ,  
            Ow_Um_Map,  
            Ow_UmNbDsTi_Map ,
            Ow_UmNbDs_Map   ,
            Payd_Slots_Map,
            Free_Slots_Map  ,
            )

        __err_panic(err)
        //fmt.Println("keys_Ti_Sl =",keys_Ti_Sl)

        //!!!!!!!!!!!!!!

    } // for i := 1; 

    return keys_Ti

} // func Get_keys_Ti

