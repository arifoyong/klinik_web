import {useState, useEffect } from 'react'
import Modal from './Modal'
import {getToday, yyyy_mm_dd} from '../../helpers/formatDate'
import {API} from '../../config'
import {toast} from '../Alert/Alert'


const blankPatient = {
  firstname: "",
  lastname: "", 
  ic: "", 
  dob: "",
  email: "", 
  phone: "", 
  address: "", 
  id: ""
}

const SearchPatientModal = ({ closeModal}) => {
  const [curPatient, setCurPatient] = useState(blankPatient)

  const [patients, setPatients] = useState([])
  const [suggestionActive, setSuggestionActive] = useState(false)
  const [suggestionList, setSuggestionList] = useState([])
  const [cursor, setCursor] = useState(0)

  const getPatients = async (query) => {
      const res = await fetch(`${API}/patients/name/${query}`)
      const data = await res.json()
      
     setPatients(data.data)
     setSuggestionList(filterItems(data.data, query))
     setSuggestionActive(true)
  }
  

  const filterItems = (arr, query) => {
    return arr && arr.filter( (el) => {
      return el.firstname.toLowerCase().indexOf(query.toLowerCase()) !== -1
    })
  }

  const handleChange = async (e) => {
    setCurPatient({...curPatient, [e.target.name]: e.target.value})

    if (e.target.name === "firstname") {
      (e.target.value.length === 1) &&  getPatients(e.target.value)
      
      if (e.target.value.length >= 1) {
        setSuggestionList(filterItems(patients, e.target.value))
        setSuggestionActive(true)
      } else {
        setSuggestionActive(false)
        setCurPatient(blankPatient)
        setCursor(0)
      }
    }
  }

  const Suggestions = () => (
    <ul className="absolute top-8 left-0 
                    w-full bg-white shadow-lg">
      {suggestionList.map((p,i) => (
        <li  key={i} 
          className={`${cursor === i ? "bg-blue-100" : "bg-white"} border border-blue-100 py-1 px-4 hover:bg-blue-100`}
          onClick={() => {
            setCurPatient(p)
            setSuggestionActive(false)
          }}
         >{p.firstname} {p.lastname}</li>
      ))  }
    </ul>
  )


  const handleKeyDown = (e) => {
    if (e.keyCode === 13) {
      setCurPatient(suggestionList[cursor])
      setSuggestionActive(false)
    }
    if (e.keyCode === 38 && cursor > 0) {
      setCursor(cursor-1)
    }

    if (e.keyCode === 40 && cursor < suggestionList.length) {
      setCursor(cursor+1)
    }
  }

  const handleSubmit = async () => {
    const curDate = new Date()
    const res = await fetch(`${API}/visits`, {
      body: JSON.stringify({
        date: curDate.toISOString(),
        patient_id: curPatient.id
      }),
      headers: {'Content-Type': 'application/json'},
      method: 'POST'
    })

    const result = await res.json()
    console.log(result)
    if (result.error) {
      toast.notify(result.error, "error")
    } 

    closeModal()
  }
  

  const test = () => {
    console.log("Patient", patients)
    // setSuggestionActive(!suggestionActive)
    // console.log(Object.keys(patients))
    // patients.map(p => {
    //   console.log(p)
    // })
  }
  
 return (
  <Modal onClose = {closeModal} 
          title={`Patient details - ${curPatient.id}`}> 
    <button onClick = {() => test()}>Test</button>
    <div autoComplete="off" className="w-full max-w-sm flex-col" >
      <div className="md:flex md:items-center mb-6">
          <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="firstname">
            First name
          </label>
          <div className="md:w-2/3 flex flex-col relative">
            <input autoComplete="off" name="firstname" type="text" placeholder="John" value={curPatient.firstname} required
                onChange={handleChange} 
                onKeyDown={handleKeyDown}
                className="bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
      
           
            { suggestionActive && <Suggestions /> }
          </div>
      </div>

      <div className="md:flex md:items-center mb-6">
          <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="lastname">
            Last name
          </label>
          <input name="lastname" type="text" placeholder="Doe" value={curPatient.lastname} required 
                onChange={handleChange} 
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
      </div>
      <div className="md:flex md:items-center mb-6">
          <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="dob">
            Date of birth
          </label>
          <input name="dob" type="date" placeholder="31/12/1990" max={getToday()} value={yyyy_mm_dd(curPatient.dob)} required
                onChange={handleChange} disabled
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
      </div>
      <div className="md:flex md:items-center mb-6">
          <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="email">
            Email
          </label>
          <input name="email" type="email" placeholder="johndoe@email.com" value={curPatient.email} required
                onChange={handleChange} disabled
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
      </div>
      <div className="md:flex md:items-center mb-6">
          <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="ic">
            IC no
          </label>
          <input name="ic" type="text" placeholder="B71637182" value={curPatient.ic} required
                onChange={handleChange} disabled
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
      </div>
      <div className="md:flex md:items-center mb-6">
          <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="phone">
            Phone no
          </label>
          <input name="phone" type="text" placeholder="85228191" value={curPatient.phone} required
                onChange={handleChange} disabled
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
      </div>
      <div className="md:flex md:items-center mb-6">
          <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="address">
            Address
          </label>
          <textarea name="address" type="text" placeholder="81 Burlington Avenue 5, Singapore" value={curPatient.address}
                onChange={handleChange} disabled
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-base text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500">
          </textarea>
      </div>

      
      <div className="md:flex md:items-center mb-6">
        <div className="md:w-1/3"></div>
        <div className="md:w-2/3">
          <button onClick={() => handleSubmit()}
                disabled={curPatient.id === ""}
                className="shadow bg-blue-500 text-white rounded py-2 px-4 
                          disabled: bg-blue-300
                          hover:bg-blue-700 focus:shadow-outline focus:outline-none">
            Register</button>
        </div>
      </div>
    </div>
   
  </Modal>
  )
}

export default SearchPatientModal