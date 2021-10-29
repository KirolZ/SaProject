export interface ScreeningInterface {

    ID: string, 
    Symptoms : string,
    Weight : number,
    Height: number,
    Temperature: number,
    PulseRate : number,
    RespirationRate: number,
    Savedate: Date | null,

    //fk
    MedRecID: number, 
    MedRec: MedicalRecordInterface
    DiseaseID: number, 
    Disease: DiseaseInterface
    NurseID: number, 
    Nurse : NurseInterface,
    SeveTime :Date,
}

export interface NurseInterface {

    ID: string, 
    Name : string,
    Email : string,
    Pass: string,
}

export interface MedicalRecordInterface {

    ID: string, 
    Hospital_number : string,
    Personal_ID : string,
    Patient_Name: string,
    Patient_Age: string,
    Patient_dob: string,
    Patient_Tel: string,

}

export interface DiseaseInterface {

    ID: string, 
    Name : string,
    Description : string,

}


