package lib

import (   
    "fmt"
    "os" 
    "github.com/boltdb/bolt"
    // "errors"
    "strconv"
    // "io/ioutil"
    "github.com/valyala/fastjson"

    //"bytes"

    //"runtime"
    "encoding/json"
    //"encoding/gob"
    
    "time"
    "math/rand"

    S "cds_go_1/config"
    //L "cds_go_1/lib"

)

//----------------------------------------------
func Gen_Ds_Bucket(byteValues  []byte, 
                       bucket_Name string, 
                       data        map[string]map[string]string) error {

    var ds                    S.Digital_Signage_STC      
    var i_ID_Digital_Signage  int 
    var ID_Digital_Signage    string 

    var ID_Owner              string 
    //var Owner_Name            string 
    
    var CnCtNb                S.Neighborhoods_KEY

    var  max_rand  int
    //var  temp_rand int
    var  nn        int



    dbFileName, err := GetDbName(byteValues);  __err_panic(err) 
    if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
        __err_panic(err)

    }
    //fmt.Println(" ----------  L.Buckets Stat IsNotExist ---------- ")
    
    db, err := bolt.Open(dbFileName, 0600, nil)
    if err != nil {
        __err_panic(err)
    }
    defer db.Close()
    

    digital_Signage_Cost_Max := fastjson.GetInt(byteValues, "Base", "Digital_Signage_Cost_Max")
    digital_Signage_Cost_Min := fastjson.GetInt(byteValues, "Base", "Digital_Signage_Cost_Min")

    // fmt.Printf("digital_Signage_Cost_Max = %d\n", digital_Signage_Cost_Max)
    // fmt.Printf("digital_Signage_Cost_Min = %d\n", digital_Signage_Cost_Min)

    d_Sign_Counter := fastjson.GetInt(byteValues, "Base", "D_Sign_Counter")
    //fmt.Printf("d_Sign_Counter = %d\n", d_Sign_Counter)

    Owner_Counter := fastjson.GetInt(byteValues, "Base", "Owner_Counter")
    //fmt.Printf("Owner_Counter = %d\n", Owner_Counter)

    len_data_Owner := len(data["Owner"])
    //fmt.Printf("len_data_Owner = %d\n", len_data_Owner)

    // !!!!!!!!!!!!!!!!!!!!!!!!!
    len_data_Owner =  Owner_Counter

    keys_Owner := []string{}
    for k := range data["Owner"] {
        keys_Owner = append(keys_Owner, k)
    }

    //fmt.Println("keys_Owner =",keys_Owner)

    //fmt.Println("keys_Owner[0] =",keys_Owner[0])

    err = db.Update(func(tx *bolt.Tx) error {

        b, err := tx.CreateBucketIfNotExists([]byte(bucket_Name))
        __err_panic(err)
        //fmt.Println(  "bucket_Name b =",bucket_Name,b)

        nn = 0

        //i_ID_Digital_Signage = 0

        for key, value := range data["Neighborhoods"] {
  
            //fmt.Println("Key =", key, "Value =", value)

            byt := []byte(key)
            err = json.Unmarshal(byt, &CnCtNb)
            if err != nil {
                fmt.Println("There was an error:", err)
            }
            //fmt.Println(  "CnCtNb =",CnCtNb)
            //fmt.Println(  "CnCtNb.CnCt =",CnCtNb.CnCt)
            // fmt.Println(  "CnCtNb.CnCt.ID_City =",CnCtNb.CnCt.ID_City)
            // fmt.Println(  "CnCtNb.CnCt.ID_Country =",CnCtNb.CnCt.ID_Country)
            // fmt.Println(  "CnCtNb.ID_Neighborhoods =",CnCtNb.ID_Neighborhoods)

            ds.CnCtNbDs.CnCtNb = CnCtNb

            //max_rand = Random(1,rand.Intn(d_Sign_Counter) )
    
            //temp_rand = Random(1, d_Sign_Counter)
            //max_rand = temp_rand


            rand.Seed(int64(time.Now().Nanosecond()))
            max_rand = rand.Intn(d_Sign_Counter)
            max_rand  += 1
            //fmt.Println(  "max_rand =",max_rand)


            //....................................................
            for i := 1; i <= max_rand; i++ {
            //for i := 1; i <= d_Sign_Counter; i++ {

                //fmt.Println(  "nn =",nn)

                // temp_rand = Random(1, d_Sign_Counter)
                // fmt.Println(  "temp_rand =",temp_rand)

                rand.Seed(int64(time.Now().Nanosecond()))
                rand_Owner_Key := rand.Intn(len_data_Owner)
                //fmt.Println("rand_Owner_Key =",rand_Owner_Key)
                ID_Owner = keys_Owner[rand_Owner_Key]
                //fmt.Println("ID_Owner =",ID_Owner)

                // Owner_Name = data["Owner"][ID_Owner]
                // fmt.Println("Owner_Name =",Owner_Name)



                // DS_Cost = random.randint(\
                //     int(data['Base']["Digital_Signage_Cost_Min"]),\
                //     int(data['Base']["Digital_Signage_Cost_Max"]))*1000
                // DS_Perc_Quality = 100 + DS_Cost/100


                rand.Seed(int64(time.Now().Nanosecond()))
                rand_DS_Cost    := (digital_Signage_Cost_Min + rand.Intn(digital_Signage_Cost_Max))*1000
                DS_Cost         := strconv.Itoa(rand_DS_Cost)
                DS_Perc_Quality := strconv.Itoa(100 + rand_DS_Cost/100)

                //fmt.Println("DS_Cost         =",DS_Cost        )
                //fmt.Println("DS_Perc_Quality =",DS_Perc_Quality)


                i_ID_Digital_Signage = i
                ID_Digital_Signage    = strconv.Itoa(i_ID_Digital_Signage)
                //ID_Digital_Signage    = strconv.Itoa(nn)

                ds.CnCtNbDs.ID_Digital_Signage = ID_Digital_Signage
                ds.CnCtNbDs.ID_Owner           = ID_Owner

                ds.DsVal.Dig_Sign        = value + "_" + strconv.Itoa(i_ID_Digital_Signage)
                ds.DsVal.DS_Cost         = DS_Cost        
                ds.DsVal.DS_Perc_Quality = DS_Perc_Quality

                //fmt.Println("i_ID_Digital_Signage =", i_ID_Digital_Signage)
                //fmt.Println("ID_Digital_Signage   =", ID_Digital_Signage)

                //fmt.Println("ds =", ds)

                encoded, err := json.Marshal(ds.CnCtNbDs)
                __err_panic(err)

                enc_val, err := json.Marshal(ds.DsVal)
                __err_panic(err)


                // fmt.Println("encoded =", string(encoded[:]),
                //                          string(enc_val[:]))


                err = b.Put([]byte(encoded), []byte(enc_val))
                __err_panic(err)

                //fmt.Println("encoded =", string(encoded[:]),
                //ds.DsVal.ID_Owner, 
                //ds.DsVal.Dig_Sign,
                //ds.DsVal.DS_Cost         ,
                //ds.DsVal.DS_Perc_Quality )

            }

            // if nn >= 20 {
            //     break	
            // }



            // break
        }

            // for index_Row, row := range sheet.Rows {
            // 
            //     if index_Row == 0 {
            //         continue	
            //     }
            // 
            //     nb.CnCtNb.CnCt.ID_Country  = row.Cells[0].String()
            //     nb.CnCtNb.CnCt.ID_City     = row.Cells[1].String()
            //     nb.CnCtNb.ID_Neighborhoods = row.Cells[2].String()
            //     nb.Neighborhoods           = row.Cells[3].String()
            // 
            //     encoded, err := json.Marshal(nb.CnCtNb)
            //     __err_panic(err)
            // 
            //     err = b.Put([]byte(encoded), []byte(nb.Neighborhoods))
            //     __err_panic(err)
            // 
            // 
            // }



            return err

    })  // err = db.Update(func(tx *bolt.Tx) error 



    return err

}


