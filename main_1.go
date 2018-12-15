package main

import (   
    "os" 
    //	"errors"
    "log"
    "fmt"
    "github.com/valyala/fastjson"
    "io/ioutil"
    //"path"
    "path/filepath"

    L "cds_go_1/lib"
    X "cds_go_1/lib_xcls"
//    A "cds_go_1/lib_alloc"
    G "cds_go_1/lib_gen"

    //"encoding/json"
    //"encoding/gob"
    //S "cds_go_1/config"

)

//---------------------------------------------------------------
func __err_panic(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

// //---------------------------------------------------------------
// func handlingCountryLine(sheet_Name string,
//                          index_Row int,
//                          index_Cells int, 
//                          text_Cells string) (err error) {
// 
// 
//     fmt.Printf("%s %d %d %s\n",
//         sheet_Name,
//         index_Row,
//         index_Cells, 
//         text_Cells)
// 
//     return 
// }

//---------------------------------------------------------------
func main() {

    p := fmt.Println


    var data            = map[string]map[string]string{}
    var Ow_Um_Map       = make(map[string]float64)
    var Ow_UmNbDsTi_Map = make(map[string]float64)
    var Ow_UmNbDs_Map   = make(map[string]float64)
    var Payd_Slots_Map  = make(map[string]float64)
    var Free_Slots_Map  = make(map[string]float64)


    proj_dir, err := os.Getwd();  __err_panic(err) 
    p("proj_dir =", proj_dir)

    // dir_proj_dir :=  filepath.Dir (proj_dir) 
    // p("dir_proj_dir =", dir_proj_dir)

	p(filepath.Join(proj_dir,"config", "t_4.json"))

    //jsonFilePath := "c:\\Users\\user\\go\\src\\data_xcls\\config\\t_4.json"
    jsonFilePath := filepath.Join(proj_dir,"config", "t_4.json")

    //fmt.Println("jsonFilePath =", jsonFilePath)


    //============================================  
    //return

    jsonFile, err := os.Open(jsonFilePath)
    if err != nil {
    fmt.Println(err)
    }
    defer jsonFile.Close()

    // read our opened xmlFile as a byte array.
    byteValues, _ := ioutil.ReadAll(jsonFile)

    dbFileName, err := L.GetDbName(byteValues);  __err_panic(err) 

    //fmt.Println("call L.Demo() START")
    //L.Demo(dbFileName,byteValues)
    //L.DemoGenDB(byteValues)

    //------------------------------------------------------------------------------
    if fastjson.GetString(byteValues, "Base", "CASE_XCLS_2_DB") == "Y" {
        //fmt.Println("call X.Demo_Xcls(jsonFilePath) START")
        // X.Load_Xcls(proj_dir, byteValues)
        X.Load_Xcls( byteValues)
    }

    //------------------------------------------------------------------------------
    CASE_DB_PRINT_ALL := fastjson.GetString(byteValues, "Base", "CASE_DB_PRINT_ALL")
    fmt.Printf("CASE_DB_PRINT_ALL = %s\n", CASE_DB_PRINT_ALL)

    if CASE_DB_PRINT_ALL == "Y" {

        fmt.Println(" ----------  L.Buckets(dbFileName)  ---------- ")
    	L.Buckets(dbFileName)

    }

    //------------------------------------------------------------------------------
    CASE_LOAD_DICT := fastjson.GetString(byteValues, "Base", "CASE_LOAD_DICT")
    fmt.Printf("CASE_LOAD_DICT = %s\n", CASE_LOAD_DICT)
    if CASE_LOAD_DICT == "Y" {

        fmt.Println(" ---------- L.LoadDict(dbFileName)  ---------- ")
    	//L.LoadDict(dbFileName,data)
    	L.LoadDict(byteValues,data)

    }

    //------------------------------------------------------------------------------
    CASE_GEN_ALL_FILES := fastjson.GetString(byteValues, "Base", "CASE_GEN_ALL_FILES")
    fmt.Printf("CASE_GEN_ALL_FILES = %s\n", CASE_GEN_ALL_FILES)
    if CASE_GEN_ALL_FILES == "Y" {

        //L.GenAllFiles(byteValues)

        //.................................................
        err = L.Gen_Ds_Bucket(byteValues, "Digital_Signage",data); __err_panic(err)
        err = L.LoadDict2(byteValues, data, "Digital_Signage");    __err_panic(err)

        err = L.Gen_Um_Bucket(byteValues, "User_Media", data);     __err_panic(err)
        err = L.LoadDict2(    byteValues, data, "User_Media");     __err_panic(err)

        //.................................................
        //err = L.Print_DB_Bucket(byteValues, "Price");             __err_panic(err)
        //err = L.Print_DB_Bucket(byteValues, "Time_Interval");     __err_panic(err)
        //err = L.Print_DB_Bucket(byteValues, "Media");             __err_panic(err)
        //err = L.Print_DB_Bucket(byteValues, "User_Media");        __err_panic(err)


        data["Free_Slots"] = make(map[string]string)
        err = G.Gen_Lvl_Um(byteValues,data,);  __err_panic(err)

        //p("data[Free_Slots] =", data["Free_Slots"])
        err = L.Save_Data_Map(byteValues, "Free_Slots"  , data ); __err_panic(err)
        //err = L.Print_DB_Bucket(byteValues, "Free_Slots");             __err_panic(err)

        //!!!!!!!!!!!!!!!!!!!!!!!!!!!
        //return 


        // L.Gen_UmNbDtTi_Bucket(byteValues, 
        //     data            ,
        //     Ow_Um_Map       ,
        //     Ow_UmNbDsTi_Map ,
        //     Ow_UmNbDs_Map   ,
        //     Payd_Slots_Map  ,
        //     Free_Slots_Map  ,
        //     )
        // __err_panic(err)


        // fmt.Println("main Ow_Um_Map =", Ow_Um_Map)
        // fmt.Println("main Ow_UmNbDs_Map =", Ow_UmNbDs_Map)
        // fmt.Println("main Ow_UmNbDsTi_Map =", Ow_UmNbDsTi_Map)

        // fmt.Println("main Payd_Slots_Map =", Payd_Slots_Map)


        // for Ow_UmNbDs_Map_Key, v := range Ow_UmNbDs_Map {
        //     fmt.Println("main Ow_UmNbDs_Map_Key =", Ow_UmNbDs_Map_Key)
        //     fmt.Println("main v =", v)
        // } // for Ow_UmNbDs_Map_Key


        err = L.Save_Map(byteValues, "Ow_Um"       , Ow_Um_Map      ); __err_panic(err)
        err = L.Save_Map(byteValues, "Ow_UmNbDs"   , Ow_UmNbDs_Map  ); __err_panic(err)
        err = L.Save_Map(byteValues, "Ow_UmNbDsTi" , Ow_UmNbDsTi_Map); __err_panic(err)
        err = L.Save_Map(byteValues, "Payd_Slots"  , Payd_Slots_Map ); __err_panic(err)
        err = L.Save_Map(byteValues, "Free_Slots"  , Free_Slots_Map ); __err_panic(err)

                  
        //err = L.LoadDict_Dbg(byteValues, data, "Ow_Um")          
        err = L.LoadDict2(byteValues, data, "Ow_Um"      );  __err_panic(err)          
        err = L.LoadDict2(byteValues, data, "Ow_UmNbDs"  );  __err_panic(err)      
        err = L.LoadDict2(byteValues, data, "Ow_UmNbDsTi");  __err_panic(err)   
        err = L.LoadDict2(byteValues, data, "Payd_Slots" );  __err_panic(err)     
        err = L.LoadDict2(byteValues, data, "Free_Slots" );  __err_panic(err)     

        // fmt.Println("main data[Ow_Um]       =", data["Ow_Um"]);       __err_panic(err)
        // fmt.Println("main data[Ow_UmNbDs]   =", data["Ow_UmNbDs"]);   __err_panic(err)
        // fmt.Println("main data[Ow_UmNbDsTi] =", data["Ow_UmNbDsTi"]); __err_panic(err)


    } // if CASE_GEN_ALL_FILES == "Y" {

    //------------------------------------------------------------------------------
    CASE_UM_ALLOCATION := fastjson.GetString(byteValues, "Base", "CASE_UM_ALLOCATION")
    fmt.Printf("CASE_UM_ALLOCATION = %s\n", CASE_UM_ALLOCATION)

    if CASE_UM_ALLOCATION == "Y" {

        err = L.LoadDict2(byteValues, data, "Digital_Signage");   __err_panic(err)
        err = L.LoadDict2(byteValues, data, "User_Media"     );   __err_panic(err)

        //err = L.LoadDict2(byteValues, data, "Ow_Um"      ); __err_panic(err)          
        //err = L.LoadDict2(byteValues, data, "Ow_UmNbDs"  ); __err_panic(err)      
        //err = L.LoadDict2(byteValues, data, "Ow_UmNbDsTi"); __err_panic(err)   
        //err = L.LoadDict2(byteValues, data, "Payd_Slots" ); __err_panic(err)     
        err = L.LoadDict2(byteValues, data, "Free_Slots" ); __err_panic(err)     

        //fmt.Println("main data[Ow_Um]       =", data["Ow_Um"]);       __err_panic(err)
        //fmt.Println("main data[Ow_UmNbDs]   =", data["Ow_UmNbDs"]);   __err_panic(err)
        //fmt.Println("main data[Ow_UmNbDsTi] =", data["Ow_UmNbDsTi"]); __err_panic(err)
        //fmt.Println("main data[Free_Slots]  =", data["Free_Slots"]);  __err_panic(err)

        //err = A.Alloc_Um(byteValues, 
        //                 data ,
        //    ); __err_panic(err) 



    } // if CASE_UM_ALLOCATION == "Y" {

    //------------------------------------------------------------------------------
    //fmt.Println(data["City"])

    //for key, value := range data["Neighborhoods"] {

    // for key, value := range data["City"] {
    //     fmt.Println("Key:", key, "Value:", value)
    // }


    // for key, value := range data["Neighborhoods"] {
    //     fmt.Println("Key:", key, "Value:", value)
    // 
    // 
    //     var CnCtNb S.Neighborhoods_KEY
    //     err = json.Unmarshal([]byte(key)  , &CnCtNb)
    //     if err != nil {
    //         fmt.Println("There was an error:", err)
    //     }
    //     //fmt.Println(  "CnCtNb =",CnCtNb)
    //     //fmt.Println(  "CnCtNb.CnCt =",CnCtNb.CnCt)
    //     fmt.Println(  "CnCtNb.CnCt.ID_Country =",CnCtNb.CnCt.ID_Country)
    //     fmt.Println(  "CnCtNb.CnCt.ID_City =",CnCtNb.CnCt.ID_City)
    //     fmt.Println(  "CnCtNb.ID_Neighborhoods =",CnCtNb.ID_Neighborhoods)
    //     fmt.Println(  "Neighborhoods =",value)
    // 
    // 
    // }



    return




}
