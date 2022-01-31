import {useState} from 'react'
import Layout from '../components/Layout'
import Modal from '../components/Modal'
import { PencilIcon, TrashIcon } from '@heroicons/react/solid'

const formatDate = (strDate) => {
  const monthNames =["Jan","Feb","Mar","Apr",
                      "May","Jun","Jul","Aug",
                      "Sep", "Oct","Nov","Dec"];
  const dt = new Date(strDate)

  const day = dt.getDate()
  const mth = monthNames[dt.getMonth()]
  const yr = dt.getFullYear()

  return `${day}-${mth}-${yr}`
}

const Patient = ({data}) => {
  const [showModal, setShowModal] = useState(false)


  return (
  <Layout>
    <main className="px-4">
      <p className="text-3xl py-4">Patient page</p>
      <p onClick={() => setShowModal(true)}>add</p>
      <Modal onClose = {() => setShowModal(false) } show={showModal}> 
        <p>This is from the children</p>
      </Modal>

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
            <td className="px-2 py-2"></td>
          </tr>
        </thead>
        <tbody> 
          {data.map((patient, idx) => (
            <tr key={idx}>
              <td className="px-2 py-2">{patient.firstname}</td>
              <td className="px-2 py-2">{patient.lastname}</td>
              <td className="px-2 py-2">{patient.ic}</td>
              <td className="px-2 py-2">{formatDate(patient.dob)}</td>
              <td className="px-2 py-2">{patient.email}</td>
              <td className="px-2 py-2">{patient.phone}</td>
              <td className="px-2 py-2">{patient.address}</td>
              <td className="px-2 py-2">
                <div className="flex space-x-1">
                  <PencilIcon className="w-4 h-4"/>
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
  const res = await fetch(`http://localhost:8000/patients`)
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