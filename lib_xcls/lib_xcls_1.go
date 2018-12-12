package lib_xcls

import (   
    "fmt"
	 "github.com/boltdb/bolt"
     "os" 
//	 "errors"
	 "log"
     "github.com/tealeg/xlsx"
     //"github.com/valyala/fastjson"
//	 "io/ioutil"

    //"bytes"
    //"runtime"
    "encoding/json"
    //"encoding/gob"
//    "path/filepath"

    L "cds_go_1/lib"
    S "cds_go_1/config"

)

//---------------------------------------------------------------
func __err_panic(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}


//------------------------------------------------------------
//func Load_Xcls(proj_dir string, byteValues []byte) {
func Load_Xcls( byteValues []byte) {

    p := fmt.Println
    fmt.Println("Load_Xcls Started")

    var ti   S.Time_Interval_STC
    var nb   S.Neighborhoods_STC
    var md   S.Media_STC
    var ct   S.City_STC

    var ID_Country         string
    //var ID_City            string
    //var ID_Neighborhoods   string
    var ID_User            string
    var ID_Owner           string
    //var ID_Media           string
    //var Type_Media         string

    var Country            string
    //var City               string
    //var Neighborhoods      string
    var Nic_User           string
    var Owner_Name         string
    // var Slots              string

    var key                string
    var val                string


    //project_name := fastjson.GetString(byteValues, "Base", "project_name")
    //fmt.Printf("project_name = %s\n", project_name)




    // proj_dir, err := os.Getwd();  __err_panic(err) 
    // p("proj_dir proj_dir =", proj_dir)
    // excelFileName := filepath.Join(proj_dir,"tbl", excel_Name)



    //dbFileName := fastjson.GetString(byteValues, "Base", "dbFileName")
    dbFileName, err := L.GetDbName(byteValues);  __err_panic(err) 
    p("dbFileName =", dbFileName)

    excelFileName, err := L.GetExcelFileName(byteValues);  __err_panic(err) 
    //fmt.Printf("proj_dir excelFileName = %s\n", excelFileName)


    fmt.Println(" ----------  CASE_XCLS_2_DB ---------- ")
    xlFile, err := xlsx.OpenFile(excelFileName)
    __err_panic(err)


    _, err = os.Stat(dbFileName)
    if err == nil {
        log.Printf("file %s exists", dbFileName)   
        err := os.Remove(dbFileName)
        __err_panic(err)
    } else {
        log.Printf("file %s stat error: %v", dbFileName, err)
    }

    // fmt.Println("call L.Demo() START")
    // L.Demo(dbFileName,byteValues)

    //db, err := bolt.Open(dbFile, 0600, nil)
    //db, err := bolt.Open(dbFileName, 0600, nil)
    db, err :=L.SetupDB(dbFileName)
    __err_panic(err)

    //defer db.Close()
    //__err_panic(err)



    // var buf bytes.Buffer        // Stand-in for a network connection
    // enc := gob.NewEncoder(&buf) // Will write to network.
    // dec := gob.NewDecoder(&buf) // Will read from network.

    err = db.Update(func(tx *bolt.Tx) error {

            for _, sheet := range xlFile.Sheets {
                sheet_Name := sheet.Name
                //fmt.Printf(" *********  %s  ************* \n", sheet_Name)

                b, err := tx.CreateBucketIfNotExists([]byte(sheet_Name))
                __err_panic(err)

                switch sheet_Name {

                case "Time_Interval":

                    for index_Row, row := range sheet.Rows {

                        if index_Row == 0 {
                            continue	
                        }
                               
                        ti.ID_Time_Interval      = row.Cells[0].String() 
                        ti.TiVl.Price            = row.Cells[1].String() 
                        ti.TiVl.D_Sign_People    = row.Cells[2].String() 
                        ti.TiVl.Slots            = row.Cells[3].String() 

                        encoded, err := json.Marshal(ti.TiVl)
                        __err_panic(err)
                                       
                        err = b.Put([]byte(ti.ID_Time_Interval), encoded)

                    }

                //................................................
                case "Neighborhoods":

                    for index_Row, row := range sheet.Rows {

                        if index_Row == 0 {
                            continue	
                        }

                        nb.CnCtNb.CnCt.ID_Country  = row.Cells[0].String()
                        nb.CnCtNb.CnCt.ID_City     = row.Cells[1].String()
                        nb.CnCtNb.ID_Neighborhoods = row.Cells[2].String()
                        nb.Neighborhoods           = row.Cells[3].String()
     
                        encoded, err := json.Marshal(nb.CnCtNb)
                        __err_panic(err)

                        err = b.Put([]byte(encoded), []byte(nb.Neighborhoods))
                        __err_panic(err)


                    }

                //................................................
                case "City":

                    for index_Row, row := range sheet.Rows {

                        if index_Row == 0 {
                            continue	
                        }

                        ct.CnCt.ID_Country = row.Cells[0].String()
                        ct.CnCt.ID_City    = row.Cells[1].String()
                        ct.City            = row.Cells[2].String()
     
                        encoded, err := json.Marshal(ct.CnCt)
                        __err_panic(err)

                        err = b.Put([]byte(encoded), []byte(ct.City))
                        __err_panic(err)

                    }

                //................................................
                case "Country":

                    for index_Row, row := range sheet.Rows {

                        if index_Row == 0 {
                            continue	
                        }

                        ID_Country  = row.Cells[0].String()
                        Country     = row.Cells[1].String()

                        key =   ID_Country

                        val = Country

                        err = b.Put([]byte(key), []byte(val))
                        __err_panic(err)


                    }
                //................................................
                case "User":

                    for index_Row, row := range sheet.Rows {

                        if index_Row == 0 {
                            continue	
                        }

                        ID_User  = row.Cells[0].String()
                        Nic_User = row.Cells[1].String()

                        key =   ID_User 
                        val   = Nic_User

                        err = b.Put([]byte(key), []byte(val))
                        __err_panic(err)


                    }
                //................................................

                case "Owner":

                    for index_Row, row := range sheet.Rows {

                        if index_Row == 0 {
                            continue	
                        }

                        ID_Owner   = row.Cells[0].String()
                        Owner_Name = row.Cells[1].String()

                        key =   ID_Owner  
                        val   = Owner_Name

                        err = b.Put([]byte(key), []byte(val))
                        __err_panic(err)


                    }
                //................................................

                case "Media":

                    for index_Row, row := range sheet.Rows {

                        if index_Row == 0 {
                            continue	
                        }

                        md.ID_Media   = row.Cells[0].String()
                        md.MdV.Type_Media = row.Cells[1].String()
                        md.MdV.Slots      = row.Cells[2].String()
     
                        encoded, err := json.Marshal(md.MdV)
                        __err_panic(err)

                        err = b.Put([]byte(md.ID_Media), []byte(encoded))
                        __err_panic(err)


                    }
                //................................................



                } // switch sheet_Name {
            } // for _, sheet := range xlFile.Sheets


            return err
        })  // err = db.Update(func(tx *bolt.Tx) error 


    db.Close()
    

}
