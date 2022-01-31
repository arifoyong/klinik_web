import {useState} from 'react'
import Layout from '../components/Layout'
import Modal from '../components/Modal'
import { PencilIcon, TrashIcon, PlusCircleIcon } from '@heroicons/react/solid'
import {getToday, dd_mmm_yyyy} from '../helpers/formatDate'
import {API} from '../config'



const Patient = ({data}) => {
  const [showModal, setShowModal] = useState(false)
  const [curPatient, setCurPatient] = useState(null)


  const registerPatient = async (e) => {
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
    if (result.data !== "success") {
      alert("error in adding data")
    } 
    setShowModal(false)
  }

  const dispModal = (dt) => {
    setCurPatient(dt)
    setShowModal(true)
  }

  const PatientModal = ({curPatient}) => (
    <Modal onClose = {() => setShowModal(false)} 
            show={showModal}
            title="Patient details"> 
      <form className="w-full max-w-sm flex-col" onSubmit={registerPatient} >
        <div className="md:flex md:items-center mb-6">
            <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="firstname">
              First name
            </label>
            <input id="firstname" type="text" placeholder="John" value={curPatient && curPatient.firstname} required 
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
        </div>
        <div className="md:flex md:items-center mb-6">
            <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="lastname">
              Last name
            </label>
            <input id="lastname" type="text" placeholder="Doe" value={curPatient && curPatient.lastname} required
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
        </div>
        <div className="md:flex md:items-center mb-6">
            <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="dob">
              Date of birth
            </label>
            <input id="dob" type="date" placeholder="31/12/1990" max={getToday()} value={curPatient && curPatient.dob} required
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
        </div>
        <div className="md:flex md:items-center mb-6">
            <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="email">
              Email
            </label>
            <input id="email" type="email" placeholder="johndoe@email.com" value={curPatient && curPatient.email} required
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
        </div>
        <div className="md:flex md:items-center mb-6">
            <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="ic">
              IC no
            </label>
            <input id="ic" type="text" placeholder="B71637182" value={curPatient && curPatient.ic} required
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
        </div>
        <div className="md:flex md:items-center mb-6">
            <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="phone">
              Phone no
            </label>
            <input id="phone" type="text" placeholder="85228191" value={curPatient && curPatient.phone} required
                className="md:w-2/3 bg-gray-200 py-1 px-4 text-gray-700 border-gray-200 border-2 rounded focus:bg-white focus:border-blue-500"/>
        </div>
        <div className="md:flex md:items-center mb-6">
            <label className="md:w-1/3 text-gray-500  block pr-4 md:text-right" htmlFor="address">
              Address
            </label>
            <textarea id="address" type="text" placeholder="81 Burlington Avenue 5, Singapore" value={curPatient && curPatient.address}
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

  return (
  <Layout>
    <main className="px-4">
      <h2 className="text-3xl py-2">List of Patients</h2>
      <PatientModal curPatient={curPatient}/>
      <table className="table-auto w-full">
        <thead className="bg-blue-700 text-white">
          <tr>
            <td className="px-2 py-2">Firstname</td>
            <td className="px-2 py-2">Lastname</td>
            <td className="px-2 py-2">IC</td>
            <td className="px-2 py-2">DOB</td>
            <td className="px-2 py-2">Email</td>
            <td className="px-2 py-2">Phone</td>
            <td className="px-2 py-2">Address</td>
            <td className="px-2 py-2">
              <PlusCircleIcon className="w-6 h-6 text-white 
                            hover:cursor-pointer hover:text-blue-200"
                onClick={() => dispModal(null)} />
            </td>
          </tr>
        </thead>
        <tbody> 
          {data.map((patient, idx) => (
            <tr key={idx}>
              <td className="px-2 py-2">{patient.firstname}</td>
              <td className="px-2 py-2">{patient.lastname}</td>
              <td className="px-2 py-2">{patient.ic}</td>
              <td className="px-2 py-2">{dd_mmm_yyyy(patient.dob)}</td>
              <td className="px-2 py-2">{patient.email}</td>
              <td className="px-2 py-2">{patient.phone}</td>
              <td className="px-2 py-2">{patient.address}</td>
              <td className="px-2 py-2">
                <div className="flex space-x-1">
                  <PencilIcon className="w-4 h-4" onClick={() => dispModal(patient)}/>
                  <TrashIcon className="w-4 h-4"/>
                </div>
                </td>
            </tr>
            ))}
        </tbody>
      </table>
    </main>
  </Layout>
  )
}

export async function getStaticProps(context) {
  const res = await fetch(`${API}/patients`)
  const data =await res.json()

  if (!data) {
    return {
      notFound : true
    }
  }

  return {
    props: data
  }
}

export default Patient