package lib_gen

import (   
    "fmt"
////    "log"
////	"sort"
    "github.com/valyala/fastjson"
////    "os" 
////    "github.com/boltdb/bolt"
//        "errors"
//
    "strconv"
////  // "io/ioutil"
////  //"bytes"
////  //"runtime"
    "encoding/json"
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
    Ds         S.Digital_Signage_STC,
    c_time     time.Time,
    Ti         S.Time_Interval_STC,
    )  (err error) {


    //p := fmt.Println

    // p("c_time =",c_time)
    // 
    // p(int(c_time.Year()))
    // p(int(c_time.Month()))
    // p(c_time.Day())
    // p(c_time.Hour())


    //p(ps.UmNbDsTiSl_Key.NbDsTiSl_key.NbDsTi_key)
    //p(ps.UmNbDsTiSl_Key.NbDsTiSl_key.Year)

    Code_Free_Users  := fastjson.GetInt(byteValues, "Base", "Code_Free_Users")
    Code_Free_Clips  := fastjson.GetInt(byteValues, "Base", "Code_Free_Clips")


    DS_Perc_Quality, _ := strconv.Atoi(Ds.DsVal.DS_Perc_Quality)
    //fmt.Println("DS_Perc_Quality =", DS_Perc_Quality)

    TI_Price, _ := strconv.Atoi(Ti.TiVl.Price)
    //fmt.Println("TI_Price =", TI_Price)


    DS_TI_Price  := float64(TI_Price)*float64(DS_Perc_Quality)/100
    //fmt.Println("DS_TI_Price =", DS_TI_Price)



    ps.UmNbDsTiSl_Key.UsMd.ID_User                = strconv.Itoa(Code_Free_Users)
    ps.UmNbDsTiSl_Key.UsMd.ID_Media               = strconv.Itoa(Code_Free_Clips)

    ps.UmNbDsTiSl_Key.NbDsTiSl_key.Year  =c_time.Year()
    ps.UmNbDsTiSl_Key.NbDsTiSl_key.Month =int(c_time.Month()) 
    ps.UmNbDsTiSl_Key.NbDsTiSl_key.Day   =c_time.Day()   
    ps.UmNbDsTiSl_Key.NbDsTiSl_key.Hour  =c_time.Hour()  

    Slots, _ := strconv.Atoi(Ti.TiVl.Slots)
    
    for i := 1; i <= Slots; i++ {

        //price := string(data["Price"][strconv.Itoa(i)])
        //price_str := data["Price"][strconv.Itoa(i)]
        price, _ := strconv.ParseFloat(data["Price"][strconv.Itoa(i)], 64)
        //p("price =",strconv.ParseFloat (price, 64))
        // p("price_str =",price_str)
        // 
        // price, _ := strconv.ParseFloat(price_str, 64)
        // 
        // //price :=    fmt.Sprintf("%f", price_str)
        //p("price =",price)
        //p("price*DS_TI_Price =",price*DS_TI_Price)

        ps.UmNbDsTiSl_Key.NbDsTiSl_key.Index =i 

        ps.Slot_Price  = price*DS_TI_Price 
        //p("          sl ps =", ps)

        enc_UmNbDsTiSl_Key, err := json.Marshal(ps.UmNbDsTiSl_Key); __err_panic(err)
        data["Free_Slots"][string(enc_UmNbDsTiSl_Key)] = fmt.Sprintf("%f", ps.Slot_Price)  

    } // for i


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

