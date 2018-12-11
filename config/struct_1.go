package strct


//.............................
type Ow_Um_KEY   struct {
    UsMd                User_Media_KEY
    ID_Owner            string
}

type Ow_Um_STC   struct {
    OwUm_Key            Ow_Um_KEY
    Total_Cost          float64
}

type Ow_UmNbDs_KEY   struct {
    UsMd                User_Media_KEY
    CnCtNbDs            Digital_Signage_KEY
}

type Ow_UmNbDs_STC   struct {
    OwUmNbDs_Key       Ow_UmNbDs_KEY
    Total_Cost         float64
}

//.............................
type Ow_UmNbDsTi_KEY   struct {
    UsMd                User_Media_KEY
    NbDsTi_key          NbDsTi_KEY
}

type Ow_UmNbDsTi_STC   struct {
    OwUmNbDsTi_Key     Ow_UmNbDsTi_KEY
    Total_Cost         float64
}
//.............................
type NbDsTi_KEY   struct {
    CnCtNbDs          Digital_Signage_KEY
    ID_Time_Interval  string
}

//.............................
//.............................
type NbDsTiSl_KEY   struct {
    NbDsTi_key  NbDsTi_KEY
    ID_Slot     string
}

type NbDsTiSl_STC   struct {
     NbDsTiSl_Key   NbDsTiSl_KEY
     Slot_Price     float64
}

//.............................
type Um_NbDsTiSl_KEY   struct {
    UsMd          User_Media_KEY
    NbDsTiSl_key  NbDsTiSl_KEY
}

type Um_NbDsTiSl_STC   struct {
    UmNbDsTiSl_Key     Um_NbDsTiSl_KEY
    Slot_Price         float64
}


//.............................
type User_Media_KEY   struct {
    ID_User     string
    ID_Media    string
}

type User_Media_VAL   struct {
    Media_Name  string
    Media_Cost  string
    Media_Slots string
}

type User_Media_STC   struct {
    UsMd        User_Media_KEY
    UsMdVl      User_Media_VAL
}

//.............................
type Media_STC   struct {
    ID_Media    string
    MdV         Md_VAL
}

type Md_VAL   struct {
    Type_Media  string
    Slots       string
}

//.............................
type Time_Interval_VAL   struct {
    Price            string
    D_Sign_People    string
    Slots            string
}

//.............................
type Time_Interval_STC   struct {
    ID_Time_Interval string
    TiVl             Time_Interval_VAL
}

//.............................
type Neighborhoods_STC   struct {
    CnCtNb           Neighborhoods_KEY
    Neighborhoods    string
}

//.............................
type Neighborhoods_KEY   struct {
    CnCt             City_KEY
    ID_Neighborhoods string
}
 
//.............................
type Digital_Signage_KEY   struct {
    CnCtNb              Neighborhoods_KEY
    ID_Digital_Signage  string
    ID_Owner            string
}
 
//.............................
type Digital_Signage_VAL   struct {
    Dig_Sign         string
    DS_Cost          string
    DS_Perc_Quality  string
}
 
//.............................
type Digital_Signage_STC   struct {
    CnCtNbDs  Digital_Signage_KEY
    DsVal     Digital_Signage_VAL
}

//............................
type City_STC   struct {
    CnCt   City_KEY
    City   string
}

//.............................
type City_KEY   struct {
    ID_Country       string
    ID_City          string
}


