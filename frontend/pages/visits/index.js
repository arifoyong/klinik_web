import {useState} from 'react'
import Link from 'next/link'
import Layout from '../../components/Layout'
import SearchPatientModal from '../../components/Modal/SearchPatientModal'
import { PencilIcon, TrashIcon, PlusCircleIcon } from '@heroicons/react/solid'
import { dd_mmm_yyyy} from '../../helpers/formatDate'
import {API} from '../../config'
import {toast} from '../../components/Alert/Alert'


const Visit = ({data}) => {
  const [patients, setPatients] = useState(data)
  const [showModal, setShowModal] = useState(false)

  const closeModal = async () => {
    setShowModal(false)

    const res = await fetch(`${API}/visits`)
    const jsonRes =await res.json()
    setPatients(jsonRes.data)
  }

  return (
  <Layout>
      <h2 className="text-3xl py-2">Visits</h2>
      {showModal && <SearchPatientModal closeModal={closeModal}/>  }
      <table className="table-auto w-full">
        <thead className="bg-blue-700 text-white">
          <tr>
          <td className="px-2 py-2">Status</td>
            <td className="px-2 py-2">Name</td>
            <td className="px-2 py-2">IC</td>
            <td className="px-2 py-2">Visit date</td>
            <td className="px-2 py-2">Problems</td>
            <td className="px-2 py-2">Diagnosis</td>
            <td className="px-2 py-2">Prescription ID</td>
            <td className="px-2 py-2">
              <PlusCircleIcon className="w-6 h-6 text-white 
                            hover:cursor-pointer hover:text-blue-200"
                onClick={() => setShowModal(true)} />
            </td>
          </tr>
        </thead>
        <tbody> 
          {patients && patients.map((visit) => (
            <tr key={visit.id}>
              <td className="px-2 py-2">{visit.status}</td>
              <td className="px-2 py-2">{visit.firstname} {visit.lastname}</td>
              <td className="px-2 py-2">{visit.ic}</td>
              <td className="px-2 py-2">{dd_mmm_yyyy(visit.date)}</td>
              <td className="px-2 py-2">{visit.problems}</td>
              <td className="px-2 py-2">{visit.diagnosis}</td>
              <td className="px-2 py-2">{visit.prescription_id}</td>
              <td className="px-2 py-2">
                <div className="flex space-x-1">
                  <Link href = {`/visits/${visit.id}`}>
                    <PencilIcon className="w-4 h-4 hover:cursor-pointer hover:text-blue-700" 
                                />
                  </Link>
                  <TrashIcon className="w-4 h-4 hover:cursor-pointer hover:text-blue-700"
                              />
                </div>
                </td>
            </tr>
            ))}
        </tbody>
      </table>
    
  </Layout>
  )
}

export async function getStaticProps(context) {
  const res = await fetch(`${API}/visits`)
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

export default Visit