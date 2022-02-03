import {useState } from 'react'

import Modal from './Modal'
import {getToday, yyyy_mm_dd} from '../../helpers/formatDate'
import {API} from '../../config'
import {toast} from '../Alert/Alert'

const PatientModal = ({currentPatient,  closeModal}) => {
  const [patient, setPatient] = useState(currentPatient)

  const handleChange = (e) => {
    setPatient({...patient, [e.target.name]: e.target.value})
  }

  const savePatient = async (e) => {
    e.preventDefault()

    const dob = new Date(e.target.dob.value)
    const res = await fetch(
      `${API}/patients`,
      {
        body: JSON.stringify({
          firstname: e.target.firstname.value,
          lastname: e.target.lastname.value,
          ic: e.target.ic.value,
          dob: dob.toISOString(),
          email: e.target.email.value,
          phone: e.target.phone.value,
          address: e.target.address.value
        }),
        headers: {
          'Content-Type': 'application/json'
        },
        method: 'POST'
      }
    )

    const result = await res.json()
    console.log(result)
    if (result.error) {
      toast.notify(result.error, "error")
    } 

    closeModal()
  }

  const updatePatient = async (e) => {
    e.preventDefault()

    const dob = new Date(e.target.dob.value)
    const res = await fetch(
      `${API}/patients/${patient.id}`,
      {
        body: JSON.stringify({
          firstname: e.target.firstname.value,
          lastname: e.target.lastname.value,
          ic: e.target.ic.value,
          dob: dob.toISOString(),
          email: e.target.email.value,
          phone: e.target.phone.value,
          address: e.target.address.value
        }),
        headers: {
          'Content-Type': 'application/json'
        },
        method: 'PATCH'
      }
    )

    const result = await res.json()
    console.log(result)
    if (result.error) {
      toast.notify(result.error, "error")
    } 

    closeModal()
  }

 return (
  <Modal onClose = {closeModal} 
          title={`Patient details - ${patient.id}`}> 
    <form className="w-full max-w-sm flex-col" onSubmit={patient.id === -1 ? savePatient : updatePatient } >
      <div className="md:flex md:items-center mb-6">
          <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="firstname">
            First name
          </label>
          <input name="firstname" type="text" placeholder="John" value={patient.firstname} required
                onChange={handleChange} 
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
      </div>
      <div className="md:flex md:items-center mb-6">
          <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="lastname">
            Last name
          </label>
          <input name="lastname" type="text" placeholder="Doe" value={patient.lastname} required
                onChange={handleChange} 
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
      </div>
      <div className="md:flex md:items-center mb-6">
          <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="dob">
            Date of birth
          </label>
          <input name="dob" type="date" placeholder="31/12/1990" max={getToday()} value={yyyy_mm_dd(patient.dob)} required
                onChange={handleChange} 
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
      </div>
      <div className="md:flex md:items-center mb-6">
          <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="email">
            Email
          </label>
          <input name="email" type="email" placeholder="johndoe@email.com" value={patient.email} required
                onChange={handleChange} 
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
      </div>
      <div className="md:flex md:items-center mb-6">
          <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="ic">
            IC no
          </label>
          <input name="ic" type="text" placeholder="B71637182" value={patient.ic} required
                onChange={handleChange} 
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
      </div>
      <div className="md:flex md:items-center mb-6">
          <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="phone">
            Phone no
          </label>
          <input name="phone" type="text" placeholder="85228191" value={patient.phone} required
                onChange={handleChange} 
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
      </div>
      <div className="md:flex md:items-center mb-6">
          <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="address">
            Address
          </label>
          <textarea name="address" type="text" placeholder="81 Burlington Avenue 5, Singapore" value={patient.address}
                onChange={handleChange} 
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-base text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500">
          </textarea>
      </div>

      
      <div className="md:flex md:items-center mb-6">
        <div className="md:w-1/3"></div>
        <div className="md:w-2/3">
          <button type="submit"
                className="shadow bg-blue-500 text-white rounded py-2 px-4 hover:bg-blue-700 focus:shadow-outline focus:outline-none">
            Save</button>
        </div>
      </div>
    </form>
   
  </Modal>
  )
}

export default PatientModal