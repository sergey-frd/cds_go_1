package lib

import (   
    "fmt"
    "github.com/valyala/fastjson"
    "math"

//    "os" 
//    "github.com/boltdb/bolt"
//    "errors"

//    "strconv"    
//    "encoding/json"

//  // "io/ioutil"
//  //"bytes"
//  //"runtime"
//  //"encoding/gob"
//  
//  "time"
//  "math/rand"
//
//    S "cds_go_1/config"

)


//----------------------------------------------
func Get_Ti_Price(
    byteValues  []byte,
    ti_type     int,
    base        float64, 
    max         int,
) {

    var res     float64
    //var result  float64
    var cur     int
    var steps   int


    cur = 1

    x        := fastjson.GetFloat64 (byteValues, "Base", "Dig_Sign_Max_Time_Interv_Price")
    //koef_min := fastjson.GetFloat64 (byteValues, "Base", "Dig_Sign_Time_Int_SQR_min")
    koef_max := fastjson.GetFloat64 (byteValues, "Base", "Dig_Sign_Time_Int_SQR_max")

    fmt.Println("==============================")
    fmt.Println("ti_type = ", ti_type)
    fmt.Println("base    = ", base   )
    fmt.Println("max     = ", max    )

    // fmt.Println("x        = ", x       )
    // fmt.Println("koef_max = ", koef_max)

    // fmt.Println("koef_min = ", koef_min)


    koef := koef_max

    switch ti_type {
    //switch data {

    case 1:
         koef = .9
    case 2:
         koef = .88
    case 3:
         koef = .86
    case 4:
         koef = .8
    }

    steps  = max - cur
    res    = x
    //result = res

    for i := 1; i <= steps+1; i++ {
        //fmt.Println("i =",i)
        
        res = math.Pow(res, koef) + base
        //fmt.Println(i, res + base)

    } // for i := 1; 
    //fmt.Println("steps     = ", steps+1    )

} // func Gen_Um_Bucket(

//----------------------------------------------
func Gen_nb_us_ds_md_ti_detail(byteValues  []byte) {

    Get_Ti_Price(byteValues,4,48,10)
    Get_Ti_Price(byteValues,3,36,20)
    Get_Ti_Price(byteValues,2,24,30 )
    Get_Ti_Price(byteValues,1,12,40 )

} // func Gen_Um_Bucket(


//----------------------------------------------
func Get_TiPrice(
    byteValues  []byte,
    ti_type     int,
    base        float64, 
    max         int,
) []float64 {


    var res_Price     []float64

    var res     float64
    var cur     int
    var steps   int

    cur = 1

    x        := fastjson.GetFloat64 (byteValues, "Base", "Dig_Sign_Max_Time_Interv_Price")
    //koef_min := fastjson.GetFloat64 (byteValues, "Base", "Dig_Sign_Time_Int_SQR_min")
    koef_max := fastjson.GetFloat64 (byteValues, "Base", "Dig_Sign_Time_Int_SQR_max")

    //fmt.Println("==============================")
    //fmt.Println("ti_type = ", ti_type)
    //fmt.Println("base    = ", base   )
    //fmt.Println("max     = ", max    )

    // fmt.Println("x        = ", x       )
    // fmt.Println("koef_max = ", koef_max)

    // fmt.Println("koef_min = ", koef_min)


    koef := koef_max

    switch ti_type {
    //switch data {

    case 1:
         koef = .9
    case 2:
         koef = .88
    case 3:
         koef = .86
    case 4:
         koef = .8
    }

    steps  = max - cur
    res    = x
    //result = res

    for i := 1; i <= steps+1; i++ {
        //fmt.Println("i =",i)
        
        res = math.Pow(res, koef) + base
        //fmt.Println(i, res + base)
        res_Price = append(res_Price, res)

    } // for i := 1; 
    //fmt.Println("steps     = ", steps+1    )
     return res_Price
} // func Gen_Um_Bucket(


