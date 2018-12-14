package lib

import (   
    "fmt"
    "github.com/boltdb/bolt"
    "os" 
    "errors"
    "log"
    "github.com/valyala/fastjson"
    "path/filepath"
    "math"
    "strconv"
    "time"

    // "github.com/tealeg/xlsx"
	// "io/ioutil"

    //"encoding/gob"

    //L "cds_go_1/lib"

    //"encoding/json"
    //S "cds_go_1/config"

    // "math/rand"

)

// type Time_Interval struct {
//     ID_Time_Interval string
//     Price            string
//     D_Sign_People    string
//     Slots            string
// }

//---------------------------------------------------------------
// func Random(min, max int) int {
//     rand.Seed(time.Now().Unix())
//     return rand.Intn(max - min) + min
// }

//------------------------------------------------------------
func GetExcelFileName( byteValues []byte) (excelFileName string, err error) {

    excel_Name := fastjson.GetString(byteValues, "Base", "excel_Name")
    proj_dir, err := os.Getwd();  __err_panic(err) 
    excelFileName = filepath.Join(proj_dir,"tbl", excel_Name)

	return excelFileName, err

}

//------------------------------------------------------------
func GetDbName( byteValues []byte) (dbFileName string, err error) {

    db_Name := fastjson.GetString(byteValues, "Base", "db_Name")
    proj_dir, err := os.Getwd();  __err_panic(err) 
    dbFileName = filepath.Join(proj_dir,"db", db_Name)

	return dbFileName, err

}

//---------------------------------------------------------------
func __err_panic(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}
//------------------------------------------------------------
func Demo(dbname string, byteValues []byte) {
    fmt.Println("lib_1.go HI",dbname)
    //fmt.Println("byteValues =",byteValues)

    //project_name := fastjson.GetString(byteValues, "Base", "project_name")
    //fmt.Printf("Demo project_name = %s\n", project_name)

}

//------------------------------------------------------------
func SetupDB(dbname string) (*bolt.DB, error) {
	db, err := bolt.Open(dbname, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db=%s, %v",dbname, err)
	}
	fmt.Println("DB Setup Done",dbname)
	return db, nil
}

// //------------------------------------------------------------
// //func AddPair(tx *bolt.Tx, sheet_Name string, k string, v string) error {
// func AddPair(b *Bucket, k string, v string) error {
// 
// 	//b, err := tx.CreateBucketIfNotExists([]byte(sheet_Name))
// 	//__err_panic(err)
// 
// 	err := b.Put([]byte(k), []byte(v))
// 	__err_panic(err)
// 
// 	return err
// }


//func Db_View(tx *bolt.Tx) error {
//	db.View(func(tx *bolt.Tx) error {
//		return tx.ForEach(func(nm []byte, b *bolt.Bucket) error {
//			bb, err := readBucket(b)
//			if err == nil {
//				bb.name = string(nm)
//				bb.expanded = false
//				memBolt.buckets = append(memBolt.buckets, *bb)
//				return nil
//			}
//			return err
//		})
//	})
//	return memBolt
//}

// Buckets prints a list of all buckets.
func Buckets(path string) {

    if _, err := os.Stat(path); os.IsNotExist(err) {
        fmt.Println(err)
        return
    }
    //fmt.Println(" ----------  L.Buckets Stat IsNotExist ---------- ")

    db, err := bolt.Open(path, 0600, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer db.Close()

    //fmt.Println(" ----------  Open ---------- ")
    err = db.View(func(tx *bolt.Tx) error {
        return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {

			sheet_Name := string(name)
            //fmt.Println(string(name))
			fmt.Printf(" *********  %s  ************* \n", sheet_Name)

			err = db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(sheet_Name))
				c := b.Cursor()
				if b == nil {
					return errors.New("Bucket not found")
				}
				//log.Printf("Value of %s is %s", "answer", b.Get([]byte("answer")))

				for k, v := c.First(); k != nil; k, v = c.Next() {
					//log.Printf("Value of %s is %s", k, v)
					log.Printf(" %s => %s", k, v)
				}
				return nil
			})

            return nil
        })
    })
    if err != nil {
        fmt.Println(err)
        return
    }
}

//----------------------------------------------
func LoadDict(byteValues []byte, data map[string]map[string]string) error {


    dbFileName, err := GetDbName(byteValues);  __err_panic(err) 
    if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
        fmt.Println(err)
        return err
    }
    //fmt.Println(" ----------  L.Buckets Stat IsNotExist ---------- ")

    db, err := bolt.Open(dbFileName, 0600, nil)
    if err != nil {
        fmt.Println(err)
        return err
    }
    defer db.Close()

    //fmt.Println(" ----------  Open ---------- ")
    err = db.View(func(tx *bolt.Tx) error {
        return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {

			sheet_Name := string(name)
            //fmt.Println(string(name))
			//fmt.Printf(" *** LoadDict %s *** \n", sheet_Name)
				 
			data[sheet_Name] = make(map[string]string)
			err = db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(sheet_Name))
				c := b.Cursor()
				if b == nil {
					return errors.New("Bucket not found")
				}
				//log.Printf("Value of %s is %s", "answer", b.Get([]byte("answer")))
			
				for k, v := c.First(); k != nil; k, v = c.Next() {

    				data[sheet_Name][string(k[:])] = string(v[:])


				}
				return nil
			})


            return nil
        })
    })

    if err != nil {
        fmt.Println(err)
        return err
    }
    return err

}

//----------------------------------------------
func LoadDict2(byteValues []byte, data map[string]map[string]string, bucket_Name string) error {


    dbFileName, err := GetDbName(byteValues);  __err_panic(err) 
    if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
        fmt.Println(err)
        return err
    }
    //fmt.Println(" ----------  L.Buckets Stat IsNotExist ---------- ")

    db, err := bolt.Open(dbFileName, 0600, nil)
    if err != nil {
        fmt.Println(err)
        return err
    }
    defer db.Close()

    //fmt.Println(" ----------  Open ---------- ")
    err = db.View(func(tx *bolt.Tx) error {

			sheet_Name := string(bucket_Name)
            //fmt.Println(string(name))
			fmt.Printf(" *** %s *** \n", sheet_Name)
				 
			data[sheet_Name] = make(map[string]string)
			err = db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(sheet_Name))
				c := b.Cursor()
				if b == nil {
					return errors.New("Bucket not found")
				}
				//log.Printf("Value of %s is %s", "answer", b.Get([]byte("answer")))
			
				for k, v := c.First(); k != nil; k, v = c.Next() {

    				data[sheet_Name][string(k[:])] = string(v[:])
                    //fmt.Println( string(k[:]), "=>",string(v[:]))


				}
				return nil
			})


            return nil
    })

    if err != nil {
        fmt.Println(err)
        return err
    }
    return err

}


//----------------------------------------------
func LoadDict_Dbg(byteValues []byte, 
    data map[string]map[string]string, 
    bucket_Name string,
    ) error {


    fmt.Println("LoadDict_Dbg bucket_Name =", bucket_Name)

    dbFileName, err := GetDbName(byteValues);  __err_panic(err) 
    if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
        fmt.Println(err)
        return err
    }
    //fmt.Println(" ----------  L.Buckets Stat IsNotExist ---------- ")

    db, err := bolt.Open(dbFileName, 0600, nil)
    if err != nil {
        fmt.Println(err)
        return err
    }
    defer db.Close()

    fmt.Println(" ----------  Open ---------- ")
    err = db.View(func(tx *bolt.Tx) error {

			sheet_Name := string(bucket_Name)
            //fmt.Println(string(name))
			fmt.Printf(" *** %s *** \n", sheet_Name)
				 
			data[sheet_Name] = make(map[string]string)

			err = db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte(sheet_Name))
				c := b.Cursor()
				if b == nil {
					return errors.New("Bucket not found")
				}
			
				for k, v := c.First(); k != nil; k, v = c.Next() {

    				data[sheet_Name][string(k[:])] = string(v[:])
                    fmt.Println( string(k[:]), "=>",string(v[:]))


				}
				return nil
			}) // db.View(func(tx *bolt.Tx)


            return nil
    })

    if err != nil {
        fmt.Println(err)
        return err
    }
    return err

}


//------------------------------------------------------------------------------
func Diff(a, b time.Time) (year, month, day, hour, min, sec int) {
    if a.Location() != b.Location() {
        b = b.In(a.Location())
    }
    if a.After(b) {
        a, b = b, a
    }
    y1, M1, d1 := a.Date()
    y2, M2, d2 := b.Date()

    h1, m1, s1 := a.Clock()
    h2, m2, s2 := b.Clock()

    year = int(y2 - y1)
    month = int(M2 - M1)
    day = int(d2 - d1)
    hour = int(h2 - h1)
    min = int(m2 - m1)
    sec = int(s2 - s1)

    // Normalize negative values
    if sec < 0 {
        sec += 60
        min--
    }
    if min < 0 {
        min += 60
        hour--
    }
    if hour < 0 {
        hour += 24
        day--
    }
    if day < 0 {
        // days in month:
        t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
        day += 32 - t.Day()
        month--
    }
    if month < 0 {
        month += 12
        year--
    }

    return
}

//------------------------------------------------------------------------------
func Plural(count int, singular string) (result string) {
        if (count == 1) || (count == 0) {
         result = strconv.Itoa(count) + " " + singular + " "
        } else {
         result = strconv.Itoa(count) + " " + singular + "s "
        }
 return
}

func SecondsToHuman(input int) (result string) {
        years := math.Floor(float64(input) / 60 / 60 / 24 / 7 / 30 / 12)
        seconds := input % (60 * 60 * 24 * 7 * 30 * 12)
        months := math.Floor(float64(seconds) / 60 / 60 / 24 / 7 / 30)
        seconds = input % (60 * 60 * 24 * 7 * 30)
        weeks := math.Floor(float64(seconds) / 60 / 60 / 24 / 7)
        seconds = input % (60 * 60 * 24 * 7)
        days := math.Floor(float64(seconds) / 60 / 60 / 24)
        seconds = input % (60 * 60 * 24)
        hours := math.Floor(float64(seconds) / 60 / 60)
        seconds = input % (60 * 60)
        minutes := math.Floor(float64(seconds) / 60)
        seconds = input % 60

        if years > 0 {
                 result = Plural(int(years), "year") + Plural(int(months), "month") + Plural(int(weeks), "week") + Plural(int(days), "day") + Plural(int(hours), "hour") + Plural(int(minutes), "minute") + Plural(int(seconds), "second")
        } else if months > 0 {
                 result = Plural(int(months), "month") + Plural(int(weeks), "week") + Plural(int(days), "day") + Plural(int(hours), "hour") + Plural(int(minutes), "minute") + Plural(int(seconds), "second")
        } else if weeks > 0 {
                 result = Plural(int(weeks), "week") + Plural(int(days), "day") + Plural(int(hours), "hour") + Plural(int(minutes), "minute") + Plural(int(seconds), "second")
        } else if days > 0 {
                 result = Plural(int(days), "day") + Plural(int(hours), "hour") + Plural(int(minutes), "minute") + Plural(int(seconds), "second")
        } else if hours > 0 {
                 result = Plural(int(hours), "hour") + Plural(int(minutes), "minute") + Plural(int(seconds), "second")
        } else if minutes > 0 {
                 result = Plural(int(minutes), "minute") + Plural(int(seconds), "second")
        } else {
                 result = Plural(int(seconds), "second")
        }

 return
}

