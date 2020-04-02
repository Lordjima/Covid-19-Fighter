pragma solidity >=0.4.21 <0.7.0;
contract Covid {

  uint8 constant DATA_UNUSE = 0x01;
  uint8 constant DATA_IN_USE = 0x02;

  address public owner;

  struct Doctor {
    string nom;
    string IDnumber;
  }

  struct Patient {
    string nom;
    string mail;
    uint256 age;
    uint gender;
  }

  struct Treatment {
    string description;
    string medication;
  }

  mapping (address => Doctor) private doctors;
  address[] public ListDocs;
  mapping (address => Patient) private patients;
  address[] public ListPatient;
  mapping (address => Treatment) private treatments;
  address[] public ListTreatments;
  
  constructor() public {
    owner = msg.sender;
  }

  modifier onlyPatient (address _address) {
    //require (msg.sender == patients[_address].validPatient);
    _;
  }

  modifier onlyDoctor (address _address) {
    //require (msg.sender == doctors[_id].doctor);
    _;
  }

  function setDoctor (address _doctorAddress, string memory _nom, string memory _IDnumber) public {
    Doctor memory newDoctor = Doctor(_nom, _IDnumber);
    doctors[_doctorAddress] = newDoctor;
    ListDocs.push(_doctorAddress);
  }

  function getDoctor (address _address) public view returns (string memory, string memory){
    return (doctors[_address].nom, doctors[_address].IDnumber);
  }
  
  function getDoctorCount() public view returns(uint doctorCount) {
      return ListDocs.length;
  }

  function setPatient (address _patientAddress, string memory _nom, string memory _mail, uint256  _age, uint  _gender) public {
    Patient memory newPatient = Patient(_nom, _mail,_age, _gender);
    patients[_patientAddress] = newPatient;
    ListPatient.push(_patientAddress);
  }
  
  function setPatientTreatment (address _patientAddress, string memory _description, string memory _medication ) public {
    Treatment memory newTreatment = Treatment(_description, _medication);
    treatments[_patientAddress] = newTreatment;
    ListTreatments.push(_patientAddress);
  }
  
  function getPatient (address _address) public view returns (string memory, string memory, uint256, uint) {
      return (patients[_address].nom, patients[_address].mail, patients[_address].age, patients[_address].gender);
  }
  
  function getPatientTreatment (address _address) public view returns (string memory, string memory) {
      return (treatments[_address].description, treatments[_address].medication);
  }
  
  function getPatientCount() public view returns(uint patientCount) {
      return ListPatient.length;
  }
}