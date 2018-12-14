package lib_gen

import (   
    "fmt"
    "log"
	"sort"
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
//func Alloc_Um(byteValues  []byte, 
func Gen_Lvl_Um(byteValues  []byte, 
    data            map[string]map[string]string,
    ) error {


    var err         error
    //var total       float64

    var um   S.User_Media_STC
    var ps   S.Um_NbDsTiSl_STC  // Payd_Slots

    var keys []string
    for k, _ := range data["User_Media"] {
        keys = append(keys, k)
    } // for k, v
    sort.Strings(keys)  // or sort.Ints(keys), sort.Sort(...), etc., per <K>
    for _, k := range keys  {
        v := data["User_Media"][k]


        //.................................................
        byt_k := []byte(k)
        //err := json.Unmarshal(byt_ds, &k_Ds)
        err = json.Unmarshal(byt_k, &ps.UmNbDsTiSl_Key.UsMd)
        err = json.Unmarshal(byt_k, &um.UsMd)
        if err != nil {
            fmt.Println("There was an error:", err)
        }

        byt_v := []byte(v)
        err = json.Unmarshal(byt_v, &um.UsMdVl)
        if err != nil {
            fmt.Println("There was an error:", err)
        }
        //fmt.Println("um =", um)


        //fmt.Println("um =", um)
        fmt.Println("um ps =", ps)
        ps.UmNbDsTiSl_Key.UsMd = um.UsMd
        err = Gen_Lvl_Ct(byteValues,data,ps,um,);  __err_panic(err)


        //!!!!!!!!!!!!!!!!!!!!
        break

    }

    return  err

} // func alloc_ow

// //----------------------------------------------
// func Gen_Lvl_Um_1(byteValues  []byte, 
//     data            map[string]map[string]string,
//     ) error {
// 
//     var err    error
//     var um     S.User_Media_STC
//     var ps     S.Um_NbDsTiSl_STC  // Payd_Slots
// 
//     //............................................................
//     dbFileName, err := L.GetDbName(byteValues);  __err_panic(err) 
//     if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
//         fmt.Println(err)
//         return err
//     }
// 
//     db, err := bolt.Open(dbFileName, 0600, nil)
//     if err != nil {
//         fmt.Println(err)
//         return err
//     }
//     defer db.Close()
// 
//     //fmt.Println(" ----------  Open ---------- ")
//     var sheet_Name string     = "User_Media" 
//     err = db.View(func(tx *bolt.Tx) error {
// 
// 			//sheet_Name := string(bucket_Name)
//             //fmt.Println(string(name))
// 			//fmt.Printf(" *** Print_DB_Bucket %s *** \n", sheet_Name)
// 				 
// 			//data[sheet_Name] = make(map[string]string)
// 			err = db.View(func(tx *bolt.Tx) error {
// 				b := tx.Bucket([]byte(sheet_Name))
// 				c := b.Cursor()
// 				if b == nil {
// 					return errors.New("Bucket not found")
// 				}
// 				//log.Printf("Value of %s is %s", "answer", b.Get([]byte("answer")))
// 			
//                 nnn := 0
// 				for k, v := c.First(); k != nil; k, v = c.Next() {
// 				//for k, _ := c.First(); k != nil; k, _ = c.Next() {
//                     nnn  += 1
// 
//                     fmt.Println(nnn, string(k[:]), "=>",string(v[:]))
// 
//                     //.................................................
//                     byt_k := []byte(k)
//                     //err := json.Unmarshal(byt_ds, &k_Ds)
//                     err = json.Unmarshal(byt_k, &ps.UmNbDsTiSl_Key.UsMd)
//                     if err != nil {
//                         fmt.Println("There was an error:", err)
//                     }
// 
//                     byt_v := []byte(v)
//                     err = json.Unmarshal(byt_v, &um.UsMdVl)
//                     if err != nil {
//                         fmt.Println("There was an error:", err)
//                     }
//                     fmt.Println("ps =", ps)
// 
//                     //.................................................
//                     // //keys_City := Get_keys_City(byteValues, 
//                     // Get_keys_City(byteValues, 
//                     //     data,
//                     //     um,
//                     //     Ow_Um_Map,
//                     //     Ow_UmNbDsTi_Map   ,
//                     //     Ow_UmNbDs_Map   ,
//                     //     Payd_Slots_Map,
//                     //     Free_Slots_Map  ,
//                     //     )
//                     // 
//                     // //fmt.Println("keys_City =", keys_City)
//                     // 
//                     // //!!!!!!!!!!
//                     // //break
// 
// 
// 				}
// 				return err
// 			})
// 
// 
//             return err
//     })
// 
// 
// 
//     return err
// 
// } //func Gen_Um(byteValues  []byte, 
// 
// 
