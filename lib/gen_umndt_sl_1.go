package lib

import (   
//    "fmt"
    "encoding/json"
    "github.com/valyala/fastjson"
//    "os" 
//    "github.com/boltdb/bolt"
//    "errors"
    "strconv"
//  // "io/ioutil"
//  //"bytes"
//  //"runtime"
//  //"encoding/gob"
//  
//  "time"
//  "math/rand"
//
    S "cds_go_1/config"

)

//----------------------------------------------
 func Get_keys_Ti_Slot(byteValues  []byte, 
        um               S.User_Media_STC      ,
        ct               S.City_STC            ,
        Nb               S.Neighborhoods_STC   ,
        Ds               S.Digital_Signage_STC ,
        Ti               S.Time_Interval_STC   ,
        data             map[string]map[string]string,
        Ow_Um_Map        map[string]float64,
        Ow_UmNbDsTi_Map  map[string]float64,
        Ow_UmNbDs_Map    map[string]float64,
        Payd_Slots_Map   map[string]float64,
        Free_Slots_Map   map[string]float64,
      ) []string {


    //fmt.Println("Get_keys_Ti_Slot Ow_Um_Map =", Ow_Um_Map)

    var ps    S.Um_NbDsTiSl_STC  // Payd_Slots
    var dsfs  S.NbDsTiSl_STC     // Free Ds Slots

    var ou       S.Ow_Um_STC
    var ound     S.Ow_UmNbDs_STC
    var oundt    S.Ow_UmNbDsTi_STC

    ou.OwUm_Key.UsMd       = um.UsMd
    ou.OwUm_Key.ID_Owner   = Ds.CnCtNbDs.ID_Owner

    enc_OwUm_Key, err := json.Marshal(ou.OwUm_Key)
    __err_panic(err)

    ound.OwUmNbDs_Key.UsMd       = um.UsMd
    ound.OwUmNbDs_Key.CnCtNbDs   = Ds.CnCtNbDs

    enc_OwUmNbDs_Key, err := json.Marshal(ound.OwUmNbDs_Key)
    __err_panic(err)

    oundt.Total_Cost                = 0
    oundt.OwUmNbDsTi_Key.UsMd       = um.UsMd

    oundt.OwUmNbDsTi_Key.NbDsTi_key.CnCtNbDs           = Ds.CnCtNbDs
    oundt.OwUmNbDsTi_Key.NbDsTi_key.ID_Time_Interval = Ti.ID_Time_Interval

    enc_OwUmNbDsTi_Key, err := json.Marshal(oundt.OwUmNbDsTi_Key)
    __err_panic(err)

    //var fs    S.Um_NbDsTiSl_STC  // Free Um Slots

    var res_Price []float64

    //fmt.Println("Get_keys_Ti_Slot um =", um)
    //fmt.Println("Get_keys_Ti_Slot ct =", ct)
    //fmt.Println("Get_keys_Ti_Slot Nb =", Nb)
    //fmt.Println("Get_keys_Ti_Slot Ds =", Ds)
    //fmt.Println("Get_keys_Ti_Slot Ti =", Ti)

    //var Ti S.Time_Interval_STC
    //var tiv S.Time_Interval_VAL
    //var err             error
    var DS_Perc_Quality    int
    var TI_Price           int
    var ID_Time_Interval   int
    var TI_Slots           int
    //var iCounter           int

    keys_Ti_Sl  := []string{}

    // DS_TI_Price  = (int(TI_Price)*DS_Perc_Quality/100)
    // #DS_TI_Price  = (int(TI_Price) * DS_Cost_Perc /100)/10*10
    //
    // TI_D_Sign_People = (int(TI_D_Sig_Ppl_Base)*DS_Perc_Quality/100)
    //
    // ti_dict = Get_Ti_Price(data,\
    //     int(ID_Time_Interval),\
    //     DS_TI_Price,\
    //     int(TI_Slots))

    ID_Time_Interval, _ = strconv.Atoi(Ti.ID_Time_Interval)
    //fmt.Println("ID_Time_Interval =", ID_Time_Interval)

    TI_Slots, _ = strconv.Atoi(Ti.TiVl.Slots)
    //fmt.Println("TI_Slots =", TI_Slots)

    DS_Perc_Quality, _ = strconv.Atoi(Ds.DsVal.DS_Perc_Quality)
    //fmt.Println("DS_Perc_Quality =", DS_Perc_Quality)

    TI_Price, _ = strconv.Atoi(Ti.TiVl.Price)
    //fmt.Println("TI_Price =", TI_Price)


    DS_TI_Price  := float64(TI_Price*DS_Perc_Quality/100)
    //fmt.Println("DS_TI_Price =", DS_TI_Price)


    Max_Total := fastjson.GetFloat64 (byteValues, "Base", "Dig_Sign_Max_Total_Price")


    Not_Used_Max_Slot_Counter := fastjson.GetInt(byteValues, "Base", "Not_Used_Max_Slot_Counter")
    //fmt.Println("Not_Used_Max_Slot_Counter =", Not_Used_Max_Slot_Counter)


    Code_Other_Users := fastjson.GetInt(byteValues, "Base", "Code_Other_Users")
    Code_Other_Clips := fastjson.GetInt(byteValues, "Base", "Code_Other_Clips")


    res_Price = Get_TiPrice(byteValues,
        ID_Time_Interval,
        DS_TI_Price,
        TI_Slots)

    //fmt.Println("res_Price =", res_Price)

    len_res_Price := len(res_Price)
    //fmt.Printf("len_res_Price = %d\n", len_res_Price)

    All_Slot_Busy := random(0, 4)
    //fmt.Println("All_Slot_Busy =",All_Slot_Busy)


    iSlot_Busy_Counter := random(1, len_res_Price-1)
    //fmt.Println("***** iSlot_Busy_Counter =", iSlot_Busy_Counter)

    // if All_Slot_Busy == 1 {
    //     iCounter = Time_Interval_Counter
    // } else {
    //     iCounter = random(1, Time_Interval_Counter)
    // }

    for i := 0; i < len_res_Price; i++ {
        n := TI_Slots - i - 1
        //fmt.Println(i,n)
        // fmt.Println(i,n, res_Price[n])


        //dsfs.NbDsTiSl_Key.ID_Slot                     = strconv.Itoa(n + 1000)
        dsfs.NbDsTiSl_Key.ID_Slot                     = strconv.Itoa(i + 1001)

        dsfs.NbDsTiSl_Key.NbDsTi_key.CnCtNbDs         = Ds.CnCtNbDs
        dsfs.NbDsTiSl_Key.NbDsTi_key.ID_Time_Interval = Ti.ID_Time_Interval
        //dsfs.Slot_Price                               = fmt.Sprintf("%f", res_Price[n])
        dsfs.Slot_Price                               = res_Price[n]

        enc_NbDsTiSl_Key, err := json.Marshal(dsfs.NbDsTiSl_Key); __err_panic(err)


        ps.UmNbDsTiSl_Key.NbDsTiSl_key                = dsfs.NbDsTiSl_Key
        ps.Slot_Price                                 = dsfs.Slot_Price 

        ps.UmNbDsTiSl_Key.UsMd.ID_User                = strconv.Itoa(Code_Other_Users)
        ps.UmNbDsTiSl_Key.UsMd.ID_Media               = strconv.Itoa(Code_Other_Clips)

        enc_UmNbDsTiSl_Key, err := json.Marshal(ps.UmNbDsTiSl_Key); __err_panic(err)



        if n <= Not_Used_Max_Slot_Counter {
                //fmt.Println("+++++dsfs =",i,n,dsfs)
                Free_Slots_Map[string(enc_NbDsTiSl_Key)]     = res_Price[n]

                if res_Price[n] < Max_Total {
                    Ow_Um_Map[string(enc_OwUm_Key)]             += res_Price[n]
                    Ow_UmNbDs_Map[string(enc_OwUmNbDs_Key)]     += res_Price[n]
                    Ow_UmNbDsTi_Map[string(enc_OwUmNbDsTi_Key)] += res_Price[n]
                } //if res_Price[n] < Max_Total 

        } else {

            if All_Slot_Busy == 1 {
                //fmt.Println("--ps =",i,n,ps)
                Payd_Slots_Map[string(enc_UmNbDsTiSl_Key)] = res_Price[n]
            } else {

                if n > iSlot_Busy_Counter {
                    //ps = fs
                    //fmt.Println("--ps =",i,n,ps)
                    Payd_Slots_Map[string(enc_OwUmNbDs_Key)]     = res_Price[n]
                } else {
                    //fmt.Println("++dsfs =",i,n,dsfs)
                    Free_Slots_Map[string(enc_NbDsTiSl_Key)]     = res_Price[n]

                    if res_Price[n] < Max_Total {
                        Ow_Um_Map[string(enc_OwUm_Key)]             += res_Price[n]
                        Ow_UmNbDs_Map[string(enc_OwUmNbDs_Key)]     += res_Price[n]
                        Ow_UmNbDsTi_Map[string(enc_OwUmNbDsTi_Key)] += res_Price[n]
                    } //if res_Price[n] < Max_Total 
                }
             }
         }
    } //for i

    return keys_Ti_Sl
 } // func Get_keys_Ti_Slot

