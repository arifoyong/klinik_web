import {useState} from 'react'
import Layout from '../components/Layout'
import PatientModal from '../components//Modal/PatientModal'
import { PencilIcon, TrashIcon, PlusCircleIcon } from '@heroicons/react/solid'
import { dd_mmm_yyyy} from '../helpers/formatDate'
import {API} from '../config'
import {toast} from '../components/Alert/Alert'

const blankPatient = {
  id: -1,
  firstname: "",
  lastname:"",
  ic:"",
  dob:"",
  email:"",
  phone:"",
  address:""
}



const Patient = ({data}) => {
  const [showModal, setShowModal] = useState(false)
  const [allPatients, setAllPatients] = useState(data)
  const [selectedPatient, setselectedPatient] = useState(blankPatient)

  const openModal = (patient) => {
    setselectedPatient(patient)
    setShowModal(true)
  }

  const closeModal = async () => {
    setShowModal(false)

    const res = await fetch(`${API}/patients`)
    const jsonResult =await res.json()
    setAllPatients(jsonResult.data)
  }

  const deletePatient = async (patient) => {

    var confirmDelete = confirm(`Delete ${patient.id} - ${patient.firstname} ${patient.lastname}. Are you sure?`);
    if (confirmDelete) {
      //Logic to delete the item
      const res = await fetch(
        `${API}/patient/${patient.id}`,
        {
          headers: {
            'Content-Type': 'application/json'
          },
          method: 'DELETE'
        }
      )
  
      const result = await res.json()
      console.log(result)
      if (result.error) {
        toast.notify(result.error)
      }
  
      const allRes = await fetch(`${API}/patients`)
      const jsonResult =await allRes.json()
      setAllPatients(jsonResult.data)    
    }

  }

  return (
  <Layout>
    <main>
      <h2 className="text-3xl py-2">List of Patients</h2>
      {showModal && <PatientModal currentPatient={selectedPatient} closeModal={closeModal}/> }
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
                onClick={() => openModal(blankPatient)} />
            </td>
          </tr>
        </thead>
        <tbody> 
          {allPatients.map((patient) => (
            <tr key={patient.id}>
              <td className="px-2 py-2">{patient.firstname}</td>
              <td className="px-2 py-2">{patient.lastname}</td>
              <td className="px-2 py-2">{patient.ic}</td>
              <td className="px-2 py-2">{dd_mmm_yyyy(patient.dob)}</td>
              <td className="px-2 py-2">{patient.email}</td>
              <td className="px-2 py-2">{patient.phone}</td>
              <td className="px-2 py-2">{patient.address}</td>
              <td className="px-2 py-2">
                <div className="flex space-x-1">
                  <PencilIcon className="w-4 h-4 hover:cursor-pointer hover:text-blue-700" 
                              onClick={() => openModal(patient)}/>
                  <TrashIcon className="w-4 h-4 hover:cursor-pointer hover:text-blue-700"
                              onClick={() => deletePatient(patient)}/>
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