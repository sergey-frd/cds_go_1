package lib_gen

import (   
    "fmt"
        "time"
//    "log"
//	"sort"
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

    //L "cds_go_1/lib"
    S "cds_go_1/config"

)

//----------------------------------------------

func Gen_Lvl_Ti(byteValues  []byte, 
    data       map[string]map[string]string,
    ps         S.Um_NbDsTiSl_STC,
    um         S.User_Media_STC,
    Ds         S.Digital_Signage_STC,
    c_time     time.Time,
    )  (err error) {

    //var total       float64

    //p := fmt.Println

    var Ti          S.Time_Interval_STC
    //var iCounter int


    Ti.ID_Time_Interval = strconv.Itoa(c_time.Hour())
    Ti_Val := data["Time_Interval"][Ti.ID_Time_Interval]

    //err = json.Unmarshal([]byte(Ti_Val) , &Ti.TiVl);  __err_panic(err)
    //!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
    err = json.Unmarshal([]byte(Ti_Val) , &Ti.TiVl);
    if err != nil {
        fmt.Println("There was an error:", err)
        fmt.Println("        Ti Ti_Val =", Ti_Val)
        //__err_panic(err)
        panic(err)
    }

    ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.ID_Time_Interval = Ti.ID_Time_Interval

    //fmt.Println("        Ti Ti =", Ti)
    //fmt.Println("        Ti ps =", ps,Ti.TiVl)
    err = Gen_Lvl_Sl(byteValues,data,ps,um,Ds,c_time,Ti);  __err_panic(err)

    return  err

} // func alloc_ow

