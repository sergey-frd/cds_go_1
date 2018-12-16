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



    //enc_OwUm_Key, err       := json.Marshal(ou.OwUm_Key);          __err_panic(err)
    //enc_OwUmNbDs_Key, err   := json.Marshal(ound.OwUmNbDs_Key);    __err_panic(err)
    //enc_OwUmNbDsTi_Key, err := json.Marshal(oundt.OwUmNbDsTi_Key); __err_panic(err)
    //enc_NbDsTiSl_Key, err   := json.Marshal(dsfs.NbDsTiSl_Key);    __err_panic(err)




    Code_Free_Users  := fastjson.GetInt(byteValues, "Base", "Code_Free_Users")
    Code_Free_Clips  := fastjson.GetInt(byteValues, "Base", "Code_Free_Clips")


    DS_Perc_Quality, _ := strconv.Atoi(Ds.DsVal.DS_Perc_Quality)
    //fmt.Println("DS_Perc_Quality =", DS_Perc_Quality)

    TI_Price, _ := strconv.Atoi(Ti.TiVl.Price)
    //fmt.Println("TI_Price =", TI_Price)


    DS_TI_Price  := float64(TI_Price)*float64(DS_Perc_Quality)/100
    //fmt.Println("DS_TI_Price =", DS_TI_Price)



    ps.UmNbDsTiSl_Key.UsMd.ID_User  = strconv.Itoa(Code_Free_Users)
    ps.UmNbDsTiSl_Key.UsMd.ID_Media = strconv.Itoa(Code_Free_Clips)

    ps.UmNbDsTiSl_Key.NbDsTiSl_key.Year  =c_time.Year()
    ps.UmNbDsTiSl_Key.NbDsTiSl_key.Month =int(c_time.Month()) 
    ps.UmNbDsTiSl_Key.NbDsTiSl_key.Day   =c_time.Day()   
    ps.UmNbDsTiSl_Key.NbDsTiSl_key.Hour  =c_time.Hour()  

    Slots, _ := strconv.Atoi(Ti.TiVl.Slots)
    
    n := 0
    for i := 1; i <= Slots; i++ {

        n++

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

        //data["Ow_Um"]       = make(map[string]string)
        //data["Ow_UmNbDs"]   = make(map[string]string)
        //data["Ow_UmNbDsTi"] = make(map[string]string)
        //data["Payd_Slots"][string(enc_UmNbDsTiSl_Key)] = "0"  


    } // for i

    return  err

} // func alloc_ow

