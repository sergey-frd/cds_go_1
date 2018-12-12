package lib

import (   
    "fmt"
	"github.com/boltdb/bolt"
    "os" 
	"errors"
	"log"
    "github.com/valyala/fastjson"
    // "github.com/tealeg/xlsx"
	// "io/ioutil"

    //"encoding/gob"

    //L "cds_go_1/lib"

    //"encoding/json"
    //S "cds_go_1/config"

    // "time"
    // "math/rand"

)

// type Time_Interval struct {
//     ID_Time_Interval string
//     Price            string
//     D_Sign_People    string
//     Slots            string
// }

//---------------------------------------------------------------
// func random(min, max int) int {
//     rand.Seed(time.Now().Unix())
//     return rand.Intn(max - min) + min
// }


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


    dbFileName := fastjson.GetString(byteValues, "Base", "dbFileName")
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


    dbFileName := fastjson.GetString(byteValues, "Base", "dbFileName")
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

    dbFileName := fastjson.GetString(byteValues, "Base", "dbFileName")
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
