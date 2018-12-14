package lib_gen

import (   
    "fmt"
////    "log"
////	"sort"
//    "github.com/valyala/fastjson"
////    "os" 
////    "github.com/boltdb/bolt"
//        "errors"
//
//    "strconv"
////  // "io/ioutil"
////  //"bytes"
////  //"runtime"
//    "encoding/json"
////  //"encoding/gob"
////  
      "time"
////  "math/rand"
//
//    L "cds_go_1/lib"
    S "cds_go_1/config"

)


//----------------------------------------------
//func Alloc_Um(byteValues  []byte, 
func Gen_Lvl_Sl(byteValues  []byte, 
    data       map[string]map[string]string,
    ps         S.Um_NbDsTiSl_STC,
    um         S.User_Media_STC,
    c_time     time.Time,
    )  (err error) {


    p := fmt.Println

    p(c_time.Year())
    p(c_time.Month())
    p(c_time.Day())
    p(c_time.Hour())
    p(c_time.Minute())
    p(c_time.Second())
    p(c_time.Nanosecond())
    p(c_time.Location())

//     var err         error
//     //var total       float64
// 
//     var Ti          S.Time_Interval_STC
//     //var um   S.User_Media_STC
//     var iCounter int
// 
// 
//     Code_Free_Users  := fastjson.GetInt(byteValues, "Base", "Code_Free_Users"")
//     Code_Free_Clips  := fastjson.GetInt(byteValues, "Base", "Code_Free_Clips"")
//     Code_Other_Users := fastjson.GetInt(byteValues, "Base", "Code_Other_Users")
//     Code_Other_Clips := fastjson.GetInt(byteValues, "Base", "Code_Other_Clips")
// 
// 
//     Time_Interval_Counter := fastjson.GetInt(byteValues, "Base", "Time_Interval_Counter")
//     //fmt.Println("Time_Interval_Counter =",Time_Interval_Counter)
// 
//     Clip_4_ALL_Ti := L.Random(0, 2)
//     //fmt.Println("Clip_4_ALL_Ti =",Clip_4_ALL_Ti)
// 
// 
//     if Clip_4_ALL_Ti == 1 {
//         iCounter = Time_Interval_Counter 
//     } else {
//         iCounter = L.Random(1, Time_Interval_Counter)
//     }
// 
//     for i := 1; i <= iCounter; i++ {
//         //fmt.Println("i =",i)
//         
//         Ti.ID_Time_Interval = strconv.Itoa(i)
// 
// 
//         //fmt.Println("Get_keys_Ti_Slot ID_Time_Interval =",ID_Time_Interval)
//         //fmt.Println("data[Time_Interval] =",data["Time_Interval"])
// 
//         //Ti.TiVl := data["Time_Interval"][ID_Time_Interval]
//         Ti_Val := data["Time_Interval"][Ti.ID_Time_Interval]
// 
//         //fmt.Println("Ti_Val =",Ti_Val)
//         ////fmt.Println(  string(Ti_Val))
// 
// 
//         err = json.Unmarshal([]byte(Ti_Val) , &Ti.TiVl)
//         if err != nil {
//             fmt.Println("There was an error:", err)
//         }
//         ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key.ID_Time_Interval = Ti.ID_Time_Interval
// 
//         //fmt.Println("        Ti Ti =", Ti)
//         fmt.Println("        Ti ps =", ps,Ti.TiVl)
//         err = Gen_Lvl_Sl(byteValues,data,ps,Ti);  __err_panic(err)
// 
//         // keys_Ti  = Get_keys_Ti_Slot(byteValues,
//         //     um       ,
//         //     ct       ,  
//         //     Nb       ,  
//         //     Ds       ,  
//         //     Ti       ,  
//         //     data     ,  
//         //     Ow_Um_Map,  
//         //     Ow_UmNbDsTi_Map ,
//         //     Ow_UmNbDs_Map   ,
//         //     Payd_Slots_Map,
//         //     Free_Slots_Map  ,
//         //     )
// 
//         // __err_panic(err)
//         //fmt.Println("keys_Ti_Sl =",keys_Ti_Sl)
// 
//         //!!!!!!!!!!!!!!
//         //break
// 
//     } // for i := 1; 
    return  err

} // func alloc_ow

