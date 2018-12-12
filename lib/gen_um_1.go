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

)

//----------------------------------------------
func Gen_Um_Bucket(byteValues  []byte, 
                       bucket_Name string, 
                       data        map[string]map[string]string) error {

    // var ds                    S.Digital_Signage_STC      
    // var i_ID_Digital_Signage  int 
    // var ID_Digital_Signage    string 
    // 
    // var ID_Owner              string 
    // //var Owner_Name            string 
    // 
    // var CnCtNb                S.Neighborhoods_KEY
    // 
    // var  max_rand  int
    // //var  temp_rand int
    var  nn        int
    var ID_User    string 
    var ID_Media    string 
    
    var User_Name  string 
    var Media_Cost string 
    var Media_Key  string 
    var MdVal     S.Md_VAL
    var um        S.User_Media_STC

    Clip_Budget_Min := fastjson.GetInt(byteValues, "Base", "Clip_Budget_Min")
    Clip_Budget_Max := fastjson.GetInt(byteValues, "Base", "Clip_Budget_Max")

    User_Counter := fastjson.GetInt(byteValues, "Base", "User_Counter")
    //fmt.Printf("User_Counter = %d\n", User_Counter)

    len_data_User := len(data["User"])
    //fmt.Printf("len_data_User = %d\n", len_data_User)


    // !!!!!!!!!!!!!!!!!!!!!!!!!
    len_data_User =  User_Counter

    keys_User := []string{}
    for k := range data["User"] {
        keys_User = append(keys_User, k)
    }
    //fmt.Println("keys_User =",keys_User)

    rand.Seed(int64(time.Now().Nanosecond()))
    rand_User_Key := rand.Intn(len_data_User)
    //fmt.Println("rand_User_Key =",rand_User_Key)
    ID_User = keys_User[rand_User_Key]
    //fmt.Println("ID_User =",ID_User)

    User_Name = data["User"][ID_User]
    //fmt.Println("User_Name =",User_Name)




    len_data_Media := len(data["Media"])
    //fmt.Printf("len_data_Media = %d\n", len_data_Media)
    keys_Media := []string{}
    for k := range data["Media"] {
        keys_Media = append(keys_Media, k)
    }
    //fmt.Println("keys_Media =",keys_Media)



    dbFileName := fastjson.GetString(byteValues, "Base", "dbFileName")
    if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
        __err_panic(err)

    }
    //fmt.Println(" ----------  L.Buckets Stat IsNotExist ---------- ")
    
    db, err := bolt.Open(dbFileName, 0600, nil)
    if err != nil {
        __err_panic(err)
    }
    defer db.Close()

    err = db.Update(func(tx *bolt.Tx) error {

        b, err := tx.CreateBucketIfNotExists([]byte(bucket_Name))
        __err_panic(err)
        //fmt.Println(  "bucket_Name b =",bucket_Name,b)

        nn = 0
        for key, _ := range data["User"] {
            nn  += 1
            //fmt.Println(nn, "Key =", key, "Value =", value)
            //fmt.Println(nn, "Key =", key)

            ID_User   = key
            User_Name = data["User"][key]
            //fmt.Println("User_Name =",User_Name)


            rand.Seed(int64(time.Now().Nanosecond()))
            rand_Media_Key := rand.Intn(len_data_Media) 
            
            Media_Key = keys_Media[rand_Media_Key]
            //ID_Media = strconv.Itoa(Media_Key)
            ID_Media = Media_Key
            //fmt.Println("ID_Media =",ID_Media)


            Media_Dict_Val := data["Media"][Media_Key]
            //fmt.Println("Media_Dict_Val =",Media_Dict_Val)


            byt := []byte(Media_Dict_Val)
            err = json.Unmarshal(byt, &MdVal)
            if err != nil {
                fmt.Println("There was an error:", err)
            }
            //fmt.Println(  "MdVal =",MdVal)
            // fmt.Println(  "MdVal.Type_Media =",MdVal.Type_Media)
            // fmt.Println(  "MdVal.Slots      =",MdVal.Slots     )

            //Media_Cost = random.randint(\
            //    data['Base']['Clip_Budget_Min'],\
            //    data['Base']['Clip_Budget_Max'],\
            //    )*1000

            rand.Seed(int64(time.Now().Nanosecond()))
            //rand_Budget_Min := rand.Intn(Clip_Budget_Min)
            rand_Budget     := rand.Intn(Clip_Budget_Max)
            if rand_Budget < Clip_Budget_Min {
                rand_Budget = Clip_Budget_Min
            }

            Media_Cost = strconv.Itoa(rand_Budget*1000)

            um.UsMd.ID_User  = ID_User
            um.UsMd.ID_Media = ID_Media

            um.UsMdVl.Media_Name  = User_Name +"_"+ MdVal.Type_Media +"_"+ ID_Media
            um.UsMdVl.Media_Cost  = Media_Cost
            um.UsMdVl.Media_Slots = ID_Media


            //fmt.Println("um =",um)

            encoded, err := json.Marshal(um.UsMd)
            __err_panic(err)

            enc_val, err := json.Marshal(um.UsMdVl)
            __err_panic(err)

            err = b.Put([]byte(encoded), []byte(enc_val))
            __err_panic(err)


            if nn >= User_Counter {
                break	
            }
        }


        return err

    })  // err = db.Update(func(tx *bolt.Tx) error 

    return err

}


