APIs


Auth
  /auth/send-otp
  /auht/verify

Patient
  [GET] -> { JWT } /patient -> [ Aadhar Data ]
  [GET] -> { JWT } /appoinments -> [ Array<{ Doctor Name, Dept., Notes: Array<String>  }> ]
  [GET] -> { JWT } /history -> [ Array<{ Hospital Info, Array<Prescription>, Array<Lab Reports>, Date, Title, Doctor name, Remarks }> ]
  <!-- [GET] -> { JWT } /medications [  ] -->